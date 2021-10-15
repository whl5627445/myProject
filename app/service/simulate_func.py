import datetime
from library.mat import DyMatFile
from config.omc import omc
from app.model.Simulate.SimulateRecord import SimulateRecord
from app.model.Simulate.SimulateResult import SimulateResult
from config.DB_config import session
import time
import socket
from library.file_operation import FileOperation
import json


# 获取本机ip, 用于访问docker服务
hostname = socket.gethostname()
ip = socket.gethostbyname(hostname)


def SimulateDataHandle(SRecord: object, result_file_path, username, model_name, simulate_result_str):
    SRecord.simulate_result_str = simulate_result_str
    SRecord.simulate_model_result_path = result_file_path + "result.mat",  # omc会为结果文件添加"_res.mat"后缀
    mat_file_data = DyMatFile(result_file_path + "result.mat")
    SRecord.simulate_nametree = mat_file_data.nameTree()
    for k, v in mat_file_data.vars.items():
        model_variable_data_abscissa = mat_file_data.abscissa(k, True).tolist()
        SResult = SimulateResult(
                username=username,
                simulate_model_name=model_name,
                simulate_record_id=SRecord.id,
                model_variable_name=k,
                variable_description=v[0],
                model_variable_data=v[1],
                model_variable_data_abscissa=model_variable_data_abscissa
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
            time.time())[:10] + '/'
    msg = {
        "mo_path": mo_path,
        "modelname": model_name,
        "ncp": simulate_parameters_data["numberOfIntervals"],  # 结果间隔
        "result_file_path": result_file_path,  # 结果文件名字
        "rtol": simulate_parameters_data["tolerance"],  # 相对公差
    }
    file_operation = FileOperation()
    file_operation.make_dir(result_file_path)
    client = socket.socket()

    client.connect((ip, 56789))
    client.send(json.dumps(msg).encode())
    data = json.loads(client.recv(1024))
    client.close()
    SRecord.simulate_end_time = datetime.datetime.now()
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
            time.time())[:10] + '/'
    if file_path:
        Load_result = omc.loadFile(file_path)
    else:
        Load_result = omc.loadModel(model_name)
    if Load_result != 'false\\n':
        SRecord.simulate_end_time = datetime.datetime.now()
        simulate_result_str = omc.simulate(model_name, result_file_path, simulate_parameters_data)
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
