# -- coding: utf-8 --
from app.model.ModelsPackage.ModelsInformation import ModelsInformationAll, ModelsInformation
from config.DB_config import DBSession
from config.omc import omc
import os
import datetime
from app.service.load_model_file import LoadModelFile
session = DBSession()



def GetClassNames(name_dict, package_name, data_dict):
    data_d = {}
    for name, name_data in name_dict.items():
        cmd = "getClassNames(" + name + ")"
        data = omc.sendExpression(cmd)
        parent_name = name_data["parent_name"]
        if data == ['']:
            data = []
        data_dict[name] = {
            "package_name": package_name,
            "parent_name": parent_name,
            "child_name": data,
            "has_child": True,
            "model_name": name_data["model_name"]
        }
        if not data or data == ['']:
            data_dict[name]["has_child"] = False
            continue
        else:
            for i in data:
                data_d[name + '.' + i] = {
                    "parent_name": name,
                    "model_name": i,
                }
    if not data_d:
        return data_dict
    return GetClassNames(data_d, package_name, data_dict)


def GetIconsData(name):
    data = omc.sendExpression("getInheritedClasses(" + name + ")", parsed=False)
    cmd_data = data[1:-2]
    if data == "{}\n" or "\n":
        cmd_data = name
    cmd = "getIconAnnotation(" + cmd_data + ")"
    icons_data = omc.sendExpression(cmd, parsed=False)
    return icons_data


def SaveClassNames(mo_path=None, init_name="Modelica", sys_or_user="sys", package_id=""):
    data_dict = {}
    res = False
    model_root_data = {
        init_name: {
                "parent_name": "",
                "model_name": init_name,  # root节点名称与包名取自初始名称
                }
            }
    if mo_path:
        loadFile_result = omc.loadFile(mo_path)
        # LoadModelFile(init_name, path)
        if not loadFile_result:
            return res, None
    ClassNames = GetClassNames(model_root_data, init_name, data_dict)
    M_id = None
    for k, v in ClassNames.items():
        # icons_data = GetIconsData(k)
            if not v["parent_name"]:
                if package_id:
                    M = session.query(ModelsInformation).filter_by(id=package_id).first()
                    M.haschild = v["has_child"]
                    M.child_name = v["child_name"]
                    M.file_path = mo_path
                    M.update_time = datetime.datetime.now()
                else:
                    M = ModelsInformation(
                            package_name=v["package_name"],
                            model_name=v["model_name"],
                            haschild=v["has_child"],
                            child_name=v["child_name"],
                            sys_or_user=sys_or_user,
                            file_path=mo_path,
                            # image=icons_data,
                    )
                    session.add(M)
                session.flush()
                M_id = M.id
                break
    session.query(ModelsInformationAll).filter_by(package_id=package_id).delete(synchronize_session=False)
    session.flush()
    for k, v in ClassNames.items():
        # icons_data = GetIconsData(k)
        if v["parent_name"]:
            MA = ModelsInformationAll(
                    package_name=v["package_name"],
                    package_id=M_id,
                    model_name=v["model_name"],
                    model_name_all=k,
                    parent_name=v["parent_name"],
                    child_name=v["child_name"],
                    haschild=v["has_child"],
                    sys_or_user=sys_or_user,
                    # image=icons_data,
                    )
            session.add(MA)
    session.flush()
    session.close()
    res = True
    return res, M_id


if __name__ == '__main__':
    # print(SaveClassNames(mo_path="public/UserFiles/ENN.mo", init_name="ENN", sys_or_user="tom"))
    SaveClassNames()
