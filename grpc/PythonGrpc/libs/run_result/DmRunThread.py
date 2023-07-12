import threading
import time
from datetime import datetime
from libs.function.defs import zip_folders
import requests
import pandas as pd
import os
import configparser
from libs.function.defs import update_app_pages_records, dymola_convert_list, \
    page_preview_component_freeze, result_step
from libs.function.grpc_log import log
import shutil

config = configparser.ConfigParser()
config.read('./config/grpc_config.ini')
DymolaSimulationConnect = config['dymola']['DymolaSimulationConnect']
adsPath = "/home/simtek/code/"


class DmRunThread(threading.Thread):
    def __init__(self, request):
        threading.Thread.__init__(self)
        self.state = "init"
        self.uuid = request.pageId
        self.request = request
        self.inputValData = request.inputValData
        self.input_data = dymola_convert_list(self.inputValData, self.request.pageId)
        update_app_pages_records(self.request.pageId, mul_sim_state=1)

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
                    return False, '上传文件失败', 0
            except Exception as e:
                log.info("(Dymola)" + str(e))
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
            log.info("(Dymola)系统模型只加载系统库:" + str(dymola_libraries))
        if uploadResult or self.request.packageFilePath == "":
            fileName = ""
            if self.request.packageFilePath != "":
                fileName = params_url + "/" + "/".join(self.request.packageFilePath.split("/")[5:])
            # array_initial_values = list(self.request.inputValData.values())
            # input_data_length = len(array_initial_values[0].inputObjList)

            # 多轮仿真
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
                "arrayInitialValues": self.input_data,
                "finalNames": ["time"] + list(self.request.outputValNames),

            }
            log.info("(Dymola)仿真请求体：" + str(simulateReqData))
            simulateRes = requests.request('post',
                                           DymolaSimulationConnect + "/dymola/simulateMulti",
                                           json=simulateReqData)
            simulateResData = simulateRes.json()

            log.info("(Dymola)dymola仿真结果：" + str(simulateResData["code"]))
            mul_output_path = adsPath + self.request.mulResultPath + "preview/"
            if self.request.mulResultPath is None:
                return False, "mulResultPath为空", ''
            log.info("(Dymola)结果路径：" + mul_output_path)
            update_app_pages_records(self.request.pageId,
                                     mul_sim_message_read=False,
                                     mul_sim_err=simulateResData.get("log"))
            if simulateResData.get("code") == 200:
                csv_data = simulateResData["data"]
                log.info(type(csv_data))
                if os.path.exists(mul_output_path):
                    shutil.rmtree(mul_output_path)
                # 创建新的文件夹
                os.mkdir(mul_output_path)
                for i in range(len(csv_data["time"])):
                    temp = {}
                    for key, value in csv_data.items():
                        temp[key] = result_step(value[i])
                    df = pd.DataFrame(pd.DataFrame.from_dict(temp, orient='index').values.T,
                                      columns=list(temp.keys()))
                    csv_file_name = ""
                    for s in self.input_data[i]:
                        s = round(s, 6)
                        csv_file_name = csv_file_name + "_" + str(s)
                    df.to_csv(mul_output_path + '{}.csv'.format(csv_file_name), index=False)

                return True, None, simulateResData["code"]
            else:
                return False, "多轮仿真失败", simulateResData["code"]

    def run(self):
        self.state = "running"
        log.info("(Dymola)开启dymola仿真")
        update_app_pages_records(self.request.pageId, mul_sim_state=2)
        res, err, code = self.send_request()
        log.info("(Dymola)send_request返回" + str(res) + str(err) + str(code))
        if res:
            update_app_pages_records(self.request.pageId,
                                     mul_sim_state=4,
                                     is_preview=1,  # 是否可以预览
                                     is_mul_simulate=1,  # 是否多轮仿真过
                                     naming_order=list(self.inputValData.keys()))
            # update_app_spaces_records(self.request.pageId)
            page_preview_component_freeze(self.request.pageId)
        else:
            update_app_pages_records(self.request.pageId, mul_sim_state=3)
        update_app_pages_records(self.request.pageId,
                                 mul_sim_time=time.time(),
                                 mul_sim_message_read=False,)
        log.info("(Dymola)仿真线程执行完毕")
        # delete_item_from_json(self.request.uuid)
        self.state = "stopped"
