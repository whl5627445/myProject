package omc

import (
	"fmt"

	"github.com/pebbe/zmq4"
)

var OMC, _ = Connect("127.0.0.1", "23456")

func Connect(addr string, port string) (*omcZMQ, error) {
	obj, _ := zmq4.NewSocket(zmq4.REQ)
	omcObj := &omcZMQ{obj}
	cmd := "tcp://" + addr + ":" + port
	err := obj.Connect(cmd)
	fmt.Println("OMC connect " + addr + ":" + port + ".")
	if err != nil {
		return omcObj, err
	}
	// _, err = obj.Send("loadModel(Modelica, {\"3.2.3\"},true,\"\",false)", 0)
	_, _ = obj.Send("setCommandLineOptions(\"-d=nogen,noevalfunc,newInst,nfAPI\")", 0)
	_, _ = obj.Send("loadFile(\"/home/simtek/yssim-go/Applications.mo\")", 0)
	_, _ = obj.Recv(0)
	if err != nil {
		return omcObj, err
	}
	return omcObj, nil
}
