package service

import (
	"encoding/json"
	"reflect"
	"yssim-go/app/DataBaseModel"

	"github.com/bytedance/sonic"
)

type ComponentParameterList struct {
	Data []*ComponentParameter `json:"public,omitempty"`
}

type ComponentParameter struct {
	Name       string       `json:"name,omitempty"`
	Parameters []*Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	Name       string `json:"name,omitempty"`
	Value      any    `json:"value,omitempty"` // Value是string或是valueObject结构体
	ExtendName string `json:"extend_name,omitempty"`
}

type valueObject struct {
	Value   string `json:"value,omitempty"`
	IsFixed string `json:"isFixed,omitempty"`
}

// GetValueObject 判断Value是否是结构体，并且返回结构体数据
func (p *Parameter) GetValueObject() (*valueObject, bool) {
	v, ok := p.Value.(*valueObject)
	if !ok {
		return nil, false
	}

	return v, true
}

// GetValueString 判断Value是否是字符串，并且返回字符串数据
func (p *Parameter) GetValueString() (string, bool) {
	v, ok := p.Value.(string)
	if !ok {
		return "", false
	}

	return v, true
}

// 判断2个实验参数是否完全相同
func IsComponentParameterListSame(list1, list2 *ComponentParameterList) bool {
	return reflect.DeepEqual(list1, list2)
}

// 判断2个组件参数是否完全相同
func IsComponentParameterSame(cp1, cp2 *ComponentParameter) bool {
	return reflect.DeepEqual(cp1, cp2)
}

// GetExperimentIdList 查询数据库中仿真结果对应的的experiment记录, 按rawRecordIdList中的顺序返回对应的experiment id
func GetExperimentIdList(rawRecordIdList []string, simulateRecordList []DataBaseModel.YssimSimulateRecord) []string {
	experimentIdList := []string{}
	for _, simulateRecordId := range rawRecordIdList {
		for _, simulateRecord := range simulateRecordList {
			if simulateRecordId == simulateRecord.ID {
				experimentIdList = append(experimentIdList, simulateRecord.ExperimentId)
			}
		}
	}

	return experimentIdList
}

// UnmarshalExperimentRecordList 将数据库实验记录中的组件参数ModelVarData(json类型)转化为ComponentParameterList结构体
func UnmarshalExperimentRecordList(experimentRecordList []DataBaseModel.YssimExperimentRecord) (map[string]*ComponentParameterList, error) {
	experimentRecordMap := map[string]*ComponentParameterList{}
	for _, experimentRecord := range experimentRecordList {
		componentParams := ComponentParameterList{}
		if err := json.Unmarshal(experimentRecord.ModelVarData, &componentParams.Data); err != nil {
			return nil, err
		}
		experimentRecordMap[experimentRecord.ID] = &componentParams
	}

	return experimentRecordMap, nil
}

// GetComponentAndParametersFromRecord 获取数据库实验记录保存的所有组件名称和组件参数
func GetComponentAndParametersFromRecord(experimentRecordMap map[string]*ComponentParameterList) (map[string]bool, map[string][]string) {
	components := make(map[string]bool)
	componentParameters := make(map[string][]string)

	for experimentId := range experimentRecordMap {
		length := len(experimentRecordMap[experimentId].Data)
		for index := 0; index < length; {
			// 如果实验记录中某组件不包含parameters字段，则直接删除该组件数据
			name, parameters := experimentRecordMap[experimentId].Data[index].Name, experimentRecordMap[experimentId].Data[index].Parameters
			if parameters == nil {
				experimentRecordMap[experimentId].Data = append(experimentRecordMap[experimentId].Data[:index], experimentRecordMap[experimentId].Data[index+1:]...)
				length = len(experimentRecordMap[experimentId].Data)
				continue
			}

			// 获取实验记录中组件的参数名称列表
			componentParameters[name] = make([]string, 0)
			for _, parameter := range parameters {
				componentParameters[name] = append(componentParameters[name], parameter.Name)
			}

			components[name] = true
			index += 1
		}
	}

	return components, componentParameters
}

// GetComponentParameterFromResultXml 获取仿真结果文件中所有组件名称和组件参数
func GetComponentParameterFromResultXml(experimentIdList []string, components map[string]bool, simulateResultMap map[string]string) map[string]map[string]map[string]any {
	// 与仿真结果文件交互获取每一个实验中每个组件的参数数据
	parametersXmlMap := map[string]*ComponentParameterList{}
	for _, experimentId := range experimentIdList {
		for component := range components {
			parameterXml := SimulateResultParameters(simulateResultMap[experimentId], component, "")
			parameterXmlByte, err := json.Marshal(parameterXml)
			if err != nil {
				return nil
			}
			componentParameter := &ComponentParameter{Name: component, Parameters: make([]*Parameter, 0)}
			if err := sonic.Unmarshal(parameterXmlByte, &(componentParameter.Parameters)); err != nil {
				return nil
			}

			if _, ok := parametersXmlMap[experimentId]; !ok {
				parametersXmlMap[experimentId] = &ComponentParameterList{Data: make([]*ComponentParameter, 0)}
			}
			parametersXmlMap[experimentId].Data = append(parametersXmlMap[experimentId].Data, componentParameter)
		}
	}

	// 分解出所有参数
	compareData := map[string]map[string]map[string]any{}
	for experimentId := range parametersXmlMap {
		if _, ok := compareData[experimentId]; !ok {
			compareData[experimentId] = make(map[string]map[string]any)
		}

		for _, component := range parametersXmlMap[experimentId].Data {
			if _, ok := compareData[experimentId][component.Name]; !ok {
				compareData[experimentId][component.Name] = make(map[string]any)
			}
			for _, parameter := range component.Parameters {
				compareData[experimentId][component.Name][parameter.Name] = parameter.Value
			}
		}
	}

	return compareData
}

// GetTableColumn 生成发送给前端的实验对比数据表头
func GetTableColumn(experimentIdList []string, experimentMap map[string]DataBaseModel.YssimExperimentRecord) []map[string]string {
	tableColumns := []map[string]string{}
	for _, experimentId := range experimentIdList {
		data := map[string]string{"key": experimentId, "name": experimentMap[experimentId].ExperimentName}
		tableColumns = append(tableColumns, data)
	}

	return tableColumns
}

// GetDifferentParameter 对比出不同的参数
func GetDifferentParameter(components map[string]bool, componentParameters map[string][]string, experimentRecordList []DataBaseModel.YssimExperimentRecord, compareData map[string]map[string]map[string]any) []map[string]any {
	tableData := []map[string]any{}
	for component := range components {
		singleComponentAllDifferentParameters := map[string]any{}
		for _, key := range componentParameters[component] {
			isDuplicated := false
			for i := 0; i < len(experimentRecordList)-1; i++ {
				if !reflect.DeepEqual(compareData[experimentRecordList[i].ID][component][key], compareData[experimentRecordList[i+1].ID][component][key]) {
					isDuplicated = true
					singleComponentAllDifferentParameters["name"] = component
				}
			}

			if isDuplicated {
				for j := 0; j < len(experimentRecordList); j++ {
					if _, ok := singleComponentAllDifferentParameters[experimentRecordList[j].ID]; !ok {
						singleComponentAllDifferentParameters[experimentRecordList[j].ID] = map[string]any{}
					}
					singleComponentAllDifferentParameters[experimentRecordList[j].ID].(map[string]any)[key] = compareData[experimentRecordList[j].ID][component][key]
				}
			}
		}
		if len(singleComponentAllDifferentParameters) != 0 {
			tableData = append(tableData, singleComponentAllDifferentParameters)
		}
	}

	return tableData
}

// GetComponentAndClassFromDB 获取所有组件名称和组件类型
func GetComponentAndClassFromDB(experimentRecordMap map[string]*ComponentParameterList) (map[string]map[string]bool, map[string]string) {
	components := make(map[string]map[string]bool)
	componentClasses := make(map[string]string)
	for experimentId := range experimentRecordMap {
		length := len(experimentRecordMap[experimentId].Data)
		for index := 0; index < length; {
			// 如果实验参数中某组件不包含parameters字段，则删除该组件数据，后续再从omc中获取
			name, parameters := experimentRecordMap[experimentId].Data[index].Name, experimentRecordMap[experimentId].Data[index].Parameters
			if parameters == nil {
				experimentRecordMap[experimentId].Data = append(experimentRecordMap[experimentId].Data[:index], experimentRecordMap[experimentId].Data[index+1:]...)
				length = len(experimentRecordMap[experimentId].Data)
				continue
			}

			if _, ok := components[name]; !ok {
				components[name] = map[string]bool{}
			}
			components[name][experimentId] = true
			componentClasses[name] = parameters[0].ExtendName
			index += 1
		}
	}

	return components, componentClasses
}

// GetComponentParameterFromOMC 与omc交互补全任一实验中缺失的组件及参数并分解出所有参数
func GetComponentParameterFromOMC(modelName string, components map[string]map[string]bool, experimentRecordMap map[string]*ComponentParameterList, componentClasses map[string]string) (map[string]map[string]map[string]any, map[string]bool) {
	// 与omc交互补全每个实验中缺失的组件及参数
	for experimentId := range experimentRecordMap {
		for component := range components {
			if _, ok := components[component][experimentId]; !ok {
				var parameterOMC []any
				if component == modelName {
					parameterOMC = GetModelParameters(modelName, "", componentClasses[component], "")
				} else {
					parameterOMC = GetModelParameters(modelName, component, componentClasses[component], "")
				}
				parameterOMCByte, err := json.Marshal(parameterOMC)
				if err != nil {
					return nil, nil
				}
				componentParameter := &ComponentParameter{Name: component, Parameters: make([]*Parameter, 0)}
				if err := sonic.Unmarshal(parameterOMCByte, &(componentParameter.Parameters)); err != nil {
					return nil, nil
				}
				experimentRecordMap[experimentId].Data = append(experimentRecordMap[experimentId].Data, componentParameter)
			}
		}
	}

	// 分解出所有参数
	compareData := map[string]map[string]map[string]any{}
	keys := map[string]bool{}
	for experimentId := range experimentRecordMap {
		if _, ok := compareData[experimentId]; !ok {
			compareData[experimentId] = make(map[string]map[string]any)
		}

		for _, component := range experimentRecordMap[experimentId].Data {
			if _, ok := compareData[experimentId][component.Name]; !ok {
				compareData[experimentId][component.Name] = make(map[string]any)
			}
			for _, parameter := range component.Parameters {
				compareData[experimentId][component.Name][parameter.Name] = parameter.Value
				keys[parameter.Name] = true
			}
		}
	}

	return compareData, keys
}

// GetDifferentParameterOld 对比出不同的参数
func GetDifferentParameterOld(components map[string]map[string]bool, keys map[string]bool, experimentRecordList []DataBaseModel.YssimExperimentRecord, compareData map[string]map[string]map[string]any) []map[string]any {
	tableData := []map[string]any{}

	// 依次比较所有的参数，找出不同
	for component := range components {
		singleComponentAllDifferentParameters := map[string]any{}
		for key := range keys {
			isDuplicated := false
			for i := 0; i < len(experimentRecordList)-1; i++ {
				if !reflect.DeepEqual(compareData[experimentRecordList[i].ID][component][key], compareData[experimentRecordList[i+1].ID][component][key]) {
					isDuplicated = true
					singleComponentAllDifferentParameters["name"] = component
				}
			}

			if isDuplicated {
				for j := 0; j < len(experimentRecordList); j++ {
					if _, ok := singleComponentAllDifferentParameters[experimentRecordList[j].ID]; !ok {
						singleComponentAllDifferentParameters[experimentRecordList[j].ID] = map[string]any{}
					}
					singleComponentAllDifferentParameters[experimentRecordList[j].ID].(map[string]any)[key] = compareData[experimentRecordList[j].ID][component][key]
				}
			}
		}
		if len(singleComponentAllDifferentParameters) != 0 {
			tableData = append(tableData, singleComponentAllDifferentParameters)
		}
	}

	return tableData
}
