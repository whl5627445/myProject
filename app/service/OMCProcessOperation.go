package service

import "yssim-go/library/omc"

func StartOMC() bool {
	if omc.OMCInstance.Start {
		return true
	}
	result := make(chan bool)
	go omc.StartOMC(result)
	return <-result
}

func StopOMC() {
	omc.StopOMC()
}
