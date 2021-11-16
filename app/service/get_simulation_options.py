# -- coding: utf-8 --
from config.omc import omc


def GetSimulationOptions(model_name_list, model_file_path=None):
    data_list = []
    if model_file_path:
        omc.loadFile(model_file_path)
    for model_name in model_name_list:
        data = omc.sendExpression("getSimulationOptions(" + model_name + ")")
        data_dict = {
            "startTime": data[0],
            "stopTime": data[1],
            "tolerance": data[2],
            "numberOfIntervals": data[3],
            "interval": data[4]
        }
        data_list.append(data_dict)
    return data_list
