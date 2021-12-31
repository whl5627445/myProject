# -- coding: utf-8 --
import requests
from datetime import datetime
import json

def DymolaFmuExport(token: str, username: str, package_name: str, model_name: str, fmu_name: str, storeResult, includeSource, fmiVersion, includeImage, fmiType, file_name: str = "", model_str: str = None, file_path: str = None):
    modelName = model_name.replace(".", "-")
    data = {
        "username": username,
        "storeResult": storeResult,
        "includeSource": includeSource,
        "fmiVersion": str(fmiVersion),
        "includeImage": includeImage,
        "fmiType": fmiType,
        "modelName": fmu_name,
        # "modelName": modelName,
        "fileName": "",
        "modelToOpen": model_name,
        "token": "Bearer " + token,
        }
    res = {"msg": "", "status": 0, "err": "", "data": []}
    if file_path:
        files = {
            "file": (file_name, model_str),
            }
        url = package_name + "/" + model_name.replace(".", "-") + "/" + str(
            datetime.now().strftime('%Y%m%d%H%M%S%f')) + ""
        file_data = {"url":username + "/" + url}
        res_upload_file = requests.post("http://121.37.183.103:8060/file/upload", data=file_data, files=files)
        upload_file_data = res_upload_file.json()
        if upload_file_data["code"] == 200:
            data["fileName"] = url + "/" + file_name
            res_export_fmu = requests.post("http://121.37.183.103:8060/dymola/translateModelFMU", json=data)
            export_fmu_data = res_export_fmu.json()
            if export_fmu_data["code"] == 200:
                res["msg"] = "导出成功"
                return res
            else:
                res["err"] = export_fmu_data["msg"]
                res["status"] = 1
        else:
            res["err"] = "导出失败"
            res["status"] = 1
    else:
        res_export_fmu = requests.post("http://121.37.183.103:8060/dymola/translateModelFMU", json=data)
        export_fmu_data = res_export_fmu.json()
        if export_fmu_data["code"] == 200:
            res["msg"] = "导出成功"
            return res
        else:
            res["err"] = export_fmu_data["msg"]
            res["status"] = 1
    return res
