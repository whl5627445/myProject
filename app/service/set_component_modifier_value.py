# -- coding: utf-8 --
from config.omc import omc
# from app.service.load_model_file import LoadModelFile


def SetComponentModifierValue(className, parameter_value):
    result = "Ok"
    for k, v in parameter_value.items():
        result = omc.setComponentModifierValue(className, k, v)
        if result != "Ok":
            break
    return result
