from library.mat import DyMatFile
from config.omc import omc
from app.model.Simulate.SimulateRecord import SimulateRecord
from app.model.Simulate.SimulateResult import SimulateResult
from config.DB_config import DBSession
from datetime import datetime
import socket
from library.file_operation import FileOperation
import json
import os
session = DBSession()

# 获取本机ip, 用于访问docker服务
hostname = socket.gethostname()
ip = socket.gethostbyname(hostname)


def SimulateDataHandle(SRecord: object, result_file_path, username, model_name, simulate_result_str):
    SRecord.simulate_result_str = simulate_result_str
    SRecord.simulate_model_result_path = result_file_path + "result_res.mat",  # omc会为结果文件添加"_res.mat"后缀
    mat_file_data = DyMatFile(result_file_path + "result_res.mat")
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
                variable_description=v[0],
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
    SRecord = session.query(SimulateRecord).filter_by(id=SRecord_id).first()
    result_file_path = "public/UserFiles/ModelResult" + '/' + username + '/' + \
                       model_name.split('.')[
                           -1] + '/' + str(
            datetime.now().strftime('%Y%m%d%H%M%S%f')) + '/'
    if file_path:
        Load_result = omc.loadFile(file_path)
    else:
        Load_result = omc.loadModel(model_name)
    file_operation = FileOperation()
    file_operation.make_dir(result_file_path)
    if Load_result != 'false\\n':
        simulate_result_str = omc.simulate(className=model_name, fileNamePrefix=result_file_path, simulate_parameters_data=simulate_parameters_data)
        SRecord.simulate_end_time = datetime.now()
        err = omc.getErrorString()
        if err != " ":
            SimulateDataHandle(SRecord, result_file_path, username, model_name, simulate_result_str)
        else:
            SRecord.simulate_status = "仿真失败"
    else:
        SRecord.simulate_status = "仿真失败"
    session.flush()
    session.close()


def Simulate(SRecord_id, username: str, model_name: str, s_type="OM", file_path: str = None, simulate_parameters_data = None):
    if s_type == "OM":
        OpenModelicaSimulate(SRecord_id, username, model_name, file_path, simulate_parameters_data)
    elif s_type == "JM":
        JModelicaSimulate(SRecord_id, username, model_name, file_path, simulate_parameters_data)
    else:
        return "暂不支持此仿真类型"