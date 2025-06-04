package WorkSpace

import (
	"context"
	"errors"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"yssim-go/config"
	smc "yssim-go/grpc/SMC"
)

type WorkSpace struct {
	Name     string
	LastTime time.Time
	Id       string
	SMC      smc.SMCClient
}

type Management struct {
	ContainerMap map[string]*WorkSpace
	mu           *sync.Mutex
	api          WorkSpaceClient
}

var WS = newWorkSpaceGateway()

func newWorkSpaceGateway() *Management {
	conn, err := grpc.NewClient(config.WorkSpaceContainerManagerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(
		grpc.MaxCallSendMsgSize(1024*1024*50),
		grpc.MaxCallRecvMsgSize(1024*1024*50),
	))
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	conn.Connect()
	c := NewWorkSpaceClient(conn)
	return &Management{
		ContainerMap: map[string]*WorkSpace{},
		mu:           &sync.Mutex{},
		api:          c,
	}
}

func newWorkSpace(id, port string) *WorkSpace {
	s := smc.NewSMC(id + ":" + port)
	return &WorkSpace{
		Name:     id,
		SMC:      s,
		Id:       id,
		LastTime: time.Now(),
	}
}

func (w *Management) Get(id string) *WorkSpace {
	if workSpace, ok := w.ContainerMap[id]; ok {
		return workSpace
	}
	return nil
}

func (w *Management) Create(id string) (bool, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.remove(id)
	result, err := w.api.Create(context.Background(), &CreateRequest{Name: id, Port: "23790", ImageName: "smc"})
	if err != nil {
		return false, err
	}
	wk := newWorkSpace(result.GetId(), "23790")
	w.ContainerMap[result.GetId()] = wk
	return result.GetResult(), errors.New(result.Message)
}

func (w *Management) Inquire() {
	w.mu.Lock()
	defer w.mu.Unlock()
	result, err := w.api.Inquire(context.Background(), &InquireRequest{})
	if err != nil {
		return
	}
	for _, workspace := range result.GetWorkspaces() {
		wk := newWorkSpace(workspace.GetId(), "23790")
		w.ContainerMap[workspace.GetId()] = wk
	}
}

func (w *Management) Remove(id string) bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.remove(id)
}

func (w *Management) remove(id string) bool {
	remove, err := w.api.Remove(context.Background(), &RemoveRequest{Id: id})
	if err != nil {
		return false
	}
	if remove.GetResult() {
		delete(w.ContainerMap, id)
	}
	return remove.GetResult()
}
