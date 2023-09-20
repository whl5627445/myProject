import shutil
import threading
import time
import psutil
import json
import os
import subprocess
import DyMat
import pandas as pd
from config.redis_config import R
from libs.OMPython.OMCSessionZMQ import OMCSessionZMQ
from libs.function.defs import update_compile_records, sendMessage, del_data_sources_records
from libs.function.grpc_log import log
from libs.function.defs import update_app_pages_records, omc_convert_dict_to_list, \
    update_app_spaces_records, page_preview_component_freeze, result_step
from libs.function.xml_input import calibration_write_xml
from libs.function.defs import update_parameter_calibration_records
from config.db_config import Session, ParameterCalibrationRecord


# 编译状态 1创建任务中  2 编译中 3编译失败 4编译成功

class CalibrationCompileThread(threading.Thread):
    def __init__(self, request, port):
        self.state = "init"
        self.port = port
        self.run_pid = None
        self.uuid = request.uuid
        self.request = request
        threading.Thread.__init__(self)
        self.omc_obj = OMCSessionZMQ(port=port)

    def load_dependent_library(self):
        for key, val in self.request.envModelData.items():
            # 初始化加载用户模型，key是名称，val是mo文件地址
            # log.info("初始化:" + str(self.omc_obj.sendExpression("loadModel(" + key + ", {\"" + val + "\"},true,\"\",false)")))
            if val.startswith("/"):
                log.info("(calibration)" + key + "初始化:" + str(self.omc_obj.loadFile(val)))

            else:
                log.info("(calibration)" + key + "初始化:" + str(self.omc_obj.loadFile("/home/simtek/code/" + val)))
            # 获取注释中的包名和版本号
        # name = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "name")
        # version = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "version")
        # log.info("(calibration)" + name + "初始化:" +
        #          str(self.omc_obj.sendExpression("loadModel(" + name + ", {\"" + version + "\"},true,\"\",false)")))

    def run(self):
        # omc准备操作
        log.info(self.uuid)
        update_parameter_calibration_records(uuid=self.uuid, compile_status="6", compile_start_time=int(time.time()))

        self.omc_obj.sendExpression('setCommandLineOptions("-d=initialization,NLSanalyticJacobian")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")')
        self.omc_obj.sendExpression(
            'setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection")')
        self.omc_obj.sendExpression(
            "setModelicaPath(\"/usr/lib/omlibrary\")")

        # 加载模型的依赖
        self.load_dependent_library()
        self.state = "compiling"  # 编译中
        log.info("(calibration)开始编译")
        # 编译

        absolute_path = r"/home/simtek/code/" + self.request.resultFilePath + "/"
        log.info("(calibration)仿真结果地址:" + absolute_path)
        log.info("(calibration)仿真模型名：" + self.request.simulateModelName)
        buildModelRes = self.omc_obj.buildModel(className=self.request.simulateModelName,
                                                fileNamePrefix=absolute_path,
                                                simulate_parameters_data=self.request.simulationPraData
                                                )
        log.info("(calibration)编译结果:" + str(buildModelRes))
        parent_proc = psutil.Process(self.omc_obj.omc_process.pid)
        for child_proc in parent_proc.children(recursive=True):
            os.kill(child_proc.pid, 9)
        os.kill(self.omc_obj.omc_process.pid, 9)
        if isinstance(buildModelRes, list) and buildModelRes != ["", ""]:
            log.info("(calibration)编译成功")
            update_parameter_calibration_records(uuid=self.uuid, compile_status="4", compile_stop_time=int(time.time()))
        else:
            # 改数据库状态为3
            if self.state != "stopped":
                update_parameter_calibration_records(uuid=self.uuid, compile_status="5", compile_stop_time=int(time.time()))
                return
            log.info("(calibration)编译失败")
            update_parameter_calibration_records(uuid=self.uuid, compile_status="3", compile_stop_time=int(time.time()))
            self.state = "stopped"
            return
        # 编译完成，通知omc进程退出，杀死父进程
        log.info("(calibration)编译完成，杀死omc进程：" + str(self.omc_obj.omc_process.pid))
        self.state = "stopped"


class CalibrationSimulateThread(threading.Thread):
    # 仿真之前需要进行额定工况参数写入，拟合系数参数写入，
    def __init__(self, request, port):
        self.state = "init"
        self.port = port
        self.run_pid = None
        self.uuid = request.uuid
        self.request = request
        threading.Thread.__init__(self)
        self.omc_obj = OMCSessionZMQ(port=port)
        self.compile_Dependencies = {}
        self.simulate_current = 0
        self.message = ""
        with (Session() as session):
            self.record = session.query(ParameterCalibrationRecord
                                   ).filter(ParameterCalibrationRecord.id == self.uuid).first()
            if self.record is None:
                self.state = "stopped"
                log.info("记录已被删除，本次仿真结束")
                return
        self.parameters_name = []
        self.parameters_value = []
        for condition in self.record.condition_parameters:
            self.parameters_name.append(condition["name"])
            self.parameters_value.append(condition["value"])

        self.parameters_value_list = []
        for p in range(len(self.parameters_value[0])):
            d = []
            for v in self.parameters_value:
                d.append(v[p])
            self.parameters_value_list.append(d)
        self.simulate_total = len(self.parameters_value_list)

    def load_dependent_library(self):
        for key, val in self.compile_Dependencies.items():
            # 初始化加载用户模型，key是名称，val是mo文件地址
            if val.startswith("/"):
                log.info("(calibration)" + key + "初始化:" + str(self.omc_obj.loadFile(val)))
            else:
                log.info("(calibration)" + key + "初始化:" + str(self.omc_obj.loadFile("/home/simtek/code/" + val)))
            # 获取注释中的包名和版本号

    def __del__(self):
        log.info("对象被删除，执行删除后操作")
        if self.record.simulate_status == "4":
            return
        status = "3"
        if self.state != "stopped":
            status = "5"
        for s in range(self.simulate_current, self.simulate_total):
            self.record.percentage[str(s)] = status
            update_parameter_calibration_records(uuid=self.uuid, percentage=self.record.percentage)
        update_parameter_calibration_records(uuid=self.uuid, simulate_status=status, compile_stop_time=int(time.time()),
                                             simulate_result_str=self.message)

    def run(self):
        # omc准备操作

        self.compile_Dependencies = self.record.compile_Dependencies
        update_parameter_calibration_records(uuid=self.uuid, simulate_status="6", compile_start_time=int(time.time()))

        self.omc_obj.sendExpression('setCommandLineOptions("-d=initialization,NLSanalyticJacobian")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")')
        self.omc_obj.sendExpression('setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")')
        self.omc_obj.sendExpression(
            'setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection")')
        self.omc_obj.sendExpression(
            "setModelicaPath(\"/usr/lib/omlibrary\")")

        # 加载模型的依赖
        self.load_dependent_library()
        self.state = "compiling"  # 编译中
        log.info("(calibration)开始编译")
        # 编译

        absolute_path = r"/home/simtek/code/" + self.record.compile_path+"/"
        # log.info("(calibration)仿真结果地址:" + absolute_path)
        # log.info("(calibration)仿真模型名：" + app_pages_record.model_name)
        buildModelRes = self.omc_obj.buildModel(className=self.record.model_name,
                                                fileNamePrefix=absolute_path,
                                                )
        # log.info("(calibration)编译结果:" + str(buildModelRes))
        parent_proc = psutil.Process(self.omc_obj.omc_process.pid)
        for child_proc in parent_proc.children(recursive=True):
            os.kill(child_proc.pid, 9)
        os.kill(self.omc_obj.omc_process.pid, 9)
        if isinstance(buildModelRes, list) and buildModelRes != ["", ""]:
            log.info("(calibration)编译成功")
        else:
            # 改数据库状态为3
            log.info("(calibration)编译失败")
            # update_parameter_calibration_records(uuid=self.uuid, simulate_status="3", compile_stop_time=int(time.time()))
            self.state = "stopped"
            self.__del__()
            return
        # 编译完成，通知omc进程退出，杀死父进程
        log.info("(calibration)编译完成，杀死omc进程：" + str(self.omc_obj.omc_process.pid))

        update_parameter_calibration_records(uuid=self.uuid, simulate_status="2")
        self.state = "running"
        log.info("仿真计算进行中。。。")
        for i in range(0, self.simulate_total):
            self.simulate_current = i
            percentage = self.record.percentage
            # if percentage is not None and str(i) in percentage and percentage[str(i)] == 4:
            #     log.info("本轮仿真已跳过：" + str(i))
            #     continue
            percentage[i] = "2"
            update_parameter_calibration_records(uuid=self.uuid, percentage=percentage)

            # 修改xml文件
            var = {}
            for n in range(len(self.parameters_name)):
                var[self.parameters_name[n]] = self.parameters_value_list[i][n]
            if calibration_write_xml(absolute_path, var):
                # 解析文件失败
                self.message = "仿真失败，模型由于未知原因损坏，请重新编译"
                break
            r = absolute_path+str(i)
            if not os.path.exists(r):
                os.mkdir(r)
            stepSize = '%.1g' % ((float(self.record.stop_time) - float(self.record.start_time)) / 500)
            override = ("-override=" + "startTime=" + self.record.start_time + ",stopTime=" +
                        self.record.stop_time + ",stepSize="+stepSize + ",tolerance=1e-05")

            # 运行可执行文件result
            cmd = [absolute_path + "result", override, "-r=" + r + "/result_res.mat"]
            process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            self.run_pid = process.pid
            # 获取命令行输出结果
            output, error = process.communicate()
            if error:
                log.info("(calibration)多轮仿真出错" + str(error))
                self.message = str(error)
                break

            run_result_str = output.decode('utf-8')
            self.message = str(run_result_str)
            if "successfully" in run_result_str:
                log.info("(calibration)successfully")
                # 从mat中读取数据
                d = DyMat.DyMatFile(r + "/" + "result_res.mat")
                result_dict = {"time": list(d.abscissa("2", True))}
                if self.record.result_parameters in [None, {}]:
                    log.info("结果参数为空，无法提取仿真结果，本次仿真终止")
                    break
                for j in self.record.result_parameters:
                    result_name = j["result_name"]
                    d_data = list(d.data(result_name))
                    if len(d_data) == 2 and d_data[0] == d_data[1]:
                        d_data = [d_data[0] for i in range(len(result_dict["time"]))]
                    result_dict[result_name] = d_data

                if self.record.simulate_result is None:
                    self.record.simulate_result = {i: result_dict}
                else:
                    self.record.simulate_result[i] = result_dict
                percentage[i] = "4"
                update_parameter_calibration_records(uuid=self.uuid, percentage=percentage,simulate_result=self.record.simulate_result)
            else:
                self.state = "stopped"
                self.__del__()
        update_parameter_calibration_records(uuid=self.uuid, simulate_status="4", simulate_result_str=self.message)
        self.state = "stopped"
