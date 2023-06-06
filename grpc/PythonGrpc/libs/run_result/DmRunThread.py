import threading
import time
import zipfile
from datetime import datetime
from libs.function.defs import zip_folders
from libs.function.run_result_json import update_item_to_json, delete_item_from_json
import requests
import pandas as pd
from config.redis_config import R
import json
import os
import configparser
from libs.function.defs import update_app_pages_records,convert_list
from libs.function.grpc_log import log

config = configparser.ConfigParser()
config.read('./config/grpc_config.ini')
DymolaSimulationConnect = config['dymola']['DymolaSimulationConnect']
adsPath = "/home/simtek/code/"


# adsPath = "/home/xuqingda/GolandProjects/YssimGoService/"

class DmRunThread(threading.Thread):
    def __init__(self, request):
        threading.Thread.__init__(self)
        self.state = "init"
        self.request = request
        update_app_pages_records(self.request.pageId, release_state=1)

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
        log.info("(Dymola)压缩文件:"+str(folders))
        log.info("(Dymola)需要加载的依赖:"+str(dymola_libraries))
        log.info("(Dymola)上传的压缩文件:"+uploadFileName)
        log.info("(Dymola)服务器上新建的路径:"+params_url)

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
                log.info("(Dymola)url:"+url)
                response = requests.post(url, data=params, files=files, timeout=timeout)
                uploadRes = response.json()
                log.info("(Dymola)上传文件:"+str(uploadRes))
                if uploadRes["code"] == 200:
                    uploadResult = True
                    log.info('(Dymola)上传文件成功')
                    # uploadFilePath = uploadRes["data"]
                else:
                    return False, '上传文件失败', 0
            except Exception as e:
                log.info("(Dymola)"+str(e))
                return False, str(e), 0
            # 上传完删除zip文件
            if os.path.exists(del_upload_fileName):
                os.system('rm -rf ' + del_upload_fileName)
                log.info("(Dymola)上传完成后，zip文件夹已成功删除！")
            else:
                log.info("(Dymola)zip文件夹不存在。")
        else:
            # 系统模型的仿真要去掉用户模型
            dymola_libraries = [element for element in dymola_libraries if element['userFile'] == '']
            log.info("(Dymola)系统模型只加载系统库:"+str(dymola_libraries))
        if uploadResult or self.request.packageFilePath == "":
            fileName = ""
            if self.request.packageFilePath != "":
                fileName = params_url + "/" + "/".join(self.request.packageFilePath.split("/")[5:])
            json_data = {"message": self.request.simulateModelName + " 开始多轮仿真"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
            array_initial_values = list(self.request.inputValData.values())
            input_data_length = len(array_initial_values[0].inputObjList)
            # 如果input_data_length的程度为1，这是单次仿真
            if input_data_length == 1:
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

                log.info("(Dymola)仿真接口请求体：" + str(simulateReqData))
                simulateRes = requests.request('post',
                                               DymolaSimulationConnect + "/dymola/simulate",
                                               json=simulateReqData)
                simulateResData = simulateRes.json()

                log.info("(Dymola)dymola仿真结果："+str(simulateResData))
                absolute_path = adsPath + self.request.resultFilePath
                log.info("(Dymola)结果路径："+absolute_path)
                if simulateResData.get("code") == 200:
                    resFileUrl = DymolaSimulationConnect + "/file/download?fileName=" + simulateResData["msg"]

                    downloadResFileUrl = requests.get(resFileUrl)
                    # 创建文件夹
                    os.makedirs(absolute_path, exist_ok=True)
                    with open(absolute_path + "result_res.mat", "wb") as f:
                        f.write(downloadResFileUrl.content)
                    return True, None, 0
                else:
                    return False, "单轮仿真失败", simulateResData.get("code")
            # 多轮仿真
            else:
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
                    "dymolaLibraries": dymola_libraries,
                    "initialNames": list(self.request.inputValData.keys()),
                    "arrayInitialValues": [[v.inputObjList[i] for v in array_initial_values] for i in
                                           range(len(array_initial_values[0].inputObjList))],
                    "finalNames": ["Time"] + list(self.request.outputValNames),

                }
                log.info("(Dymola)仿真请求体："+str(simulateReqData))
                simulateRes = requests.request('post',
                                               DymolaSimulationConnect + "/dymola/simulateMulti",
                                               json=simulateReqData)
                simulateResData = simulateRes.json()

                log.info("(Dymola)dymola仿真结果："+str(simulateResData))
                absolute_path = adsPath + self.request.resultFilePath
                log.info("(Dymola)结果路径："+absolute_path)
                if simulateResData.get("code") == 200:
                    csv_data = simulateResData["data"]
                    df = pd.DataFrame(csv_data)
                    # 将DataFrame对象保存为CSV文件
                    # df.to_csv(absolute_path + 'output.csv', index=False)
                    df.to_csv(r"/home/simtek/code/" + absolute_path + 'output.csv', index=False)
                    # 更新数据库
                    update_app_pages_records(self.request.pageId,
                                             single_simulation_result_path=absolute_path + 'output.csv')

                    return True, None, simulateResData["code"]

                else:
                    return False, "多轮仿真失败", simulateResData["code"]

    def run(self):
        log.info("(Dymola)开启dymola仿真")
        update_app_pages_records(self.request.pageId, release_state=2)
        res, err, code = self.send_request()
        log.info("(Dymola)返回"+str(res)+str(err)+str(code))
        if res:
            update_app_pages_records(self.request.pageId, release_state=4)
            json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        elif code == 300:
            update_app_pages_records(self.request.pageId, release_state=3)
            json_data = {"message": self.request.simulateModelName + " 结束任务"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        else:
            update_app_pages_records(self.request.pageId, release_state=3)
            json_data = {"message": self.request.simulateModelName + " 仿真失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        self.state = "stopped"
        log.info("(Dymola)仿真线程执行完毕")
        delete_item_from_json(self.request.uuid)
