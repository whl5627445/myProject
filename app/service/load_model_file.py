# -- coding: utf-8 --
from config.omc import omc
from config.DB_config import session
from app.model.ModelsPackage.ModelsInformation import ModelsInformation


def LoadModelFile(package_name, path, username="", check=True, return_res=False):
    if check:
        package_name_list = omc.getClassNames()
        if package_name not in package_name_list:
            models = session.query(ModelsInformation).filter(ModelsInformation.sys_or_user == username).all()
            for model in models:
                load_res = omc.loadFile(model.file_path)

    else:
        omc.loadFile(path)
