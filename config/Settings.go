package config

import "os"

var ModelicaKeywords = map[string]bool{"der": true, "and": true, "or": true, "not": true, "constant": true}

var USERNAME = os.Getenv("USERNAME")

var ParameterTranslation = map[string]string{
	"Initialization": "初始化",
	"General":        "通用设置",
	"Advanced":       "高级设置",
	"Attributes":     "属性设置",
	"Parameters":     "参数",
	"Modifiers":      "Modifiers",
	"Dummy":          "Dummy",
	"Component":      "组件",
	"Name":           "名称",
	"comment":        "注释",
}
