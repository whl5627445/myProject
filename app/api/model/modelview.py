# -- coding: utf-8 --
import copy
import json
import logging
import re
from datetime import datetime

from fastapi import Request, HTTPException
from app.BaseModel.respose_model import InitResponseModel, ResponseModel
from app.BaseModel.simulate import AddComponentModel, DeleteComponentModel, UpdateComponentModel, \
    UpdateConnectionAnnotationModel
from app.BaseModel.simulate import CopyClassModel, SetComponentModifierValueModel, SetComponentPropertiesModel
from app.BaseModel.simulate import DeleteConnectionModel, DeletePackageModel, GetComponentNameModel, \
    UpdateConnectionNamesModel, SetModelDocumentModel, ConvertUnitsModel
from app.model.ModelsPackage.ModelsInformation import ModelsInformation
from app.service.component_operation import AddComponent, DeleteComponent, UpdateComponent
from app.service.connection_operation import AddConnection, DeleteConnection, UpdateConnectionAnnotation, \
    UpdateConnectionNames
from app.service.unit_operation import ConvertUnits
from app.service.copy_class import SaveClass
from app.service.get_component_name import GetComponentName
from app.service.get_components import GetComponents
from app.service.get_graphics_data import GetGraphicsData
from app.service.get_model_code import GetModelCode
from app.service.get_model_parameters import GetModelParameters
from app.service.set_component_modifier_value import SetComponentModifierValue
from app.service.set_component_properties import SetComponentProperties
from app.service.check_model import CheckModel
from app.service.get_model_child import GetModelChild, GetModelHasChild
from app.service.icon_operation import GetIcon
from app.service.model_document_operation import GetModelDocument, SetModelDocument
from config.DB_config import DBSession
from config.redis_config import r
from sqlalchemy import or_,and_
from router.simulatemodel_router import router

session = DBSession()


@router.get("/listrootlibrary", response_model=ResponseModel)
async def GetRootModelView (request: Request):
    """
    # 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
    """
    res = InitResponseModel()
    data = []
    space_id = request.user.user_space
    username = request.user.username
    models_obj_list = session.query(ModelsInformation).filter(or_(ModelsInformation.sys_or_user == "sys", and_(ModelsInformation.userspace_id ==space_id, ModelsInformation.sys_or_user == username))).all()
    for i in models_obj_list:
        mn_data = {
            "package_id": i.id,
            "package_name": i.package_name,
            "sys_or_user": i.sys_or_user,
            "haschild": GetModelHasChild(i.package_name),
            "image": GetIcon(i.package_name),
        }
        if i.sys_or_user != "sys":
            mn_data["sys_or_user"] = "user"
        data.append(mn_data)
    res.data = data
    return res


@router.get("/listlibrary", response_model=ResponseModel)
async def GetListModelView (package_id: str, model_name: str, request: Request):
    """
    # 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的子节点节点列表(需用传入父节点名称，返回子节点列表)，暂时没有图标信息
    ## modelname: 模型的父节点名称
    ## return：返回此父节点下的子节点列表
    """
    res = InitResponseModel()
    space_id = request.user.user_space
    username = request.user.username
    models_obj = session.query(ModelsInformation).filter(ModelsInformation.id == package_id , ModelsInformation.userspace_id.in_([0, space_id]), ModelsInformation.sys_or_user.in_(["sys", username])).first()
    if models_obj and model_name:
        model_child_list = GetModelChild(model_name)
        for i in model_child_list:
            i["image"] = GetIcon(model_name + "." + i["model_name"])
        res.data = model_child_list
    return res


@router.get("/getgraphicsdata", response_model=ResponseModel)
async def GetGraphicsDataView (package_id: str, model_name: str, request: Request, component_name: str = None):
    """
    # 获取模型的画图数据，一次性返回
    ## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
    ## component_name: 模型的组件名称，用于获取单个组件时传入
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]), ModelsInformation.id==package_id).first()
    if not package:
        raise HTTPException(status_code=401, detail="not found")
    GraphicsData = GetGraphicsData()
    if not component_name:
        data = GraphicsData.get_data([model_name])
    else:
        data = GraphicsData.get_one_data([model_name], component_name)
    res.data = data
    return res


@router.get("/getmodelcode", response_model=ResponseModel)
async def GetModelCodeView (package_id: str, model_name: str, request: Request):
    """
    # 获取模型的源码数据，一次性返回
    ## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]),
                                                      ModelsInformation.id == package_id).first()
    if not package:
        raise HTTPException(status_code=401, detail="not found")
    data = GetModelCode(model_name)
    res.data = [data]
    return res


@router.get("/getmodelparameters", response_model=ResponseModel)
async def GetModelParametersView (request: Request, package_id: str, model_name: str, name: str="", components_name: str=""):
    """
    # 获取模型组件的参数数据，一次性返回, 注意，如果是获取整个模型的顶层参数， 只传入模型名称即可， 组件别名和组件名称都不需要传入
    ## model_name: 需要查询的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ## components_name: 需要查询模型的组件名称，全称， 例如“Modelica.Blocks.Continuous.LimPID“
    ## name: 需要查询的组件别名，全称，“PID”
    ## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]),
                                                      ModelsInformation.id == package_id).first()
    if not package:
        raise HTTPException(status_code=401, detail="not found")
    data = GetModelParameters(model_name, name, components_name, package.package_name).get_data()
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
    package_id = item.package_id
    model = session.query(ModelsInformation).filter_by(id=package_id, sys_or_user=username).first()
    if not model:
        res.err = "设置失败"
        res.status = 2
        return res
    data = SetComponentModifierValue(item.model_name, item.parameter_value)
    if data == "Ok":
        res.msg = "设置完成"
    else:
        res.err = "设置失败: 请检查参数是否正确"
        res.status = 1
    return res


@router.get("/getcomponentproperties", response_model=ResponseModel)
async def GetComponentPropertiesView (package_id: str, model_name: str, component_name: str, request: Request):
    """
    # 获取模型组件的属性数据，一次性返回
    ##  class_name: 需要查询属性数据的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ##  component_name: 需要查询的组件别名，全称，“PID”
    ##  sys_user: 需要查询的模型组件是系统还是用户模型
    """
    res = InitResponseModel()
    username = request.user.username
    model = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]),
                                                      ModelsInformation.id == package_id).first()
    if not model:
        res.err = "查询失败"
        res.status = 2
        return res
    result = GetComponents(model_name, component_name)
    if result:
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
    ## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
    ## old_component_name: 需要设置的组件名，全称，“PID”
    ## new_component_name: 需要设置的组件新名称，全称，“PID”
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
        "old_component_name": item.old_component_name,
        "new_component_name": item.new_component_name,
        "final": item.final,
        "protected": item.protected,
        "replaceable": item.replaceable,
        "variabilty": item.variabilty,
        "inner": item.inner,
        "outer": item.outer,
        "causality": item.causality,
    }
    username = request.user.username
    model = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if not model:
        res.err = "设置失败"
        res.status = 2
        return res
    result = SetComponentProperties(parameters_data)
    if result:
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
    space_id = request.user.user_space
    username = request.user.username
    package_name = item.package_name
    package = session.query(ModelsInformation).filter_by(id=item.package_id, sys_or_user=username).first()
    if package:
        model_file_path = package.file_path.split("/")
        model_file_path[-2] = datetime.now().strftime('%Y%m%d%H%M%S%f')
        model_file_path = "/".join(model_file_path)
        package.file_path = model_file_path
        file_path = model_file_path
        filename = None
    elif not item.parent_name:
        model_file_path = "public/UserFiles/UploadFile/" + request.user.username + "/" + str(
            datetime.now().strftime('%Y%m%d%H%M%S%f'))
        file_path = model_file_path + "/" + item.class_name + ".mo"
        filename = item.class_name + ".mo"
    else:
        res.err = "复制失败"
        res.status = 1
        return res
    save_result, msg = SaveClass(item.class_name, item.copied_class_name, item.parent_name, package_name,
                                 new_model_file_path=file_path, file_name=filename)
    if save_result:
        if not item.parent_name:
            model = ModelsInformation(
                    package_name=item.class_name,
                    userspace_id=space_id,
                    model_name=item.class_name,
                    sys_or_user=username,
                    file_path=file_path,
                    )
            session.add(model)
            session.flush()
            session.close()
    else:
        res.err = msg
        res.status = 1
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
    res = InitResponseModel()
    package_name = item.package_name
    package_id = item.package_id
    parent_name = item.parent_name
    class_name = item.class_name
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=package_id, sys_or_user=username).first()
    if not package:
        raise HTTPException(status_code=404, detail="not found")
    if parent_name:
        model_name_all = parent_name + "." + class_name
        if package:
            model_file_path = package.file_path.split("/")
            model_file_path[-2] = datetime.now().strftime('%Y%m%d%H%M%S%f')
            model_file_path = "/".join(model_file_path)
            save_result, msg = SaveClass(class_name=model_name_all, package_name=package_name, copy_or_delete="delete", new_model_file_path=model_file_path)
            if save_result:
                package.file_path = model_file_path
                res.msg = "删除成功"
            else:
                res.err = msg
                res.status = 1
                return res
    else:
        save_result, msg = SaveClass(class_name=class_name, copy_or_delete="delete")
        if save_result:
            res.msg = "删除成功"
        else:
            res.err = msg
            res.status = 1
            return res
        session.query(ModelsInformation).filter_by(id=package_id, sys_or_user=username).delete(synchronize_session=False)
    session.flush()
    session.close()
    return res


@router.post("/get_component_name", response_model=ResponseModel)
async def GetComponentNameView(item: GetComponentNameModel, request: Request):
    """
    # 获取模型当中的模型组件的名字
    ## package_name： 需要创建的组件在哪个包之下，例如在"NN.Examples.Scenario1_Status"模型中创建组件，包就是ENN
    ## package_id： 包id
    ## model_name_all: 需要创建的组件在哪个模型之下，例如在"NN.Examples.Scenario1_Status"模型中创建组件
    ## old_component_name: 被创建成组件的模型名称， 例如"Modelica.Blocks.Math.Abs"
    ## return: 返回json格式数据
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]),
                                                      ModelsInformation.id == item.package_id).first()
    if not package:
        raise HTTPException(status_code=401, detail="not found")
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
    if not package:
        raise HTTPException(status_code=401, detail="not found")
    result, err = AddComponent(item.new_component_name, item.old_component_name, item.model_name_all, item.origin, item.extent, item.rotation)
    if result is True:
        res.msg = "新增组件成功"
    else:
        res.err = err
        res.status = 1
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
        model_name_all = item.delete_list[0]["model_name_all"]
        for i in range(len(item.delete_list)):
            data = item.delete_list[i]
            if data["delete_type"] == "component":
                result = DeleteComponent(data["component_name"], data["model_name_all"])
            elif data["delete_type"] == "connector":
                result = DeleteConnection(data["model_name_all"], data["connect_start"], data["connect_end"])
            else:
                result = False
                break
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
        result = UpdateComponent(item.component_name, item.component_model_name, item.model_name_all, origin, extent, item.rotation)
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
        result = AddConnection(item.model_name_all, item.connect_start, item.connect_end, item.line_points, item.color)
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
        result = DeleteConnection(item.model_name_all, item.connect_start, item.connect_end)
        if result == "Ok":
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
        result = UpdateConnectionAnnotation(item.model_name_all, item.connect_start, item.connect_end, item.line_points, item.color)
        if result is True:
            res.msg = "连接拐点修改成功"
        else:
            res.err = "连接拐点修改失败: " + str(result)
            res.status = 1
    else:
        res.err = "连接修改失败"
        res.status = 2
    return res


@router.get("/exists", response_model=ResponseModel)
async def existsView(package_id: str, model_name: str, request: Request):
    """
    # 检查模型是否存在
    ## package_id： 包id
    ## model_name：模型全称
    ## return: 返回true or false
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=package_id).filter(ModelsInformation.sys_or_user.in_(["sys", username])).first()

    if package:
        res.data.append(True)
    else:
        res.data.append(False)
    return res


@router.get("/checkmodel", response_model=ResponseModel)
async def CheckModelView(package_id: str, model_name: str, request: Request):
    """
    # 检查模型是否合规
    ## package_id： 包id
    ## model_name：模型全称
    ## return: 返回对象数组，包含多种type，不同type表示信息的类型， message表示内容
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=package_id).filter(ModelsInformation.sys_or_user.in_(["sys", username])).first()
    if not package:
        raise HTTPException(status_code=401, detail="not found")
    data_list = CheckModel(model_name)
    for i in data_list:
        r.lpush(username + "_" + "notification", json.dumps(i))
    res.data = [{"message": "模型检查完成"}]
    return res


@router.get("/getcomponents", response_model=ResponseModel)
async def GetComponentsView (request: Request, model_name: str):
    """
    # 获取模型的全部组件数据，一次性返回
    ##  model_name: 需要查询属性数据的模型名称，全称，例如“Modelica.Blocks.Examples.PID_Controller”
    ##  package_id: 所属package的id值，例如“1”
    """
    res = InitResponseModel()
    result = GetComponents(model_name)
    if result:
        for i in result:
            components_data = {
                "component_model_name": i[0],
                "component_name": i[1],
                "component_description": i[2],
                }
            res.data.append(components_data)
    return res


@router.get("/getmodeldocument", response_model=ResponseModel)
async def GetModelDocumentView (request: Request, model_name: str):
    """
    # 获取模型的文档数据
    ##  model_name: 需要查询文档的模型名称，全称，例如“Modelica.Blocks.Examples.PID_Controller”
    ##  package_id: 所属package的id值，例如“1”
    """
    res = InitResponseModel()
    # username = request.user.username
    # package_name = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user.in_(["sys", username]),
    #                                                   ModelsInformation.id == package_id).first()
    # if not package_name:
    #     return HTTPException(status_code=401, detail="not found")
    result = GetModelDocument(model_name)
    res.data.append({
        "document": result,
        "model_name": model_name
        })
    return res


@router.post("/setmodeldocument", response_model=ResponseModel)
async def SetModelDocumentView (request: Request, item: SetModelDocumentModel):
    """
    # 设置模型的文档数据
    ##  model_name: 需要查询文档的模型名称，全称，例如“Modelica.Blocks.Examples.PID_Controller”
    ##  document: 文档内容
    ##  package_id: 所属package的id值，例如“1”
    """
    res = InitResponseModel()
    username = request.user.username
    package_name = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user==username,
                                                      ModelsInformation.id == item.package_id).first()
    if not package_name:
        return HTTPException(status_code=401, detail="not found")
    result = SetModelDocument(item.model_name, item.document)
    if result:
        res.msg = "文档更新成功"
    else:
        res.err = "文档更新失败, 请检查模型名称是否正确"
        res.status = 2
    return res

@router.post("/convertunits", response_model=ResponseModel)
async def ConvertUnitsView (item: ConvertUnitsModel):
    """
    # 转换单位
    ##  s1: 转换后的单位, new单位
    ##  s2: 需要转换的单位， old单位
    ##  return 单位转换后的比值,与原结果数值相乘即可
    """
    res = InitResponseModel()
    result = ConvertUnits(item.s1, item.s2)
    if result[0]:
        res.data.append(float(result[1]))
    else:
        res.data.append(1)
    return res


@router.get("/test")
async def _test (model_name: str, request: Request):
    import time
    start = time.time()
    # res1 = omc.sendExpression(model_name, parsed=False)
    # res2 = omc.sendExpression("getClassNames()")
    from app.service.save_class_names import SaveClassNames
    a = SaveClassNames(space_id=0, mo_path="/home/simtek/code/omlibrary/WindPowerSystem/package.mo", init_name="WindPowerSystem",sys_or_user="sys")
    # a = SaveClassNames(space_id=0, mo_path="/home/simtek/code/omlibrary/SolarPower/package.mo", init_name="SolarPower",sys_or_user="sys")
    # a = SaveClassNames(space_id=0, mo_path="/home/simtek/code/omlibrary/Buildings 8.0.0/package.mo", init_name="Buildings",sys_or_user="sys")
    # a = SaveClassNames(space_id=0, mo_path="/home/simtek/code/omlibrary/Modelica 3.2.3/package.mo", init_name="Modelica",sys_or_user="sys")
    return {
            # "msg": [res1,res2],
            "msg": [a],
            "user": request.user.display_name,
            "auth": request.user.is_authenticated
        }
