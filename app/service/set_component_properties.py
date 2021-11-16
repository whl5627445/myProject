# -- coding: utf-8 --
from config.omc import omc
from app.service.load_model_file import LoadModelFile
from library.file_operation import FileOperation
from app.service.get_model_code import GetModelCode

def SetComponentProperties(file_path, package_name, **kwargs):
    if file_path:
        LoadModelFile(package_name, file_path)
    result = omc.setComponentProperties(**kwargs)
    return result
