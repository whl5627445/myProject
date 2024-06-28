package serviceV2

import (
	"log"
	"strings"

	"yssim-go/library/omc"
	instance "yssim-go/library/omc/ModelInstance"

	"github.com/bytedance/sonic"
)

type ModelInstanceData struct {
	Graphics   map[string]any   `json:"graphics,omitempty"`
	Parameters []map[string]any `json:"parameters,omitempty"`
}

func GetModelInstanceData(modelName string) *ModelInstanceData {
	m := getModelInstance(modelName)
	if m == nil {
		return nil
	}

	m.DataPreprocessing()
	modelData := &ModelInstanceData{}
	modelData.Parameters = getModelElementsParameter(m)
	modelData.Graphics = map[string]any{
		"connections": getConnectionsListAll(m),
		"diagram":     getDiagramListAll(m),
		"elements":    getElementsIconList(m),
	}
	return modelData
}

// getModelElementsParameter 获取给定实例化模型的所有参数数据
func getModelElementsParameter(modelInstance *instance.ModelInstance) []map[string]any {
	modelParameterMap := map[string]map[string]*instance.Parameter{}
	return modelInstance.GetModelParameterValue(modelParameterMap, false, 0)
}

// getModelInstance 获取给定模型名字的实例化数据
func getModelInstance(modelName string) *instance.ModelInstance {
	i := omc.OMC.GetModelInstance(modelName)
	m := &instance.ModelInstance{}
	err := sonic.Unmarshal(i, m)
	if err != nil {
		log.Println("ModelInstance数据序列化失败: ", err)
		return nil
	}
	return m
}

// getConnectionsListAll 获取模型实例的全部连接信息
func getConnectionsListAll(modelInstance *instance.ModelInstance) map[string]any {

	connectionsList := make(map[string]any, 0)
	mDiagramList := modelInstance.GetConnectionsList()
	connectionsList["model"] = mDiagramList
	connectionsList["extends"] = getExtendsConnectionsList(modelInstance)
	return connectionsList
}

// getExtendsConnectionsList 获取模型实例继承模型的全部连接信息
func getExtendsConnectionsList(m *instance.ModelInstance) []map[string]any {
	diagramList := make([]map[string]any, 0)
	for i := 0; i < len(m.Elements); i++ {
		if m.Elements[i].Kind != "extends" {
			continue
		}
		diagramList = append(diagramList, getExtendsConnectionsList(m.Elements[i].BaseClass)...)
		diagramList = append(diagramList, m.Elements[i].BaseClass.GetConnectionsList()...)
	}
	return diagramList
}

// getDiagramListAll 获取模型实例的Diagram信息
func getDiagramListAll(modelInstance *instance.ModelInstance) map[string]any {
	DiagramList := make(map[string]any, 0)
	mDiagramList := append([]map[string]any{}, modelInstance.Annotation.Diagram.GetAnnotationDiagram()...)
	DiagramList["model"] = mDiagramList
	DiagramList["extends"] = getExtendsDiagramMap(modelInstance)
	return DiagramList
}

// getExtendsDiagramMap 获取模型实例继承模型的Diagram信息
func getExtendsDiagramMap(modelInstance *instance.ModelInstance) []map[string]any {
	DiagramList := make([]map[string]any, 0)
	for i := 0; i < len(modelInstance.Elements); i++ {
		if modelInstance.Elements[i].Kind != "extends" {
			continue
		}
		DiagramList = append(DiagramList, getExtendsDiagramMap(modelInstance.Elements[i].BaseClass)...)
		DiagramList = append(DiagramList, modelInstance.Elements[i].BaseClass.Annotation.Diagram.GetAnnotationDiagram()...)
	}
	return DiagramList
}

// getElementsGraphicsList 获取模型本身组件图形数据列表
func getElementsGraphicsList(modelInstance *instance.ModelInstance, parentName string) []map[string]any {
	elementsList := make([]map[string]any, 0)
	connectorSizingMap := map[string]bool{}
	for i := 0; i < len(modelInstance.Elements); i++ {
		e := modelInstance.Elements[i]
		connectorSizingMap[e.Name] = e.Annotation.Dialog.ConnectorSizing
		if (e.BaseClass != nil && e.BaseClass.BasicType && e.Kind == "extends") || e.Annotation.Placement == nil || e.Type == nil || (e.Type != nil && e.Type.BasicType) {
			continue
		}
		typeInstance := e.Type
		modelIconList := make(map[string]any, 0)
		modelIconList["type"] = ""
		if len(typeInstance.Elements) > 0 && typeInstance.Elements[0].BaseClass != nil && typeInstance.Elements[0].BaseClass.BasicType {
			modelIconList["type"] = typeInstance.Elements[0].BaseClass.Name
		}
		modelIconList["name"] = e.Name
		modelIconList["classname"] = typeInstance.Name
		modelIconList["comment"] = typeInstance.Comment
		modelIconList["restriction"] = typeInstance.Restriction
		if modelIconList["restriction"] == "expandable connector" || modelIconList["restriction"] == "connector" {
			modelIconList["type"] = getConnectorType(e.Name, typeInstance)
		}
		modelIconList["direction"] = typeInstance.Prefixes.Direction
		modelIconList["visibleList"] = nil
		if e.Kind != "extends" {
			modelIconList["visibleList"] = GetConnectionOption(e.Name, e.Type, e.Modifiers)
		}
		modelIconList["subShapes"] = typeInstance.GetIconListALL(e, true)
		modelIconList["modelName"] = modelInstance.Name
		modelIconList["outputType"] = getOutputType(connectorSizingMap, e.Dims.Absyn, e.Dims.Typed)
		modelIconList["connectors"] = getElementsConnectorList(typeInstance, e.Name)
		modelIconList["parentName"] = parentName
		modelIconList["origin"] = e.Annotation.Placement.GetElementsOrigin()
		modelIconList["extents"] = e.Annotation.Placement.GetElementsExtents()
		modelIconList["rotation"] = e.Annotation.Placement.Transformation.Rotation
		modelIconList["coordinateSystem"] = typeInstance.Annotation.Icon.GetCoordinateSystem()
		elementsList = append(elementsList, modelIconList)
	}
	return elementsList
}

// getElementsIconList 获取模型组件icon数据列表，包括模型本身的与继承过来的
func getElementsIconList(modelInstance *instance.ModelInstance) map[string]any {
	iconList := make(map[string]any, 0)
	iconList["model"] = getElementsGraphicsList(modelInstance, "")
	iconList["extends"] = getExtendsElementsGraphicsList(modelInstance, modelInstance.Name)
	return iconList
}

// getExtendsElementsGraphicsList 获取模型继承组件图形数据列表
func getExtendsElementsGraphicsList(modelInstance *instance.ModelInstance, parentName string) []map[string]any {
	elementsList := make([]map[string]any, 0)
	for i := 0; i < len(modelInstance.Elements); i++ {
		e := modelInstance.Elements[i]
		if e.BaseClass != nil && e.BaseClass.BasicType || e.Kind != "extends" {
			continue
		}
		elementsList = append(elementsList, getExtendsElementsGraphicsList(e.BaseClass, modelInstance.Name)...)
		elementsList = append(elementsList, getElementsGraphicsList(e.BaseClass, parentName)...)
	}
	return elementsList
}

// getElementsConnectorList 获取模型组件连接器数据列表
func getElementsConnectorList(modelInstance *instance.ModelInstance, parentName string) []map[string]any {
	connectorList := make([]map[string]any, 0)
	connectorSizingMap := map[string]bool{}
	connectorDumpMap := map[string]string{}
	for i := 0; i < len(modelInstance.Elements); i++ {
		e := modelInstance.Elements[i]
		connectorSizingMap[e.Name] = e.Annotation.Dialog.ConnectorSizing
		if e.BaseClass != nil && !e.BaseClass.BasicType && e.Kind == "extends" {
			extendsConnectorList := getElementsConnectorList(modelInstance.Elements[i].BaseClass, parentName)
			for _, connector := range extendsConnectorList {
				if _, ok := connectorDumpMap[connector["name"].(string)]; !ok {
					connectorList = append(connectorList, connector)
					connectorDumpMap[connector["name"].(string)] = connector["classname"].(string)
				}
			}
			continue
		}
		if e.Type == nil || (e.Type != nil && e.Type.BasicType) {
			continue
		}
		typeInstance := e.Type
		if (typeInstance.Restriction == "expandable connector" || typeInstance.Restriction == "connector") && e.Annotation.Placement != nil {
			if c, ok := e.Condition.(bool); ok && !c {
				// condition = c
				continue
			}

			if _, ok := connectorDumpMap[e.Name]; ok {
				continue
			}

			modelIconList := make(map[string]any, 0)
			modelIconList["name"] = e.Name
			modelIconList["coordinateSystem"] = typeInstance.Annotation.Diagram.GetCoordinateSystem()
			modelIconList["classname"] = typeInstance.Name
			modelIconList["comment"] = e.Comment
			modelIconList["restriction"] = typeInstance.Restriction
			modelIconList["direction"] = typeInstance.Prefixes.Direction
			modelIconList["type"] = ""
			if len(typeInstance.Elements) > 0 && typeInstance.Elements[0].BaseClass != nil && typeInstance.Elements[0].BaseClass.BasicType {
				modelIconList["type"] = typeInstance.Elements[0].BaseClass.Name
			}
			modelIconList["type"] = getConnectorType(e.Name, typeInstance)
			modelIconList["subShapes"] = typeInstance.GetIconListALL(e, false)
			modelIconList["modelName"] = modelInstance.Name
			modelIconList["outputType"] = getOutputType(connectorSizingMap, e.Dims.Absyn, e.Dims.Typed)
			modelIconList["parentName"] = parentName
			modelIconList["origin"] = e.Annotation.Placement.GetElementsOrigin()
			modelIconList["extents"] = e.Annotation.Placement.GetElementsExtents()
			modelIconList["rotation"] = e.Annotation.Placement.Transformation.Rotation
			connectorList = append(connectorList, modelIconList)
			connectorDumpMap[e.Name] = typeInstance.Name
		}
	}
	return connectorList
}

// getConnectorType 获取连接器类型
func getConnectorType(name string, typeInstance *instance.ModelInstance) *instance.TypeConnector {
	t := &instance.TypeConnector{
		ClassName:   typeInstance.Name,
		Name:        name,
		Restriction: typeInstance.Restriction,
		Direction:   typeInstance.Prefixes.Direction,
		Elements:    make([]*instance.TypeConnector, 0),
		Extends:     make([]*instance.TypeConnector, 0),
		Type:        "",
	}

	for i := 0; i < len(typeInstance.Elements); i++ {
		e := typeInstance.Elements[i]

		if e.Kind == "extends" && e.BaseClass != nil && e.BaseClass.BasicType == true {
			t.Type = e.BaseClass.Name
			continue
		}

		if e.Kind == "extends" {
			t.Extends = append(t.Extends, getConnectorType(e.Name, e.BaseClass))
		} else {
			t.Elements = append(t.Elements, getConnectorType(e.Name, e.Type))
		}
	}

	return t
}

// getOutputType 获取接口的特殊标记
func getOutputType(connectorSizingMap map[string]bool, nameList, numList []string) map[string]any {
	opt := make(map[string]any, 0)
	for i, n := range nameList {
		opt["name"] = nameList[i]
		opt["num"] = numList[i]
		opt["connectorSizing"] = connectorSizingMap[n]
		return opt
	}
	return opt
}

// GetConnectionOption 获取visibleList
func GetConnectionOption(elementName string, typeInstance *instance.ModelInstance, modifiers map[string]any) []map[string]any {
	if typeInstance.Restriction == "expandable connector" {
		return getConnectionOptionExpandableConnector(elementName, typeInstance, modifiers)
	}

	if typeInstance.Restriction == "connector" {
		return getConnectionOptionConnector(elementName, typeInstance, modifiers)
	}

	return nil
}

// getConnectionOptionExpandableConnector 获取expandable connector的visibleList
func getConnectionOptionExpandableConnector(elementName string, typeInstance *instance.ModelInstance, modifiers map[string]any) []map[string]any {
	variableList := make([]map[string]any, 0)
	connectorSizingMap := map[string]bool{}
	for _, element := range typeInstance.Elements {
		if element.Kind == "extends" {
			continue
		}
		connectorSizingMap[element.Name] = element.Annotation.Dialog.ConnectorSizing
		option := getConnectionOptionExpandableConnector(element.Name, element.Type, element.Modifiers)
		ser := map[string]any{
			"variableName": element.Name,
			"hasChild":     len(option) > 0,
			"option":       option,
			"outputType":   getOutputTypeExpandableConnector(connectorSizingMap, element.Dims.Absyn, element.Dims.Typed, modifiers),
		}
		variableList = append(variableList, ser)
	}

	return variableList
}

// getConnectionOptionConnector 获取connector的visibleList
func getConnectionOptionConnector(parentName string, typeInstance *instance.ModelInstance, modifiers map[string]any) []map[string]any {
	variableList := getSubConnectorList(typeInstance, parentName, modifiers)
	return variableList
}

// getSubConnectorList 获取子连接器列表
func getSubConnectorList(modelInstance *instance.ModelInstance, parentName string, modifiers map[string]any) []map[string]any {
	connectorList := make([]map[string]any, 0)
	connectorSizingMap := map[string]bool{}
	connectorRemoveDuplicateMap := map[string]string{}
	for i := 0; i < len(modelInstance.Elements); i++ {
		e := modelInstance.Elements[i]
		connectorSizingMap[e.Name] = e.Annotation.Dialog.ConnectorSizing
		if e.BaseClass != nil && !e.BaseClass.BasicType && e.Kind == "extends" {
			extendsConnectorList := getSubConnectorList(modelInstance.Elements[i].BaseClass, parentName, e.Modifiers)
			for _, connector := range extendsConnectorList {
				if _, ok := connectorRemoveDuplicateMap[connector["variableName"].(string)]; !ok {
					connectorList = append(connectorList, connector)
					connectorRemoveDuplicateMap[connector["variableName"].(string)] = connector["classname"].(string)
				}
			}
			continue
		}
		if e.Type == nil || (e.Type != nil && e.Type.BasicType) {
			continue
		}
		typeInstance := e.Type
		if typeInstance.Restriction == "expandable connector" || typeInstance.Restriction == "connector" {
			if c, ok := e.Condition.(bool); ok && !c {
				// condition = c
				continue
			}
			if _, ok := connectorRemoveDuplicateMap[e.Name]; ok {
				continue
			}

			modelIconList := make(map[string]any, 0)
			modelIconList["variableName"] = e.Name
			modelIconList["hasChild"] = false
			modelIconList["option"] = nil
			modelIconList["classname"] = typeInstance.Name
			modelIconList["type"] = getConnectorType(e.Name, typeInstance)
			modelIconList["outputType"] = getOutputTypeConnector(connectorSizingMap, e.Dims.Absyn, e.Dims.Typed, modifiers)

			connectorList = append(connectorList, modelIconList)
			connectorRemoveDuplicateMap[e.Name] = typeInstance.Name
		}
	}
	return connectorList
}

// getOutputTypeConnector 获取连接器的维度
func getOutputTypeConnector(connectorSizingMap map[string]bool, nameList, numList []string, modifiers map[string]any) map[string]any {
	opt := make(map[string]any, 0)
	for i, n := range nameList {
		opt["name"] = nameList[i]

		// 维度是参数而不是数字时需将参数转化为数字
		if strings.Contains(numList[i], ".") && len(modifiers) > 0 {
			strs := strings.Split(numList[i], ".")
			num := modifiers[strs[len(strs)-1]]
			opt["num"] = num
		} else {
			opt["num"] = numList[i]
		}

		opt["connectorSizing"] = connectorSizingMap[n]
		return opt
	}
	return opt
}

// getOutputTypeExpandableConnector 获取变量的维度
func getOutputTypeExpandableConnector(connectorSizingMap map[string]bool, nameList, numList []string, modifiers map[string]any) []map[string]any {
	opts := make([]map[string]any, 0)
	for i, n := range nameList {
		opt := map[string]any{}
		opt["name"] = nameList[i]

		// 维度是参数而不是数字时需将参数转化为数字
		if strings.Contains(numList[i], ".") && len(modifiers) > 0 {
			strs := strings.Split(numList[i], ".")
			num := modifiers[strs[len(strs)-1]]
			opt["num"] = num
		} else {
			opt["num"] = numList[i]
		}

		opt["connectorSizing"] = connectorSizingMap[n]
		opts = append(opts, opt)
	}
	return opts
}
