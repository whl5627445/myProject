import threading
import time
import zipfile
from datetime import datetime
from libs.function.defs import update_compile_records, new_another_name, zip_folders
import requests
from config.redis_config import R
import json
import os
import configparser

config = configparser.ConfigParser()
config.read('./config/grpc_config.ini')
DymolaSimulationConnect = config['dymola']['DymolaSimulationConnect']
adsPath = "/home/simtek/code/"


# adsPath = "/home/xuqingda/GolandProjects/YssimGoService/"

class DmTranslateThread(threading.Thread):
    def __init__(self, request):
        threading.Thread.__init__(self)
        self.state = "init"
        self.request = request

    def send_request(self):
        print("发送请求")
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
        print("folders:", folders)
        print("dymola_libraries:", dymola_libraries)
        print("uploadFileName:", uploadFileName)
        print("params_url:", params_url)

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
                print("开始上传文件")
                print("url:", url)
                response = requests.post(url, data=params, files=files, timeout=timeout)
                uploadRes = response.json()
                print("uploadRes:", uploadRes)
                if uploadRes["code"] == 200:
                    uploadResult = True
                    print('上传文件成功')
                    # uploadFilePath = uploadRes["data"]
                else:
                    return False, None, 0
            except Exception as e:
                print(e)
                return False, None, 0
            # 上传完删除zip文件
            if os.path.exists(del_upload_fileName):
                os.system('rm -rf ' + del_upload_fileName)
                print("文件夹已成功删除！")
            else:
                print("文件夹不存在。")
        else:
            # 系统模型的仿真要去掉用户模型
            dymola_libraries = [element for element in dymola_libraries if element['userFile'] == '']
            print("系统模型只加载系统库:", dymola_libraries)
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
            print('开始编译')
            print(compileReqData)
            json_data = {"message": self.request.simulateModelName + " 模型开始编译"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))

            compileRes = requests.request("post", DymolaSimulationConnect + "/dymola/translate",
                                          json=compileReqData,
                                          timeout=600)
            if compileRes.status_code != 200:
                print("dymola服务编译错误: ", compileRes.reason)
                return False, None, compileRes.status_code
            compileResData = compileRes.json()
            print("dymola服务编译结果：", compileResData)
            if compileResData["code"] == 200:
                update_compile_records(uuid=self.request.uuid, compile_status=4, compile_stop_time=int(time.time()))
                json_data = {"message": self.request.simulateModelName + " 编译成功"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                absolute_path = adsPath + self.request.resultFilePath
                print("absolute_path", absolute_path)
                fmuFileUrl = DymolaSimulationConnect + "/file/download?fileName=" + compileResData["msg"]
                # 创建文件夹
                os.makedirs(absolute_path, exist_ok=True)
                downloadFmuFileUrl = requests.get(fmuFileUrl)
                with open(absolute_path + "dymola_model.fmu.zip", "wb") as f:
                    f.write(downloadFmuFileUrl.content)
                with zipfile.ZipFile(absolute_path + "dymola_model.fmu.zip", 'r') as zip_ref:
                    zip_ref.extract('modelDescription.xml', absolute_path)
                os.rename(absolute_path + "modelDescription.xml", absolute_path + "result_init.xml")
                return True, None, 0

            else:
                return False, None, compileResData["code"]

    def run(self):
        print("开启dymola仿真")
        update_compile_records(uuid=self.request.uuid, compile_status=2, compile_start_time=int(time.time()))
        res, err, code = self.send_request()
        print("send_request返回", res, err, code)
        if res:
            update_compile_records(uuid=self.request.uuid,
                                   compile_status=4,
                                   compile_stop_time=int(time.time())
                                   )
            json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        elif code == 300:
            json_data = {"message": self.request.simulateModelName + " 结束任务"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        else:
            update_compile_records(uuid=self.request.uuid,
                                   compile_status=3,
                                   compile_stop_time=int(time.time())
                                   )
            json_data = {"message": self.request.simulateModelName + " 仿真失败"}
            R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
        self.state = "stopped"