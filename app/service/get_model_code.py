# -- coding: utf-8 --
from config.omc import omc
# from app.service.load_model_file import LoadModelFile


def GetModelCode(model_name):
    data = omc.list(model_name)
    data = data.replace('\\"', '\"')
    data = data.replace('\\"', '\"')
    data = data[1:-2]
    return data


def GetModelPath(model_name):
    data = omc.getSourceFile(model_name)
    return data
