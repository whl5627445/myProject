package service

import "yssim-go/library/omc"

func CheckModel(className string) []map[string]interface{} {
	message := omc.OMC.CheckModel(className)
	messageList := GetMessagesStringInternal()
	if message != "" {
		messageList = append(messageList, map[string]interface{}{"type": "message", "message": message})
	}

	for _, messageMap := range messageList {
		msg, ok := messageMap["message"]
		if !ok {
			continue
		}

		msgString, ok := msg.(string)
		if !ok {
			continue
		}
		messageMap["message"] = "[" + className + "]: " + msgString
	}

	return messageList
}
