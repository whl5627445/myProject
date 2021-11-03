from config.omc import omc

def LoadModelFile(package_name, path, check=True):
    if check:
        package_name_list = omc.getClassNames()
        if package_name not in package_name_list:
            load_res = omc.loadFile(path)
    else:
        omc.loadFile(path)
