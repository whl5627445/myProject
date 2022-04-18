# -- coding: utf-8 --
import logging

from app.service.get_model_code import GetModelCode
from library.file_operation import FileOperation
from config.omc import omc
# from app.service.load_model_file import LoadModelFile


def CopyClass (copied_class_name, class_name, parent_name):
    exist_class = omc.existClass(parent_name + "." + class_name)
    if not parent_name:
        parent_name = "TopLevel"
    else:
        parent_information = omc.getClassInformation(parent_name)
        logging.debug(parent_information)
        if parent_information and parent_information[0] != "package":
            return False, "Parent node is not a package or root"
    if exist_class:
        return False, "Model already exists"
    logging.debug(exist_class)
    logging.debug(parent_name)
    copy_result = omc.copyClass(copied_class_name, class_name, parent_name)
    if copy_result:
        return True, "Copy model successfully"
    else:
        return False, "Copy model failed"


def DeleteClass (class_name):
    exist = omc.existClass(class_name)
    if exist:
        result = omc.deleteClass(class_name)
        if result:
            return True, "Delete model successfully"
        else:
            return False, "Delete model failed"
    else:
        return True, "Delete model successfully"


def SaveClass (class_name, copied_class_name=None, parent_name=None, package_name=None,
               new_model_file_path=None, copy_or_delete="copy", file_name=None):
    if copy_or_delete == "copy":
        result, msg = CopyClass(copied_class_name, class_name, parent_name)
    elif copy_or_delete == "delete":
        result, msg = DeleteClass(class_name)
    else:
        result = False
        msg = "Unknown operation"
    if result:
        if parent_name:
            data = GetModelCode(package_name)
        else:
            data = GetModelCode(class_name)
        if new_model_file_path:
            if not file_name:
                file_name = new_model_file_path.split("/")[-1]
            path = "/".join(new_model_file_path.split("/")[:-1])
            FileOperation.write_file(path, file_name, data)
    logging.debug(result)
    logging.debug(msg)
    return result, msg
