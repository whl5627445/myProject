package service

import "yssim-go/library/omc"

func AddConnection(classNameAll, connectStart, connectEnd, color string, linePoints []string) bool {
	result := omc.OMC.AddConnection(classNameAll, connectStart, connectEnd, color, linePoints)
	omc.OMC.Save(classNameAll)
	return result
}

func UpdateConnectionNames(classNameAll, fromName, toName, fromNameNew, toNameNew string) bool {
	result := omc.OMC.UpdateConnectionNames(classNameAll, fromName, toName, fromNameNew, toNameNew)
	omc.OMC.Save(classNameAll)
	return result
}

func UpdateConnection(classNameAll, connectStart, connectEnd, color string, linePoints []string) bool {
	result := omc.OMC.UpdateConnectionAnnotation(classNameAll, connectStart, connectEnd, color, linePoints)
	omc.OMC.Save(classNameAll)
	return result
}

func DeleteConnection(classNameAll, connectStart, connectEnd string) bool {
	result := omc.OMC.DeleteConnection(classNameAll, connectStart, connectEnd)
	omc.OMC.Save(classNameAll)
	return result
}
