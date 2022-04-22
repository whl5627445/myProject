# -- coding: utf-8 --
from config.omc import omc
import re
from config.settings import modelica_keywords

def GetComponentName(class_name, component_name_all):
    data = omc.getComponents(class_name)
    component_name = component_name_all.split(".")[-1].lower()
    name_num = 0
    name_dict = {}

    if component_name in modelica_keywords:
        name_num += 1
    for i in data:
        expression = "^" + component_name + "\d*$"
        if len(i) >= 2:
            match_result = re.match(expression, i[1], flags=0)
            if match_result:
                name_dict[i[1]] = True
                name_num += 1
    if name_num == 0:
        return component_name
    else:
        for num in range(1, name_num+1):
            name = component_name + str(num)
            if name not in name_dict:
                component_name = name
                break
        return component_name
