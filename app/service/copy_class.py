# -- coding: utf-8 --
from app.service.get_model_code import GetModelCode
from library.file_operation import FileOperation
from config.omc import omc
from app.service.load_model_file import LoadModelFile


def CopyClass(copied_class_name, class_name, parent_name):
    exist_class = omc.existClass(parent_name + "." + class_name)
    if exist_class:
        return False
    copy_result = omc.copyClass(copied_class_name, class_name, parent_name)
    return copy_result

def DeleteClass(class_name):
    result = omc.deleteClass(class_name)
    return result

def SaveClass(class_name, copied_class_name=None, parent_name=None, package_name=None, model_file_path=None, new_model_file_path=None, copy_or_delete="copy"):
    LoadModelFile(package_name, model_file_path)
    if copy_or_delete=="copy":
        result = CopyClass(copied_class_name, class_name, parent_name)
    elif copy_or_delete=="delete":
        result = DeleteClass(class_name)
    else:
        return False
    if result:
        data = GetModelCode(package_name)
        file_name = model_file_path.split("/")[-1]
        path = "/".join(new_model_file_path.split("/")[:-1])
        FileOperation.write_file(path, file_name, data)
    else:
        return False
    return True

