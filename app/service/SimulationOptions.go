package service

import (
	"yssim-go/library/omc"
)

func GetSimulationOptions(modelName string) map[string]string {
	data := omc.OMC.GetSimulationOptions(modelName)
	return map[string]string{
		"startTime":         data[0],
		"stopTime":          data[1],
		"tolerance":         data[2],
		"numberOfIntervals": data[3],
		"interval":          data[4],
	}
}

func SetSimulationOptions(modelName, startTime, stopTime, interval string) bool {
	annotate := "experiment(StartTime=" + startTime + ",StopTime=" + stopTime + ",Interval=" + interval + ")"
	data := omc.OMC.AddClassAnnotation(modelName, annotate)
	return data
}
