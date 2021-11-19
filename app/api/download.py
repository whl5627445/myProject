# -- coding: utf-8 --
from router.download_router import router
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.models_package.ModelsInformation import ModelsInformation
from fastapi import File, UploadFile, Request
from config.DB_config import DBSession

session = DBSession()


@router.get("/getfilelist", response_model=ResponseModel)
async def GetFileListView(request: Request):
    """
    # 用户下载mo文件接口
    ## package_name: 包或模型的名字

    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    username = request.user.username
    package_list = session.query(ModelsInformation).filter_by(sys_or_user=username).all()
    for i in package_list:
        data_dict = {
            "package_id": i.id,
            "package_name": i.package_name,
            "create_time": i.create_time,
            "url": "static/" + i.file_path,
            }
        res.data.append(data_dict)
    return res



@router.get("/downloadfile", response_model=ResponseModel)
async def DownloadFileView(request: Request, package_id: str):
    """
    # 用户下载mo文件接口
    ## package_name: 包或模型的名字

    ## return: 会返回文件上传的状态
    """
    res = InitResponseModel()
    username = request.user.username
    package = session.query(ModelsInformation).filter_by(id=package_id, sys_or_user=username).first()
    path = package.file_path
    data = [path]
    res.data = data
    return res
