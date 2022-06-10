# -- coding: utf-8 --
import logging

from config.omc import omc


def UploadIcon(model_name, icon_data):
    annotate_str = "Icon(graphics = {Bitmap(origin = {0, 0}, extent = {{-100,100},{100,-100}}, imageSource = \"" + icon_data + "\")})"
    result = omc.addClassAnnotation(model_name, annotate_str)
    if result is not True:
        return False
    return True


def GetIcon(model_name):
    icon_data = omc.getIconAnnotation(model_name)
    if len(icon_data)>7:
        image_data = icon_data[7]
        if image_data[0] == "Bitmap":
            image = image_data[1][5]
            return image
    return ""
