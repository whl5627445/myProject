# -- coding: utf-8 --
import logging

from config.omc import omc


def GetComponents(class_name, component_name=None):
    Components_data = omc.getComponents(class_name)
    component_data = []
    if Components_data != "Error":
        for i in Components_data:
            if i[8] == "parameter" or  i[3] == "protected" or i[4] == "True":
                continue
            if component_name and i[1] == component_name:
                component_data = i
                break
            else:
                component_data.append(i)
    return component_data



