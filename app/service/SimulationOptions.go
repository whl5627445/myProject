package service

import "yssim-go/library/omc"

func GetSimulationOptions(modelName string) map[string]string {
	data := omc.OMC.GetSimulationOptions(modelName)
	return map[string]string{
		"startTime":         data[0],
		"stopTime":          data[1],
		"tolerance":         data[2],
		"numberOfIntervals": data[3],
		"interval":          data[4],
		"method":            data[5],
	}
}

func SetSimulationOptions(modelName, startTime, stopTime, interval, simulationFlags string) bool {
	annotate0 := "experiment(StartTime=" + startTime + ",StopTime=" + stopTime + ",Interval=" + interval + ")"
	data0 := omc.OMC.AddClassAnnotation(modelName, annotate0)
	annotate1 := "__OpenModelica_simulationFlags(solver=\"" + simulationFlags + "\")"
	data1 := omc.OMC.AddClassAnnotation(modelName, annotate1)
	return data0 && data1
}
