from config.omc import omc


def GetModelCode(model_name, model_file_path=None):
    if model_file_path:
        omc.loadFile(model_file_path)
    data = omc.sendExpression("list(" + model_name + ")")
    return data
