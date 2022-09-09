package omc

import (
	"context"
	"github.com/go-zeromq/zmq4"
	"log"
	"sync"
	"time"
)

func Connect(addr string, port string) (*ZmqObject, error) {
	obj := zmq4.NewReq(context.Background(), zmq4.WithDialerRetry(time.Second))
	err := obj.Dial("tcp://" + addr + ":" + port)
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	var mux sync.Mutex
	return &ZmqObject{obj, mux}, nil
}

//func Connect(addr string, port string) (ZmqObject, error) {
//	//obj := zmq4.NewReq(context.Background(), zmq4.WithDialerRetry(time.Second))
//	obj := goczmq.NewReqChanneler("tcp://" + addr + ":" + port)
//
//	return ZmqObject{obj}, nil
//}
