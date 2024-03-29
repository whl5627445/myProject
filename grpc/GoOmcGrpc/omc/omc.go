package omc

import (
	"github.com/sirupsen/logrus"
	"log"
	"os/exec"
	"sync"
	"time"
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
	time.Sleep(2 * time.Second)
	OMCInstance.Mu.Lock()
	if OMCInstance.Cmd != nil {
		result <- true
		logrus.Println("OMC实例已存在，无需重复启动")
		return
	}
	cmd := exec.Command("omc", "--interactive=zmq", "--locale=C", "-z=omc", "--interactivePort=23456")
	err := cmd.Start()
	if err != nil {
		result <- false
		logrus.Println("启动OMC实例失败， 错误： ", err)
		return
	}
	OMCInstance.Start = true
	OMCInstance.Cmd = cmd
	UseTime := time.Now().Local()
	OMCInstance.UseTime = &UseTime
	OMC = OmcInit()
	OMC.SetOptions()
	OMCInstance.Mu.Unlock()
	result <- true
	//libraryInitialization()
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
	OMCInstance.Mu.Lock()
	if OMCInstance.Cmd != nil {
		OMCInstance.Cmd.Process.Kill()
		OMCInstance.Cmd = nil
	}
	OMCInstance.Start = false
	OMC = nil
	OMCInstance.Mu.Unlock()
	//config.UserSpaceId = ""
	log.Printf("omc实例信息： %#v", OMCInstance)
	log.Println("omc进程已停止")
	return
}

//func libraryInitialization(LibraryMap map[string]map[string]string) {
//	OMC.SetOptions()
//	//log.Println("LibraryMap", config.LibraryMap)
//	for name, information := range LibraryMap {
//		version := information["version"]
//		ok := OMC.LoadModel(information["file"], "")
//		log.Printf("初始化模型库：%s %s  %t \n", name, version, ok)
//	}
//}
