package omc

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-zeromq/zmq4"
)

func Connect(addr string, port string) (ZmqObject, error) {
	obj := zmq4.NewReq(context.Background(), zmq4.WithDialerRetry(time.Second))
	err := obj.Dial("tcp://" + addr + ":" + port)
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	return ZmqObject{obj, sync.Mutex{}}, nil
}
