# -- coding: utf-8 --
import logging

from config.omc import omc


def LoadModel(package_name="", path="", check=True):
    load_res = True
    package_name_list = omc.getClassNames()
    if check:
        if package_name not in package_name_list:
            load_res = omc.loadFile(path)
    else:
        load_res = omc.loadFile(path)
    return load_res
