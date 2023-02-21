package service

import (
	"strings"
	"yssim-go/library/omc"
)

func SearchModel(keyWords, parentNode string) []interface{} {
	var modelNameList []interface{}
	var nodeNames []string
	searchModelMap := map[string]bool{}
	words := strings.ToLower(keyWords)
	parentNodePackageList := strings.Split(parentNode, ".")

	if parentNode == "" {
		nodeNames = omc.OMC.GetPackages()
	} else {
		nodeNames = omc.OMC.GetClassNames(parentNode, false)
	}
	for i := 0; i < len(nodeNames); i++ {
		modelNames := omc.OMC.GetClassNames(parentNode+"."+nodeNames[i], true)
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
					break
				}
			}
		}
	}
	return modelNameList
}
