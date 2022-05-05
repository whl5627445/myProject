# -- coding: utf-8 --
import logging

from config.omc import omc


def UploadIcon(model_name, icon_data):
    # annotate=Icon(icon="XXXX")
    # annotate_str = "Icon(image=" + icon_data + ")"
    annotate_str = "Icon(graphics = {Bitmap(origin = {0, 0}, extent = {{100, 100}, {-100, -100}}, imageSource = \"" + icon_data + "\")})"
    result = omc.addClassAnnotation(model_name, annotate_str)
    logging.info(result)
    if result is not True:
        return False
    return True


def GetIcon(model_name):
    icon_data = omc.getIconAnnotation(model_name)
    if len(icon_data) > 7:
        if len(icon_data[7]) > 1:
            if icon_data[7][0] == "Bitmap":
                return True, icon_data[7][1][5]
    return False, ""
