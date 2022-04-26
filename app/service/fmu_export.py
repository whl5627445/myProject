# -- coding: utf-8 --
import requests
from datetime import datetime
import json, logging
from library.file_operation import FileOperation


def DymolaFmuExport(fmu_par, token, username: str, file_name: str = "", model_str: str = None, file_path: str = None):
    data = {
        "username": username,
        "fmuPar": fmu_par.fmuPar,
        "modelName": fmu_par.fmu_name,
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
        if upload_file_data.get("code", None) == 200:
            data["fileName"] = url + "/" + file_name + ".mo"
            res_export_fmu = requests.post("http://121.37.183.103:8060/dymola/translateModelFMU", json=data)
            export_fmu_data = res_export_fmu.json()

        else:
            export_fmu_data = {}
            res["result"] = False
    else:
        res_export_fmu = requests.post("http://121.37.183.103:8060/dymola/translateModelFMU", json=data)
        export_fmu_data = res_export_fmu.json()

        if export_fmu_data.get("code", None) != 200:
            if not fmu_par.download_local:
                res["result"] = False

    if res.get("result", None):
        fmu_fileName = export_fmu_data.get("msg", "")
        fmu_file_url = "http://121.37.183.103:8061/" + fmu_fileName
        download_fmu_file = requests.get(fmu_file_url)
        fmu_file_data = download_fmu_file.content
        file_operation = FileOperation()
        result_file_path = "public/UserFiles/FmuExport" + '/' + username + '/' + \
                           fmu_par.fmu_name.split('.')[
                               -1] + '/' + str(datetime.now().strftime('%Y%m%d%H%M%S%f')) + '/'
        file_operation.write_file(result_file_path, fmu_par.fmu_name + ".fmu", fmu_file_data)
        res["file_path"] = result_file_path + fmu_par.fmu_name + ".fmu"
    else:
        res["result"] = False
    return res
