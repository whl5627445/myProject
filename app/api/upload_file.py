# -- coding: utf-8 --
from fastapi import File, UploadFile, Request
from router.upload_file_router import router
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.models_package.ModelsInformation import ModelsInformation, ModelsInformationAll
from app.service.save_class_names import SaveClassNames
from library.file_operation import FileOperation
from config.DB_config import DBSession
from datetime import datetime
from app.BaseModel.uploadfile import UploadSaveFileModel, UploadSaveModelModel
from app.service.get_model_code import GetModelCode
from app.service.create_modelica_class import CreateModelicaClass, UpdateModelicaClass
session = DBSession()


@router.post("/uploadfile", response_model=ResponseModel)
async def UploadFile(request: Request, file: UploadFile = File(...)):
    """
    # 用户上传mo文件接口
    ## file: 文件数据，bytes形式的文件流
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    file_data = await file.read()
    file_name = file.filename.removesuffix(".mo")
    UP = session.query(ModelsInformation).filter_by(package_name=file_name, sys_or_user=request.user.username).first()
    if UP:
        res.err = "文件已存在！"
        res.status = 2
        return res
    if not file.filename.endswith(".mo"):
        res.err = "文件格式不正确, 请上传以.mo为后缀的模型文件"
        res.status = 2
        return res
    file_path = "public/UserFiles/UploadFile/" + request.user.username + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f'))
    FileOperation().write_file(file_path, file.filename, file_data)
    save_result = SaveClassNames(mo_path=file_path + "/" + file.filename, init_name=file_name, sys_or_user=request.user.username)
    if save_result:
        res.msg = "模型上传成功！"
    else:
        res.status = 1
        res.err = "模型加载失败，请重新检查后上传"
    return res


@router.post("/uploadfile/savefile", response_model=ResponseModel)
async def SaveFile(request: Request, item: UploadSaveFileModel):
    """
    # 用户保存mo文件接口
    ## model_str: 文件数据，字符串形式, 如果是新建模型文件则为空字符串
    ## package_name: 包名字
    ## package_id: 包的id，如果是新建模型文件则为空字符串
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    package_name_list = item.package_name.split(".")
    parent_name = None
    package_name = package_name_list[0]
    if len(package_name_list) > 1:
        parent_name = ".".join(package_name_list[:-1])
    model_str = item.model_str
    package_id = item.package_id
    username = request.user.username
    file_path = "public/UserFiles/UploadFile/" + username + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f'))
    file_name = package_name + ".mo"
    mo_path = file_path + "/" + file_name
    if parent_name:
        model_str = "within " + parent_name + ";" + model_str
    result = UpdateModelicaClass(model_str, path=package_name)
    file_model_str = GetModelCode(package_name)
    FileOperation().write_file(file_path, file_name, file_model_str)
    if result:
        save_result, M_id = SaveClassNames(mo_path=mo_path, init_name=package_name, sys_or_user=request.user.username,
                                           package_id=package_id)
        res_model_str = GetModelCode(item.package_name)
        res.data = [{"model_str": res_model_str, "id": M_id}]
        res.msg = "保存文件成功！"
    else:
        res.status = 1
        res.err = "模型加载失败，请重新检查后上传"
    return res


@router.post("/uploadfile/savemodel", response_model=ResponseModel)
async def SaveModel(request: Request, item: UploadSaveModelModel):
    """
    # 用户创建和保存mo文件接口
    ## package_name: 包或模型的名字
    ## type: 要创建的类型
    ## model_str: str = 模型源码字符串
    ## package_id: 包的id
    ## var: {
    ##     "expand": "", 扩展
    ##     "insert_to": "", 父节点， 要插入哪个节点下
    ##     "partial": False,  部分的
    ##     "encapsulated": False, 封装
    ##     "state": False 状态
    ##     }
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    create_package_name = item.package_name
    str_type = item.str_type
    model_str = item.model_str
    package_id = item.package_id
    var = item.vars
    username = request.user.username
    package = session.query(ModelsInformation).filter(
            ModelsInformation.package_name == create_package_name).filter_by(
        sys_or_user=username).first()
    if var["insert_to"]:
        insert_package_name = var["insert_to"].split(".")[0]
        package = session.query(ModelsInformation).filter(
                ModelsInformation.id == package_id).filter_by(
                sys_or_user=username).first()
        create_package_name_all = var["insert_to"] + "." + create_package_name
        init_name = insert_package_name
        file_name = insert_package_name + ".mo"
        model_obj = session.query(ModelsInformationAll).filter_by(model_name=create_package_name, package_id=package_id,
                                                                  sys_or_user=username,
                                                                  parent_name=var["insert_to"]).first()
        path = package.file_path
    else:
        path = create_package_name
        create_package_name_all = create_package_name
        init_name = create_package_name_all
        file_name = create_package_name_all + ".mo"
        model_obj = None
    if (not var["insert_to"] and package and str_type) or (var["insert_to"] and model_obj):
        res.err = "名称已存在！"
        res.status = 2
        return res
    file_path = "public/UserFiles/UploadFile/" + username + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f'))
    if not model_str:
        result = CreateModelicaClass(create_package_name, str_type, var, create_package_name_all, path)
    else:
        result = UpdateModelicaClass(model_str, path)
    res_package_str = GetModelCode(init_name)
    FileOperation().write_file(file_path, file_name, res_package_str)
    if result:
        res_model_str = GetModelCode(create_package_name_all)
        save_result, M_id = SaveClassNames(mo_path=file_path + "/" + file_name, init_name=init_name, sys_or_user=request.user.username, package_id=package_id)
        if save_result:
            res.data = [{"model_str": res_model_str, "id": M_id}]
            res.msg = "successful！"
        else:
            res.status = 1
            res.err = "模型加载失败，请重新检查"
    else:
        res.status = 1
        res.err = "模型加载失败，请重新检查"
    return res


