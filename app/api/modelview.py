from fastapi import Request
from router.simulatemodel import router
from config.DB_config import session
from config.redis_config import r
from app.model.models_package.ModelsInformation import ModelsInformationAll, ModelsInformation
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from library.get_graphics_data import GetGraphicsData
from library.get_model_code import GetModelCode
from library.get_model_parameters import GetModelParameters
from library.set_component_modifier_value import SetComponentModifierValue
from library.set_component_properties import SetComponentProperties
from library.get_components import GetComponents
from app.BaseModel.simulate import SetComponentModifierValueModel, SetComponentPropertiesModel
import time
import json

@router.get("/listrootlibrary", response_model=ResponseModel)
async def GetRootModelView (request: Request):
    """
    # 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
    """
    res = InitResponseModel()
    try:
        data = []
        mn = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", request.user["username"]])).all()
        for i in mn:
            mn_data = {
                "package_name": i.package_name,
                "sys_or_user": i.sys_or_user,
                "image": i.image,
                "haschild": i.haschild
            }
            if i.sys_or_user != "sys":
                mn_data["sys_or_user"] = "user"
            data.append(mn_data)
        res.data = data
    except Exception as e:
        session.rollback()
        print(e)
        res.err = "获取模型列表失败，请稍后再试"
        res.status = 1
    return res


@router.get("/listlibrary", response_model=ResponseModel)
async def GetModelView (modelname: str, request: Request):
    """
    # 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的子节点节点列表(需用传入父节点名称，返回子节点列表)，暂时没有图标信息
    ## modelname: 模型的父节点名称
    ## return：返回此父节点下的子节点列表
    """
    res = InitResponseModel()
    try:
        data = []
        mn = session.query(
                ModelsInformationAll.model_name,
                ModelsInformationAll.haschild,
                ModelsInformationAll.image,
                ModelsInformationAll.sys_or_user
        ).filter_by(parent_name=modelname).filter(ModelsInformationAll.sys_or_user.in_(["sys", request.user["username"]])).all()
        for i in mn:
            mn_data = {"model_name": i[0], "haschild": i[1], "image": i[2], "sys_or_user": i[3]}
            if i[3] != "sys":
                mn_data["sys_or_user"] = "user"
            data.append(mn_data)
        res.data = data
    except Exception as e:
        session.rollback()
        print(e)
        res.err = "获取模型列表失败，请稍后再试"
        res.status = 1
    return res


@router.get("/getgraphicsdata", response_model=ResponseModel)
async def GetGraphicsDataView (modelname: str, sys_user: str, request: Request):
    """
    # 获取模型的画图数据，一次性返回， 第一次调用时间较久，有缓存机制
    ## modelname: 需要查询的模型名称，全称， 例如“ENN.Examples.Scenario1_Status”
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    try:
        username = request.user["username"]
        r_data = r.hget("GetGraphicsData_" + username, modelname)
        if r_data:
            data = r_data.decode()
        else:
            model_file_path = None
            if sys_user == "user":
                package_name = modelname.split(".")[0]
                package = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
                model_file_path = package.file_path
            data = GetGraphicsData().get_data([modelname], model_file_path)
            data = json.dumps(data)
            r.hset("GetGraphicsData_" + username, modelname, data)
        res.data = json.loads(data)
    except Exception as e:
        print(e)
        res.err = "获取数据失败，请稍后再试"
        res.status = 1
    return res


@router.get("/getmodelcode", response_model=ResponseModel)
async def GetModelCodeView (modelname: str, sys_user: str, request: Request):
    """
    # 获取模型的源码数据，一次性返回
    ## modelname: 需要查询的模型名称，全称， 例如“ENN.Examples.Scenario1_Status”
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    try:
        username = request.user["username"]
        path = None
        if sys_user == "user":
            package_name = modelname.split(".")[0]
            model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
            path = model.file_path
        data = GetModelCode(modelname, path)
        res.data = [data]
    except Exception as e:
        print(e)
        res.err = "获取数据失败，请稍后再试"
        res.status = 1
    return res


@router.get("/getmodelparameters", response_model=ResponseModel)
async def GetModelParametersView (model_name: str, sys_user: str, name: str, components_name: str, request: Request):
    """
    # 获取模型组件的参数数据，一次性返回
    ## model_name: 需要查询的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ## components_name: 需要查询模型的组件名称，全称， 例如“Modelica.Blocks.Continuous.LimPID“
    ## name: 需要查询的组件别名，全称，“PID”
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    try:
        username = request.user["username"]
        path = None

        if sys_user == "user":
            package_name = model_name.split(".")[0]
            model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
            path = model.file_path
        data = GetModelParameters(model_name, name, components_name, path)
        res.data = data
    except Exception as e:
        print(e)
        res.err = "获取数据失败，请稍后再试"
        res.status = 1
    return res


@router.post("/setmodelparameters", response_model=ResponseModel)
async def SetModelParametersView (item: SetComponentModifierValueModel, request: Request):
    """
    # 设置模型组件的参数数据，一次性返回
    ## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ## parameter_value: 需要设置的变量和新的值，全称，例如{"PI.k": "200"}， k是模型的组件别名和变量名字的组成， 类似于“别名.变量名”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    try:
        username = request.user["username"]
        package_name = item.model_name.split(".")[0]
        model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
        if not model:
            res.err = "设置失败"
            res.status = 2
            return res
        path = model.file_path
        data = SetComponentModifierValue(item.model_name, item.parameter_value, path)
        if data == "Ok":
            res.msg = "设置完成"
        else:
            res.err = "设置失败"
            res.status = 1
    except Exception as e:
        print(e)
        res.err = "设置失败"
        res.status = 1
    return res


@router.get("/getcomponentproperties", response_model=ResponseModel)
async def GetComponentPropertiesView (model_name: str, component_name: str, sys_user: str, request: Request):
    """
    # 获取模型组件的属性数据，一次性返回
        "class_name": 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
        "component_name": 需要查询的组件别名，全称，“PID”
    """
    res = InitResponseModel()

    try:
        username = request.user["username"]
        package_name = model_name.split(".")[0]
        model = None
        if sys_user == "user":
            model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
        if not model:
            res.err = "设置失败"
            res.status = 2
            return res
        result = GetComponents(model_name, component_name, model.file_path)
        data = {
            "model_name": model_name,
            "component_name": component_name,
            "path": result[0],
            "dimension": str(result[-1]).replace("['']", "[]"),
            "annotation": str(result[2]),
            "Properties": [result[4], result[3], result[7]],
            "Variability": result[8],
            "Inner/Outer": result[9],
            "Causality": result[10],
        }
        res.data = [data]
    except Exception as e:
        print(e)
        res.err = "获取数据失败"
        res.status = 1
    return res


@router.post("/setcomponentproperties", response_model=ResponseModel)
async def SetComponentPropertiesView (item: SetComponentPropertiesModel, request: Request):
    """
    # 设置模型组件的属性数据，一次性返回
        "class_name": 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
        "component_name": 需要查询的组件别名，全称，“PID”
        "final": "true" or "false",
        "protected": "public" or "protected",
        "replaceable": "true" or "false",
        "variabilty": "unspecified" or  "parameter" or "discrete" or "constant"
        "inner": "true" or "false",
        "outer": "true" or "false",
        "causality": "output" or "input"
    """
    res = InitResponseModel()
    parameters_data = {
        "class_name": item.model_name,
        "component_name": item.component_name,
        "final": item.final,
        "protected": item.protected,
        "replaceable": item.replaceable,
        "variabilty": item.variabilty,
        "inner": item.inner,
        "outer": item.outer,
        "causality": item.causality,
    }
    try:
        username = request.user["username"]
        package_name = item.model_name.split(".")[0]
        model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
        if not model:
            res.err = "设置失败"
            res.status = 2
            return res
        result = SetComponentProperties(model.file_path, **parameters_data)
        print(result)
        if result == "Ok":
            res.msg = "设置成功"
        else:
            res.err = "设置失败"
            res.status = 2
    except Exception as e:
        print(e)
        res.err = "设置失败"
        res.status = 1
    return res


@router.get("/test", response_model=ResponseModel)
async def test (modelname: str, request: Request):
    username = request.user["username"]
    r.hdel("GetGraphicsData_" + username, modelname)
    return {"msg": "ok"}
