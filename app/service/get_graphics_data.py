# -- coding: utf-8 --
import logging

from config.omc import omc as mod
import os, re
# from app.service.load_model_file import LoadModelFile


class GetGraphicsData(object):

    def __init__ (self):
        self.InheritedClasses_data = []
        self.DiagramAnnotation_data = []
        self.Components_data = []
        self.ComponentAnnotations_data = []
        self.data = [[], []]
        self.mod = mod
        self.package_name = None


    def getICList (self, name):
        data_list = name
        name_list = name
        while True:
            InheritedClassesData = self.mod.getInheritedClassesList(name_list)
            if InheritedClassesData:
                data_list.extend(InheritedClassesData)
                name_list = InheritedClassesData
            else:
                break
        return data_list

    def data_01 (self, c_data):
        data_list = []
        if c_data == [''] or c_data == "-":
            return data_list
        drawing_data_list = c_data
        try:
            for i in range(0, len(drawing_data_list), 2):
                data = {}
                drawing_data = drawing_data_list[i + 1]
                data["visible"] = drawing_data[0]
                data["originalPoint"] = ",".join(drawing_data[1])
                data["rotation"] = drawing_data[2]
                data["type"] = drawing_data_list[i]
                if drawing_data_list[i] == "Polygon":
                    data["color"] = ",".join(drawing_data[3])
                    data["fillColor"] = ",".join(drawing_data[4])
                    data["linePattern"] = drawing_data[5]
                    data["fillPattern"] = drawing_data[6]
                    data["lineThickness"] = drawing_data[7]
                    data["polygonPoints"] = [",".join(x) for x in drawing_data[8]]
                    data["smooth"] = drawing_data[9]
                elif drawing_data_list[i] == "Line":
                    data["points"] = [",".join(x) for x in drawing_data[3]]
                    data["color"] = ",".join(drawing_data[4])
                    data["linePattern"] = drawing_data[5]
                    data["lineThickness"] = drawing_data[6]
                    data["arrow"] = ",".join(drawing_data[7])
                    data["arrowSize"] = drawing_data[8]
                    data["smooth"] = drawing_data[9]
                elif drawing_data_list[i] == "Text":
                    data["color"] = ",".join(drawing_data[3])
                    data["fillColor"] = ",".join(drawing_data[4])
                    data["linePattern"] = drawing_data[5]
                    data["fillPattern"] = drawing_data[6]
                    data["lineThickness"] = drawing_data[7]
                    data["extentsPoints"] = [",".join(x) for x in drawing_data[8]]
                    if type(drawing_data[9]) is list:
                        originalTextString = drawing_data[9][0]
                    else:
                        originalTextString = drawing_data[9]
                    data["originalTextString"] = originalTextString
                    data["fontSize"] = drawing_data[10]
                    data["textColor"] = ",".join(drawing_data[11])
                    data["fontName"] = drawing_data[12]
                    data["textStyles"] = drawing_data[13]
                    data["horizontalAlignment"] = drawing_data[14]
                elif drawing_data_list[i] == "Rectangle":
                    data["color"] = ",".join(drawing_data[3])
                    data["fillColor"] = ",".join(drawing_data[4])
                    data["linePattern"] = drawing_data[5]
                    data["fillPattern"] = drawing_data[6]
                    data["lineThickness"] = drawing_data[7]
                    data["borderPattern"] = drawing_data[8]
                    data["extentsPoints"] = [",".join(x) for x in drawing_data[9]]
                    data["radius"] = drawing_data[10]
                elif drawing_data_list[i] == "Ellipse":
                    data["color"] = ",".join(drawing_data[3])
                    data["fillColor"] = ",".join(drawing_data[4])
                    data["linePattern"] = drawing_data[5]
                    data["fillPattern"] = drawing_data[6]
                    data["lineThickness"] = drawing_data[7]
                    data["extentsPoints"] = [",".join(x) for x in drawing_data[8]]
                    data["startAngle"] = drawing_data[9]
                    data["endAngle"] = drawing_data[10]
                else:
                    pass
                data_list.append(data)
            return data_list
        except Exception as e:
            print("连线数据有误: ", drawing_data_list)
            print(e)

    def data_02 (self, c_data, ca_data, is_icon=False, parent=""):
        data_list = []
        c_data_filter = []
        ca_data_filter = []
        if is_icon and c_data != [] and ca_data != []:
            # if c_data != [] and ca_data != []:
            for i in range(len(c_data)):
                if "Interfaces" in c_data[i][0].split('.'):
                    c_data_filter.append(c_data[i])
                    ca_data_filter.append(ca_data[i])
        else:
            c_data_filter = c_data
            ca_data_filter = ca_data
        if ca_data_filter == [] or c_data_filter == [] or not ca_data_filter or not c_data_filter:
            return data_list
        for i in range(len(c_data_filter)):
            namelist = self.getICList([c_data_filter[i][0]])
            Placement_index = ca_data_filter[i].index("Placement") if "Placement" in ca_data_filter[i] else None
            if Placement_index is not None:
                Components_data = self.mod.getComponentsList(namelist)
                ComponentAnnotations_data = self.mod.getComponentAnnotationsList(namelist)
                IconAnnotation_data = self.mod.getIconAnnotationList(namelist)
                caf = ca_data_filter[i][Placement_index + 1]
                rotateAngle = "0" if caf[7] == "-" else caf[7]
                data = {"type": "Transformation" , "graphType": "connecter" if "Interfaces" in c_data_filter[i][0] else "",}
                data["ID"] = str(i)
                name = c_data_filter[i][1]
                data["original_name"] = c_data_filter[i][1]
                data["name"] = name
                data["parent"] = parent
                data["classname"] = c_data_filter[i][0]
                data["visible"] = ca_data_filter[i][1][0]
                data["visible"] = caf[0]
                data["rotateAngle"] = rotateAngle
                data["originDiagram"] = ",".join([ca_data_filter[i][1][1], ca_data_filter[i][1][2]])
                data["extent1Diagram"] = ",".join([ca_data_filter[i][1][3], ca_data_filter[i][1][4]])
                data["extent2Diagram"] = ",".join([ca_data_filter[i][1][5], ca_data_filter[i][1][6]])
                data["originDiagram"] = ",".join([caf[1], caf[2]])
                data["extent1Diagram"] = ",".join([caf[3], caf[4]])
                data["extent2Diagram"] = ",".join([caf[5], caf[6]])
                data["rotation"] = rotateAngle
                data["output_type"] = c_data_filter[i][-1][1:-1]
                data["inputOutputs"] = self.data_02(Components_data, ComponentAnnotations_data, is_icon=True,
                                                    parent=data["name"])
                data["subShapes"] = self.data_01(IconAnnotation_data)
                data_list.append(data)

        return data_list

    def getNthConnection_data (self, name_list):
        ConnectionCount = self.mod.getConnectionCountList(name_list)
        for count in range(len(ConnectionCount)):
            if ConnectionCount[count] != 0:
                for i in range(ConnectionCount[count]):
                    nc_data = self.mod.getNthConnection(name_list[count], i + 1)
                    nca_data = self.mod.getNthConnectionAnnotation(name_list[count], i + 1)
                    da_data = self.data_01(nca_data)[0]
                    da_data["connectionfrom_original_name"] = nc_data[0]
                    da_data["connectionto_original_name"] = nc_data[1]
                    expression1 = r"\[\d+\]$"
                    expression2 = r"\[\d+\]\."
                    connectionfrom = re.sub(expression1, "", nc_data[0])
                    connectionto = re.sub(expression1, "", nc_data[1])
                    da_data["connectionfrom"] = re.sub(expression2, ".", connectionfrom)
                    da_data["connectionto"] = re.sub(expression2, ".", connectionto)
                    self.data[0].append(da_data)

    def get_data (self, name_list):
        self.package_name = name_list[0].split(".")[0]
        name_list = self.getICList(name_list)
        DiagramAnnotation_data = self.mod.getDiagramAnnotationList(name_list)
        Components_data = self.mod.getComponentsList(name_list)
        ComponentAnnotations_data = self.mod.getComponentAnnotationsList(name_list)
        if len(DiagramAnnotation_data) >= 8:
            data_1 = self.data_01(DiagramAnnotation_data[-1])
            self.data[0].extend(data_1)
        self.getNthConnection_data(name_list)
        data_2 = self.data_02(Components_data, ComponentAnnotations_data)
        self.data[1].extend(data_2)
        return self.data

    def get_one_data (self, name_list, component_name):
        self.package_name = name_list[0].split(".")[0]
        name_list = self.getICList(name_list)
        Components = self.mod.getComponentsList(name_list)
        ComponentAnnotations = self.mod.getComponentAnnotationsList(name_list)
        Components_data = []
        ComponentAnnotations_data = []
        for i in range(len(Components)):
            if Components[i] != ['']:
                if Components[i][1] == component_name:
                    Components_data = [Components[i]]
                    ComponentAnnotations_data = [ComponentAnnotations[i]]
        data_2 = self.data_02(Components_data, ComponentAnnotations_data)
        self.data[1].extend(data_2)
        return self.data


if __name__ == '__main__':
    import json

    a = GetGraphicsData()
    # print(a.get_data(["Modelica.ComplexBlocks.Examples.TestConversionBlock"]))
    # print(a.get_data(["ENN.Examples.Scenario1_Status"]))
    # print(a.get_data(["Modelica.Blocks.Examples.PID_Controller"]))
    print(json.dumps(a.get_data(["Modelica.Blocks.Math.Sin"])))
    # print(a.get_data(["Modelica.Blocks.Sources.KinematicPTP"]))
    # print(a.get_data(["Modelica.Blocks.Examples.ShowLogicalSources"]))
