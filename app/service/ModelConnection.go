package service

import "yssim-go/library/omc"

func AddConnection(classNameAll, connectStart, connectEnd, color string, linePoints []string) bool {
	result := omc.OMC.AddConnection(classNameAll, connectStart, connectEnd, color, linePoints)
	return result
}

func UpdateConnectionNames(classNameAll, fromName, toName, fromNameNew, toNameNew string) bool {
	result := omc.OMC.UpdateConnectionNames(classNameAll, fromName, toName, fromNameNew, toNameNew)
	return result
}

func UpdateConnection(classNameAll, connectStart, connectEnd, color string, linePoints []string) bool {
	result := omc.OMC.UpdateConnectionAnnotation(classNameAll, connectStart, connectEnd, color, linePoints)
	return result
}

func DeleteConnection(classNameAll, connectStart, connectEnd string) bool {
	result := omc.OMC.DeleteConnection(classNameAll, connectStart, connectEnd)
	return result
}
