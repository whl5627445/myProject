import threading
import time
import psutil
import json
import os
import subprocess
from config.redis_config import R
from libs.OMPython.OMCSessionZMQ import OMCSessionZMQ
from libs.function.defs import update_compile_records, sendMessage


# 编译状态 1创建任务中  2 编译中 3编译失败 4编译成功

class OmcTranslateThread(threading.Thread):
    def __init__(self, request, port):
        self.state = "init"
        self.port = port
        self.run_pid = None
        self.uuid = request.uuid
        self.request = request
        update_compile_records(self.uuid,compile_status=2, compile_start_time=int(time.time()))
        threading.Thread.__init__(self)
        self.omc_obj = OMCSessionZMQ(port=port)

    def load_dependent_library(self):
        for key, val in self.request.envModelData.items():
            if "/" in val:
                # 初始化加载用户模型，key是名称，val是mo文件地址
                print(key + "初始化:", self.omc_obj.loadFile("/home/simtek/code/" + val))
            else:
                # 初始化加载系统模型，key是名称，val是版本号
                print(key + val + "初始化:",
                      self.omc_obj.sendExpression("loadModel(" + key + ", {\"" + val + "\"},true,\"\",false)"))
        # 获取注释中的包名和版本号
        name = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "name")
        version = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "version")
        print(name + version + "初始化:",
              self.omc_obj.sendExpression("loadModel(" + name + ", {\"" + version + "\"},true,\"\",false)"))

    def run(self):
        # omc准备操作
        self.omc_obj.sendExpression('setCommandLineOptions("-d=nogen,noevalfunc,newInst,nfAPI")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")')
        self.omc_obj.sendExpression(
            'setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection")')
        self.omc_obj.sendExpression(
            "setModelicaPath(\"/usr/lib/omlibrary\")")

        # 加载模型的依赖
        self.load_dependent_library()
        self.state = "compiling"  # 编译中
        print("开始编译")
        # 编译
        json_data = {"message": self.request.simulateModelName + " 模型正在编译"}
        R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))

        absolute_path = r"/home/simtek/code/" + self.request.resultFilePath
        print("absolute_path:", absolute_path)
        print("simulateModelName", self.request.simulateModelName)
        print("simulationPraData",self.request.simulationPraData)
        buildModelRes = self.omc_obj.buildModel(className=self.request.simulateModelName,
                                                fileNamePrefix=absolute_path,
                                                simulate_parameters_data=self.request.simulationPraData
                                                )
        print("编译结果", buildModelRes)
        print("编译结果的type", type(buildModelRes))
        sendMessage(self.omc_obj, self.request.userName)
        print("消息推送完成")
        parent_proc = psutil.Process(self.omc_obj.omc_process.pid)
        for child_proc in parent_proc.children(recursive=True):
            os.kill(child_proc.pid, 9)
        os.kill(self.omc_obj.omc_process.pid, 9)
        if isinstance(buildModelRes, list) and buildModelRes != ["", ""]:
            json_data = {"message": self.request.simulateModelName + " 模型编译完成,开始仿真"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        else:
            # 改数据库状态为3
            update_compile_records(uuid=self.uuid, compile_status=3, compile_stop_time=int(time.time()))
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
            print(error)
            update_compile_records(uuid=self.uuid, compile_status=3, compile_stop_time=int(time.time()))
            json_data = {"message": self.request.simulateModelName + " 仿真失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))

        else:
            simulate_result_str = output.decode('utf-8')
            if "successfully" in simulate_result_str:
                json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                update_compile_records(uuid=self.uuid, compile_status=4, compile_stop_time=int(time.time()))

            else:
                print("仿真失败:",simulate_result_str)
                json_data = {"message": self.request.simulateModelName + " 仿真失败"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                update_compile_records(uuid=self.uuid, compile_status=3, compile_stop_time=int(time.time()))
        # 仿真
        self.state = "stopped"
