# -- coding: utf-8 --
from config.omc import omc
# from app.service.load_model_file import LoadModelFile


def GetSimulationOptions(model_name):
    data = omc.sendExpression("getSimulationOptions(" + model_name + ")")
    data_dict = {
        "startTime": data[0],
        "stopTime": data[1],
        "tolerance": data[2],
        "numberOfIntervals": data[3],
        "interval": data[4]
    }
    return data_dict
