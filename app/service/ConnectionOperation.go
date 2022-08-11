package service

import "yssim-go/library/omc"

func AddConnection() bool {
	return true
}

func UpdateConnectionNames() bool {
	return true
}

func UpdateConnectionAnnotation() bool {
	return true
}

func DeleteConnection(classNameAll, connectStart, connectEnd string) bool {
	result := omc.OMC.DeleteConnection(classNameAll, connectStart, connectEnd)
	return result
}
