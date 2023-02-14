package service

import (
	"strings"
	"yssim-go/library/omc"
)

func SearchModel(keyWords, parentNode string) []interface{} {
	modelNameList := []interface{}{}
	packageNames := []string{}
	words := strings.ToLower(keyWords)
	parentNodePackageList := strings.Split(parentNode, ".")
	if parentNode == "" {
		packageNames = omc.OMC.GetPackages()
	} else {
		packageNames = parentNodePackageList[:1]
	}
	var modelNameAll []string
	searchModelMap := map[string]bool{}
	for i := 0; i < len(packageNames); i++ {
		modelNames := omc.OMC.GetClassNames(packageNames[i], true)
		modelNameAll = append(modelNameAll, modelNames...)
	}
	for _, name := range modelNameAll {
		wordsIndex := strings.Index(strings.ToLower(name), words)
		parentIndex := strings.HasPrefix(name, parentNode)
		if wordsIndex != -1 && parentIndex {
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

				classInformation := GetClassInformation(nameParent)
				modelType := strings.TrimSpace(classInformation[0].(string))
				data := map[string]interface{}{
					"name":       shortName,
					"model_name": nameParent,
					"haschild":   false,
					"type":       modelType,
					"image":      "",
				}
				childList := omc.OMC.GetClassNames(nameParent, false)
				if len(childList) > 0 {
					data["haschild"] = true
				}
				modelNameList = append(modelNameList, data)
			}
		}
	}
	return modelNameList

}
