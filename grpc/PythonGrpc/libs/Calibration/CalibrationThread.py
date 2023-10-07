import shutil
import threading
import time
import psutil
import json
import os
import subprocess
import DyMat
import pandas as pd

# from config.redis_config import R
from libs.OMPython.OMCSessionZMQ import OMCSessionZMQ
from libs.function.defs import (
    update_compile_records,
    sendMessage,
    del_data_sources_records,
)
from libs.function.grpc_log import log
from libs.function.defs import (
    update_app_pages_records,
    omc_convert_dict_to_list,
    update_app_spaces_records,
    page_preview_component_freeze,
    result_step,
)
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
                log.info(
                    "(calibration)" + key + "初始化:" + str(self.omc_obj.loadFile(val))
                )

            else:
                log.info(
                    "(calibration)"
                    + key
                    + "初始化:"
                    + str(self.omc_obj.loadFile("/home/simtek/code/" + val))
                )
            # 获取注释中的包名和版本号
        # name = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "name")
        # version = self.omc_obj.getAnnotationModifierValue(self.request.simulateModelName, "from", "version")
        # log.info("(calibration)" + name + "初始化:" +
        #          str(self.omc_obj.sendExpression("loadModel(" + name + ", {\"" + version + "\"},true,\"\",false)")))

    def run(self):
        # omc准备操作
        log.info(self.uuid)
        update_parameter_calibration_records(
            uuid=self.uuid, compile_status="6", compile_start_time=int(time.time())
        )

        self.omc_obj.sendExpression(
            'setCommandLineOptions("-d=initialization,NLSanalyticJacobian")'
        )
        self.omc_obj.sendExpression(
            'setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")'
        )
        self.omc_obj.sendExpression(
            'setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")'
        )
        self.omc_obj.sendExpression(
            'setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection")'
        )
        self.omc_obj.sendExpression('setModelicaPath("/usr/lib/omlibrary")')

        # 加载模型的依赖
        self.load_dependent_library()
        self.state = "compiling"  # 编译中
        log.info("(calibration)开始编译")
        # 编译

        absolute_path = "/home/simtek/code/" + self.request.resultFilePath
        log.info("(calibration)仿真结果地址:" + absolute_path)
        log.info("(calibration)仿真模型名：" + self.request.simulateModelName)
        buildModelRes = self.omc_obj.buildModel(
            className=self.request.simulateModelName,
            fileNamePrefix=absolute_path,
            simulate_parameters_data=self.request.simulationPraData,
        )
        log.info("(calibration)编译结果:" + str(buildModelRes))
        parent_proc = psutil.Process(self.omc_obj.omc_process.pid)
        for child_proc in parent_proc.children(recursive=True):
            os.kill(child_proc.pid, 9)
        os.kill(self.omc_obj.omc_process.pid, 9)
        if isinstance(buildModelRes, list) and buildModelRes != ["", ""]:
            log.info("(calibration)编译成功")
            update_parameter_calibration_records(
                uuid=self.uuid, compile_status="4", compile_stop_time=int(time.time())
            )
        else:
            # 改数据库状态为3
            if self.state != "stopped":
                update_parameter_calibration_records(
                    uuid=self.uuid,
                    compile_status="5",
                    compile_stop_time=int(time.time()),
                )
                return
            log.info("(calibration)编译失败")
            update_parameter_calibration_records(
                uuid=self.uuid, compile_status="3", compile_stop_time=int(time.time())
            )
            shutil.rmtree(absolute_path)
            self.state = "stopped"
            return
        # 编译完成，通知omc进程退出，杀死父进程
        log.info("(calibration)编译完成，杀死omc进程：" + str(self.omc_obj.omc_process.pid))
        self.state = "stopped"


class CalibrationSimulateThread(threading.Thread):
    # 仿真之前需要进行额定工况参数写入，拟合系数参数写入，
    def __init__(self, request, port, calibration_simulate_task_mark_dict):
        self.state = "init"
        self.port = port
        self.run_pid = None
        self.uuid = request.uuid
        self.request = request
        self.omc_obj = OMCSessionZMQ(port=port)
        self.compile_Dependencies = {}
        self.simulate_current = 0
        self.message = ""
        self.destroy = False
        self.calibration_simulate_task_mark_dict = calibration_simulate_task_mark_dict
        self.percentage = {}
        with Session() as session:
            self.record = (
                session.query(ParameterCalibrationRecord)
                .filter(ParameterCalibrationRecord.id == self.uuid)
                .first()
            )
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
        threading.Thread.__init__(self)

    def load_dependent_library(self):
        for key, val in self.compile_Dependencies.items():
            # 初始化加载用户模型，key是名称，val是mo文件地址
            if val.startswith("/"):
                log.info(
                    "(calibration)" + key + "初始化:" + str(self.omc_obj.loadFile(val))
                )
            else:
                log.info(
                    "(calibration)"
                    + key
                    + "初始化:"
                    + str(self.omc_obj.loadFile("/home/simtek/code/" + val))
                )
            # 获取注释中的包名和版本号

    def __del__(self):
        log.info("对象被删除，执行删除后操作")

        if not self.destroy:
            parent_proc = psutil.Process(self.omc_obj.omc_process.pid)
            for child_proc in parent_proc.children(recursive=True):
                log.info("(calibration)关闭子进程：" + str(child_proc.pid))
                os.kill(child_proc.pid, 9)
            os.kill(self.omc_obj.omc_process.pid, 9)
            log.info("(calibration)关闭omc进程成功")
            try:
                if self.run_pid:
                    os.kill(self.run_pid, 9)
            except OSError:
                log.info("(calibration)运行进程不存在，跳过")
            status = "4"
            if self.state == "delete":
                status = "5"
            elif self.state == "fail":
                status = "3"
                self.state = "stop"
            for s in range(self.simulate_current, self.simulate_total):
                self.percentage[str(s)] = status
                update_parameter_calibration_records(
                    uuid=self.uuid, percentage=self.percentage
                )
            log.info("status： " + str(status))
            log.info("message： " + str(self.message))
            update_parameter_calibration_records(
                uuid=self.uuid,
                simulate_status=status,
                compile_stop_time=int(time.time()),
                simulate_result_str=self.message,
            )
            self.destroy = True

    def run(self):
        # omc准备操作

        self.compile_Dependencies = self.record.compile_Dependencies
        self.omc_obj.sendExpression(
            'setCommandLineOptions("-d=initialization,NLSanalyticJacobian")'
        )
        self.omc_obj.sendExpression(
            'setCommandLineOptions("+ignoreSimulationFlagsAnnotation=false")'
        )
        self.omc_obj.sendExpression(
            'setCommandLineOptions("+ignoreCommandLineOptionsAnnotation=false")'
        )
        self.omc_obj.sendExpression(
            'setCommandLineOptions("--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection")'
        )
        self.omc_obj.sendExpression('setModelicaPath("/usr/lib/omlibrary")')

        # 加载模型的依赖
        self.load_dependent_library()
        self.state = "compiling"  # 编译中
        log.info("(calibration)开始编译")
        # 编译
        update_parameter_calibration_records(
            uuid=self.uuid,
            simulate_status="6",
            compile_start_time=int(time.time()),
        )
        absolute_path = r"/home/simtek/code/" + self.record.compile_path + "/"
        # log.info("(calibration)仿真结果地址:" + absolute_path)
        # log.info("(calibration)仿真模型名：" + app_pages_record.model_name)
        buildModelRes = self.omc_obj.buildModel(
            className=self.record.model_name,
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
            self.join()

        # 编译完成，通知omc进程退出，杀死父进程
        log.info("(calibration)编译完成，杀死omc进程：" + str(self.omc_obj.omc_process.pid))

        update_parameter_calibration_records(uuid=self.uuid, simulate_status="2")

        self.state = "running"
        log.info("仿真计算进行中。。。")
        simulate_status = "4"
        for i in range(0, self.simulate_total):
            if self.state == "delete" or self.state == "stop":
                self.join()
            self.simulate_current = i
            self.percentage[str(i)] = "2"
            update_parameter_calibration_records(
                uuid=self.uuid, percentage=self.percentage
            )

            # 修改xml文件
            var = {}
            for n in range(len(self.parameters_name)):
                var[self.parameters_name[n]] = self.parameters_value_list[i][n]
            if calibration_write_xml(absolute_path, var):
                # 解析文件失败
                self.message = "仿真失败，模型由于未知原因损坏，请重新编译"
                break
            r = absolute_path + str(i)
            if not os.path.exists(r):
                os.mkdir(r)
            size = (float(self.record.stop_time) - float(self.record.start_time)) / 500
            if size + 0.5 > 1:
                stepSize = str(int(size + 0.5))
            else:
                stepSize = "%.1g" % size
            override = (
                "-override="
                + "startTime="
                + self.record.start_time
                + ",stopTime="
                + self.record.stop_time
                + ",stepSize="
                + stepSize
                + ",tolerance=1e-05"
            )

            # 运行可执行文件result
            cmd = [absolute_path + "result", override, "-r=" + r + "/result_res.mat"]
            process = subprocess.Popen(
                cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE
            )
            self.run_pid = process.pid
            # 获取命令行输出结果
            output, error = process.communicate()
            if error:
                log.info("(calibration)多轮仿真出错" + str(error))
                self.message = str(error)
                break

            run_result_str = output.decode("utf-8")
            self.message = str(run_result_str)
            if "The simulation finished successfully." in run_result_str:
                log.info("(calibration)successfully")
                # 从mat中读取数据
                d = DyMat.DyMatFile(r + "/" + "result_res.mat")
                result_dict = {"time": list(d.abscissa("2", True))}
                if self.record.result_parameters in [None, {}]:
                    log.info("结果参数为空，无法提取仿真结果，本次仿真终止")
                    self.join()

                self.percentage[i] = "4"
                for j in self.record.result_parameters:
                    result_name = j["result_name"]
                    d_data = list(d.data(result_name))
                    if len(d_data) == 2 and d_data[0] == d_data[1]:
                        d_data = [d_data[0] for i in range(len(result_dict["time"]))]
                    result_dict[result_name] = d_data

                    if result_name != "time" and not steady_state(d_data):
                        self.percentage[i] = "7"  # 表示本次仿真没有达到稳态
                if self.record.simulate_result is None:
                    self.record.simulate_result = {i: result_dict}
                else:
                    self.record.simulate_result[i] = result_dict

                update_parameter_calibration_records(
                    uuid=self.uuid,
                    percentage=self.percentage,
                    simulate_result=self.record.simulate_result,
                )
            else:
                log.info("仿真执行失败，本次仿真终止")

                self.state = "fail"
                self.__del__()
                break
            shutil.rmtree(r)
        update_parameter_calibration_records(
            uuid=self.uuid,
            simulate_status=simulate_status,
            simulate_result_str=self.message,
        )
        self.state = "stopped"
        self.__del__()


def steady_state(data):
    e = 0.1
    step = int(len(data) / 100) if int(len(data) / 100) > 0 else 1
    t0 = float("inf")
    for i in range(step, len(data), step):
        t1 = abs(data[step] - data[step - 1])
        if (t1 == 0 and t0 == 0) or (t1 / t0) > (1 - e):
            return True
        t0 = t1
    return False
