# -- coding: utf-8 --
import json

from router.simulate_router import router
from sqlalchemy import or_
from fastapi import HTTPException
from config.DB_config import DBSession
from app.model.Simulate.SimulateResult import SimulateResult
from app.model.Simulate.SimulateRecord import SimulateRecord
from app.model.ModelsPackage.ModelsInformation import ModelsInformation
from app.model.Simulate.ExperimentRecord import ExperimentRecord
from app.BaseModel.simulate import ExperimentCreateModel, SetSimulationOptionsModel
from app.service.set_simulation_options import SetSimulationOptions
from fastapi import Request, BackgroundTasks
from app.BaseModel.simulate import ModelSimulateModel, ModelCodeSaveModel, FmuExportModel
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.service.get_tree_data import GetTreeData
from app.service.get_model_code import GetModelCode
from library.file_operation import FileOperation
from app.service.get_simulation_options import GetSimulationOptions
from app.service.fmu_export import DymolaFmuExport
from library.HW_OBS_operation import HWOBS
from config.kafka_config import producer

import logging

session = DBSession()


@router.get("/getsimulationoptions", response_model=ResponseModel)
async def GetSimulationOptionsView(request: Request, package_id: str, model_name: str):
    """
        # 仿真参数获取接口
        ## model_name: 模型名称，
        ## return: 仿真仿真参数目前有五个
             startTime：仿真开始时间，
             stopTime：仿真结束时间，
             tolerance：积分方法使用的容差。<默认> = 1e-6，
             numberOfIntervals：间隔数，
             interval：间隔
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]), ModelsInformation.id==package_id).first()
    if not package:
        raise HTTPException(status_code=401, detail="not found")
    data = GetSimulationOptions(model_name)
    res.data.append(data)
    return res


@router.post("/setsimulationoptions", response_model=ResponseModel)
async def SetSimulationOptionsView(request: Request, item: SetSimulationOptionsModel):
    """
    # 仿真参数设置接口
    ## model_name: 模型名称， 全称
    ## package_id: 模型所在包的id
    ## package_name: 模型所在包的名称
    ## experiment: 仿真参数，对象字典类型，包含以下几个变量
        startTime：仿真开始时间，
        stopTime：仿真结束时间，
        tolerance：积分方法使用的容差，
        numberOfIntervals：间隔数，
        interval：间隔
    ## return:
    """
    # TODO: 系统模型仿真参数设置疑似无效，待确认
    res = InitResponseModel()
    experiment = item.experiment
    StartTime = experiment["startTime"]
    stopTime = experiment["stopTime"]
    tolerance = experiment["tolerance"]
    interval = experiment['interval']
    username = request.user.username
    package = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]), ModelsInformation.id==item.package_id).first()
    if package:
        result =SetSimulationOptions(model_name=item.model_name, StartTime=StartTime, StopTime=stopTime, Tolerance=tolerance, Interval=interval)
        if result is True:
            res.msg = "设置成功"
        else:
            res.err = "设置失败"
            res.status = 1
    else:
        res.msg = "设置成功" # 系统模型不允许设置到模型当中， 本消息只是提示参数仿真时可用，不会保存
    return res


@router.get("/getmodelstate", response_model=ResponseModel)
async def GetModelStateView (request: Request, package_id: str, model_name: str):
    """
    ## 1、初始状态, 仿真完成也是此状态
    ## 2、开始编译
    ## 3、编译完成
    ## 4、正在仿真
    ## model_name: 模型名称， 全称
    ## package_id: 模型所在包的id
    """
    res = InitResponseModel()
    model_record = session.query(SimulateRecord).filter_by(username=request.user.username, simulate_model_name=model_name).\
        filter(SimulateRecord.simulate_status.notin_(["仿真失败","仿真已结束"])).first()
    if model_record:
        res.data.append(4)
    else:
        res.data.append(1)
    return res


@router.post("/", response_model=ResponseModel)
async def ModelSimulateView (item: ModelSimulateModel, background_tasks: BackgroundTasks, request: Request):
    """
    # 仿真接口，用于模型的仿真计算
    ## package_id: 模型所在包的id,
    ## simulate_type: 仿真模型时使用的求解器是哪种,
    ## model_name: 仿真模型的名字,
    ## start_time: 仿真参数，仿真的开始时间，单位是整数秒。
    ## stop_time: 仿真参数，仿真的结束时间，单位是整数秒。
    ## number_of_intervals: 仿真参数， 间隔设置当中的间隔数。 与间隔参数是计算关系，
    ## method: 仿真参数， 选择求解器，默认参数是dassl(Openmodelica使用，dymola使用Dassl)。 与间隔参数是计算关系，
    ## return: 立即返回是否已经开始计算，仿真结果需用查看记录列表当中的记录状态是否为"仿真完成"
    """
    res = InitResponseModel()
    space_id = request.user.user_space
    username = request.user.username
    simulate_parameters_data = {
        "startTime": 0.0 if item.start_time == "" else float(item.start_time),
        "stopTime": 4.0 if item.start_time == "" else float(item.stop_time),
        "numberOfIntervals": 500 if item.start_time == "" else float(item.number_of_intervals),
        "tolerance": 0.000001 if item.start_time == "" else float(item.tolerance),
        "method": "dassl" if item.method == "" else item.method,
        # "interval": item.interval,
    }
    simulate_type = "OM" if item.simulate_type == "" else item.simulate_type
    if simulate_type not in ["OM", "JM", "DM"]:
        return res
    model = session.query(ModelsInformation).filter(
            ModelsInformation.sys_or_user.in_([request.user.username, "sys"]),
            ModelsInformation.id == item.package_id
    ).first()
    SRecord = SimulateRecord(
            username=request.user.username,
            package_id=item.package_id,
            userspace_id=space_id,
            simulate_model_name=item.model_name,
            simulate_status="正在排队等待",
            simulate_parameters_data=simulate_parameters_data,
    )
    session.add(SRecord)
    session.flush()
    MQ_data = {
            "space_id": space_id,
            "SRecord_id": SRecord.id,
            "model_name": item.model_name,
            "s_type": simulate_type,
            "file_path": model.file_path,
            "simulate_parameters_data": simulate_parameters_data,
        }
    future = producer.send(username + "_SIMULATE", key='SIMULATE'.encode(), value=json.dumps(MQ_data).encode(), partition=0)
    future.get(timeout=10)
    result = future.succeeded()
    if result:
    # background_tasks.add_task(SimulateTask, space_id, SRecord.id, request.user.username, item.model_name, simulate_type, model.file_path, simulate_parameters_data)
        res.msg = "仿真任务正在准备，请等待仿真完成"
        res.data = [SRecord.id]
    else:
        res.err = "仿真服务尚未开启，请稍后再试"
        res.status = 1
    return res


@router.get("/result", response_model=ResponseModel)
async def SimulateResultView (request: Request, variable: str, model_name: str, id: str):
    """
    # 仿真结果获取接口
    ## username: 用户名(已弃用，当前版本无须将用户名当做参数传入)
    ## variable: 模型变量名字，
    ## model_name: 模型名称，
    ## id: 仿真记录id值，在/simulate/record/list接口获取，
    ## return: 仿真结束后获取对于记录的仿真结果
    """
    res = InitResponseModel()
    space_id = request.user.user_space
    result_data = session.query(SimulateResult).filter_by(
            simulate_record_id=id,
            username=request.user.username,
            simulate_model_name=model_name,
            model_variable_name=variable,
            userspace_id=space_id
    ).first()
    if result_data:
        variable_data = {
            "abscissa": result_data.model_variable_data_abscissa,
            "ordinate": result_data.model_variable_data,
            "unit": result_data.unit if result_data.unit else "",
            "displayUnit": result_data.display_unit if result_data.display_unit else "",
        }
        res.data = [variable_data]
    else:
        res.msg = "没有查询到记录"
        res.status = 1
    return res


@router.get("/record/list", response_model=ResponseModel)
async def SimulateResultListView (request: Request):
    """
    # 仿真记录获取接口
    ## username: 用户名(已弃用，当前版本无须传入用户名当做参数)
    ## return: 返回对应用户的所有仿真记录
    """
    res = InitResponseModel()
    space_id = request.user.user_space
    record_list = session.query(
            SimulateRecord.id,
            SimulateRecord.simulate_model_name,
            SimulateRecord.simulate_status,
            SimulateRecord.create_time,
            SimulateRecord.simulate_start_time,
            SimulateRecord.simulate_end_time
    ).filter_by(username=request.user.username,
                userspace_id=space_id
                ).order_by(SimulateRecord.simulate_start_time.desc()).all()

    if record_list:
        data_list = []
        for i in record_list:
            data_dict = {
                "id": i[0],
                "simulate_model_name": i[1],
                "simulate_status": i[2],
                "create_time": i[3],
                "simulate_start_time": i[4],
                "simulate_end_time": i[5],
            }
            data_list.append(data_dict)
        res.data = data_list
    else:
        res.msg = "没有查询到记录"
        res.status = 1
    return res


@router.get("/result/tree", response_model=ResponseModel)
async def SimulateResultTreeView (id: str, variable_name: str = None):
    """
    # 仿真结果树接口， root节点只需要id， 其他子节点需要传变量名字
    ## id: 仿真记录id， 在/simulate/record/list接口获取
    ## variable_name: 模型变量名称
    ## return: 返回的是对应节点的所有子节点与其需要的数据。 description：描述， start：值， unit：显示单位， Variables：变量名， haschild：是否有子节点
    """
    # TODO: 需要重写， 此部分在结果文件中读取后处理
    res = InitResponseModel()
    if not variable_name:
        data = session.query(SimulateResult.model_variable_parent, SimulateResult.model_variable_name, SimulateResult.unit, SimulateResult.description,
                             SimulateResult.start, SimulateResult.display_unit).filter_by(simulate_record_id=id, level=1).all()
    else:
        data = session.query(SimulateResult.model_variable_parent, SimulateResult.model_variable_name, SimulateResult.unit, SimulateResult.description, SimulateResult.start, SimulateResult.display_unit).\
            filter_by(simulate_record_id=id).\
            filter(or_(SimulateResult.model_variable_parent.like(variable_name), SimulateResult.model_variable_parent.like(variable_name + "." + "%"))).all()
    tree_name_data = GetTreeData(data, variable_name)
    res.data = tree_name_data
    return res


@router.post("/experiment/create", response_model=ResponseModel)
async def ExperimentCreateView (request: Request, item: ExperimentCreateModel):
    """
    # 仿真实验创建记录接口，
    ## package_id: 保存的实验是属于哪个包id
    ## model_name: 实验属于哪个模型，全称，例如"Modelica.Blocks.Examples.PID_Controller"
    ## model_var_data: 模型的变量数据，修改过哪个模型变量，保存到当前数组对象
    ## simulate_var_data: 模型仿真选项数据
    ## experiment_name: 实验名称
    ## return: 返回是否成功状态
    """
    res = InitResponseModel()
    package_id = item.package_id
    model_name = item.model_name
    model_var_data = item.model_var_data
    simulate_var_data = item.simulate_var_data
    experiment_name = item.experiment_name
    username = request.user.username
    package_name = model_name.split(".")[0]
    if package_id != 1:
        enn = session.query(ModelsInformation).filter_by(id=package_id, sys_or_user=username).first()
    else:
        enn = session.query(ModelsInformation).filter_by(id=package_id).first()
    experimentation = session.query(ExperimentRecord).filter_by(experiment_name=experiment_name, username=username, package_id=item.package_id, model_name_all=item.model_name).first()
    if enn and not experimentation:
        ER = ExperimentRecord(
                package_id=package_id,
                model_name_all = model_name,
                model_var_data = model_var_data,
                simulate_var_data = simulate_var_data,
                experiment_name = experiment_name,
                username = username,
                package_name = package_name,
                )
        session.add(ER)
        session.flush()
        session.close()
        res.msg = "实验已创建"
    elif not enn:
        raise HTTPException(status_code=400, detail="not found")
    else:
        res.err = "实验名称已存在！"
        res.status = 2
    return res


@router.get("/experiment/list",response_model=ResponseModel)
async def ExperimentGetView (request: Request, package_id: str, model_name: str):
    """
    # 获取仿真实验列表接口，
    ## package_id: 获取的是哪个包当中的实验列表
    ## model_name： 哪个模型当中的实验列表，全称，例如："Modelica.Blocks.Examples.PID_Controller"
    ## return: 返回的实验记录列表
    """
    res = InitResponseModel()
    username = request.user.username
    experiment_all = session.query(ExperimentRecord).filter_by(username=username, package_id=package_id, model_name_all=model_name).all()
    for i in experiment_all:
        data = {
            "id": i.id,
            "experiment_name": i.experiment_name,
            "model_var_data": i.model_var_data,
            "simulate_var_data": i.simulate_var_data,
            }
        res.data.append(data)
    return res


@router.post("/simulate/codesave",response_model=ResponseModel)
async def ModelCodeSaveView (request: Request, item: ModelCodeSaveModel):
    """
    # 保存模型所在包的代码到.mo文件
    ## package_id: 包的id
    ## package_name： 包的名称
    ## return: 返回的实验记录列表
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    file_path = package.file_path
    if file_path:
        model_str = GetModelCode(item.package_name)
        FileOperation().write_file("/".join(file_path.split("/")[:-1]), item.package_name + ".mo", model_str)
        res.msg = "保存成功"
    else:
        res.err = "保存失败"
        res.status = 2
    return res


@router.post("/simulate/fmuexport",response_model=ResponseModel)
async def FmuExportModelView (request: Request, item: FmuExportModel):
    """
    # 导出模型的fmu文件
    ## package_id: 包的id
    ## package_name： 包的名称
    ## model_name： 模型全名
    ## fmu_name： fmu文件的名字
    ## fmu_par： fmu导出的参数
    ## download_local： 是否下载到本地
    ## return: 返回导出结果、如果下载到本地，则返回下载地址
    """
    res = InitResponseModel()
    username = request.user.username
    token = request.headers["Authorization"]
    file_name = ""
    file_path = ""
    model_str = ""
    package = session.query(ModelsInformation).filter_by(id=item.package_id, package_name=item.package_name, sys_or_user=username).first()
    if package:
        file_name = item.package_name
        file_path = package.file_path
        model_str = GetModelCode(item.package_name)
    res_dy = DymolaFmuExport(fmu_par=item,
                             token=token,
                             username=username,
                             file_name=file_name,
                             model_str=model_str,
                             file_path=file_path,
                             )
    if item.download_local:
        result = res_dy.get("result", None)
        logging.info("fmu导出：{}".format(res_dy))
        if result:
            # fmu_data = session.query(FmuAttachment).filter_by(create_user=username, file_name=item.fmu_name + ".fmu").first()
            file_path = res_dy.get("file_path", None)
            obs = HWOBS()
            HW_res = obs.putFile(file_path, file_path)
            logging.info("obs上传结果：{}".format(HW_res))
            if HW_res.get("status", None) == 200:
                res.data = [HW_res["body"]["objectUrl"]]
            else:
                res.err = "导出失败，请稍后再试"
                res.status = 1
        else:
            res.status = 2
            res.msg = "导出失败"
    else:
        # if not res_dy["result"]:
        res.status = 2
        res.err = "导出失败"
    return res

