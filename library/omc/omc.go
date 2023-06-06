package omc

import (
	"log"
	"os/exec"
	"sync"
	"time"
	"yssim-go/config"

	"github.com/sirupsen/logrus"
)

func OmcInit() *ZmqObject {
	addr := "127.0.0.1"
	port := "23456"
	omcInit, _ := Connect(addr, port)
	return &omcInit
}

var OMC *ZmqObject

type instance struct {
	Start   bool
	Cmd     *exec.Cmd
	UseTime *time.Time
	Mu      sync.Mutex
}

var OMCInstance instance

func StartOMC(result chan bool) {
	if OMCInstance.Cmd != nil {
		return
	}
	OMCInstance.Mu.Lock()
	cmd := exec.Command("omc", "--interactive=zmq", "--locale=C", "-z=omc", "--interactivePort=23456")
	err := cmd.Start()
	if err != nil {
		result <- false
		logrus.Println("启动OMC实例失败， 错误： ", err)
	}
	OMCInstance.Start = true
	OMCInstance.Cmd = cmd
	UseTime := time.Now().Local()
	OMCInstance.UseTime = &UseTime
	OMC = OmcInit()
	OMCInstance.Mu.Unlock()
	result <- true
	err = cmd.Wait()
	if err != nil {
		log.Println("omc wait 出错：", err)
	}
	StopOMC()
}

func StopOMC() {

	if OMCInstance.Start == false {
		return
	}
	if OMCInstance.Cmd != nil {
		OMCInstance.Cmd.Process.Kill()
		OMCInstance.Cmd = nil
	}
	OMCInstance.Start = false
	OMC = nil
	config.UserSpaceId = ""
	log.Println("omc进程已停止")
	return
}
