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
		"simulate_type":     data[6],
	}
}

func SetSimulationOptions(modelName, startTime, stopTime, interval, simulationFlags, SimulateType string) bool {
	annotate0 := "experiment(StartTime=" + startTime + ",StopTime=" + stopTime + ",Interval=" + interval + ")"
	data0 := omc.OMC.AddClassAnnotation(modelName, annotate0)
	annotate1 := "__OpenModelica_simulationFlags(solver=\"" + simulationFlags + "\")"
	data1 := omc.OMC.AddClassAnnotation(modelName, annotate1)
	annotate2 := "simulate_type(solver=\"" + SimulateType + "\")"
	data2 := omc.OMC.AddClassAnnotation(modelName, annotate2)
	return data0 && data1 && data2
}
