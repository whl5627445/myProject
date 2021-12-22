# -- coding: utf-8 --
from config.omc import omc
from app.service.load_model_file import LoadModelFile


class GetModelParameters(object):
    def __init__(self, class_name, name, component_name, path=None, package_name=None):
        self.model_name = None
        self.name = name
        self.class_name = class_name
        self.path = path
        self.component_name = component_name
        self.package_name = package_name
        self.class_all = None
        self.Components = None
        self.ComponentAnnotations = None
        self.Components_dict = {}

    def get_derived_class_modifier_value(self):
        DerivedClassModifierNames = omc.getDerivedClassModifierNames(self.model_name)
        if DerivedClassModifierNames != ['']:
            DerivedClassModifierValue = []
            for n in range(1, len(DerivedClassModifierNames)):
                data = omc.getDerivedClassModifierValue(self.model_name, DerivedClassModifierNames[n])
                DerivedClassModifierValue.append(data)
            return DerivedClassModifierValue

    def get_component_modifier_start_value(self, name, show_start_attribute = False):
        data = ""
        if show_start_attribute:
            data = omc.getComponentModifierValue(self.class_name, name)
            return data
        for i in self.class_all:
            data = omc.getComponentModifierValue(i, name)
            if data:
                return data
        return data


    def get_component_modifier_fixed_value(self, name):
        data = omc.getComponentModifierValue(self.class_name, name)
        return data

    def get_Parameter_value(self, name):
        data = ""
        for i in self.class_all:
            data = omc.getParameterValue(i, name)
            if data:
                return data
        return data

    def get_data(self):
        if self.path:
            LoadModelFile(self.package_name, self.path)
        data_list = []
        self.class_all = omc.getInheritedClassesListAll([self.component_name])
        self.Components = omc.getComponentsList(self.class_all)
        self.ComponentAnnotations = omc.getComponentAnnotationsList(self.class_all)
        Components_dict = self.Components_dict
        if type(self.Components) is not list:
            return data_list
        for i in range(len(self.Components)):
            self.Components[i].append(self.ComponentAnnotations[i])
            Components_dict[self.Components[i][1]] = self.Components[i]

        for i in range(len(self.Components)):
            data_default = {
                "tab": "General",
                "type": "Normal",
                "group": "Parameters"
                }
            p = Components_dict[self.Components[i][1]]
            self.model_name = p[0]
            var_name = p[1]
            if p[3] == "protected" or p[4] == "True":  # 受保护的变量和不显示的暂时丢弃
                continue
            data_default["name"] = var_name
            data_default["comment"] = p[2]
            Dialog_index = p[-1].index("Dialog") if "Dialog" in p[-1] else None
            showStartAttribute = None
            if Dialog_index is not None:
                tab_index = Dialog_index + 1
                # TODO Dialog 有可能返回error数据， 与单机版软件数据不一致，需要排查原因，暂时遇到了调过
                if len(p[-1][tab_index]) <= 1:
                    continue
                tab = p[-1][tab_index][0]
                group = p[-1][tab_index][1]
                data_default["tab"] = tab
                data_default["group"] = group
                showStartAttribute = p[-1][tab_index][3]
            ComponentModifierValue = omc.getComponentModifierValue(self.class_name, '.'.join([self.name, data_default["name"]]))
            data_default["value"] = ComponentModifierValue
            if p[8] == 'parameter':
                isEnumeration = omc.isEnumeration(self.model_name)
                if isEnumeration:
                    Literals = omc.getEnumerationLiterals(self.model_name)
                    data_default["options"] = ['.'.join([self.model_name.removeprefix("."), i]) for i in Literals]
                    data_default["type"] = "Enumeration"
                ParameterValue = self.get_Parameter_value(data_default["name"])
                data_default["defaultvalue"] = ParameterValue

                if ParameterValue in [True, False]:
                    data_default["type"] = "CheckBox"
                    if ComponentModifierValue:  # 判断CheckBox的值不以ComponentModifierValue为第一优先级，与其他类型不同，如果ComponentModifierValue不为true，就以ParameterValue的值为主
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

            else:
                ComponentModifierNames = omc.getComponentModifierNamesList(self.class_all, var_name)
                fixed_value = self.get_component_modifier_fixed_value(self.name + "." + var_name + ".fixed")
                data_default["name"] = var_name + ".start"
                data_default["unit"] = self.get_derived_class_modifier_value()
                data_default["group"] = "Initialization"
                fixed = {
                    'tab': data_default["tab"],
                    'type': "fixed",
                    'group': "Initialization",
                    'name': var_name + ".fixed",
                    'comment': data_default["comment"],
                    'defaultvalue': fixed_value,
                    'value': fixed_value,
                    "unit": self.get_derived_class_modifier_value(),
                    }
                if showStartAttribute:
                    start_value = self.get_component_modifier_start_value(self.name + "." + var_name + ".start", show_start_attribute=True)
                    data_default["defaultvalue"] = start_value
                    data_list.append(fixed)
                elif "start" in ComponentModifierNames or "stateSelect" in ComponentModifierNames:
                    start_value = self.get_component_modifier_start_value(var_name + ".start")
                    data_default["defaultvalue"] = start_value
                    data_list.append(fixed)
                else:
                    continue
            data_default["unit"] = ""
            unit = self.get_derived_class_modifier_value()
            if unit:
                data_default["unit"] = unit
            data_list.append(data_default)
        data = data_list
        return data


if __name__ == '__main__':
    name = "PI"
    # name = "kinematicPTP"
    # name = "kinematicPTP"
    # name = "inertia1"
    # name = "spring"
    c_name = "Modelica.Blocks.Continuous.LimPID"
    # m_name = "Modelica.Blocks.Sources.KinematicPTP"
    # c_name = "Modelica.Blocks.Sources.KinematicPTP"
    # c_name = "Modelica.Mechanics.Rotational.Components.Inertia"
    # c_name = "Modelica.Mechanics.Rotational.Components.SpringDamper"
    class_name = "Modelica.Blocks.Examples.PID_Controller"
    import json
    print(json.dumps(GetModelParameters(class_name, name, c_name).get_data()))
