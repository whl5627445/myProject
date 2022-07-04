# -- coding: utf-8 --
import time
from config.omc import omc
from library.file_operation import FileOperation


def GetModelCode(model_name):
    data = omc.list(model_name)
    data = data.replace('\\"', '\"')
    data = data.replace('\\"', '\"')
    data = data[1:-2]
    return data


def GetModelPath(model_name):
    data = omc.getSourceFile(model_name)
    return data


def SaveModelCode(file_path, package_name):
    code = GetModelCode(package_name)
    path = file_path
    FileOperation.write(path, code)
    return
