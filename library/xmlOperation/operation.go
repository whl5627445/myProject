package xmlOperation

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/beevik/etree"
)

func ParseXML(path string, obj any) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.New("文件打开错误: " + err.Error())
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("读取消息错误错误：" + err.Error())
	}
	err = xml.Unmarshal(data, obj)
	return nil
}

func GetVarXml(orderedVariables *etree.Element, parent string, keyWords string, id int, nameMap map[string]bool) ([]map[string]any, int, map[string]bool) {
	var dataList []map[string]any
	parentName := ""
	if parent != "" {
		parentName = parent + "."
	}
	if orderedVariables != nil {
		for _, variable := range orderedVariables.SelectElements("variable") {
			//if variable.SelectAttrValue("type", "") != "Real" {
			//	continue
			//}
			name := variable.SelectAttrValue("name", "")
			var splitName []string
			trimPrefixName := strings.TrimPrefix(name, parent+".")
			if strings.HasPrefix(name, parentName) && strings.Contains(strings.ToLower(name), strings.ToLower(keyWords)) {
				if !strings.HasPrefix(name, "der(") && !strings.HasPrefix(name, "$") {
					splitName = strings.Split(trimPrefixName, ".")
				} else {
					continue
				}
				displayUnitString := ""
				unitString := ""
				startString := ""
				if attributesValues := variable.SelectElement("attributesValues"); attributesValues != nil {
					if displayUnit := attributesValues.SelectElement("displayUnit"); displayUnit != nil {
						displayUnitString = strings.Replace(displayUnit.SelectAttrValue("string", ""), "\"", "", -1)
					}
					if unit := attributesValues.SelectElement("unit"); unit != nil {
						unitString = strings.Replace(unit.SelectAttrValue("string", ""), "\"", "", -1)
					}
				}
				if bindExpression := variable.SelectElement("bindExpression"); bindExpression != nil {
					startString = bindExpression.SelectAttrValue("string", "")
				}
				if !nameMap[splitName[0]] {
					data := map[string]any{
						"variables":    splitName[0],
						"description":  variable.SelectAttrValue("comment", ""),
						"display_unit": displayUnitString,
						"has_child":    false,
						"id":           id,
						"start":        startString,
						"unit":         unitString,
					}
					if len(splitName) > 1 {
						data["has_child"] = true
						data["unit"] = ""
						data["display_unit"] = ""
					}
					id += 1
					nameMap[splitName[0]] = true
					dataList = append(dataList, data)
				}
			}
		}
	}
	return dataList, id, nameMap
}
