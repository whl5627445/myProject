import logging
import signal
import threading
from multiprocessing import Process
from libs.OMPython.OMCSessionZMQ import OMCSessionZMQ
from config.db_config import Session, YssimSimulateRecords
import subprocess
from config.redis_config import R
import json, re, time, os




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


def update_records(uuid, simulate_status=None, simulate_result_str=None, simulate_start=None, simulate_start_time=None,
                   simulate_end_time=None, simulate_model_result_path=None, another_name=None
                   ):
    with Session() as session:
        processDetails = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.id == uuid).first()
        if simulate_status:
            processDetails.simulate_status = simulate_status  # 更改状态
        if simulate_result_str:
            processDetails.simulate_result_str = simulate_result_str  # 更改仿真结果字符串
        if simulate_start:
            processDetails.simulate_start = simulate_start  # 仿真开始标致
        if simulate_start_time:
            processDetails.simulate_start_time = simulate_start_time  # 仿真开始时间
        if simulate_end_time:
            processDetails.simulate_end_time = simulate_end_time  # 仿真结束时间
        if simulate_model_result_path:
            processDetails.simulate_model_result_path = simulate_model_result_path  # 仿真结果文件路径
        if another_name:
            processDetails.another_name = another_name  # 结果记录别名
        session.commit()


import psutil


def kill_process_by_port(port):
    for conn in psutil.net_connections(kind='tcp'):
        if conn.laddr.port == port:
            pid = conn.pid
            psutil.Process(pid).kill()
            print(f"Killed process with PID: {pid}")


class PyOmcSimulation(threading.Thread):
    def __init__(self, request, port):
        self.state = "init"
        self.port = port
        self.run_pid = None
        self.uuid = request.uuid
        self.request = request
        self.processStartTime = None
        threading.Thread.__init__(self)
        self.omc_obj = OMCSessionZMQ(port=port)

    def run(self):

        self.omc_obj.sendExpression('setCommandLineOptions("-d=nogen,noevalfunc,newInst,nfAPI")')
        self.omc_obj.sendExpression(
            "setModelicaPath(\"/usr/lib/omlibrary\")")
        # print("clearProgram:", )
        # print("setModelicaPath:", )

        # print("Buildings9.1.0初始化:", self.omc_obj.sendExpression("loadModel(Buildings, {\"9.1.0\"},true,\"\",false)"))
        # print("Modelica4.0.0初始化:", self.omc_obj.sendExpression("loadModel(Modelica, {\"4.0.0\"},true,\"\",false)"))

        # 初始化加载系统模型，key是名称，val是版本号
        for key, val in self.request.sysModel.items():
            self.omc_obj.sendExpression("loadModel(" + key + ", {\"" + val + "\"},true,\"\",false)")
            # print(key + "初始化:",
            #       self.omc_obj.sendExpression("loadModel(" + key + ", {\"" + val + "\"},true,\"\",false)"))
        # 初始化加载用户模型，key是名称，val是mo文件地址
        for key, val in self.request.userModel.items():
            self.omc_obj.loadFile("/home/simtek/code/" + val)
            # print(key, val)
            # print(key + "初始化:", self.omc_obj.loadFile("/home/simtek/code/" + val))
        # # 测试omc buildModelFmu能不能用
        # fmuPath = self.omc_obj.buildModelFmu(className="Modelica.Blocks.Examples.PID_Controller", fileNamePrefix="xxx")
        # print("测试omc buildModelFmu能不能用PID_Controller:", fmuPath)
        # print("初始化omc完毕！端口号：", self.port)

        self.processStartTime = time.time()
        self.state = "compiling"  # 编译中
        # print("进程开始时间:", self.processStartTime)

        # 编译
        update_records(uuid=self.uuid, simulate_status="6", simulate_start_time=time.time())
        json_data = {"message": self.request.simulateModelName + " 模型正在编译"}
        R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        # print("resultFilePath:", self.request.resultFilePath)
        buildModelRes = self.omc_obj.buildModel(className=self.request.simulateModelName,
                                                fileNamePrefix=self.request.resultFilePath,
                                                simulate_parameters_data=self.request.simulationPraData
                                                )
        # print("buildModelRes:", buildModelRes)
        sendMessage(self.omc_obj, self.request.userName)
        print("消息推送完成")
        json_data = {"message": self.request.simulateModelName + " 模型编译完成，准备仿真"}
        R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        parent_proc = psutil.Process(self.omc_obj.omc_process.pid)
        for child_proc in parent_proc.children(recursive=True):
            print(child_proc.name)
            print(child_proc.pid)
            os.kill(child_proc.pid, 9)
        os.kill(self.omc_obj.omc_process.pid, 9)
        if buildModelRes != ["", ""]:
            # print("改数据库状态为2")
            update_records(uuid=self.uuid, simulate_status="2", simulate_start="1")
        else:
            # print("改数据库状态为3")
            update_records(uuid=self.uuid, simulate_status="3", simulate_start="0", simulate_result_str="编译失败")
            json_data = {"message": self.request.simulateModelName + " 模编译失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            return
        # 编译完成，通知omc进程退出，杀死父进程
        print(self.omc_obj.omc_process.pid, "编译完成，杀死omc进程")

        # 仿真

        self.state = "running"
        cmd = [self.request.resultFilePath + "result"]
        process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        self.run_pid = process.pid
        # 获取命令行输出结果
        output, error = process.communicate()
        if error:
            # print("subprocess Popen Error（仿真运行错误）: ", error)
            # print("仿真进程被kill")
            update_records(uuid=self.uuid,
                           simulate_status="3",
                           simulate_result_str="编译失败",
                           simulate_start="0",
                           simulate_start_time=str(self.processStartTime),
                           simulate_end_time=str(time.time())
                           )

        else:
            simulate_result_str = output.decode('utf-8')
            if "successfully" in simulate_result_str:
                # print("运行正常结束。")
                json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                update_records(uuid=self.uuid,
                               simulate_model_result_path=self.request.relativePath,
                               simulate_result_str=simulate_result_str,
                               simulate_status="4",
                               simulate_start="0",
                               simulate_start_time=str(self.processStartTime),
                               simulate_end_time=str(time.time()),
                               another_name=new_another_name(self.request.userName,
                                                             self.request.simulateModelName,
                                                             self.request.userSpaceId)
                               )

            else:
                update_records(uuid=self.uuid,
                               simulate_result_str=simulate_result_str,
                               simulate_status="3",
                               simulate_start="0",
                               simulate_start_time=str(self.processStartTime),
                               simulate_end_time=str(time.time())
                               )
        self.state = "stopped"


def sendMessage(omc_obj, username):
    message_str = omc_obj.getMessagesStringInternal()
    data_list = message_str.split(";,")
    message_list = []
    for i in data_list:
        dl = i.strip().split(",\n")
        message_dict = {}
        for p in dl:
            pl = p.strip()
            if "MODELICAPATH" in pl or "installPackage" in pl or "Downloaded" in pl:
                continue
            # elif "Automatically " in pl or "Lexer " in pl:
            #     continue
            elif pl.startswith("message"):
                mes = pl.replace("message = ", "", -1)
                message_dict["message"] = mes[1:-1]
                # print("mes", mes)
            elif pl.startswith("level"):
                level = pl.split(".")
                message_dict["type"] = level[-1]
                # print("level", level)
        if len(message_dict) > 1:
            message_list.append(message_dict)
    for i in message_list:
        message_notice(username, i)
    return message_list


def message_notice(username, mes):
    R.lpush(username + "_" + "notification", json.dumps(mes))
