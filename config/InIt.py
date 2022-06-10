# -- coding: utf-8 --
import time

from config.DB_config import DBSession
from app.model.User.User import UserSpace
from app.model.ModelsPackage.ModelsInformation import ModelsInformation
from config.settings import USERNAME
from app.service.clear_program import ClearProgram
from app.service.load_model_file import LoadModel
session = DBSession()

def InIt():
    space_obj = session.query(UserSpace).filter(UserSpace.username==USERNAME).order_by(UserSpace.last_login_time.desc()).first()
    if space_obj:
        models_obj_list = session.query(ModelsInformation).filter(ModelsInformation.userspace_id == space_obj.id,
                                                                  ModelsInformation.sys_or_user == USERNAME).all()
        ClearProgram()
        for i in models_obj_list:
            path = i.file_path
            LoadModel(path=path, check=False)

