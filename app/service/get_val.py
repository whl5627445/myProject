# -- coding: utf-8 --

from config.omc import omc


def GetVal(variable_name:str, time_point: str ="0" , file_name: str=""):
    data = omc.val(variable_name, time_point, file_name)
    return data
