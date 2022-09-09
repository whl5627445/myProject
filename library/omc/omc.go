package omc

var omcInit, _ = Connect("127.0.0.1", "23456")

var OMC = omcInit

//var OMC = func() ZmqObject {
//	omcInit, _ := Connect("127.0.0.1", "23456")
//	return omcInit
//}()
