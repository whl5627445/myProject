import threading
import time
import zipfile
from datetime import datetime
from libs.function.defs import update_simulate_records, new_another_name, zip_folders
import requests
from config.redis_config import R
import json
import os
import configparser
from libs.function.grpc_log import log

config = configparser.ConfigParser()
config.read('./config/grpc_config.ini')
DymolaSimulationConnect = config['dymola']['DymolaSimulationConnect']
adsPath = "/home/simtek/code/"


# adsPath = "/home/xuqingda/GolandProjects/YssimGoService/"

class DmSimulation(threading.Thread):
    def __init__(self, request):
        threading.Thread.__init__(self)
        self.state = "init"
        self.uuid = request.uuid
        self.request = request
        self.processStartTime = None
        self.percentage = 0

    def update_percentage(self):
        while self.state != "stopped":
            if 30 <= self.percentage <= 98:
                # 仿真进度到达30之后,每一秒增加增加1,增加到99停止
                self.percentage += 1
                update_simulate_records(uuid=self.request.uuid, percentage=self.percentage)
                log.info("更新数据库进度为:{}".format(self.percentage))
            time.sleep(1)

    def send_request(self):
        log.info("(Dymola)发送请求")
        folders = []  # 所有要加载的用户模型的包文件地址
        uploadResult = False
        # uploadFilePath = ""
        dymola_libraries = []
        now = datetime.now()
        timestamp = now.strftime("%Y%m%d%H%M%S")
        # 新建zip文件地址
        uploadFileName = adsPath + "static/tmp/" + timestamp + "/" + self.request.packageName + ".zip"
        del_upload_fileName = adsPath + "static/tmp/" + timestamp
        # dymola服务器上新建的路径
        params_url = self.request.userName + "/" + self.request.packageName + "/" + timestamp

        for key, val in self.request.envModelData.items():
            # 初始化加载用户模型，key是名称，val是mo文件地址
            if "/" in val:
                # 将路径分割  "static/UserFiles/UploadFile/songyi/20230427165917/Applications1/Applications1/package.mo"
                front_path = "/".join(
                    val.split("/")[:6])  # static/UserFiles/UploadFile/songyi/20230427165917/Applications1
                behind_path = "/".join(val.split("/")[5:])  # Applications1/Applications1/package.mo
                # 插入到zip文件列表
                folders.append(adsPath + front_path)
                dymola_libraries.append({
                    "libraryName": "",
                    "libraryVersion": "",
                    "userFile": params_url + "/" + behind_path
                })
            else:
                # 初始化加载系统模型，key是名称，val是版本号
                dymola_libraries.append({
                    "libraryName": key,
                    "libraryVersion": val,
                    "userFile": ""
                })
        # Modelica必须放在第一个加载
        for d in dymola_libraries:
            if d['libraryName'] == 'Modelica':
                dymola_libraries.remove(d)
                dymola_libraries.insert(0, d)
                break
        log.info("(Dymola)压缩文件:" + str(folders))
        log.info("(Dymola)需要加载的依赖:" + str(dymola_libraries))
        log.info("(Dymola)上传的压缩文件:" + uploadFileName)
        log.info("(Dymola)服务器上新建的路径:" + params_url)

        if self.request.packageFilePath != "":
            # 不是系统模型,上传文件
            zip_folders(folders, uploadFileName)
            url = DymolaSimulationConnect + "/file/upload"
            params = {"url": params_url}
            timeout = 600
            files = {
                "file": (self.request.packageName + ".zip", open(uploadFileName, "rb"))
            }
            try:
                log.info("(Dymola)开始上传文件")
                log.info("(Dymola)url:" + url)
                response = requests.post(url, data=params, files=files, timeout=timeout)
                uploadRes = response.json()
                log.info("(Dymola)上传文件:" + str(uploadRes))
                if uploadRes["code"] == 200:
                    uploadResult = True
                    log.info('(Dymola)上传文件成功')
                    # uploadFilePath = uploadRes["data"]
                else:
                    return False, None, 0
            except Exception as e:
                log.info("(Dymola)" + str(e))
                return False, None, 0
            # 上传完删除zip文件
            if os.path.exists(del_upload_fileName):
                # os.remove(del_upload_fileName)
                os.system('rm -rf ' + del_upload_fileName)
                log.info("(Dymola)上传完成后，zip文件夹已成功删除！")
            else:
                log.info("(Dymola)ip文件夹不存在。")
        else:
            # 系统模型的仿真要去掉用户模型
            dymola_libraries = [element for element in dymola_libraries if element['userFile'] == '']
            log.info("(Dymola)系统模型只加载系统库:" + str(dymola_libraries))
        if uploadResult or self.request.packageFilePath == "":
            fileName = ""
            if self.request.packageFilePath != "":
                fileName = params_url + "/" + "/".join(self.request.packageFilePath.split("/")[5:])

            compileReqData = {
                "userName": self.request.userName,
                "fileName": fileName,
                "modelName": self.request.simulateModelName,
                "taskId": self.request.uuid,
                "dymolaLibraries": dymola_libraries
            }
            log.info('(Dymola)开始编译')
            log.info("(Dymola)编译请求体：" + str(compileReqData))
            json_data = {"message": self.request.simulateModelName + " 模型开始编译"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            update_simulate_records(uuid=self.request.uuid,
                                    simulate_status="6",
                                    simulate_start_time=time.time(),
                                    percentage=20)
            self.percentage = 20
            compileRes = requests.request("post", DymolaSimulationConnect + "/dymola/translate",
                                          json=compileReqData,
                                          timeout=600)
            if compileRes.status_code != 200:
                log.info("(Dymola)服务编译错误: " + str(compileRes.reason))
                return False, None, compileRes.status_code
            compileResData = compileRes.json()
            log.info("(Dymola)服务编译结果：" + str(compileResData))
            if compileResData["code"] == 200:
                update_simulate_records(uuid=self.request.uuid,
                                        simulate_status="2",
                                        simulate_start="1",
                                        percentage=30)
                self.percentage = 30
                json_data = {"message": self.request.simulateModelName + " 编译成功，开始仿真"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                simulateReqData = {
                    "startTime": self.request.simulationPraData["startTime"],
                    "stopTime": self.request.simulationPraData["stopTime"],
                    "numberOfIntervals": self.request.simulationPraData["numberOfIntervals"],
                    "outputInterval": 0.0,
                    "method": self.request.simulationPraData["method"],
                    "tolerance": self.request.simulationPraData["tolerance"],
                    "fixedStepSize": 0.0,
                    "resultFile": "dsres",
                    "fileName": fileName,
                    "modelName": self.request.simulateModelName,
                    "userName": self.request.userName,
                    "taskId": self.request.uuid,
                    "dymolaLibraries": dymola_libraries
                }

                log.info("(Dymola)仿真请求体：" + str(simulateReqData))
                simulateRes = requests.request('post',
                                               DymolaSimulationConnect + "/dymola/simulate",
                                               json=simulateReqData)
                simulateResData = simulateRes.json()

                log.info("(Dymola)仿真结果：" + str(simulateResData))
                absolute_path = adsPath + self.request.resultFilePath
                log.info("(Dymola)结果地址：" + absolute_path)
                if simulateResData.get("code") == 200:
                    resFileUrl = DymolaSimulationConnect + "/file/download?fileName=" + simulateResData["msg"]
                    fmuFileUrl = DymolaSimulationConnect + "/file/download?fileName=" + compileResData["msg"]

                    downloadResFileUrl = requests.get(resFileUrl)
                    # 创建文件夹
                    os.makedirs(absolute_path, exist_ok=True)
                    with open(absolute_path + "result_res.mat", "wb") as f:
                        f.write(downloadResFileUrl.content)

                    downloadFmuFileUrl = requests.get(fmuFileUrl)
                    with open(absolute_path + "dymola_model.fmu.zip", "wb") as f:
                        f.write(downloadFmuFileUrl.content)

                    with zipfile.ZipFile(absolute_path + "dymola_model.fmu.zip", 'r') as zip_ref:
                        zip_ref.extract('modelDescription.xml', absolute_path)

                    os.rename(absolute_path + "modelDescription.xml", absolute_path + "result_init.xml")

                    return True, None, 0
                else:
                    return False, None, simulateResData.get("code")
            else:
                return False, None, compileResData["code"]

    def run(self):
        self.state = "running"
        log.info("(Dymola)开启dymola仿真")
        update_simulate_records(uuid=self.request.uuid,
                                simulate_start_time=time.time(),
                                simulate_start=True,
                                percentage=10)
        self.percentage = 10
        res, err, code = self.send_request()
        # 开启假的进度增加线程
        fake_update_percentage = threading.Thread(target=self.update_percentage)
        fake_update_percentage.start()

        log.info("(Dymola)返回" + str(res) + str(err) + str(code))
        if res:
            update_simulate_records(uuid=self.request.uuid,
                                    simulate_status="4",
                                    simulate_result_str="DM",
                                    simulate_start="0",
                                    simulate_end_time=str(time.time()),
                                    another_name=new_another_name(self.request.userName, self.request.simulateModelName,
                                                                  self.request.simulatePackageId,
                                                                  self.request.userSpaceId),
                                    percentage=100
                                    )
            self.percentage = 100
            json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        elif code == 300:
            update_simulate_records(uuid=self.request.uuid,
                                    simulate_status="3",
                                    simulate_result_str="DM",
                                    simulate_start="0",
                                    simulate_end_time=str(time.time())
                                    )
            json_data = {"message": self.request.simulateModelName + " 结束任务"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        else:
            update_simulate_records(uuid=self.request.uuid,
                                    simulate_status="3",
                                    simulate_result_str="DM",
                                    simulate_start="0",
                                    simulate_end_time=str(time.time())
                                    )
            json_data = {"message": self.request.simulateModelName + " 仿真失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        log.info("(Dymola)仿真线程执行完毕")
        self.state = "stopped"
