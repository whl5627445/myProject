# -- coding: utf-8 --
import datetime

from fastapi import Request, HTTPException
from router.user_router import router
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.ModelsPackage.ModelsInformation import ModelsInformation
from app.model.User.User import UserSpace
from app.BaseModel.user import UserSpaceModel
from sqlalchemy import or_, and_
from app.service.load_model_file import LoadModel
import logging
from config.DB_config import DBSession
from config.settings import EXAMPLES
from app.service.clear_program import ClearProgram
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
    space_filter = session.query(UserSpace)
    space__filter = space_filter.filter(UserSpace.username == username).all()
    space_name_filter = space_filter.filter(UserSpace.spacename == space_name, UserSpace.username == username).first()
    if space_name_filter or len(space__filter) >= 5:
        res.err = "空间名称已存在或数量超过5个"
        res.status = 2
        return res
    space_obj = UserSpace(username=username, spacename=space_name)
    session.add(space_obj)
    session.flush()
    res.data.append({"id": space_obj.id, "name": space_obj.spacename})
    res.msg = "创建成功"
    session.close()
    return res


@router.post("/deleteuserspace", response_model=ResponseModel)
async def DeleteUserSpaceView (item: UserSpaceModel, request: Request):
    """
    # 删除用户空间
    # space_id: 用户空间id
    """
    res = InitResponseModel()
    username = request.user.username
    space_id = item.space_id
    space_filter = session.query(UserSpace).filter(UserSpace.id==space_id, UserSpace.username==username).first()
    if not space_filter:
        raise HTTPException(status_code=401, detail="not found")
    session.query(UserSpace).filter(UserSpace.id == space_id).delete(synchronize_session=False)
    session.query(ModelsInformation).filter(ModelsInformation.userspace_id == space_id).delete(synchronize_session=False)
    session.flush()
    res.msg = "删除成功"
    session.close()
    return res



@router.post("/loginuserspace", response_model=ResponseModel)
async def LoginUserSpaceView (item: UserSpaceModel, request: Request):
    """
    # 进入用户空间
    ## space_id: 空间id，在getuserspace接口获取
    """
    res = InitResponseModel()
    username = request.user.username
    space_obj = session.query(UserSpace).filter(UserSpace.id == item.space_id, UserSpace.username==username).first()
    if not space_obj:
        raise HTTPException(status_code=401, detail="not found")
    models_obj_list = session.query(ModelsInformation).filter(ModelsInformation.userspace_id ==item.space_id, ModelsInformation.sys_or_user == username).all()
    models_list = []
    load_res = True
    ClearProgram()
    for i in models_obj_list:
        logging.info("加载模型库：{0}".format(i.package_name))
        path = i.file_path
        load_res = LoadModel(path=path, check=False)
    if load_res:
        space_obj.last_login_time = datetime.datetime.now()
        session.flush()
        session.close()
        res.data.append({"models": models_list, "examples": EXAMPLES})
        res.msg = "初始化完成"
    else:
        res.status = 2
        res.msg = "初始化失败"
    return res
