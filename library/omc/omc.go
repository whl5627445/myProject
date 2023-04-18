package omc

func OmcInit() ZmqObject {
	addr := "127.0.0.1"
	port := "23456"
	omcInit, _ := Connect(addr, port)
	return omcInit
}

var OMC = OmcInit()
