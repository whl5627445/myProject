package service

import (
	"strings"
	"yssim-go/app/DataBaseModel"
	"yssim-go/library/omc"
)

func SearchModel(model DataBaseModel.YssimModels, keyWords, parentNode string) []map[string]interface{} {
	var modelNameList []map[string]interface{}
	var nodeNames []string
	searchModelMap := map[string]bool{}
	words := strings.ToLower(keyWords)
	parentNodePackageList := strings.Split(parentNode, ".")

	if parentNode == "" {
		nodeNames = append(nodeNames, model.PackageName)
	} else {
		nodeNames = omc.OMC.GetClassNames(parentNode, false)
	}
	for i := 0; i < len(nodeNames); i++ {
		n := nodeNames[i]
		if parentNode != "" {
			n = parentNode + "." + nodeNames[i]
		}
		modelNames := omc.OMC.GetClassNames(n, true)
		for _, name := range modelNames {
			wordsIndex := strings.Index(strings.ToLower(name), words)
			if wordsIndex != -1 {
				nameListAll := strings.Split(name, ".")
				shortName := nameListAll[0]
				nameParent := shortName
				if parentNode != "" {
					nameParent = strings.Join(nameListAll[:len(parentNodePackageList)+1], ".")
					shortName = nameListAll[len(parentNodePackageList)]
				}
				_, ok := searchModelMap[nameParent]
				if !ok {
					searchModelMap[nameParent] = true
					modelType := omc.OMC.GetClassRestriction(nameParent)
					data := map[string]interface{}{
						"name":         shortName,
						"model_name":   nameParent,
						"haschild":     false,
						"type":         modelType,
						"image":        "",
						"package_id":   model.ID,
						"package_name": model.PackageName,
						"sys_user":     model.SysUser,
						//"package_version": model.Version,
					}
					childList := omc.OMC.GetClassNames(nameParent, false)
					if len(childList) > 0 {
						data["haschild"] = true
					}
					modelNameList = append(modelNameList, data)
					break
				}
			}
		}
	}
	return modelNameList
}

func SearchFunctionType(parentNode string) []map[string]interface{} {
	var modelNameList []map[string]interface{}
	var nodeNames []string
	searchModelMap := map[string]bool{}
	parentNodePackageList := strings.Split(parentNode, ".")

	nodeNames = omc.OMC.GetClassNames(parentNode, false)
	for i := 0; i < len(nodeNames); i++ {
		n := nodeNames[i]
		if parentNode != "" {
			n = parentNode + "." + nodeNames[i]
		}
		modelNames := omc.OMC.GetClassNames(n, true)
		for _, name := range modelNames {
			modelType := omc.OMC.GetClassRestriction(name)
			if modelType == "function" {
				nameListAll := strings.Split(name, ".")
				shortName := nameListAll[0]
				nameParent := shortName
				if parentNode != "" {
					nameParent = strings.Join(nameListAll[:len(parentNodePackageList)+1], ".")
					shortName = nameListAll[len(parentNodePackageList)]
				}
				_, ok := searchModelMap[nameParent]
				if !ok {
					searchModelMap[nameParent] = true

					data := map[string]interface{}{
						"name":       shortName,
						"model_name": nameParent,
						"haschild":   false,
						"type":       modelType,
						"image":      "",
						//"package_version": model.Version,
					}
					childList := omc.OMC.GetClassNames(nameParent, false)
					if len(childList) > 0 {
						data["haschild"] = true
					}
					modelNameList = append(modelNameList, data)

					break
				}
			}
		}
	}
	return modelNameList
}

func SearchPackage(model DataBaseModel.YssimModels, keyWords string) map[string]interface{} {
	words := strings.ToLower(keyWords)
	modelNames := omc.OMC.GetClassNames(model.PackageName, true)
	for _, name := range modelNames {
		wordsIndex := strings.Index(strings.ToLower(name), words)
		if wordsIndex != -1 {
			nameListAll := strings.Split(name, ".")
			shortName := nameListAll[0]
			nameParent := shortName
			modelType := omc.OMC.GetClassRestriction(nameParent)
			data := map[string]interface{}{
				"name":            shortName,
				"model_name":      nameParent,
				"haschild":        false,
				"type":            modelType,
				"image":           "",
				"package_id":      model.ID,
				"package_name":    model.PackageName,
				"sys_user":        model.SysUser,
				"package_version": model.Version,
			}
			childList := omc.OMC.GetClassNames(nameParent, false)
			if len(childList) > 0 {
				data["haschild"] = true
			}
			return data
		}
	}
	return nil
}
