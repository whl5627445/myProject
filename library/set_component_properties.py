from config.omc import omc

def SetComponentProperties(file_path, **kwargs):
    omc.loadFile(file_path)
    result = omc.setComponentProperties(**kwargs)
    return result
