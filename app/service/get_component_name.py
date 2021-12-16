# -- coding: utf-8 --
from config.omc import omc
from app.service.load_model_file import LoadModelFile
import re


def GetComponentName(class_name, component_name_all, file_path=None, package_name=None):
    # if file_path:
    #     LoadModelFile(package_name, file_path)
    data = omc.getComponents(class_name)
    component_name = component_name_all.split(".")[-1].lower()
    name_num = 0
    for i in data:
        expression = "^" + component_name + "\d*$"
        match_result = re.match(expression, i[1], flags=0)
        if match_result:
            name_num += 1
    if name_num == 0:
        return component_name.lower()
    else:
        return component_name.lower() + str(name_num)
