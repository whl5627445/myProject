from router.simulate import router
from config.DB_config import DBSession
from app.model.Simulate.SimulateResult import SimulateResult
from app.model.Simulate.SimulateRecord import SimulateRecord
from app.model.models_package.ModelsInformation import ModelsInformationAll, ModelsInformation
from app.service.simulate_func import Simulate
from fastapi import Request, BackgroundTasks
from app.BaseModel.simulate import ModelSimulateModel
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.service.get_tree_data import GetTreeData
from app.service.get_simulation_options import GetSimulationOptions
session = DBSession()


@router.get("/getsimulationoptions", response_model=ResponseModel)
async def GetSimulationOptionsView(request: Request, model_name: str):
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
    try:
        MI_all = session.query(ModelsInformationAll).filter(
                ModelsInformationAll.sys_or_user.in_([request.user["username"], "sys"]),
                ModelsInformationAll.model_name_all == model_name
        ).first()
        model = session.query(ModelsInformation).filter(
                ModelsInformation.sys_or_user.in_([request.user["username"], "sys"]),
                ModelsInformation.package_name == MI_all.package_name
        ).first()
        data = GetSimulationOptions([model_name], model.file_path)
        res.data = data
    except Exception as e:
        print(e)
        res.err = "获取数据失败，请重试"
        res.status = 1
    return res


@router.post("/", response_model=ResponseModel)
async def ModelSimulateView (item: ModelSimulateModel, background_tasks: BackgroundTasks, request: Request):
    """
    # 仿真接口，用于模型的仿真计算
    ## simulate_type: 仿真模型时使用的求解器是哪种,
    ## model_name: 仿真模型的名字,
    ## start_time: 仿真参数，仿真的开始时间，单位是整数秒。
    ## stop_time: 仿真参数，仿真的结束时间，单位是整数秒。
    ## number_of_intervals: 仿真参数， 间隔设置当中的间隔数。 与间隔参数是计算关系，
    ## return: 立即返回是否已经开始计算，仿真结果需用查看记录列表当中的记录状态是否为"仿真完成"
    """
    res = InitResponseModel()
    try:
        simulate_parameters_data = {
            "startTime": item.start_time,
            "stopTime": item.stop_time,
            "numberOfIntervals": item.number_of_intervals,
            "tolerance": item.tolerance,
            # "interval": item.interval,
        }
        MI_all = session.query(ModelsInformationAll).filter(
                ModelsInformationAll.sys_or_user.in_([request.user["username"], "sys"]),
                ModelsInformationAll.model_name_all == item.model_name
        ).first()
        model = session.query(ModelsInformation).filter(
                ModelsInformation.sys_or_user.in_([request.user["username"], "sys"]),
                ModelsInformation.package_name == MI_all.package_name
        ).first()

        SRecord = SimulateRecord(
                username=request.user["username"],
                simulate_model_name=item.model_name,
                simulate_status="仿真进行中",
        )
        session.add(SRecord)
        session.flush()
        background_tasks.add_task(Simulate, SRecord.id, request.user["username"], item.model_name, item.simulate_type, model.file_path, simulate_parameters_data)
        res.msg = "仿真任务已开始，请等待仿真完成"
        res.data = [SRecord.id]
    except Exception as e:
        print(e)
        res.err = "仿真任务创建失败，请重试"
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
    result_data = session.query(SimulateResult).filter_by(
            simulate_record_id=id,
            username=request.user["username"],
            simulate_model_name=model_name,
            model_variable_name=variable
    ).first()
    if result_data:
        variable_data = {
            "abscissa": result_data.model_variable_data_abscissa,
            "ordinate": result_data.model_variable_data
        }
        res.data = [variable_data]
    else:
        res.msg = "没有查询到记录"
        res.status = 1
    return res


@router.get("/record/list", response_model=ResponseModel)
async def SimulateResultTreeView (request: Request):
    """
    # 仿真记录获取接口
    ## username: 用户名(已弃用，当前版本无须传入用户名当做参数)
    ## return: 返回对应用户的所有仿真记录
    """
    res = InitResponseModel()
    record_list = session.query(
            SimulateRecord.id,
            SimulateRecord.simulate_model_name,
            SimulateRecord.simulate_status,
            SimulateRecord.simulate_start_time,
            SimulateRecord.simulate_end_time
    ).filter_by(username=request.user["username"]).order_by(SimulateRecord.simulate_start_time.desc()).all()
    if record_list:
        data_list = []
        for i in record_list:
            data_dict = {
                "id": i[0],
                "simulate_model_name": i[1],
                "simulate_status": i[2],
                "simulate_start_time": i[3],
                "simulate_end_time": i[4],
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
    ## return: 返回的是对应节点的所有子节点
    """
    res = InitResponseModel()
    tree = session.query(SimulateRecord).filter_by(id=id).first()
    name_tree = tree.simulate_nametree
    if name_tree:
        tree_name_data = GetTreeData(name_tree, variable_name)
        res.data = tree_name_data
    else:
        res.msg = "没有查询到记录"
        res.status = 1
    return res