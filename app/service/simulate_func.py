# -- coding: utf-8 --
import logging

from library.mat import DyMatFile
from config.omc import omc
from config.redis_config import r
from app.model.Simulate.SimulateRecord import SimulateRecord
from app.model.Simulate.SimulateResult import SimulateResult
from config.DB_config import DBSession
from datetime import datetime
import socket
from library.file_operation import FileOperation
from app.service.load_model_file import LoadModelFile
import json, requests,time
from app.service.get_model_code import GetModelCode
from library.file_operation import FileOperation
import os
import xmltodict

session = DBSession()




def SimulateDataHandle(SRecord: object, result_file_path, username, model_name, simulate_result_str):
    SRecord.simulate_result_str = simulate_result_str
    SRecord.simulate_model_result_path = result_file_path + "result_res.mat"  # omc会为结果文件添加"_res.mat"后缀
    try:
        mat_file_data = DyMatFile(result_file_path + "result_res.mat")
        with open(result_file_path + "result_init.xml", 'r') as f:
            data = f.read()
        xml_file_data = xmltodict.parse(data)
        xml_scalar_variable_data = {}
        for i in xml_file_data.get("fmiModelDescription", {}).get("ModelVariables", {}).get("ScalarVariable", {}):
            xml_scalar_variable_data[i.get("@name", "")] = i
        DefaultExperiment = xml_file_data["fmiModelDescription"].get("DefaultExperiment", {})
        SRecord.fmi_version = xml_file_data["fmiModelDescription"].get("@fmiVersion", "")
        SRecord.description = xml_file_data["fmiModelDescription"].get("@description", "")
        SRecord.start_time = DefaultExperiment.get("@startTime", "")
        SRecord.stop_time = DefaultExperiment.get("@stopTime", "")
        SRecord.step_size = DefaultExperiment.get("@stepSize", "")
        SRecord.tolerance = DefaultExperiment.get("@tolerance", "")
        SRecord.solver = DefaultExperiment.get("@solver", "")
        SRecord.output_format = DefaultExperiment.get("@outputFormat", "")
        SRecord.variable_filter = DefaultExperiment.get("@variableFilter", "")

        for k, v in mat_file_data.vars.items():
            model_variable_data_abscissa = mat_file_data.abscissa(k, True).tolist()
            data_len = len(v[1])
            step_size_f = data_len / 500
            if step_size_f < 1:
                step_size = 1
            else:
                step_size = int(step_size_f)
            variable_data = xml_scalar_variable_data.get(k, {})
            if variable_data:
                var_type = list(variable_data)[-1]
            else:
                var_type = ""
            var_type_data = variable_data.get(var_type, {})
            if var_type_data is None:
                var_type_data = {}
            model_variable_parent = k
            k_list = k.split(".")
            if len(k_list) > 1 and "der(" not in k:
                model_variable_parent_list = k.split(".")[:-1]
                model_variable_parent = ".".join(model_variable_parent_list)
                level = len(model_variable_parent_list)
            SResult = SimulateResult(
                    username=username,
                    simulate_model_name=model_name,
                    simulate_record_id=SRecord.id,
                    model_variable_name=k,
                    model_variable_parent=model_variable_parent,
                    level=level,
                    variable_description=v[0][:128],
                    model_variable_data=v[1][::step_size],
                    model_variable_data_abscissa=model_variable_data_abscissa[::step_size],
                    value_reference=variable_data.get("@valueReference", ""),
                    description=variable_data.get("@description", ""),
                    variability=variable_data.get("@variability", ""),
                    is_discrete=variable_data.get("@isDiscrete", ""),
                    causality=variable_data.get("@causality", ""),
                    is_value_changeable=variable_data.get("@isValueChangeable", ""),
                    alias=variable_data.get("@alias", ""),
                    class_index=variable_data.get("@classIndex", ""),
                    class_type=variable_data.get("@classType", ""),
                    is_protected=variable_data.get("@isProtected", ""),
                    hide_result=variable_data.get("@hideResult", ""),
                    file_name=variable_data.get("@fileName", ""),
                    start_line=variable_data.get("@startLine", ""),
                    start_column=variable_data.get("@startColumn", ""),
                    end_line=variable_data.get("@endLine", ""),
                    end_column=variable_data.get("@endColumn", ""),
                    file_writable=variable_data.get("@fileWritable", ""),
                    var_type=var_type,
                    fixed=var_type_data.get("@fixed", ""),
                    start=var_type_data.get("@start", ""),
                    use_nominal=var_type_data.get("@useNominal", ""),
                    unit=var_type_data.get("@unit", ""),
                    )
            session.add(SResult)
        SRecord.simulate_status = "仿真已结束"
        session.flush()  # 提交数据
    except Exception as e:
        print(e)
        SRecord.simulate_status = "仿真失败"
        session.flush()
        return


def JModelicaSimulate(SRecord:object, result_file_path: str, model_name: str, mo_path: str, simulate_parameters_data = None, username=None):
    res = False
    res_str = ""
    if not mo_path:
        mo_path = "omlibrary/Modelica 3.2.3"
    msg = {
        "start_time": simulate_parameters_data["startTime"],
        "final_time": simulate_parameters_data["stopTime"],
        "mo_path": "/" + mo_path,
        "result_name": "result_res.mat",
        "modelname": model_name,
        "ncp": simulate_parameters_data["numberOfIntervals"],  # 结果间隔
        "result_file_path": "/" + result_file_path ,  # 结果文件名字
        "rtol": simulate_parameters_data["tolerance"],  # 相对公差
        "type": "compile",  # 是编译还是计算， 默认是编译
    }
    file_operation = FileOperation()
    file_operation.make_dir(result_file_path)
    client = socket.socket()
    logging.debug(msg)
    logging.debug("开始连接")
    client.connect(("119.3.155.11", 56789))
    logging.debug("连接成功")
    try:
        client.send(json.dumps(msg).encode())
        compile_data = client.recv(1024).decode()
        logging.debug(compile_data)
        if compile_data == "ok":
            msg["type"] = "simulate"
            client.send(json.dumps(msg).encode())
            simulate_data = client.recv(1024).decode()
            logging.debug(simulate_data)
            if simulate_data == "ok":
                SRecord.simulate_end_time = datetime.now()
                model_name = model_name.replace(".", "_")
                with open(result_file_path + model_name + ".fmu", "rb") as f:
                    fmu_data = f.read()
                file_operation.write(result_file_path + "fmu.zip", fmu_data)
                file_operation.un_zip(result_file_path + "fmu.zip", result_file_path)
                os.rename(result_file_path + "modelDescription.xml", result_file_path + "result_init.xml")
                res = True
                res_str = "ok"
        else:
            SRecord.simulate_status = "仿真失败"
            SRecord.simulate_result_str = str(res_str)

    except Exception as e:
        SRecord.simulate_status = "仿真失败"
        SRecord.simulate_result_str = e
    client.close()
    session.flush()
    return res, res_str

def OpenModelicaSimulate(SRecord:object, result_file_path: str, model_name: str, file_path: str = None, simulate_parameters_data = None, username=None):
    res = False
    if file_path:
        package_name = model_name.split('.')[0]
        LoadModelFile(package_name, file_path)
    FileOperation().make_dir(result_file_path)
    # simulate_result_str = omc.simulate(className=model_name, fileNamePrefix=result_file_path, simulate_parameters_data=simulate_parameters_data)
    buildModel_res = omc.buildModel(className=model_name, fileNamePrefix=result_file_path, simulate_parameters_data=simulate_parameters_data)
    if buildModel_res:
        r.lpush(username + "_" + "notification", str(datetime.now().strftime('%Y-%m-%d %H:%M:%S; ')) + model_name + " 编译成功，开始仿真")
        simulate_result_str = os.popen(result_file_path + "result").read()
        if "successfully" in simulate_result_str:
            res = True
        else:
            SRecord.simulate_status = "仿真失败"
    else:
        SRecord.simulate_status = "仿真失败"
        simulate_result_str = "编译失败"
    SRecord.simulate_end_time = datetime.now()
    session.flush()
    return res, simulate_result_str

def DymolaSimulate(SRecord: object, username, model_name, file_path=None, simulate_parameters_data=None, result_file_path = None):
    res = False
    res_str = ""
    file_operation = FileOperation()
    package_name = model_name.split('.')[0]
    url = package_name + "/" + model_name.replace(".", "-") + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f')) + ""
    url_dict = {"url": username + "/" + url}
    file_name = package_name + ".mo"

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
                fmu_fileName = username + "/" + url + "/" + package_name + "/dymola_model.fmu"
            else:
                var_fileName = username + "/" + model_name.split('.')[-1] + "/" + r_simulate_data.get("msg")
                fmu_fileName = username + "/" + model_name.split('.')[-1] + "/".join(r_simulate_data.get("msg").split("/")[:-1]) + "/dymola_model.fmu"
            res_file_url = "http://121.37.183.103:8060/file/download/?fileName=" + var_fileName
            fmu_file_url = "http://121.37.183.103:8060/file/download/?fileName=" + fmu_fileName
            download_result_file = requests.get(res_file_url)
            download_fmu_file = requests.get(fmu_file_url)
            result_file_data = download_result_file.content
            fmu_file_data = download_fmu_file.content

            file_operation.write_file(result_file_path, "result_res.mat", result_file_data)
            file_operation.write_file(result_file_path, "dymola_model.fmu", fmu_file_data)
            with open(result_file_path + "dymola_model.fmu", "rb") as f:
                fmu_data = f.read()
            file_operation.write(result_file_path + "fmu.zip", fmu_data)
            file_operation.un_zip(result_file_path + "fmu.zip", result_file_path)
            os.rename(result_file_path + "modelDescription.xml", result_file_path + "result_init.xml")
            res = True
            res_str = "DM"
        else:
            SRecord.simulate_status = "仿真失败"
    else:

        SRecord.simulate_status = "仿真失败"
    session.flush()
    return res, res_str


def Simulate(SRecord_id, username: str, model_name: str, s_type="OM", file_path: str = None, simulate_parameters_data = None):
    logging.debug(s_type)
    package_name = model_name.split(".")[0]
    result_file_path = "public/UserFiles/ModelResult" + '/' + username + '/' + \
                       model_name.split('.')[
                           -1] + '/' + str(datetime.now().strftime('%Y%m%d%H%M%S%f')) + '/'
    SRecord = session.query(SimulateRecord).filter_by(id=SRecord_id).first()
    if file_path:
        model_str = GetModelCode(package_name, file_path, package_name)
        FileOperation().write_file("/".join(file_path.split("/")[:-1]), package_name + ".mo", model_str)
    r.lpush(username + "_" + "notification", str(datetime.now().strftime('%Y-%m-%d %H:%M:%S; ')) + model_name + " 模型开始编译")
    if s_type == "OM":
        s_result, s_str = OpenModelicaSimulate(SRecord, result_file_path, model_name, file_path, simulate_parameters_data, username)
    elif s_type == "JM":
        s_result, s_str = JModelicaSimulate(SRecord, result_file_path, model_name, file_path, simulate_parameters_data, username)
    elif s_type == "DM":
        s_result, s_str = DymolaSimulate(SRecord, username, model_name, file_path, simulate_parameters_data)
    else:
        return "暂不支持此仿真类型"
    if s_result:
        time.sleep(8)
        SimulateDataHandle(SRecord, result_file_path, username, model_name, simulate_result_str=s_str)
        r.lpush(username + "_" + "notification", str(datetime.now().strftime('%Y-%m-%d %H:%M:%S; ')) + model_name + " 模型仿真完成")
    else:
        r.lpush(username + "_" + "notification", str(datetime.now().strftime('%Y-%m-%d %H:%M:%S; ')) + model_name + " 编译失败")
    session.close()



# def Simulate(SRecord_id, username: str, model_name: str, s_type="OM", file_path: str = None, simulate_parameters_data = None):
#     data = [json.dumps({
#             "SRecord_id": SRecord_id,
#             "username": username,
#             "model_name": model_name,
#             "s_type": s_type,
#             "file_path": file_path,
#             "simulate_parameters_data": simulate_parameters_data,
#         })]
#     response = SimulateServiceRpc.simulate(SimulateRequest(data=data))
