# -- coding: utf-8 --
import logging

from config.omc import omc
from app.service.icon_operation import GetIcon
import time


def GetPackageNodeTree(package_name, sys_or_user="user"):
    TREE_dict = {
        'model_name': package_name,
        'model_full_name': package_name,
        'child': [],
        'image': "",
        'haschild': False,
        'sys_or_user ': sys_or_user
        }
    name_list = omc.getClassNames(package_name)
    child_list = []
    for name in name_list:
        if name:
            model_full_name = ".".join([package_name, name])
            name_dict = {
                'model_name': name,
                'model_full_name': model_full_name,
                'child': GetPackageNodeTree(model_full_name),
                'image': "",
                'haschild': True,
                'sys_or_user ': sys_or_user
        }
            child_list.append(name_dict)
    TREE_dict['child'] = child_list
    return TREE_dict

