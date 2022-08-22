package service

import (
	"strings"
	"yssim-go/library/omc"
)

func CheckModel(className string) []map[string]string {
	message := omc.OMC.CheckModel(className)
	messageList := GetMessagesStringInternal()
	messageList = append(messageList, map[string]string{"type": "message", "message": message})
	return messageList
}

func GetMessagesStringInternal() []map[string]string {
	data := omc.OMC.GetMessagesStringInternal()
	if len(data) < 1 {
		return []map[string]string{}
	}
	dataList := strings.Split(data, ";,")
	var messageList []map[string]string
	for i := 0; i < len(dataList); i++ {
		dl := strings.Split(strings.TrimSpace(dataList[i]), ",\n")
		messageMap := make(map[string]string)
		for j := 0; j < len(dl); j++ {
			jl := strings.TrimSpace(dl[j])
			switch true {
			case strings.Index(jl, "MODELICAPATH") != -1 || strings.Index(jl, "installPackage") != -1:
				continue
			case strings.HasPrefix(jl, "message"):
				mes := strings.ReplaceAll(jl, "message = ", "")
				messageMap["message"] = mes[1 : len(mes)-1]
			case strings.HasPrefix(jl, "level"):
				level := strings.Split(jl, ".")
				messageMap["type"] = level[len(level)-1]
			}
		}
		if len(messageMap) > 0 {
			messageList = append(messageList, messageMap)
		}
	}
	return messageList
}
