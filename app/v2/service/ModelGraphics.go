package service

import (
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"yssim-go/library/omc"
	instance "yssim-go/library/omc/ModelInstance"
)

type ModelInstanceData struct {
	Graphics   map[string]any   `json:"graphics,omitempty"`
	Parameters []map[string]any `json:"parameters,omitempty"`
}

func GetModelInstanceData(modelName string) *ModelInstanceData {
	s := time.Now().Local().UnixMilli()
	m := getModelInstance(modelName)
	fmt.Printf("模型实例化用时：%d", time.Now().Local().UnixMilli()-s)
	ss := time.Now().Local().UnixMilli()
	m.DataPreprocessing()
	fmt.Printf("数据预处理用时：%d", time.Now().Local().UnixMilli()-ss)
	sss := time.Now().Local().UnixMilli()
	modelData := &ModelInstanceData{}
	modelData.Parameters = getModelElementsParameter(m)
	modelData.Graphics = map[string]any{
		"connections": getConnectionsListAll(m),
		"diagram":     getDiagramListAll(m),
		"elements":    getElementsIconList(m),
	}
	i := map[string]any{"graphics": graphics, "parameters": make(map[string]any, 0)}
	fmt.Printf("逻辑处理用时：%d", time.Now().Local().UnixMilli()-sss)
	return modelData
}

// getModelElementsParameter 获取给定实例化模型的所有参数数据
func getModelElementsParameter(modelInstance *instance.ModelInstance) []map[string]any {
	modelParameterMap := map[string]map[string]*instance.Parameter{}

	return modelInstance.GetModelParameterValue(modelParameterMap, 0)
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
	for i := 0; i < len(modelInstance.Elements); i++ {
		e := modelInstance.Elements[i]
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

		modelIconList["visibleList"] = e.GetConnectionOption()

		// modelIconList["connectorSizing"] = e.Dims.Absyn
		modelIconList["subShapes"] = typeInstance.GetIconListALL(e)
		modelIconList["modelName"] = modelInstance.Name
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
	for i := 0; i < len(modelInstance.Elements); i++ {
		e := modelInstance.Elements[i]
		connectorSizingMap[e.Name] = e.Annotation.Dialog.ConnectorSizing
		if e.BaseClass != nil && !e.BaseClass.BasicType && e.Kind == "extends" {
			connectorList = append(connectorList, getElementsConnectorList(modelInstance.Elements[i].BaseClass, parentName)...)
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
			modelIconList["type"] = typeInstance.Name
			if typeInstance.Elements[0].BaseClass != nil && typeInstance.Elements[0].BaseClass.BasicType {
				modelIconList["type"] = typeInstance.Elements[0].BaseClass.Name
			}
			// modelIconList["connectorSizing"] =
			modelIconList["subShapes"] = typeInstance.GetIconListALL(e)
			modelIconList["modelName"] = modelInstance.Name
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

// geOutputType 获取接口的特殊标记
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
