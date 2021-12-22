# -- coding: utf-8 --
from library.mat import DyMatFile
from config.omc import omc, OmcFactory
from app.model.Simulate.SimulateRecord import SimulateRecord
from app.model.Simulate.SimulateResult import SimulateResult
from config.DB_config import DBSession
from datetime import datetime
import socket
from library.file_operation import FileOperation
from app.service.load_model_file import LoadModelFile
import json, requests,time
from app.service.get_model_code import GetModelCode
import os
session = DBSession()

# 获取本机ip, 用于访问docker服务
hostname = socket.gethostname()
ip = socket.gethostbyname(hostname)


def SimulateDataHandle(SRecord: object, result_file_path, username, model_name, simulate_result_str):
    SRecord.simulate_result_str = simulate_result_str
    SRecord.simulate_model_result_path = result_file_path + "result_res.mat"  # omc会为结果文件添加"_res.mat"后缀
    try:
        mat_file_data = DyMatFile(result_file_path + "result_res.mat")
    except Exception as e:
        print("打开结果文件失败", e)
        SRecord.simulate_status = "仿真失败"
        session.flush()
        return
    SRecord.simulate_nametree = mat_file_data.nameTree()
    for k, v in mat_file_data.vars.items():
        model_variable_data_abscissa = mat_file_data.abscissa(k, True).tolist()
        data_len = len(v[1])
        step_size_f = data_len / 500
        if step_size_f < 1:
            step_size = 1
        else:
            step_size = int(step_size_f)
        SResult = SimulateResult(
                username=username,
                simulate_model_name=model_name,
                simulate_record_id=SRecord.id,
                model_variable_name=k,
                variable_description=v[0][:128],
                model_variable_data=v[1][::step_size],
                model_variable_data_abscissa=model_variable_data_abscissa[::step_size]
        )
        session.add(SResult)
    SRecord.simulate_status = "仿真已结束"
    session.flush()  # 提交数据

def JModelicaSimulate(SRecord_id, username: str, model_name: str, mo_path: str, simulate_parameters_data = None):
    if not mo_path:
        mo_path = "/omlibrary/Modelica 3.2.3"
    SRecord = session.query(SimulateRecord).filter_by(id=SRecord_id).first()
    result_file_path = "public/UserFiles/ModelResult" + '/' + username + '/' + \
                       model_name.split('.')[
                           -1] + '/' + str(
            datetime.now().strftime('%Y%m%d%H%M%S%f')) + '/'
    msg = {
        "start_time": simulate_parameters_data["startTime"],
        "final_time": simulate_parameters_data["stopTime"],
        "mo_path": mo_path,
        "result_name": "result_res.mat",
        "modelname": model_name,
        "ncp": simulate_parameters_data["numberOfIntervals"],  # 结果间隔
        "result_file_path": result_file_path ,  # 结果文件名字
        "rtol": simulate_parameters_data["tolerance"],  # 相对公差
    }
    file_operation = FileOperation()
    file_operation.make_dir(result_file_path)
    client = socket.socket()
    client.connect((ip, 56789))
    client.send(json.dumps(msg).encode())
    data = client.recv(1024).decode()
    client.close()
    SRecord.simulate_end_time = datetime.now()
    if data == "ok":
        SimulateDataHandle(SRecord, result_file_path, username, model_name, "ok")
    else:
        SRecord.simulate_status = "仿真失败"
    session.flush()
    session.close()

def OpenModelicaSimulate(SRecord_id, username: str, model_name: str, file_path: str = None, simulate_parameters_data = None):
    omc_once = OmcFactory()
    SRecord = session.query(SimulateRecord).filter_by(id=SRecord_id).first()
    result_file_path = "public/UserFiles/ModelResult" + '/' + username + '/' + \
                       model_name.split('.')[
                           -1] + '/' + str(
            datetime.now().strftime('%Y%m%d%H%M%S%f')) + '/'
    if file_path:
        package_name = model_name.split('.')[0]
        LoadModelFile(package_name, file_path)
    file_operation = FileOperation()
    file_operation.make_dir(result_file_path)
    simulate_result_str = omc_once.simulate(className=model_name, fileNamePrefix=result_file_path, simulate_parameters_data=simulate_parameters_data)
    SRecord.simulate_end_time = datetime.now()
    err = OmcFactory().getErrorString()
    if err == '':
        SimulateDataHandle(SRecord, result_file_path, username, model_name, simulate_result_str)
    else:
        SRecord.simulate_status = "仿真失败"
    session.flush()
    session.close()

def DymolaSimulate(SRecord_id, username, model_name, file_path=None, simulate_parameters_data=None):
    package_name = model_name.split('.')[0]
    url = package_name + "/" + model_name.replace(".", "-") + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f')) + ""
    url_dict = {"url": username + "/" + url}
    file_name = package_name + ".mo"
    SRecord = session.query(SimulateRecord).filter_by(id=SRecord_id).first()
    data = {"code": 200}
    model_str = GetModelCode(package_name, file_path, package_name)
    if file_path:
        files = {
            "file": (file_name, model_str),
            }
        r_upload_file = requests.post("http://121.37.183.103:8060/file/upload", data=url_dict, files=files)
        data = r_upload_file.json()

    if data["code"] == 200 or file_path is None:
        fileName = ""
        if file_path:
            fileName = url + "/" + file_name
        json_data_dict = {
            "id": 0,
            "fileName": fileName,
            "modelName": model_name,
            "userName": username,
            "startTime": simulate_parameters_data["startTime"],
            "stopTime": simulate_parameters_data["stopTime"],
            "numberOfIntervals": simulate_parameters_data["numberOfIntervals"],
            "outputInterval": 0.0,
            "method": "Dassl",
            "tolerance": simulate_parameters_data["tolerance"],
            "fixedStepSize": 0.0,
            "resultFile": "dsres",
            "initialNames": "",
            "initialValues": "",
            "finalNames": "",
            }
        r_simulate = requests.post("http://121.37.183.103:8060/dymola/simulate", json=json_data_dict)
        r_simulate_data = r_simulate.json()
        SRecord.simulate_end_time = datetime.now()
        if r_simulate_data.get("code", None) == 200:
            if file_path:
                var_fileName = username + "/" + url + "/" + package_name + "/dsres.mat"
            else:
                var_fileName = username + "/" + model_name.split('.')[-1] + "/" + r_simulate_data.get("msg")
            file_url = "http://121.37.183.103:8060/file/download/?fileName=" + var_fileName
            download_result_file = requests.get(file_url)
            file_data = download_result_file.content
            result_file_path = "public/UserFiles/ModelResult" + '/' + url + '/'
            FileOperation.write_file(result_file_path, "result_res.mat", file_data)
            SimulateDataHandle(SRecord, result_file_path, username, model_name, simulate_result_str="DM")
        else:
            SRecord.simulate_status = "仿真失败"
    else:
        SRecord.simulate_status = "仿真失败"
    session.flush()
    session.close()


def Simulate(SRecord_id, username: str, model_name: str, s_type="OM", file_path: str = None, simulate_parameters_data = None):
    package_name = model_name.split(".")[0]
    if file_path:
        model_str = GetModelCode(package_name, file_path, package_name)
        FileOperation().write_file("/".join(file_path.split("/")[:-1]), package_name + ".mo", model_str)

    if s_type == "OM":
        OpenModelicaSimulate(SRecord_id, username, model_name, file_path, simulate_parameters_data)
    elif s_type == "JM":
        JModelicaSimulate(SRecord_id, username, model_name, file_path, simulate_parameters_data)
    elif s_type == "DM":
        DymolaSimulate(SRecord_id, username, model_name, file_path, simulate_parameters_data)
    else:
        return "暂不支持此仿真类型"
