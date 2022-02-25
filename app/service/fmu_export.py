# -- coding: utf-8 --
import requests
from datetime import datetime
import json

def DymolaFmuExport(fmu_par, token, username: str, file_name: str = "", model_str: str = None, file_path: str = None):
    modelName = fmu_par.model_name.replace(".", "-")
    data = {
        "username": username,
        "storeResult": fmu_par.storeResult,
        "includeSource": fmu_par.includeSource,
        "fmiVersion": str(fmu_par.fmiVersion),
        "includeImage": fmu_par.includeImage,
        "fmiType": fmu_par.fmiType,
        "modelName": fmu_par.fmu_name,
        # "modelName": modelName,
        "fileName": "",
        "modelToOpen": fmu_par.model_name,
        "token": token,
        }
    res = {"result": True}
    url = fmu_par.package_name + "/" + fmu_par.model_name.replace(".", "-") + "/" + str(
            datetime.now().strftime('%Y%m%d%H%M%S%f')) + ""
    if file_path:
        files = {
            "file": (file_name + ".mo", model_str),
            }
        file_data = {"url":username + "/" + url}
        res_upload_file = requests.post("http://121.37.183.103:8060/file/upload", data=file_data, files=files)
        upload_file_data = res_upload_file.json()
        if upload_file_data["code"] == 200:
            data["fileName"] = url + "/" + file_name + ".mo"
            res_export_fmu = requests.post("http://121.37.183.103:8060/dymola/translateModelFMU", json=data)
            export_fmu_data = res_export_fmu.json()
            if not export_fmu_data["code"] == 200:
                res["result"] = False
        else:
            res["result"] = False
    else:
        res_export_fmu = requests.post("http://121.37.183.103:8060/dymola/translateModelFMU", json=data)
        export_fmu_data = res_export_fmu.json()
        if export_fmu_data["code"] == 200:
            if not fmu_par.download_local:
                res["result"] = False
        else:
            res["result"] = False
    return res
