import json
import os
import subprocess
import threading
import time
import psutil
from config.redis_config import R
from libs.OMPython.OMCSessionZMQ import OMCSessionZMQ
from libs.function.defs import update_simulate_records, new_another_name, sendMessage
from libs.function.grpc_log import log


class OmcSimulation(threading.Thread):
    def __init__(self, request, port):
        self.state = "init"
        self.port = port
        self.run_pid = None
        self.uuid = request.uuid
        self.request = request
        self.processStartTime = None
        threading.Thread.__init__(self)
        update_simulate_records(uuid=self.uuid, simulate_status="6", simulate_start_time=int(time.time()))
        self.omc_obj = OMCSessionZMQ(port=port)

    def load_dependent_library(self):
        for key, val in self.request.envModelData.items():
            log.info(key + val)
            if "/" in val:
                # 初始化加载用户模型，key是名称，val是mo文件地址
                log.info("(OMC)" + key + "初始化:" + str(self.omc_obj.loadFile("/home/simtek/code/" + val)))
            else:
                # 初始化加载系统模型，key是名称，val是版本号
                log.info("(OMC)" + key + "初始化:" +
                         str(self.omc_obj.sendExpression("loadModel(" + key + ", {\"" + val + "\"},true,\"\",false)")))
        # 获取注释中的包名和版本号
        name = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "name")
        version = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "version")
        if "not Found" not in name:
            log.info("(OMC)" + name + "初始化:" +
                     str(self.omc_obj.sendExpression("loadModel(" + name + ", {\"" + version + "\"},true,\"\",false)")))

    def run(self):

        # self.omc_obj.sendExpression('setCommandLineOptions("-d=nfAPI,initialization")')
        # self.omc_obj.sendExpression('setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")')
        # self.omc_obj.sendExpression('setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")')
        # self.omc_obj.sendExpression(
        #     'setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection")')
        # self.omc_obj.sendExpression(
        #     "setModelicaPath(\"/usr/lib/omlibrary\")")
        self.omc_obj.sendExpression('setCommandLineOptions("--simCodeTarget=C")')
        self.omc_obj.sendExpression('setCommandLineOptions("--target=gcc")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")')
        self.omc_obj.sendExpression('setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection --allowNonStandardModelica=reinitInAlgorithms -d=nfAPI,initialization,NLSanalyticJacobian")')
        self.omc_obj.sendExpression(
            "setModelicaPath(\"/usr/lib/omlibrary\")")
        # 加载模型的依赖
        self.load_dependent_library()

        # self.processStartTime = time.time()
        self.state = "compiling"  # 编译中
        log.info("(OMC)开始编译")
        # 编译

        absolute_path = r"/home/simtek/code/" + self.request.resultFilePath + "result"
        log.info("(OMC)仿真结果地址:" + absolute_path)
        log.info("(OMC)仿真模型名：" + self.request.simulateModelName)
        log.info("(OMC)仿真参数：" + str(self.request.simulationPraData))
        log.info("(OMC)OMC进程参数：" + str(self.omc_obj.__dict__))
        translateModelRes = self.omc_obj.translateModel(className=self.request.simulateModelName,
                                                        fileNamePrefix=absolute_path,
                                                        translate_parameters_data=self.request.simulationPraData
                                                        )
        log.info("(OMC)转换结果:" + str(translateModelRes))
        sendMessage(self.omc_obj, self.request.userName)
        log.info("(OMC)消息推送完成")
        parent_proc = psutil.Process(self.omc_obj.omc_process.pid)
        for child_proc in parent_proc.children(recursive=True):
            os.kill(child_proc.pid, 9)
        os.kill(self.omc_obj.omc_process.pid, 9)

        if translateModelRes:
            json_data = {"message": self.request.simulateModelName + " 模型代码转换c代码完成，准备编译"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        else:
            update_simulate_records(uuid=self.uuid, simulate_status="3",
                                    simulate_start="0",
                                    simulate_result_str="编译失败",
                                    # simulate_start_time=str(self.processStartTime),
                                    simulate_end_time=int(time.time()))
            self.state = "stopped"
            json_data = {"message": self.request.simulateModelName + " 模型代码转换c代码出现错误"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            return
        json_data = {"message": self.request.simulateModelName + " 模型正在编译"}
        R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        cmd = "make -j4 -f " + absolute_path + ".makefile"
        make_result = subprocess.getoutput(cmd)
        if "make: ***" not in make_result and translateModelRes == True:
            # 改数据库状态为2
            log.info("(OMC)编译成功")
            update_simulate_records(uuid=self.uuid, simulate_status="2", simulate_start="1")
            json_data = {"message": self.request.simulateModelName + " 模型编译完成，准备仿真"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        else:
            # 改数据库状态为3
            log.info("(OMC)编译失败")
            update_simulate_records(uuid=self.uuid, simulate_status="3",
                                    simulate_start="0",
                                    simulate_result_str="编译失败",
                                    # simulate_start_time=str(self.processStartTime),
                                    simulate_end_time=int(time.time()))
            json_data = {"message": self.request.simulateModelName + " 模编译失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            self.state = "stopped"
            return
        # 编译完成，通知omc进程退出，杀死父进程
        log.info("(OMC)编译完成，杀死omc进程：" + str(self.omc_obj.omc_process.pid))

        # 仿真
        self.state = "running"
        time1 = time.time()
        cmd = [absolute_path]
        process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        self.run_pid = process.pid
        # 获取命令行输出结果
        output, error = process.communicate()
        time2 = time.time()
        if error:
            log.info("(OMC)仿真失败,error:" + str(error))
            update_simulate_records(uuid=self.uuid,
                                    simulate_status="3",
                                    simulate_result_str="仿真失败",
                                    simulate_start="0",
                                    # simulate_start_time=str(self.processStartTime),
                                    simulate_end_time=int(time.time())
                                    )
            json_data = {"message": self.request.simulateModelName + " 仿真失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))

        else:
            simulate_result_str = output.decode('utf-8')
            if "successfully" in simulate_result_str:
                log.info("(OMC)模型仿真成功完成")
                json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                update_simulate_records(uuid=self.uuid,
                                        simulate_model_result_path=self.request.resultFilePath,
                                        simulate_result_str=simulate_result_str,
                                        simulate_status="4",
                                        simulate_start="0",
                                        result_run_time=time2 - time1,
                                        simulate_end_time=int(time.time()),
                                        another_name=new_another_name(self.request.userName,
                                                                      self.request.simulateModelName,
                                                                      self.request.simulatePackageId,
                                                                      self.request.userSpaceId)
                                        )

            else:
                log.info("(OMC)仿真失败:" + str(simulate_result_str))
                json_data = {"message": self.request.simulateModelName + " 仿真失败"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                update_simulate_records(uuid=self.uuid,
                                        simulate_result_str=simulate_result_str,
                                        simulate_status="3",
                                        simulate_start="0",
                                        # simulate_start_time=str(self.processStartTime),
                                        simulate_end_time=int(time.time())
                                        )
        log.info("(OMC)仿真线程执行完毕")
        self.state = "stopped"
