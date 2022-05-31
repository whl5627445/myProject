# -- coding: utf-8 --
import logging

from config.omc import omc


def CheckModel(model_name):
    omc.checkModel(model_name)
    data_list = GetMessagesStringInternal()
    return True, data_list


def GetMessagesStringInternal():
    res = omc.getMessagesStringInternal()
    res = res[1:-1]
    res = res.split(';,')
    message = []
    level = []
    data_list = []
    for i in res:
        data = i.strip().split(',\n')
        for j in data:
            j_data = j.strip()
            if "MODELICAPATH" in j_data  or "installPackage" in j_data:
                continue
            if j_data.startswith('message'):
                message.append(j_data.replace('message = ', '')[1:-1])
            if j_data.startswith('level'):
                level.append(j_data.split('.')[-1])
    for i in range(len(message)):
        data_list.append({
            "type": level[i],
            "message": message[i],
            })
    return data_list


def CheckUsesPackage(package_name):
    result = omc.getUses(package_name)
    classnames = omc.getClassNames()

    missing_library = []
    for i in result:
        if type(i) is list:
            if i[0] not in classnames:
                missing_library.append(i)
    return missing_library
