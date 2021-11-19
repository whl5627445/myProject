# -- coding: utf-8 --
from config.omc import omc
from app.service.load_model_file import LoadModelFile


def AddConnection(model_name_all, connect_start, connect_end, line_points, color, file_path, package_name):
    # LoadModelFile(package_name, file_path)
    result = omc.addConnection(model_name_all, connect_start, connect_end, line_points, color)
    return result


def UpdateConnectionAnnotation(model_name_all, connect_start, connect_end, line_points, color, file_path, package_name):
    # LoadModelFile(package_name, file_path)
    result = omc.updateConnectionAnnotation(model_name_all, connect_start, connect_end, line_points, color)
    return result


def DeleteConnection(model_name_all, connect_start, connect_end, file_path, package_name):
    # LoadModelFile(package_name, file_path)
    result = omc.deleteConnection(model_name_all, connect_start, connect_end)
    return result
