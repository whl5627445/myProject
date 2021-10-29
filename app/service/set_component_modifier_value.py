from config.omc import omc


def SetComponentModifierValue(className, parameter_value, path):
    # omc.loadFile(path)
    result = "Ok"
    for k, v in parameter_value.items():
        try:
            result = omc.setComponentModifierValue(className, k, v)
        except Exception as e:
            result = "err"
    return result
