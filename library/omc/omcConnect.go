package omc

import (
	"context"
	"github.com/go-zeromq/zmq4"
	"log"
)

func Connect(addr string, port string) (*omcZMQ, error) {
	obj := zmq4.NewReq(context.Background())
	err := obj.Dial("tcp://" + addr + ":" + port)
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	return &omcZMQ{obj}, nil
}
