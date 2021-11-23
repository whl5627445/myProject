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
    # 用户获取mo文件信息接口， 可以根据url进行下载
    ## return: 包名， 上传时间， 下载路径
    """
    res = InitResponseModel()
    username = request.user.username
    package_list = session.query(ModelsInformation).filter_by(sys_or_user=username).all()
    for i in range(len(package_list)):
        data_dict = {
            "id": i,
            "package_name": package_list[i].package_name,
            "create_time": package_list[i].create_time,
            "url": "static/" + package_list[i].file_path,
            }
        res.data.append(data_dict)
    return res

