package service

import (
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"yssim-go/library/omc"
	instance "yssim-go/library/omc/ModelInstance"
)

func GetModelInstance(modelName string) map[string]any {
	s := time.Now().Local().UnixMilli()
	m := getModelInstance(modelName)
	fmt.Printf("模型实例化用时：%d", time.Now().Local().UnixMilli()-s)
	ss := time.Now().Local().UnixMilli()
	m.DataPreprocessing()
	fmt.Printf("数据预处理用时：%d", time.Now().Local().UnixMilli()-ss)
	sss := time.Now().Local().UnixMilli()
	graphics := map[string]any{
		"connections": GetConnectionsListAll(m),
		"diagram":     GetDiagramListAll(m),
		"elements":    GetElementsIconList(m),
	}
	i := map[string]any{"graphics": graphics, "parameters": make(map[string]any, 0)}
	fmt.Printf("逻辑处理用时：%d", time.Now().Local().UnixMilli()-sss)
	return i
}

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

// GetConnectionsListAll 获取模型实例的全部连接信息
func GetConnectionsListAll(m *instance.ModelInstance) map[string]any {

	connectionsList := make(map[string]any, 0)
	mDiagramList := m.GetConnectionsList()
	connectionsList["model"] = mDiagramList
	connectionsList["extends"] = GetExtendsConnectionsList(m)
	return connectionsList
}

// GetExtendsConnectionsList 获取模型实例继承模型的全部连接信息
func GetExtendsConnectionsList(m *instance.ModelInstance) []map[string]any {
	diagramList := make([]map[string]any, 0)
	for i := 0; i < len(m.Elements); i++ {
		if m.Elements[i].Kind != "extends" {
			continue
		}
		mDiagramList := m.Elements[i].BaseClass.GetConnectionsList()
		diagramList = append(diagramList, mDiagramList...)
		GetExtendsConnectionsList(m.Elements[i].BaseClass)
	}
	return diagramList
}

// GetDiagramListAll 获取模型实例的Diagram信息
func GetDiagramListAll(m *instance.ModelInstance) map[string]any {
	DiagramList := make(map[string]any, 0)
	mDiagramList := m.Annotation.Diagram.GetAnnotationDiagram()
	DiagramList["model"] = mDiagramList
	DiagramList["extends"] = GetExtendsDiagramList(m)
	return DiagramList
}

// GetExtendsDiagramList 获取模型实例继承模型的Diagram信息
func GetExtendsDiagramList(m *instance.ModelInstance) []map[string]any {
	DiagramList := make([]map[string]any, 0)
	for i := 0; i < len(m.Elements); i++ {
		if m.Elements[i].Kind != "extends" {
			continue
		}
		mDiagramList := m.Elements[i].BaseClass.Annotation.Diagram.GetAnnotationDiagram()
		DiagramList = append(DiagramList, mDiagramList)
		GetExtendsDiagramList(m.Elements[i].BaseClass)
	}
	return DiagramList
}

// getElementsGraphicsList 获取模型本身组件图形数据列表
func getElementsGraphicsList(m *instance.ModelInstance, parentName string) []map[string]any {
	elementsList := make([]map[string]any, 0)
	for i := 0; i < len(m.Elements); i++ {
		e := m.Elements[i]
		if (e.BaseClass != nil && e.BaseClass.BasicType && e.Kind == "extends") || e.Annotation.Placement == nil || e.Type == nil || (e.Type != nil && e.Type.BasicType) {
			continue
		}
		typeInstance := e.Type
		modelIconList := make(map[string]any, 0)
		modelIconList["name"] = e.Name
		modelIconList["classname"] = typeInstance.Name
		modelIconList["comment"] = typeInstance.Comment
		modelIconList["restriction"] = typeInstance.Restriction
		modelIconList["direction"] = typeInstance.Prefixes.Direction
		// modelIconList["connectorSizing"] = e.Dims.Absyn
		modelIconList["subShapes"] = typeInstance.Annotation.Icon.GetIconList(e)
		modelIconList["modelName"] = m.Name
		modelIconList["connectors"] = getElementsConnectorList(typeInstance, e.Name)
		// modelIconList["outputType"] = e.Dims.Absyn
		modelIconList["parentName"] = parentName
		modelIconList["origin"] = e.Annotation.Placement.Transformation.Origin
		modelIconList["extents"] = e.Annotation.Placement.Transformation.Extents
		modelIconList["rotation"] = e.Annotation.Placement.Transformation.Rotation
		modelIconList["coordinateSystem"] = typeInstance.Annotation.Icon.GetCoordinateSystem()
		elementsList = append(elementsList, modelIconList)
	}
	return elementsList
}

// GetElementsIconList 获取模型组件icon数据列表，包括模型本身的与继承过来的
func GetElementsIconList(m *instance.ModelInstance) map[string]any {
	iconList := make(map[string]any, 0)
	iconList["model"] = getElementsGraphicsList(m, "")
	iconList["extends"] = getExtendsElementsGraphicsList(m, m.Name)
	return iconList
}

// getExtendsElementsGraphicsList 获取模型继承组件图形数据列表
func getExtendsElementsGraphicsList(m *instance.ModelInstance, parentName string) []map[string]any {
	elementsList := make([]map[string]any, 0)
	for i := 0; i < len(m.Elements); i++ {
		e := m.Elements[i]
		if e.BaseClass != nil && e.BaseClass.BasicType || e.Kind != "extends" {
			continue
		}
		elementsList = append(elementsList, getElementsGraphicsList(e.BaseClass, parentName)...)
		getExtendsElementsGraphicsList(e.BaseClass, m.Name)
	}
	return elementsList
}

// getElementsConnectorList 获取模型组件连接器数据列表
func getElementsConnectorList(m *instance.ModelInstance, parentName string) []map[string]any {
	connectorList := make([]map[string]any, 0)
	connectorSizingMap := map[string]bool{}
	for i := 0; i < len(m.Elements); i++ {
		e := m.Elements[i]
		connectorSizingMap[e.Name] = e.Annotation.Dialog.ConnectorSizing
		if e.BaseClass != nil && !e.BaseClass.BasicType && e.Kind == "extends" {
			connectorList = append(connectorList, getElementsConnectorList(m.Elements[i].BaseClass, parentName)...)
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
			modelIconList := make(map[string]any, 0)
			modelIconList["name"] = e.Name
			// modelIconList["condition"] = condition
			modelIconList["coordinateSystem"] = typeInstance.Annotation.Diagram.GetCoordinateSystem()
			modelIconList["classname"] = typeInstance.Name
			modelIconList["comment"] = e.Comment
			modelIconList["restriction"] = typeInstance.Restriction
			modelIconList["direction"] = typeInstance.Prefixes.Direction
			// modelIconList["connectorSizing"] =
			modelIconList["subShapes"] = typeInstance.Annotation.Icon.GetIconList(e)
			modelIconList["modelName"] = m.Name
			modelIconList["outputType"] = geOutputType(connectorSizingMap, e.Dims.Absyn, e.Dims.Typed)
			modelIconList["parentName"] = parentName
			modelIconList["origin"] = e.Annotation.Placement.Transformation.Origin
			modelIconList["extents"] = e.Annotation.Placement.Transformation.Extents
			modelIconList["rotation"] = e.Annotation.Placement.Transformation.Rotation
			connectorList = append(connectorList, modelIconList)
		}
	}
	return connectorList
}

func geOutputType(connectorSizingMap map[string]bool, nameList, numList []string) map[string]any {
	opt := make(map[string]any, 0)
	for i, n := range nameList {
		if c, ok := connectorSizingMap[n]; ok && c {
			opt["name"] = nameList[i]
			opt["num"] = numList[i]
			opt["connectorSizing"] = true
			return opt
		}
	}
	return opt
}
