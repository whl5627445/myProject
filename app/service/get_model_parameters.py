# -- coding: utf-8 --
import logging

from config.omc import omc
# from app.service.load_model_file import LoadModelFile


class GetModelParameters(object):
    def __init__(self, class_name, name, component_name, package_name=None):
        self.model_name = ""
        self.name = name
        self.class_name = class_name
        self.component_name = component_name
        self.package_name = package_name
        self.class_all = None
        self.Components = None
        self.ComponentAnnotations = None
        self.Components_dict = {}

    def get_derived_class_modifier_value(self):
        DerivedClassModifierNames = omc.getDerivedClassModifierNames(self.model_name)

        if DerivedClassModifierNames and DerivedClassModifierNames != ['']:
            logging.debug("model_name: " + str(self.model_name))
            logging.debug("DerivedClassModifierNames: " + str(DerivedClassModifierNames))
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
            if data != "":
                return data
        return data

    def get_extends_modifier_names_and_value(self):
        data_name_list = []
        data_value_list = []
        data_final_list = []
        n = 0
        m = n + 1
        while n < len(self.class_all) - 1:
            name_data = omc.getExtendsModifierNames(self.class_all[n], self.class_all[m])
            if name_data and name_data != 'Error' and name_data != ['']:
                data_name_list.extend(name_data)
                value_data = omc.getExtendsModifierValue(self.class_all[n], self.class_all[m], name_data[0])
                data_value_list.append(value_data)
                final_data = omc.isExtendsModifierFinal(self.class_all[n], self.class_all[m], name_data[0].split(".")[0])
                data_final_list.append(final_data)
            m += 1
            if m == len(self.class_all):
                n += 1
                m = n + 1
        return data_name_list, data_value_list, data_final_list

    def get_component_modifier_fixed_value(self, name):
        data = omc.getComponentModifierValue(self.class_name, name)
        return data

    def get_Parameter_value(self, name):
        data = ""
        for i in self.class_all:
            data = omc.getParameterValue(i, name)
            if data != "":
                return data
        return data

    def get_data(self):
        data_list = []
        self.class_all = omc.getInheritedClassesListAll([self.component_name])
        self.Components = omc.getElementsList(self.class_all)
        self.ComponentAnnotations = omc.getElementAnnotationsList(self.class_all)
        Components_dict = self.Components_dict
        if type(self.Components) is not list:
            return data_list
        for i in range(len(self.Components)):
            self.Components[i].append(self.ComponentAnnotations[i])
            Components_dict[self.Components[i][3]] = self.Components[i]

        for i in range(len(self.Components)):
            data_default = {
                "tab": "General",
                "type": "Normal",
                "group": ""
                }
            p = Components_dict[self.Components[i][3]]
            if p[2] != "-":
                self.model_name = p[2]
            else:
                self.model_name = ""
            var_name = p[3]
            data_default["name"] = var_name
            data_default["comment"] = p[4]
            Dialog_index = p[-1].index("Dialog") if "Dialog" in p[-1] else None
            showStartAttribute = None
            if Dialog_index is not None:
                tab_index = Dialog_index + 1
                # TODO Dialog 有可能返回error数据， 与单机版软件数据不一致，需要排查原因，暂时遇到了跳过
                if len(p[-1][tab_index]) <= 1:
                    continue
                tab = p[-1][tab_index][0]
                group = p[-1][tab_index][1]
                data_default["tab"] = tab
                data_default["group"] = group
                showStartAttribute = p[-1][tab_index][3]

            ComponentModifierValue = omc.getComponentModifierValue(self.class_name, '.'.join([self.name, data_default["name"]]))

            data_default["value"] = ComponentModifierValue
            if (p[10] != "parameter" and data_default["group"] != "Parameters" and p[9] != "True") or p[5] == "protected" or p[6] == "True":
                continue
            if p[10] == "parameter" or data_default["group"] == "Parameters" or p[9] == "True":
                data_default["group"]= "Parameters"
                isEnumeration = omc.isEnumeration(self.model_name)

                if isEnumeration:
                    Literals = omc.getEnumerationLiterals(self.model_name)
                    data_default["options"] = ['.'.join([self.model_name.removeprefix("."), i]) for i in Literals]
                    data_default["type"] = "Enumeration"
                ParameterValue = self.get_Parameter_value(data_default["name"])
                data_default["defaultvalue"] = ParameterValue
                # if p[13] != "$Any" or  p[9] == "True":
                #     all_subtype = omc.getAllSubtypeOf(p[13], self.component_name)
                #     data_default["options"] = ["redeclare " + i + " " + p[3] for i in all_subtype] if all_subtype else []
                #     data_default["type"] = "Enumeration"
                #     data_default["defaultvalue"] = "replaceable " + p[2] +  " " + p[3]
                if p[2] == 'Boolean':
                    data_default["type"] = "CheckBox"
                    if ComponentModifierValue in [True, False]:
                        if ComponentModifierValue:
                            ComponentModifierValue = "true"
                        else:
                            ComponentModifierValue = "false"
                        data_default["value"] = ComponentModifierValue
                        data_default["checked"] = ComponentModifierValue
                        data_default["defaultvalue"] = ComponentModifierValue

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
        extend_modifier_name, extend_modifier_value, extend_modifier_final = self.get_extends_modifier_names_and_value()
        if extend_modifier_name and extend_modifier_value:
            for i in range(len(extend_modifier_name)):
                var_name = extend_modifier_name[i].removesuffix(".start")
                self.model_name = Components_dict[var_name][2]
                data_default = {
                    "tab": "General", "type": "Normal", "group": "Initialization", "name": var_name + ".start",
                    "unit": self.get_derived_class_modifier_value(),
                    'comment': Components_dict[var_name][3],
                    'defaultvalue': extend_modifier_value[i],
                    'value': "",
                    }
                fixed = {
                    'tab': data_default["tab"],
                    'type': "fixed",
                    'group': "Initialization",
                    'name': var_name + ".fixed",
                    'comment': Components_dict[var_name][3],
                    'defaultvalue': extend_modifier_final[i],
                    'value': extend_modifier_final[i],
                    "unit": self.get_derived_class_modifier_value(),
                    }
                data_list.append(fixed)
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
