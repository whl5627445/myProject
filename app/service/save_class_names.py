# -- coding: utf-8 --
import logging

from app.model.ModelsPackage.ModelsInformation import ModelsInformation
from config.DB_config import DBSession
from config.omc import omc
from app.service.check_model import GetMessagesStringInternal, CheckUsesPackage
import os
from app.service.load_model_file import LoadModel
import datetime
session = DBSession()


def SaveClassNames(space_id, mo_path=None, init_name="Modelica", sys_or_user="sys", package_id=""):
    res = False
    data_list = []
    if mo_path:
        LoadModel(path=mo_path, check=False)
        use_package = CheckUsesPackage(init_name)
        if use_package:
            data_list.append({"type": "error", "message": "Minssing " + ",".join(
                        [i[0] + "(" + i[1] + ")" for i in use_package]) + " and other libraries"})
            return res, None, data_list
        loadFile_result = omc.loadFile(mo_path)
        data_list = GetMessagesStringInternal()
        if not loadFile_result:
            return res, None, data_list
    if package_id:
        M = session.query(ModelsInformation).filter_by(id=package_id).first()
        M.file_path = mo_path
        M.update_time = datetime.datetime.now()
    else:
        M = ModelsInformation(
                package_name=init_name,
                model_name=init_name,
                sys_or_user=sys_or_user,
                file_path=mo_path,
                userspace_id=space_id,
        )
        session.add(M)
    M_id = M.id
    session.flush()
    session.close()
    res = True
    return res, M_id, data_list


if __name__ == '__main__':
    print(SaveClassNames(mo_path="/home/simtek/code/omlibrary/Buildings 8.0.0/package.mo", init_name="Buildings", sys_or_user="sys"))
    # print(SaveClassNames(space_id=None))
