# -- coding: utf-8 --
import logging

from fastapi import File, UploadFile, Request, HTTPException, Form
from router.upload_file_router import router
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.ModelsPackage.ModelsInformation import ModelsInformation
from app.service.save_class_names import SaveClassNames
from library.file_operation import FileOperation
from config.DB_config import DBSession
from config.redis_config import r
from datetime import datetime
from app.BaseModel.uploadfile import UploadSaveModelModel
from app.service.get_model_code import GetModelCode, GetModelPath
from app.service.create_modelica_class import CreateModelicaClass, UpdateModelicaClass
from app.service.icon_operation import UploadIcon
from library.HW_OBS_operation import HWOBS
import os, re, json, urllib.parse, base64
session = DBSession()


@router.post("/", response_model=ResponseModel)
async def UploadFileView(request: Request, file: UploadFile = File(...), package_id : str = Form(...),):
    """
    # 用户上传mo文件接口
    ## file: 文件数据，bytes形式的文件流
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    space_id = request.user.user_space
    file_data = await file.read()
    filename = file.filename
    username = request.user.username
    file_path = "public/UserFiles/UploadFile/" + username + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f'))
    fo = FileOperation()
    fo.write_file(file_path, filename, file_data)
    save_result_list = []
    notice_data = []
    if filename.endswith(".mo"):
        package_name = file.filename.removesuffix(".mo")
        mo_path = file_path + "/" + package_name + ".mo"
        UP = session.query(ModelsInformation).filter_by(userspace_id=space_id, sys_or_user=request.user.username, package_name=package_name).first()
        if UP:
            res.err = "文件已存在！"
            res.status = 2
            return res
        save_result, M_id, notice_data = SaveClassNames(space_id=space_id, mo_path=mo_path, init_name=package_name, sys_or_user=request.user.username)
        if save_result:
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
            UP = session.query(ModelsInformation).filter_by(userspace_id=space_id, package_name=i["package_name"],
                                                            sys_or_user=request.user.username).first()
            if UP:
                os.remove(i["file_path"])
                res.err = i["package_name"] + "， 已存在相同名字的包！"
                res.status = 2
                return res
            save_result, M_id, notice_data = SaveClassNames(space_id=space_id, mo_path=i["file_path"], init_name=i["package_name"],
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
        res.err = "模型上传失败！"
    for i in notice_data:
        r_data = {
            "message": i["message"],
            "type": i["type"],
            }
        r.lpush(username + "_" + "notification", json.dumps(r_data))
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
    space_id = request.user.user_space
    item = await request.json()
    package_name_list = item.get("package_name", "").split(".")
    parent_name = None
    package_name = package_name_list[0]  # 更新模型的话，是模型包的名字 item里的是模型全名
    if len(package_name_list) > 1:
        parent_name = ".".join(package_name_list[:-1])
    model_str = urllib.parse.unquote(item.get("model_str", ""))  # html和JavaScript标签被过滤，选择转码后解码
    package_id = item.get("package_id", None)
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(userspace_id=space_id, sys_or_user=request.user.username, id=package_id).first()
    if not package:
        raise HTTPException(status_code=404, detail="not found")
    item_package_name = item.get("package_name", "")
    model_path = GetModelPath(item_package_name)
    if parent_name:
        model_str = "within " + parent_name + ";" + model_str
    result, notice_data = UpdateModelicaClass(model_str, path=package_name)
    if result is True:
        if package.file_path.endswith("package.mo"):
            file_model_str = GetModelCode(item_package_name)
        else:
            file_model_str = GetModelCode(package_name)
        FileOperation().write(model_path, file_model_str)
        save_result, M_id, notice_data = SaveClassNames(space_id=space_id, mo_path=model_path, init_name=package_name, sys_or_user=request.user.username,
                                           package_id=package_id)
        if save_result:
            res_model_str = GetModelCode(item_package_name)
            res.data = [{"model_str": res_model_str, "id": M_id}]
            res.msg = "保存文件成功！"
        else:
            res.err = "语法错误，请重新检查后上传"
            res.status = 1
    else:
        res.err = "语法错误，请重新检查后上传"
        res.status = 1
    for i in notice_data:
        r_data = {
            "message": i["message"],
            "type": i["type"],
            }
        r.lpush(username + "_" + "notification", json.dumps(r_data))
    return res


@router.post("/uploadfile/savemodel", response_model=ResponseModel)
async def SaveModelView(request: Request, item: UploadSaveModelModel):
    """
    # 用户创建和保存mo文件接口
    ## package_name: 包或模型的名字
    ## type: 要创建的类型
    ## model_str: str = 模型源码字符串
    ## package_id: 包的id
    ## vars: {
    ##     "expand": "", 扩展
    ##     "comment": "", 注释
    ##     "insert_to": "", 父节点， 要插入哪个节点下
    ##     "partial": False,  部分的
    ##     "encapsulated": False, 封装
    ##     "state": False 状态
    ##     }
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    notice_data = []
    space_id = request.user.user_space
    create_package_name = item.package_name
    if not re.match("^[a-zA-Z_]", create_package_name):
        res.err = "名称请以子母和下划线开头"
        res.status = 1
        return res
    str_type = item.str_type
    model_str = item.model_str
    package_id = item.package_id
    var = item.vars
    username = request.user.username
    package = session.query(ModelsInformation).filter(
            ModelsInformation.package_name == create_package_name, ModelsInformation.userspace_id==space_id).filter_by(
        sys_or_user=username).first()
    if var["insert_to"]:
        insert_package_name = var["insert_to"].split(".")[0]
        package = session.query(ModelsInformation).filter(
                ModelsInformation.id == package_id).filter_by(
                sys_or_user=username).first()
        create_package_name_all = var["insert_to"] + "." + create_package_name
        init_name = insert_package_name
        file_name = insert_package_name + ".mo"
        path = package.file_path
    else:
        path = create_package_name
        create_package_name_all = create_package_name
        init_name = create_package_name_all
        file_name = create_package_name_all + ".mo"
    if not var["insert_to"] and package and str_type:
        res.err = "名称已存在！"
        res.status = 2
        return res
    file_path = "public/UserFiles/UploadFile/" + username + "/" + str(datetime.now().strftime('%Y%m%d%H%M%S%f'))
    if not model_str:
        result, notice_data = CreateModelicaClass(create_package_name, str_type, var, create_package_name_all, path, item.comment)
    else:
        result, notice_data = UpdateModelicaClass(model_str, path)
    res_package_str = GetModelCode(init_name)
    FileOperation().write_file(file_path, file_name, res_package_str)
    if result:
        res_model_str = GetModelCode(create_package_name_all)
        save_result, M_id, notice_data = SaveClassNames(space_id=space_id, mo_path=file_path + "/" + file_name, init_name=init_name, sys_or_user=request.user.username, package_id=package_id)
        if save_result:
            res.data = [{"model_str": res_model_str, "id": M_id}]
            res.msg = "successful！"
        else:
            res.status = 1
            res.err = "语法错误，请重新检查"
    else:
        res.status = 1
        res.err = "语法错误，请重新检查"
    for i in notice_data:
        r_data = {
            "message": i["message"],
            "type": i["type"],
            }
        r.lpush(username + "_" + "notification", json.dumps(r_data))
    return res


@router.post("/uploadicon", response_model=ResponseModel)
async def UploadIconView(request: Request, model_name: str = Form(...), package_id: str = Form(...), file: UploadFile = File(...)):
    """
    # 用户上传模型图标接口
    ## file: 文件数据，bytes形式的文件流
    ## package_name: 包名称
    ## model_name: 模型名称
    ## package_id: 包id
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    username = request.user.username
    model = session.query(ModelsInformation).filter(ModelsInformation.id == package_id, ModelsInformation.sys_or_user==username).first()
    if not model:
        res.err = "暂时只能更换用户区域模型图标"
    suffix_list = ["jpg", "png", "jpeg", "svg"]
    file_suffix = file.filename.split(".")[-1]
    if file_suffix not in suffix_list:
        res.err = "暂只支持jpg,png,jpeg,svg格式的图片"
        res.status = 1
        return res
    file_data = await file.read()
    encoded_string = base64.b64encode(file_data)
    result = UploadIcon(model_name, encoded_string.decode())
    if result:
        res.msg = "图标上传成功"
    else:
        res.err = "上传失败，请重新上传"
        res.status = 2
    return res
