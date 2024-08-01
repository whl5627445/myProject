import time
from fmpy.fmi2 import fmi2Warning
from config.db_config import Session, YssimSimulateRecords
from fmpy import *
import zarr
from multiprocessing import Process
from config.redis_config import R
import json, re


def new_another_name(username: str, simulate_model_name: str, userspace_id: str) -> str:
    another_name_list = []
    with Session() as session:
        record_list = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.username == username,
            YssimSimulateRecords.simulate_model_name == simulate_model_name,
            YssimSimulateRecords.userspace_id == userspace_id,
            YssimSimulateRecords.simulate_status == "4",
            YssimSimulateRecords.deleted_at.is_(None),
        ).all()

    for record in record_list:
        another_name_list.append(record.another_name)
    max_suffix = 0
    suffix_pattern = re.compile(r"\s(\d+)\s*$")
    for another_name in another_name_list:
        matches = suffix_pattern.findall(another_name)
        if len(matches) > 0:
            suffix = int(matches[0])
            if suffix > max_suffix:
                max_suffix = suffix

    return "结果 " + str(max_suffix + 1)


def TimeStampToTime(timestamp):
    timeStruct = time.localtime(timestamp)
    return time.strftime('%Y-%m-%d %H:%M:%S', timeStruct)


def saveZarr(path, ojb):
    zarr.save(path, ojb)


class MyProcess(Process):
    def __init__(self, request, managerResDict):
        Process.__init__(self)
        self.newFmuPath = "/home/simtek/code/" + request.resPath + request.className.replace(".", "_") + ".fmu"
        self.uuid = request.uuid
        self.request = request
        self.processStartTime = 0
        self.outputs = []
        self.AllLogTxt = ""
        self.progress1 = 0
        self.simulateRes = None
        self.managerResDict = managerResDict
        self.resPath = request.resPath + "zarr_res.zarr"
        with Session() as session:
            processDetails = session.query(YssimSimulateRecords).filter(YssimSimulateRecords.id == self.uuid).first()
            if processDetails:
                processDetails.simulate_status = "1"
                session.commit()

    def stepFinished(self, running_time, recorder):
        progress2 = int((running_time / self.request.stopTime) * 100)
        if progress2 > self.progress1:
            self.simulateRes = recorder.result()
            self.managerResDict[self.uuid] = recorder.result()
            self.progress1 = progress2
            print(self.progress1, r"%", end=" ")
        return True

    # 信息日志输出
    def logFMUMessage(self, *args):
        # works with FMI 1.0, 2.0, and 3.0
        status = args[-3]
        message = args[-1]
        if status == fmi2Warning:
            level = '(warning)'
        elif status > fmi2Warning:
            level = '(error)'
        else:
            level = '(info)'
        logTxt = level + "  " + message.decode('utf-8')
        self.AllLogTxt += logTxt

    def run(self):
        self.processStartTime = int(time.time())
        try:
            print("开始仿真,仿真结束时间{}，仿真间隔{}".format(self.request.stopTime, self.request.outputInterval))
            json_data = {"message": self.request.className + " FmPy开始仿真"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            time1 = time.time()
            self.outputs = [v.name for v in read_model_description(self.newFmuPath).modelVariables]
            time2 = time.time()
            print("读取变量耗时：", time2 - time1)
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == self.uuid).first()
                if processDetails:
                    processDetails.simulate_result_str = self.AllLogTxt
                    processDetails.simulate_status = "2"
                    processDetails.simulate_start_time = self.processStartTime
                    processDetails.simulate_end_time = time.time()
                    session.commit()
            simulate_fmu(self.newFmuPath,
                         start_time=self.request.startTime,
                         stop_time=self.request.stopTime,
                         output_interval=self.request.outputInterval,
                         start_values=dict(self.request.params),
                         output=self.outputs,
                         step_finished=self.stepFinished,
                         # relative_tolerance=self.request.tolerance,
                         logger=self.logFMUMessage
                         )

        except Exception as e:
            log = "(Exception)" + str(e)
            print(log)
            json_data = {"message": self.request.className + self.AllLogTxt + log}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == self.uuid).first()
                if processDetails:
                    processDetails.simulate_result_str = self.AllLogTxt + log
                    processDetails.simulate_status = "3"
                    processDetails.simulate_start = "0"
                    processDetails.simulate_start_time = str(self.processStartTime)
                    processDetails.simulate_end_time = str(time.time())
                    session.commit()

        else:
            print("运行正常结束。")
            json_data = {"message": self.request.className + " FmPy模型仿真完成"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == self.uuid).first()
                if processDetails:
                    processDetails.simulate_model_result_path = self.request.resPath
                    processDetails.simulate_result_str = self.AllLogTxt
                    processDetails.simulate_status = "4"
                    processDetails.another_name = new_another_name(self.request.userName, self.request.className,
                                                                   self.request.userSpaceId)
                    processDetails.simulate_start = "0"
                    processDetails.simulate_start_time = str(self.processStartTime)
                    processDetails.simulate_end_time = str(time.time())
                    session.commit()

        finally:
            if self.simulateRes is not None:
                saveZarr("/home/simtek/code/" + self.resPath, self.simulateRes)
            if self.uuid in self.managerResDict:
                del self.managerResDict[self.uuid]
        return "end"
