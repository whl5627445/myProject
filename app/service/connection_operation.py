# -- coding: utf-8 --
from config.omc import omc


def AddConnection(class_name_all, connect_start, connect_end, line_points, color):
    result = omc.addConnection(class_name_all, connect_start, connect_end, line_points, color)
    return result


def UpdateConnectionNames(class_name_all, from_name, to_name, from_name_new, to_name_new):
    result = omc.updateConnectionNames(class_name_all, from_name, to_name, from_name_new, to_name_new)
    return result


def UpdateConnectionAnnotation(class_name_all, connect_start, connect_end, line_points, color):
    result = omc.updateConnectionAnnotation(class_name_all, connect_start, connect_end, line_points, color)
    return result


def DeleteConnection(class_name_all, connect_start, connect_end):
    result = omc.deleteConnection(class_name_all, connect_start, connect_end)
    return result
