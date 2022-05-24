# -- coding: utf-8 --
import logging

from config.omc import omc


def GetComponents(class_name, component_name=None):
    Components_data = omc.getComponents(class_name)
    component_data = []
    if Components_data != "Error":
        try:
            for i in Components_data:
                if component_name and i[1] == component_name:
                    component_data = i
                    break
                elif not (i[8] == "parameter" or  i[3] == "protected" or i[4] == "True") and not component_name:
                    component_data.append(i)
        except Exception as e:
            logging.error(e)
    return component_data



