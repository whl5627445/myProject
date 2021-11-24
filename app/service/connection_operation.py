# -- coding: utf-8 --
from config.omc import omc
from app.service.load_model_file import LoadModelFile


def AddConnection(class_name_all, connect_start, connect_end, line_points, color, file_path, package_name):
    # LoadModelFile(package_name, file_path)
    result = omc.addConnection(class_name_all, connect_start, connect_end, line_points, color)
    return result


def UpdateConnectionNames(class_name_all, from_name, to_name, from_name_new, to_name_new):
    # LoadModelFile(package_name, file_path)
    result = omc.updateConnectionNames(class_name_all, from_name, to_name, from_name_new, to_name_new)
    return result


def UpdateConnectionAnnotation(class_name_all, connect_start, connect_end, line_points, color, file_path, package_name):
    # LoadModelFile(package_name, file_path)
    result = omc.updateConnectionAnnotation(class_name_all, connect_start, connect_end, line_points, color)
    return result


def DeleteConnection(class_name_all, connect_start, connect_end, file_path, package_name):
    # LoadModelFile(package_name, file_path)
    result = omc.deleteConnection(class_name_all, connect_start, connect_end)
    return result
