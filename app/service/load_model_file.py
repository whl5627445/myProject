# -- coding: utf-8 --
from config.omc import omc

def LoadModelFile(package_name, path, check=True, return_res=False):
    if check:
        package_name_list = omc.getClassNames()
        if package_name not in package_name_list:
            load_res = omc.loadFile(path)
            if return_res:
                return load_res

    else:
        omc.loadFile(path)
