# -- coding: utf-8 --
import logging

from config.omc import omc
# from app.service.load_model_file import LoadModelFile


def ComponentsVerification(class_name, component_name, model_name_all):
    class_information = omc.getClassInformation(class_name)
    if class_information:
        class_type = class_information[0]
        if class_type not in ["model", "class", "connector", "block"]:
            return False, "不能插入：" + class_name + ", 这是一个 \"" + class_type + " \"类型。组件视图层只允许有model、class、connector或者block。"
    else:
        return False, ""
    data = omc.getElements(model_name_all)
    if data != "Error":
        for i in data:
            if component_name in i:
                return False, "组件名称已存在"
    else:
        return False, ""
    return True, ""


def AddComponent(new_component_name, old_component_name, model_name_all, origin, extent, rotation):
    v, err = ComponentsVerification(old_component_name, new_component_name, model_name_all)
    if not v:
        return False, err
    result = omc.addComponent(new_component_name, old_component_name, model_name_all, origin, extent, rotation)
    return result, err


def DeleteComponent(component_name, model_name_all):
    result = omc.deleteComponent(component_name, model_name_all)
    return result


def UpdateComponent(component_name, component_model_name, class_name_all, origin, extent, rotation):
    result = omc.updateComponent(component_name, component_model_name, class_name_all, origin, extent, rotation)
    return result
