# -- coding: utf-8 --
from config.omc import omc
from app.service.load_model_file import LoadModelFile
from app.service.get_model_code import GetModelCode


def CreateModelicaClass(package_name, str_type, var, create_package_name_all, path=""):
    insert_to = var.get("insert_to", "")
    expand = var.get("expand", "")
    partial = var.get("partial", False)
    encapsulated = var.get("encapsulated", False)
    state = var.get("state", False)
    insert_package_name = ""
    if expand:
        expand = " extends " + expand + ";"
    model_str_base = str_type + " " + package_name + expand + " end " + package_name + ";"
    model_str = ""
    if insert_to:
        insert_package_name = insert_to.split(".")[0]
        LoadModelFile(insert_package_name, path)
        model_str = "within " + insert_to + "; "
    if encapsulated:
        model_str = model_str + "encapsulated "
    if partial :
        model_str = model_str + "partial "
    model_str = model_str + model_str_base
    res = omc.loadString(model_str, path)
    if state:
        res = omc.addClassAnnotation(create_package_name_all, annotate_str="Icon(graphics={Text(extent={{-100,100},{100,-100}},textString=\"%name\")})")
        res = omc.addClassAnnotation(create_package_name_all, annotate_str="annotate=__Dymola_state(true)")
        res = omc.addClassAnnotation(create_package_name_all, annotate_str="singleInstance(true)")
    return res


def UpdateModelicaClass(model_str, path="", merge="true"):
    load_string = omc.loadString(model_str, path, merge=merge)
    return load_string

if __name__ == '__main__':
    CreateModelicaClass("q123456", "model", {
        "insert_to": "ENN.Examples.PID_Controller",
        "expand": "ENN.Examples.Scenario1_Status",
        "partial": True,
        "encapsulated": True,
        "state": True,
        }, "/home/simtek/dev/public/UserFiles/ENN.mo")

