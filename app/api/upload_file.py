import time
from fastapi import File, UploadFile, Request
from router.upload_file import router
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.models_package.ModelsInformation import ModelsInformation, ModelsInformationAll
from app.service.save_class_names import SaveClassNames
from library.file_operation import FileOperation
from config.DB_config import DBSession
session = DBSession()


@router.post("/", response_model=ResponseModel)
async def UploadFile(request: Request, file: UploadFile = File(...)):
    """
    # 用户上传mo文件接口
    ## file: 文件数据，bytes形式的文件流
    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    file_data = await file.read()
    file_path = "public/UserFiles/UploadFile/" + request.user["username"] + "/" + str(time.time())
    file_name = file.filename.removesuffix(".mo")
    UP = session.query(ModelsInformation).filter_by(package_name=file_name).first()
    if UP:
        res.err = "文件已存在！"
        res.status = 2
        return res
    fo = FileOperation()
    if not file.filename.endswith(".mo"):
        res.err = "文件格式不正确, 请上传以.mo为后缀的模型文件"
        res.status = 2
        return res
    fo.write_file(file_path, file.filename, file_data)
    save_result = SaveClassNames(mo_path=file_path + "/" + file.filename, init_name=file_name, sys_or_user=request.user["username"])
    if save_result:
        res.msg = "模型上传成功！"
    else:
        res.status = 1
        res.err = "模型加载失败，请重新检查后上传"
    return res


@router.delete("/delete_package", response_model=ResponseModel)
async def UploadFile(request: Request, package_name: str):
    session.query(ModelsInformation).filter_by(package_name=package_name, sys_or_user=request.user["username"]).delete(synchronize_session=False)
    session.query(ModelsInformationAll).filter_by(package_name=package_name, sys_or_user=request.user["username"]).delete(synchronize_session=False)
    res = InitResponseModel()
    session.flush()
    session.close()
    res.msg = "删除成功"
    return res
