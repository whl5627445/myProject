# -- coding: utf-8 --
from fastapi import File, UploadFile, Request, HTTPException
from router.upload_file_router import router
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.ModelsPackage.ModelsInformation import ModelsInformation, ModelsInformationAll
from app.service.save_class_names import SaveClassNames
from library.file_operation import FileOperation
from config.DB_config import DBSession
from datetime import datetime
from app.BaseModel.uploadfile import UploadSaveFileModel, UploadSaveModelModel
from app.service.get_model_code import GetModelCode, GetModelPath
from app.service.create_modelica_class import CreateModelicaClass, UpdateModelicaClass
import os
import urllib.parse
session = DBSession()


@router.post("/uploadfile", response_model=ResponseModel)
async def UploadFileView(request: Request, file: UploadFile = File(...)):
    """
    # 用户上传mo文件接口
    ## file: 文件数据，bytes形式的文件流
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    file_data = await file.read()
    filename = file.filename
    file_path = "public/UserFiles/UploadFile/" + request.user.username + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f'))
    fo = FileOperation()
    fo.write_file(file_path, filename, file_data)
    save_result_list = []
    if filename.endswith(".mo"):
        package_name = file.filename.removesuffix(".mo")
        mo_path = file_path + "/" + package_name + ".mo"
        UP = session.query(ModelsInformation).filter_by(package_name=package_name,
                                                        sys_or_user=request.user.username).first()
        if UP:
            res.err = "文件已存在！"
            res.status = 2
            return res
        save_result, M_id = SaveClassNames(mo_path=mo_path, init_name=package_name, sys_or_user=request.user.username)
        save_result_list.append({
                "filename": filename,
                "result": save_result,
            })
    elif filename.endswith(".rar") or filename.endswith(".zip") or filename.endswith(".7z"):
        un_file_res, err = fo.un_file(file_path + "/" + filename, file_path)
        if not un_file_res:
            res.err = err
            res.status = 1
            return res
        for i in un_file_res:
            UP = session.query(ModelsInformation).filter_by(package_name=i["package_name"],
                                                            sys_or_user=request.user.username).first()
            if UP:
                os.remove(i["file_path"])
                res.err = i["package_name"] + "， 已存在相同名字的包！"
                res.status = 2
                return res
            save_result, M_id = SaveClassNames(mo_path=i["file_path"], init_name=i["package_name"],
                                               sys_or_user=request.user.username)
            if save_result:
                save_result_list.append({
                    "filename": filename,
                    "result": save_result,
                    })

    else:
        res.err = "文件格式不正确, 请上传以.mo为后缀的模型文件，或者是rar、zip、7z三种格式的压缩文件"
        res.status = 2
        return res

    if save_result_list:
        res.msg = "模型上传成功！"
        res.data = save_result_list
    else:
        res.status = 1
        res.err = "模型加载失败，请重新检查后上传"
    return res


@router.post("/uploadfile/savefile", response_model=ResponseModel)
async def SaveFileView(request: Request):
    """
    # 用户保存mo文件接口
    ## model_str: 文件数据，字符串形式, 如果是新建模型文件则为空字符串
    ## package_name: 包名字
    ## package_id: 包的id，如果是新建模型文件则为空字符串
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    item = await request.json()
    package_name_list = item.get("package_name", "").split(".")
    parent_name = None
    package_name = package_name_list[0]  # 更新模型的话，是模型包的名字 item里的是模型全名
    if len(package_name_list) > 1:
        parent_name = ".".join(package_name_list[:-1])
    model_str = urllib.parse.unquote(item.get("model_str", ""))  # html和JavaScript标签被过滤，无奈选择转码后解码
    package_id = item.get("package_id", None)
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(sys_or_user=username, id=package_id).first()
    if not package:
        raise HTTPException(status_code=404, detail="not found")
    mo_path = package.file_path
    item_package_name = item.get("package_name", "")
    model_path = GetModelPath(item_package_name)
    if parent_name:
        model_str = "within " + parent_name + ";" + model_str
    result = UpdateModelicaClass(model_str, path=package_name)
    if result is True:
        # model_path = GetModelPath(item.package_name)
        if package.file_path.endswith("package.mo"):
            file_model_str = GetModelCode(item_package_name)
        else:
            file_model_str = GetModelCode(package_name)
        FileOperation().write(model_path, file_model_str)
        save_result, M_id = SaveClassNames(mo_path=mo_path, init_name=package_name, sys_or_user=request.user.username,
                                           package_id=package_id)
        res_model_str = GetModelCode(item_package_name)
        res.data = [{"model_str": res_model_str, "id": M_id}]
        res.msg = "保存文件成功！"
    else:
        res.status = 1
        res.err = "模型加载失败，请重新检查后上传"
    return res


@router.post("/uploadfile/savemodel", response_model=ResponseModel)
async def SaveModelView(request: Request, item: UploadSaveModelModel):
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


