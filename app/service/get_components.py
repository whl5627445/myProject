from config.omc import omc


def GetComponents(class_name, component_name=None, file_path=None):
    # if file_path:
    #     omc.loadFile(file_path)
    data = omc.getComponents(class_name)
    if component_name:
        for i in data:
            if i[1] == component_name:
                data = i
    return data


if __name__ == '__main__':
    from app.service.set_component_properties import SetComponentProperties
    # setComponentProperties(ENN.Examples.Scenario1_Status,PID,{true,false,false,true}, {""}, {true,true}, {""})
    class_name = "ENN.Examples.Scenario1_Status"
    component_name = "PID"
    file_path = "/home/simtek/dev/public/UserFiles/UploadFile/tom/1631690039.291318/ENN.mo"
    data = {
        "class_name": class_name,
        "component_name": component_name,
        "final": "true",
        "protected": "true",
        "replaceable": "true",
        "variabilty": "parameter",
        "inner": "true",
        "outer": "true",
        "causality": "input",
    }
    print(SetComponentProperties(file_path, **data))
    print(GetComponents(class_name, component_name, file_path))