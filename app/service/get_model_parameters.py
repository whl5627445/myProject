from config.omc import omc
import json


def GetModelParameters(class_name, name, component_name, path=None):
    if path:
        omc.loadFile(path)
    data_list = []
    Components = omc.getComponents(component_name)
    ComponentAnnotations = omc.getComponentAnnotations(component_name)
    ParameterNames = omc.getParameterNames(component_name)
    ClassComment = omc.getClassComment(component_name)
    Components_dict = {}
    if type(Components) is not list:
        return data_list
    for i in range(len(Components)):
        Components[i].append(ComponentAnnotations[i])
        Components_dict[Components[i][1]] = Components[i]

    for i in range(len(ParameterNames)):
        data_default = {
            "tab": "General", #
            "type": "Normal", #
            "group": "Parameters" #
        }
        p = Components_dict[ParameterNames[i]]
        data_default["name"] = p[1]
        data_default["comment"] = p[2]
        Dialog_index = p[-1].index("Dialog") if "Dialog" in p[-1] else None
        if 'HideResult=true' in p[-1]:
            continue
        if Dialog_index is not None:
            tab_index = Dialog_index + 1
            tab = p[-1][tab_index][0]
            group = p[-1][tab_index][1]
            data_default["tab"] = tab
            data_default["group"] = group

        isEnumeration = omc.isEnumeration(p[0])
        if isEnumeration == 'True':
            Literals = omc.getEnumerationLiterals(p[0])
            data_default["options"] = ['.'.join([p[0].removeprefix("."), i]) for i in Literals]
            data_default["type"] = "Enumeration"
        ParameterValue = omc.getParameterValue(component_name, p[1])
        data_default["defaultvalue"] = ParameterValue
        ComponentModifierValue = omc.getComponentModifierValue(class_name,'.'.join(
                    [name, p[1]]))
        data_default["value"] = ComponentModifierValue
        if p[0] == "Boolean":
            data_default["type"] = "CheckBox"
            if ComponentModifierValue:     # 判断CheckBox的值不以ComponentModifierValue为第一优先级，与其他类型不同，如果ComponentModifierValue不为true，就以ParameterValue的值为主
                data_default["value"] = "true"
                data_default["checked"] = "true"
                data_default["defaultvalue"] = "true"
            else:
                if ParameterValue:
                    ParameterValue = "true"
                else:
                    ParameterValue = "false"
                data_default["value"] = ParameterValue
                data_default["checked"] = ParameterValue
                data_default["defaultvalue"] = ParameterValue
        elif p[0] == "Real":
            pass
        else:
            DerivedClassModifierNames = omc.getDerivedClassModifierNames(p[0])
            if DerivedClassModifierNames != ['']:
                DerivedClassModifierValue = omc.getDerivedClassModifierValue(p[0], DerivedClassModifierNames[1])
                data_default["unit"] = DerivedClassModifierValue
        data_list.append(data_default)
    data = data_list
    return data


if __name__ == '__main__':
    name = "PI"
    m_name = "Modelica.Blocks.Continuous.LimPID"
    class_name = "Modelica.Blocks.Examples.PID_Controller"
    print(GetModelParameters(class_name, name, m_name))
