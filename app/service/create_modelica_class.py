# -- coding: utf-8 --
import logging

from config.omc import omc
# from app.service.load_model_file import LoadModelFile
from app.service.get_model_code import GetModelCode
from app.service.check_model import GetMessagesStringInternal
import json


def CreateModelicaClass(package_name, str_type, var, create_package_name_all, path=""):
    insert_to = var.get("insert_to", "")
    expand = var.get("expand", "")
    comment = var.get("comment", "")
    partial = var.get("partial", False)
    encapsulated = var.get("encapsulated", False)
    state = var.get("state", False)
    if expand:
        expand = " extends " + expand + ";"
    if comment:
        comment = " \\\"" + comment + "\\\""
    model_str_base = str_type + " " + package_name + comment + expand + " end " + package_name + ";"
    model_str = ""
    if insert_to:
        insert_package_name = insert_to.split(".")[0]
        model_str = "within " + insert_to + "; "
    if encapsulated:
        model_str = model_str + "encapsulated "
    if partial :
        model_str = model_str + "partial "
    model_str = model_str + model_str_base
    res = omc.loadString(model_str, path)
    logging.debug("CreateModelicaClass: " + model_str)
    logging.debug(res)
    if state:
        res = omc.addClassAnnotation(create_package_name_all, annotate_str="Icon(graphics={Text(extent={{-100,100},{100,-100}},textString=\"%name\")})")
        res = omc.addClassAnnotation(create_package_name_all, annotate_str="annotate=__Dymola_state(true)")
        res = omc.addClassAnnotation(create_package_name_all, annotate_str="singleInstance(true)")
    if res is True:
        return True, ""
    else:
        err_data = GetMessagesStringInternal()
        return False, err_data


def UpdateModelicaClass(model_str, path="", merge="false"):
    m_str = json.dumps(model_str, ensure_ascii=False)
    m_str = m_str[1:-1]
    m_str = m_str.replace("&lt;", "<")
    m_str = m_str.replace("&gt;", ">")
    m_str = m_str.replace("&amp;", "&")
    load_string = omc.loadString(m_str, path, merge=merge)
    if not load_string:
        err_data = GetMessagesStringInternal()
        return False, err_data
    return load_string, ""

if __name__ == '__main__':
    CreateModelicaClass("q123456", "model", {
        "insert_to": "ENN.Examples.PID_Controller",
        "expand": "ENN.Examples.Scenario1_Status",
        "partial": True,
        "encapsulated": True,
        "state": True,
        }, "/home/simtek/dev/public/UserFiles/ENN.mo")

