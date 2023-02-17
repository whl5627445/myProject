import time
from fmpy.fmi2 import fmi2Warning
from db_config.config import Session, YssimSimulateRecords
from fmpy import *
import zarr
from multiprocessing import Process


def TimeStampToTime(timestamp):
    timeStruct = time.localtime(timestamp)
    return time.strftime('%Y-%m-%d %H:%M:%S', timeStruct)


# def updateDb(uuid, log, state, processStartTime, processRunTime):
#     with Session() as session:
#         processDetails = session.query(YssimSimulateRecords).filter(YssimSimulateRecords.uuid == uuid).first()
#         if processDetails:
#             # processDetails.progress = progress
#             # processDetails.exception = exception
#             processDetails.simulate_result_str = log
#             processDetails.simulate_status = state
#             processDetails.simulate_start_time = processStartTime
#             processDetails.simulate_end_time = processRunTime
#             session.commit()


def saveZarr(path, ojb):
    zarr.save(path, ojb)


class MyProcess(Process):
    def __init__(self, request, managerResDict):
        Process.__init__(self)
        self.newFmuPath = "/yssim-go/" + request.resPath + request.className.replace(".", "_") + ".fmu"
        self.uuid = request.uuid
        self.request = request
        self.processStartTime = 0
        self.outputs = []
        self.AllLogTxt = ""
        self.progress1 = 0
        self.simulateRes = None
        self.managerResDict = managerResDict

        self.resPath = request.resPath+"zarr_res.zarr"
        # updateDb(uuid=self.uuid, progress=self.progress1, exception=0, log=self.AllLogTxt,
        #          state="初始化", processStartTime=None, processRunTime=None)
        with Session() as session:
            processDetails = session.query(YssimSimulateRecords).filter(YssimSimulateRecords.id == self.uuid).first()
            if processDetails:
                processDetails.simulate_status = "1"
                session.commit()

    def stepFinished(self, running_time, recorder):
        self.simulateRes = recorder.result()
        self.managerResDict[self.uuid] = recorder.result()
        # saveZarr(os.path.join(self.resFilePath, "result_res.zarr"), self.simulateRes)

        progress2 = int((running_time / self.request.stopTime) * 100)

        if progress2 > self.progress1:
            self.progress1 = progress2
            # with open("log.txt", "a+") as f:
            #     f.write(str(self.progress1) + "%" + '\t')
            #     if progress2 == 100:
            #         f.write('\n')

        # if self.progress1 == 100:
        #     stateString = "运行结束"
        # else:
        #     stateString = "正在运行"

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
        # with open("log.txt", "a+") as ff:
        #     ff.write(logTxt)
        self.AllLogTxt += logTxt

    def run(self):
        self.processStartTime = time.time()
        try:
            print("开始仿真")
            time1 = time.time()
            self.outputs = [v.name for v in read_model_description(self.newFmuPath).modelVariables]
            time2 = time.time()
            print("读取变量耗时：", time2 - time1)
            # updateDb(uuid=self.uuid, progress=self.progress1, exception=0, log=self.AllLogTxt,
            #          state="正在运行", processStartTime=TimeStampToTime(self.processStartTime),
            #          processRunTime=TimeStampToTime(time.time()))
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == self.uuid).first()
                if processDetails:
                    processDetails.simulate_result_str = self.AllLogTxt
                    processDetails.simulate_status = "2"
                    processDetails.simulate_start_time = self.processStartTime
                    processDetails.simulate_end_time = time.time()
                    session.commit()
            print(self.newFmuPath,
                  self.request.startTime,
                  self.request.stopTime,
                  "间隔:", self.request.outputInterval,
                  self.request.tolerance)
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
            log = "(error)" + str(e)
            print(log)
            # updateDb(uuid=self.uuid, progress=self.progress1, exception=1, log=self.AllLogTxt + log,
            #          state="运行结束", processStartTime=TimeStampToTime(self.processStartTime),
            #          processRunTime=TimeStampToTime(time.time()))

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
            # updateDb(uuid=self.uuid, progress=self.progress1, exception=0, log=self.AllLogTxt,
            #          state="运行结束", processStartTime=TimeStampToTime(self.processStartTime),
            #          processRunTime=TimeStampToTime(time.time()))
            print("运行正常结束。")
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == self.uuid).first()
                if processDetails:
                    processDetails.simulate_model_result_path = self.request.resPath
                    processDetails.simulate_result_str = self.AllLogTxt
                    processDetails.simulate_status = "4"
                    processDetails.simulate_start = "0"
                    processDetails.simulate_start_time = str(self.processStartTime)
                    processDetails.simulate_end_time = str(time.time())
                    session.commit()

        finally:
            if self.simulateRes is not None:
                print(self.resPath )
                saveZarr("/yssim-go/"+self.resPath, self.simulateRes)
            if self.uuid in self.managerResDict:
                del self.managerResDict[self.uuid]
        return "end"
