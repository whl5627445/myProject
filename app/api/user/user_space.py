# -- coding: utf-8 --
from fastapi import Request, HTTPException
from router.user_router import router
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.ModelsPackage.ModelsInformation import UserSpace, ModelsInformation
from sqlalchemy import or_, and_
from app.service.load_model_file import LoadModel
import logging
from config.DB_config import DBSession
session = DBSession()


@router.get("/getuserspace", response_model=ResponseModel)
async def GetUserSpaceView (request: Request):
    """
    # 获取用户所有的用户空间条目
    ## return: 用户空间列表
    """
    res = InitResponseModel()
    username = request.user.username
    space_all = session.query(UserSpace).filter(UserSpace.username == username).all()
    space_default = {
        "id": 0,
        "name": "默认空间",
        }
    res.data.append(space_default)
    if space_all:
        for space in space_all:
            space_dict = {
                "id": space.id,
                "name": space.spacename,
            }
            res.data.append(space_dict)
    return res


@router.post("/createuserspace", response_model=ResponseModel)
async def CreateUserSpaceView (space_name: str, request: Request):
    """
    # 创建用户空间
    # space_name: 用户空间名称
    """
    res = InitResponseModel()
    username = request.user.username
    space_filter = session.query(UserSpace).filter(UserSpace.spacename == space_name, UserSpace.username == username).first()
    if space_filter:
        res.msg = "空间名已存在"
        res.status = 2
        return res
    space_obj = UserSpace(username=username, spacename=space_name)
    session.add(space_obj)
    session.flush()
    res.data.append({"id": space_obj.id, "name": space_obj.spacename})

    session.close()
    return res


@router.post("/loginuserspace", response_model=ResponseModel)
async def LoginUserSpaceView (space_id: str, request: Request):
    """
    # 进入用户空间
    ## space_id: 空间id，在getuserspace接口获取
    """
    res = InitResponseModel()
    username = request.user.username
    logging.debug(request.user.__dict__)
    space_obj = session.query(UserSpace).filter(UserSpace.id == space_id, UserSpace.username==username).first()
    if not space_obj:
        raise HTTPException(status_code=401, detail="not found")
    models_obj_list = session.query(ModelsInformation).filter(or_(and_(ModelsInformation.userspace_id ==space_id, ModelsInformation.sys_or_user == username), ModelsInformation.sys_or_user == "sys")).all()
    models_list = [i.package_name for i in models_obj_list]
    load_res = LoadModel(models_obj_list)
    if load_res:
        res.data.append({"models": models_list})
        res.msg = "初始化完成"
    else:
        res.status = 2
        res.msg = "初始化失败"
    return res
