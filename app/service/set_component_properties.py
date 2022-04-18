# -- coding: utf-8 --
import logging

from config.omc import omc
# from app.service.load_model_file import LoadModelFile
from config.modelica_config import modelica_keywords


def RenameComponentInClass(class_name, old_component_name, new_component_name):
    data = omc.getElements(class_name)
    if data:
        for name in data:
            if name[3] == new_component_name or new_component_name in modelica_keywords:
                return False
    rename_result = omc.renameComponentInClass(class_name, old_component_name, new_component_name) # 重命名组件
    return rename_result

def SetComponentProperties(parameters_data):

    class_name = parameters_data.get('class_name')
    new_component_name = parameters_data.get('new_component_name')
    old_component_name = parameters_data.get('old_component_name')
    final = parameters_data.get('final')
    protected = parameters_data.get('protected')
    replaceable = parameters_data.get('replaceable')
    variabilty = parameters_data.get('variabilty')
    inner = parameters_data.get('inner')
    outer = parameters_data.get('outer')
    causality = parameters_data.get('causality')
    logging.debug("Set_class_name: %s" % class_name)
    rename_result = RenameComponentInClass(class_name, old_component_name, new_component_name)
    logging.debug("rename_result: %s" % rename_result)
    SCP_result = omc.setComponentProperties(class_name, new_component_name, final, protected, replaceable, variabilty, inner, outer, causality) #设置组件属性
    logging.debug("SetComponentProperties: %s" % SCP_result)
    if SCP_result == "Ok" and rename_result:
        return True
    else:
        return False
