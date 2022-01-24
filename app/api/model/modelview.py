# -- coding: utf-8 --
import copy
import json
import re
from datetime import datetime

from fastapi import Request, HTTPException
from app.BaseModel.respose_model import InitResponseModel, ResponseModel
from app.BaseModel.simulate import AddComponentModel, DeleteComponentModel, UpdateComponentModel, \
    UpdateConnectionAnnotationModel
from app.BaseModel.simulate import CopyClassModel, SetComponentModifierValueModel, SetComponentPropertiesModel
from app.BaseModel.simulate import DeleteConnectionModel, DeletePackageModel, GetComponentNameModel, \
    UpdateConnectionNamesModel
from app.model.ModelsPackage.ModelsInformation import ModelsInformation, ModelsInformationAll
from app.service.component_operation import AddComponent, DeleteComponent, UpdateComponent
from app.service.connection_operation import AddConnection, DeleteConnection, UpdateConnectionAnnotation, \
    UpdateConnectionNames
from app.service.copy_class import SaveClass
from app.service.get_component_name import GetComponentName
from app.service.get_components import GetComponents
from app.service.get_graphics_data import GetGraphicsData
from app.service.get_model_code import GetModelCode
from app.service.get_model_parameters import GetModelParameters
from app.service.set_component_modifier_value import SetComponentModifierValue
from app.service.set_component_properties import SetComponentProperties
from config.DB_config import DBSession
from config.omc import omc
from router.simulatemodel_router import router

session = DBSession()


@router.get("/listrootlibrary", response_model=ResponseModel)
async def GetRootModelView (request: Request):
    """
    # 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
    """
    res = InitResponseModel()
    data = []
    mn = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", request.user.username])).all()
    for i in mn:
        mn_data = {
            "package_id": i.id,
            "package_name": i.package_name,
            "sys_or_user": i.sys_or_user,
            "haschild": i.haschild
        }
        if i.sys_or_user != "sys":
            mn_data["sys_or_user"] = "user"
        data.append(mn_data)
    res.data = data
    return res


@router.get("/listlibrary", response_model=ResponseModel)
async def GetListModelView (model_name: str, request: Request):
    """
    # 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的子节点节点列表(需用传入父节点名称，返回子节点列表)，暂时没有图标信息
    ## modelname: 模型的父节点名称
    ## return：返回此父节点下的子节点列表
    """
    res = InitResponseModel()
    data = []
    username = request.user.username
    ma = session.query(
            ModelsInformationAll.model_name,
            ModelsInformationAll.haschild,
            ModelsInformationAll.sys_or_user
    ).filter_by(parent_name=model_name).filter(ModelsInformationAll.sys_or_user == username).all()
    if not ma:
        ma = session.query(
            ModelsInformationAll.model_name,
            ModelsInformationAll.haschild,
            ModelsInformationAll.sys_or_user
    ).filter_by(parent_name=model_name).filter(ModelsInformationAll.sys_or_user == "sys").all()
    for i in ma:
        mn_data = {
            "model_name": i[0],
            "haschild": i[1],
            "sys_or_user": i[2]
            }
        if i[2] != "sys":
            mn_data["sys_or_user"] = "user"
        data.append(mn_data)
    res.data = data
    return res


@router.get("/getgraphicsdata", response_model=ResponseModel)
async def GetGraphicsDataView (model_name: str, request: Request, component_name: str = None):
    """
    # 获取模型的画图数据，一次性返回， 第一次调用时间较久，有缓存机制，redis
    ## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
    ## component_name: 模型的组件名称，用于获取单个组件时传入
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    username = request.user.username
    # r_data = r.hget("GetGraphicsData_" + username, model_name)
    # if r_data:
    #     G_data = r_data.decode()
    # else:
    model_file_path = None
    # if sys_user == "user":
    package_name = model_name.split(".")[0]
    package = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
    if package:
        model_file_path = package.file_path
    if not component_name:
        data = GetGraphicsData().get_data([model_name], model_file_path)
    else:
        data = GetGraphicsData().get_one_data([model_name], component_name, model_file_path)
    # G_data = json.dumps(data)
    # r.hset("GetGraphicsData_" + username, model_name, G_data)
    # res.data = json.loads(G_data)
    res.data = data
    return res


@router.get("/getmodelcode", response_model=ResponseModel)
async def GetModelCodeView (model_name: str, sys_user: str, request: Request):
    """
    # 获取模型的源码数据，一次性返回
    ## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    # username = request.user.username
    # package_name = model_name.split(".")[0]
    # if sys_user == "user":
    #     package = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
    #     if not package:
    #         raise HTTPException(status_code=400, detail="not found")
    data = GetModelCode(model_name)
    res.data = [data]
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
    username = request.user.username
    path = None
    package_name = model_name.split(".")[0]
    if sys_user == "user":
        model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
        path = model.file_path
    data = GetModelParameters(model_name, name, components_name, path, package_name).get_data()
    res.data = data
    return res


@router.post("/setmodelparameters", response_model=ResponseModel)
async def SetModelParametersView (item: SetComponentModifierValueModel, request: Request):
    """
    # 设置模型组件的参数数据，一次性返回
    ## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ## parameter_value: 需要设置的变量和新的值，全称，例如{"PID.k": "200"}， k是模型的组件别名和变量名字的组成， 类似于“别名.变量名”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    username = request.user.username
    package_name = item.model_name.split(".")[0]
    model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
    if not model:
        res.err = "设置失败"
        res.status = 2
        return res
    path = model.file_path
    data = SetComponentModifierValue(item.model_name, item.parameter_value, path, package_name)
    if data == "Ok":
        res.msg = "设置完成"
    else:
        res.err = "设置失败: " + data
        res.status = 1
    return res


@router.get("/getcomponentproperties", response_model=ResponseModel)
async def GetComponentPropertiesView (model_name: str, component_name: str, sys_user: str, request: Request):
    """
    # 获取模型组件的属性数据，一次性返回
    ##  class_name: 需要查询属性数据的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ##  component_name: 需要查询的组件别名，全称，“PID”
    ##  sys_user: 需要查询的模型组件是系统还是用户模型
    """
    res = InitResponseModel()
    username = request.user.username
    package_name = model_name.split(".")[0]
    file_path = None
    if sys_user == "user":
        model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
        if not model:
            res.err = "查询失败"
            res.status = 2
            return res
        file_path = model.file_path
    result = GetComponents(model_name, component_name, file_path, package_name)
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
    return res


@router.post("/setcomponentproperties", response_model=ResponseModel)
async def SetComponentPropertiesView (item: SetComponentPropertiesModel, request: Request):
    """
    # 设置模型组件的属性数据，一次性返回
    ## class_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ## component_name: 需要查询的组件别名，全称，“PID”
    ## final: "true" or "false",
    ## protected: "true" or "false",
    ## replaceable: "true" or "false",
    ## variabilty: "unspecified" or  "parameter" or "discrete" or "constant"
    ## inner: "true" or "false",
    ## outer: "true" or "false",
    ## causality: "output" or "input"
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
    username = request.user.username
    package_name = item.model_name.split(".")[0]
    model = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
    if not model:
        res.err = "设置失败"
        res.status = 2
        return res
    result = SetComponentProperties(model.file_path, package_name, **parameters_data)
    if result == "Ok":
        res.msg = "设置成功"
    else:
        res.err = "设置失败"
        res.status = 2
    return res


@router.post("/copyclass", response_model=ResponseModel)
async def CopyClassView (item: CopyClassModel, request: Request):
    """
    # 复制模型
    ## parent_name: 需要复制到哪个父节点之下，例如“ENN.Examples”
    ## package_name: 被复制的模型在哪个包之下，例如“ENN”
    ## class_name: 复制之后的模型名称，例如“Scenario1_Status_test”
    ## copied_class_name: 被复制的模型全称，例如“ENN.Examples.Scenario1_Status”
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package_name = item.package_name
    package = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
    if package:
        model = session.query(ModelsInformationAll).filter(
                ModelsInformationAll.package_name == package_name,
                ModelsInformationAll.package_id == package.id,
                ModelsInformationAll.model_name == item.class_name,
                ModelsInformationAll.sys_or_user == username,
                ).first()
        if model:
            res.err = "模型名称已存在"
            res.status = 1
            return res
        model_file_path = package.file_path.split("/")
        model_file_path[-2] = datetime.now().strftime('%Y%m%d%H%M%S%f')
        model_file_path = "/".join(model_file_path)
        save_result = SaveClass(item.class_name, item.copied_class_name, item.parent_name, package_name, model_file_path=package.file_path, new_model_file_path=model_file_path)
        if save_result:
            package.file_path = model_file_path
            model = session.query(ModelsInformationAll).filter(
                            ModelsInformationAll.model_name_all == item.copied_class_name,
                            ModelsInformationAll.package_name == package_name
                        ).first()
            if model:
                child_name = model.child_name
                haschild = model.haschild
            else:
                child_name = []
                haschild = False
            ModelsInformationAll_new = ModelsInformationAll(
                    package_name=package_name,
                    package_id=package.id,
                    model_name=item.class_name,
                    parent_name=item.parent_name,
                    child_name=child_name,
                    haschild=haschild,
                    model_name_all=item.parent_name + "." + item.class_name,
                    sys_or_user=username
            )
            session.add(ModelsInformationAll_new)
            model_parent = session.query(ModelsInformationAll).filter(
                    ModelsInformationAll.model_name_all == item.parent_name,
                    ModelsInformationAll.package_name == package_name,
                    ModelsInformationAll.sys_or_user == username
                    ).first()
            if not model_parent:
                model_parent = session.query(ModelsInformation).filter(
                    ModelsInformation.package_name == package_name,
                    ModelsInformation.sys_or_user == username).first()
            m_child_name = copy.deepcopy(model_parent.child_name)
            m_child_name.append(item.class_name)
            model_parent.child_name = m_child_name
            model_parent.haschild = True
            res.msg = "复制成功"
            session.flush()
            session.close()
        else:
            res.err = "复制失败"
            res.status = 1
    else:
        res.err = "复制失败"
        res.status = 2
    return res


@router.post("/delete_package", response_model=ResponseModel)
async def DeletePackageAndModelView(request: Request, item: DeletePackageModel):
    """
    # 删除模型包或包中的模型
    ## parent_name: 需要删除的模型在哪个父节点之下，例如“ENN.Examples”
    ## package_name: 被删除的模型在哪个包之下，例如“ENN”，如果删除的是包，则就是包的名字，
    ## class_name: 被删除的的模型名称，例如“Scenario1_Status_test”
    ## return: 返回json格式数据,告知是否成功
    """
    package_name = item.package_name
    package_id = item.package_id
    parent_name = item.parent_name
    class_name = item.class_name
    username = request.user.username
    if parent_name:
        model_name_all = parent_name + "." + class_name
        package = session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=username).first()
        package_id = package.id
        if package:
            model_file_path = package.file_path.split("/")
            model_file_path[-2] = datetime.now().strftime('%Y%m%d%H%M%S%f')
            model_file_path = "/".join(model_file_path)
            save_result = SaveClass(class_name=model_name_all, package_name=package_name , model_file_path=package.file_path, copy_or_delete="delete", new_model_file_path=model_file_path)
            if save_result:
                package.file_path = model_file_path
                session.query(ModelsInformationAll).filter_by(model_name_all=model_name_all, package_id=package_id, sys_or_user=username).delete(synchronize_session=False)
                model_parent = session.query(ModelsInformationAll).filter(
                        ModelsInformationAll.model_name_all == parent_name,
                        ModelsInformationAll.package_id == package_id,
                        ModelsInformationAll.sys_or_user == username,
                        ).first()
                if model_parent:
                    child_name = copy.deepcopy(model_parent.child_name)
                    model_parent.child_name = [i for i in child_name if i != class_name]
                    if not model_parent.child_name:
                        model_parent.haschild = False
                else:
                    child_name = copy.deepcopy(package.child_name)
                    package.child_name = [i for i in child_name if i != class_name]
                    if not package.child_name:
                        package.haschild = False
    else:
        session.query(ModelsInformation).filter_by(id=package_id, sys_or_user=username).delete(synchronize_session=False)
        session.query(ModelsInformationAll).filter_by(package_id=package_id, sys_or_user=username).delete(synchronize_session=False)
    res = InitResponseModel()
    session.flush()
    session.close()
    res.msg = "删除成功"
    return res


@router.post("/get_component_name", response_model=ResponseModel)
async def GetComponentNameView(item: GetComponentNameModel, request: Request):
    """
    # 创建模型当中的模型组件
    ## package_name： 需要创建的组件在哪个包之下，例如在"NN.Examples.Scenario1_Status"模型中创建组件，包就是ENN
    ## package_id： 包id
    ## model_name_all: 需要创建的组件在哪个模型之下，例如在"NN.Examples.Scenario1_Status"模型中创建组件
    ## old_component_name: 被创建成组件的模型名称， 例如"Modelica.Blocks.Math.Abs"
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    name = GetComponentName(item.model_name_all, item.old_component_name)
    res.data = [name]
    return res


@router.post("/add_component", response_model=ResponseModel)
async def AddModelComponentView (item: AddComponentModel, request: Request):
    """
    # 创建模型当中的模型组件
    ## package_name： 需要创建的组件在哪个包之下，例如在"NN.Examples.Scenario1_Status"模型中创建组件，包就是ENN
    ## package_id： 包id
    ## model_name_all: 需要创建的组件在哪个模型之下，例如在"NN.Examples.Scenario1_Status"模型中创建组件
    ## new_component_name: 新创建的组件名称，例如"abs1"
    ## old_component_name: 被创建成组件的模型名称， 例如"Modelica.Blocks.Math.Abs"
    ## origin: 原点， 例如"0,0"
    ## extent: 范围坐标, 例如["-10,-10", "10,10"]
    ## rotation: 旋转角度, 例如"0"，不旋转
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if package:
        result, err = AddComponent(item.new_component_name, item.old_component_name, item.model_name_all, item.origin, item.extent, item.rotation, package.file_path, package.package_name)
        if result is True:
            res.msg = "新增组件成功"
        else:
            res.err = "新增组件失败，名称为" + item.new_component_name + " 的组件已经存在或者是 Modelica 关键字。 请选择其他名称。"
            res.status = 1
    else:
        res.err = "新增组件失败"
        res.status = 2
    return res


@router.post("/delete_component", response_model=ResponseModel)
async def DeleteModelComponentView(item: DeleteComponentModel, request: Request):
    """
    # 删除模型当中的模型组件
    ## package_name： 包名称
    ## package_id： 包id
    ## component_name： 要删除的组件名称
    ## model_name_all： 要删除的组件在哪个模型当中，是模型的全名
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if package:
        result = True
        for i in item.delete_list:
            if i["delete_type"] == "component":
                result = DeleteComponent(i["component_name"], i["model_name_all"], package.file_path, package.package_name)
            elif i["delete_type"] == "connector":
                result = DeleteConnection(i["model_name_all"], i["connect_start"], i["connect_end"], package.file_path, package.package_name)
            else:
                result = False
                break
        # result = DeleteComponent(item.component_name, item.model_name_all, package.file_path, package.package_name)

        if result:
            res.msg = "删除组件成功"
        else:
            res.err = "删除组件失败"
            res.status = 1
    else:
        res.err = "删除组件失败"
        res.status = 2
    return res


@router.post("/update_component", response_model=ResponseModel)
async def UpdateModelComponentView(item: UpdateComponentModel, request: Request):
    """
    # 更新模型当中的模型组件
    ## package_name： 包名称
    ## package_id： 包id
    ## component_name: 需要更新的组件名称，例如"limPID"，
    ## component_model_name: 需要更新的组件原本模型名称，例如"Modelica.Blocks.Continuous.LimPID"
    ## model_name_all: 需要更新的组件在哪个模型当中， 例如"Modelica.Blocks.Examples.PID_Controller"
    ## origin: 原点， 例如"0,0"
    ## extent: 范围坐标, 例如["-10,-10", "10,10"]
    ## rotation: 旋转角度, 例如"0"，不旋转
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    origin = item.origin
    extent = item.extent
    if origin == "-,-":
        p1 = list(map(float, item.extent[0].split(',')))
        p2 = list(map(float, item.extent[1].split(',')))
        left_p = (p1[0] + p2[0]) / 2
        right_p = (p1[1] + p2[1]) / 2
        origin = ",".join([str(left_p), str(right_p)])
        extent_1 = [str(p1[0] - left_p), str(p1[1] - right_p)]
        extent_2 = [str(p2[0] - left_p), str(p2[1] - right_p)]
        extent = [",".join(extent_1), ",".join(extent_2)]

    if package:
        result = UpdateComponent(item.component_name, item.component_model_name, item.model_name_all, origin, extent, item.rotation, package.file_path, package.package_name)
        if result is True:
            res.msg = "更新组件成功"
        else:
            res.err = "更新组件失败" + str(result)
            res.status = 1
    else:
        res.err = "更新组件失败"
        res.status = 2
    return res


@router.post("/create_connection_annotation", response_model=ResponseModel)
async def CreateConnectionAnnotationView(item: UpdateConnectionAnnotationModel, request: Request):
    """
    # 创建模型当中的组件连线
    ## package_name： 包名称
    ## package_id： 包id
    ## model_name_all：在哪个模型创建，模型全称
    ## connect_start：连线起点， 输出点， 例如"sum1.y"
    ## connect_end：连线终点， 输入点， 例如"ChebyshevI.u"
    ## color：连线颜色， 例如"0,0,127"
    ## line_points：连线拐点坐标，包含起始点坐标，从起点开始到终点 例如["213,-38","-163.25,-38","-163.25,-4"]
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if package:
        result = AddConnection(item.model_name_all, item.connect_start, item.connect_end, item.line_points, item.color, package.file_path, package.package_name)
        if result == "Ok":
            res.msg = "连接成功"
            expression_inout = r"\[\d+\]$"
            expression_component = r"\[\d+\]\."
            connect_start_inout = re.sub(expression_inout, "", item.connect_start)
            connect_end_inout = re.sub(expression_inout, "", item.connect_end)
            res.data = [{
                "arrow": "Arrow.None,Arrow.None",
                "arrowSize": "3",
                "color": "0,0,127",
                "connectionfrom": re.sub(expression_component, "..", connect_start_inout),
                "connectionfrom_original_name": item.connect_start,
                "connectionto": re.sub(expression_component, "..", connect_end_inout),
                "connectionto_original_name": item.connect_end,
                "linePattern": "LinePattern.Solid",
                "lineThickness": "0.25",
                "originalPoint": "0.0,0.0",
                "points": item.line_points,
                "rotation": "0",
                "smooth": "Smooth.None",
                "type": "Line",
                "visible": "true",
                        }]
        else:
            res.err = "连接失败, err: " + result
            res.status = 1
    else:
        res.err = "连接失败"
        res.status = 2
    return res


@router.post("/update_connection_names", response_model=ResponseModel)
async def UpdateConnectionNamesView(item: UpdateConnectionNamesModel, request: Request):
    """
    # 创建模型当中的组件连线
    ## package_name： 包名称
    ## package_id： 包id
    ## model_name_all：在哪个模型修改，模型全称
    ## from_name：连线起点， 输出点， 例如"sum1.y"
    ## to_name：连线终点， 输入点， 例如"sum2.y"
    ## from_name_new：连线起点， 输出点， 例如"sum1new.y"
    ## to_name_new：连线终点， 输入点， 例如"sum2new.y"
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if package:
        result = UpdateConnectionNames(item.model_name_all, item.from_name, item.to_name, item.from_name_new, item.to_name_new)
        if result == "Ok":
            res.msg = "连接起止点更新成功"
        else:
            res.err = "连接起止点更新失败, err: " + result
            res.status = 1
    else:
        res.err = "连接起止点更新失败"
        res.status = 2
    return res


@router.post("/delete_connection_annotation", response_model=ResponseModel)
async def DeleteConnectionAnnotationView(item: DeleteConnectionModel, request: Request):
    """
    # 删除模型当中的删除组件连线
    ## package_name： 包名称
    ## package_id： 包id
    ## model_name_all： 在哪个模型当中删除连线
    ## connect_start： 连线起始位置
    ## connect_end： 连线终止位置
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if package:
        result = DeleteConnection(item.model_name_all, item.connect_start, item.connect_end, package.file_path, package.package_name)
        if result  == "Ok":
            res.msg = "删除成功"
        else:
            res.err = "删除失败: " + result
            res.status = 1
    else:
        res.err = "删除失败"
        res.status = 2

    return res


@router.post("/update_connection_annotation", response_model=ResponseModel)
async def UpdateConnectionAnnotationView(item: UpdateConnectionAnnotationModel, request: Request):
    """
    # 更新模型当中的组件连线
    ## package_name： 包名称
    ## package_id： 包id
    ## model_name_all：在哪个模型中更新，模型全称
    ## connect_start：连线起点， 输出点， 例如"sum1.y"
    ## connect_end：连线终点， 输入点， 例如"ChebyshevI.u"
    ## color：连线颜色， 例如"0,0,127"
    ## line_points：连线拐点坐标，包含起始点坐标，从起点开始到终点 例如["213,-38","-163.25,-38","-163.25,-4"]
    ## return: 返回json格式数据,告知是否成功
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if package:
        result = UpdateConnectionAnnotation(item.model_name_all, item.connect_start, item.connect_end, item.line_points, item.color, package.file_path, package.package_name)
        if result is True:
            res.msg = "连接拐点修改成功"
        else:
            res.err = "连接拐点修改失败: " + str(result)
            res.status = 1
    else:
        res.err = "连接修改失败"
        res.status = 2
    return res


@router.get("/test")
async def _test (model_name: str, request: Request):

    # username = request.user.username
    # r.hdel("GetGraphicsData_" + username, model_name)
    res = omc.sendExpression(model_name)
    # res = request.auth
    return {"msg": res,
            # "user": request.user.display_name,
            "auth": request.user.is_authenticated
        }
