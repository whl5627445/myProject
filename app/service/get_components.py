# -- coding: utf-8 --
from config.omc import omc


def GetComponents(class_name, component_name=None):
    Components_data = omc.getComponents(class_name)
    if component_name and Components_data != "Error":
        component_data = []
        for i in Components_data:
            if i[1] == component_name:
                component_data = i
                break
        return component_data
    else:
        if Components_data != "Error":
            return Components_data



