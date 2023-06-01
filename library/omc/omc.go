package omc

import (
	"log"
	"os/exec"
	"time"

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
}

var OMCInstance instance

func StartOMC(result chan bool) {

	//log.Println("准备启动omc，杀死之前的残留")
	StopOMC()

	cmd := exec.Command("omc", "--interactive=zmq", "--locale=C", "-z=omc", "--interactivePort=23456")
	err := cmd.Start()
	//logrus.Printf("用户 %s 启动OMC实例", config.USERNAME)
	if err != nil {
		result <- false
		logrus.Println("启动OMC实例失败， 错误： ", err)
	}
	OMCInstance.Start = true
	OMCInstance.Cmd = cmd
	UseTime := time.Now().Local()
	OMCInstance.UseTime = &UseTime
	OMC = OmcInit()
	//logrus.Println("OMC实例连接成功")
	//logrus.Println("OMC实例启动完毕")
	result <- true
	err = cmd.Wait()
	if err != nil {
		log.Println("omc wait 出错：", err)
		return
	}
	StopOMC()
}

func StopOMC() {
	if OMCInstance.Cmd != nil {
		OMCInstance.Cmd.Process.Kill()
		OMCInstance.Cmd = nil
	}
	OMCInstance.Start = false
	OMC = nil
	log.Println("omc进程已停止")
	return
}
