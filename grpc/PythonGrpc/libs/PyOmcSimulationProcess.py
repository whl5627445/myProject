import json
import os
import subprocess
import threading
import time
import psutil
from config.redis_config import R
from libs.OMPython.OMCSessionZMQ import OMCSessionZMQ
from libs.defs import update_records, new_another_name
from config.db_config import Session, YssimSimulateRecords, YssimModels


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

    def load_dependent_library(self):
        # 通过records_id找到包id
        with Session() as session:
            # 获取包信息
            record = session.query(YssimModels, YssimSimulateRecords).filter(
                YssimModels.id == YssimSimulateRecords.package_id).filter(
                YssimSimulateRecords.id == self.request.uuid,
                YssimSimulateRecords.deleted_at.is_(None)).first()
        # 判断是不是系统模型
        if record[0].sys_or_user == 'sys':
            # 是系统模型,加载系统模型
            res = self.omc_obj.sendExpression(
                "loadModel(" + record[0].package_name + ", {\"" + record[0].version + "\"},true,\"\",false)")
            print("加载系统模型", record[0].package_name, record[0].version, res)
        else:
            # 不是系统模型,先加载mo文件
            self.omc_obj.loadFile("/home/simtek/code/" + record[0].file_path)
            # 获取依赖包
            dependent_library_list = self.omc_obj.getUses(record[0].package_name)
            print("dependent_library_list:", dependent_library_list)
            for i in dependent_library_list:
                if i != "":
                    with Session() as session:
                        # 获取依赖包信息
                        package = session.query(YssimModels).filter(
                            YssimModels.package_name == i[0],
                            YssimModels.version == i[1],
                            YssimModels.sys_or_user == self.request.userName,
                            YssimModels.userspace_id == self.request.userSpaceId,
                            YssimSimulateRecords.deleted_at.is_(None)).first()
                    # if package.sys_or_user == "sys":
                    #     print("加载系统模型",package.package_name)
                    #     self.omc_obj.sendExpression(
                    #         "loadModel(" + package.package_name + ", {\"" + package.version + "\"},true,\"\",false)")
                    # 加载模型库或者mo文件
                    if package:
                        self.omc_obj.loadFile("/home/simtek/code/" + package.file_path)

    def run(self):

        self.omc_obj.sendExpression('setCommandLineOptions("-d=nogen,noevalfunc,newInst,nfAPI")')
        self.omc_obj.sendExpression('setCommandLineOptions("-d=nogen,noevalfunc,newInst,nfAPI")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")')
        self.omc_obj.sendExpression(
            'setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection")')
        self.omc_obj.sendExpression(
            "setModelicaPath(\"/usr/lib/omlibrary\")")

        # 加载模型的依赖
        self.load_dependent_library()

        self.processStartTime = time.time()
        self.state = "compiling"  # 编译中
        print("开始编译")
        # 编译
        update_records(uuid=self.uuid, simulate_status="6", simulate_start_time=time.time())
        json_data = {"message": self.request.simulateModelName + " 模型正在编译"}
        R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))

        absolute_path = r"/home/simtek/code/" + self.request.resultFilePath
        print("absolute_path:", absolute_path)
        buildModelRes = self.omc_obj.buildModel(className=self.request.simulateModelName,
                                                fileNamePrefix=absolute_path,
                                                simulate_parameters_data=self.request.simulationPraData
                                                )
        print("编译结果", buildModelRes)
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
            # 改数据库状态为2
            update_records(uuid=self.uuid, simulate_status="2", simulate_start="1")
        else:
            # 改数据库状态为3
            update_records(uuid=self.uuid, simulate_status="3", simulate_start="0", simulate_result_str="编译失败")
            json_data = {"message": self.request.simulateModelName + " 模编译失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            self.state = "stopped"
            return
        # 编译完成，通知omc进程退出，杀死父进程
        print(self.omc_obj.omc_process.pid, "编译完成，杀死omc进程")

        # 仿真
        self.state = "running"
        cmd = [absolute_path + "result"]
        process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        self.run_pid = process.pid
        # 获取命令行输出结果
        output, error = process.communicate()
        if error:
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
                json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                update_records(uuid=self.uuid,
                               simulate_model_result_path=self.request.resultFilePath,
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
