# -- coding: utf-8 --
from config.omc import omc
from app.service.load_model_file import LoadModelFile


def ComponentsVerification(class_name, component_name):
    class_information = omc.getClassInformation(class_name)
    if class_information:
        class_type = class_information[0]
        if class_type in ["package"]:
            return False, "包类型无法创建组件"
    else:
        return False
    data = omc.getComponents(class_name)
    if data != "Error":
        for i in data:
            if i[1] == component_name:
                return False, "组件名称已存在"
    else:
        return False, ""
    return True, ""


def AddComponent(new_component_name, old_component_name, model_name_all, origin, extent, rotation, file_path, package_name):
    LoadModelFile(package_name, file_path)
    v, err = ComponentsVerification(model_name_all, new_component_name)
    if not v:
        return False, err
    result = omc.addComponent(new_component_name, old_component_name, model_name_all, origin, extent, rotation)
    return result, err


def DeleteComponent(component_name, model_name_all, file_path, package_name):
    LoadModelFile(package_name, file_path)
    result = omc.deleteComponent(component_name, model_name_all)
    return result


def UpdateComponent(new_component_name, old_component_name, model_name_all, origin, extent, rotation, file_path, package_name):
    LoadModelFile(package_name, file_path)
    result = omc.updateComponent(new_component_name, old_component_name, model_name_all, origin, extent, rotation)
    return result
