package omc

func OmcInit() *ZmqObject {
	addr := "127.0.0.1"
	port := "23456"
	omcInit, _ := Connect(addr, port)
	return omcInit
}

var OMC = OmcInit()

//var OMC = func() ZmqObject {
//	omcInit, _ := Connect("127.0.0.1", "23456")
//	return omcInit
//}()
