package service

import "yssim-go/library/omc"

func CheckModel(className string) []map[string]string {
	message := omc.OMC.CheckModel(className)
	messageList := GetMessagesStringInternal()
	if message != "" {
		messageList = append(messageList, map[string]string{"type": "message", "message": message})
	}
	return messageList
}
