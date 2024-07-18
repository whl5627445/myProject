package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/bytedance/sonic"

	"github.com/go-zeromq/zmq4"
)

func OmcInit(port string) *ZmqObject {
	addr := "127.0.0.1"
	omcInit, _ := Connect(addr, port)
	return &omcInit
}

func Connect(addr string, port string) (ZmqObject, error) {
	obj := zmq4.NewReq(context.Background(), zmq4.WithDialerRetry(time.Second))
	addrPort := "tcp://" + addr + ":" + port
	log.Println("连接地址：", addrPort)
	n := 0
	for {
		err := obj.Dial(addrPort)
		if err != nil {
			log.Printf("could not dial: %v", err)
			time.Sleep(1 * time.Second)
			n += 1
			log.Printf("第 %d 次重试", n)
			if n == 20 {
				log.Printf("重试%d次未成功, 放弃连接，端口号是： %s", n, port)
				return ZmqObject{obj, sync.Mutex{}}, nil
			}
			continue
		}
		log.Println("omc初始化完成")
		break
	}
	return ZmqObject{obj, sync.Mutex{}}, nil
}

type ZmqObject struct {
	zmq4.Socket
	sync.Mutex
}

func (o *ZmqObject) SendExpression(cmd string) []byte {
	//s := time.Now().UnixNano() / 1e6
	var msg []byte
	o.Lock()
	log.Println("准备向omc发送指令", cmd)
	_ = o.Send(zmq4.NewMsgString(cmd))
	data, _ := o.Recv()
	msg = data.Bytes()
	log.Println("收到omc数据", string(msg))
	o.Unlock()
	return msg
}

func main() {

	server("23456")
	//  Socket to talk to clients

}
func server(port string) {
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}
	log.Println("准备启动。。。")
}

//func server(port string) {
//	rep := zmq4.NewRep(context.Background())
//	defer rep.Close()
//	err := rep.Listen("tcp://127.0.0.1:" + port)
//	if err != nil {
//		log.Printf("could not dial: %v", err)
//	}
//	log.Println("启动成功。。。")
//	for {
//		//  Wait for next request from client
//		msg, err := rep.Recv()
//		if err != nil {
//			log.Printf("could not recv request: %v", err)
//		}
//		var cData clientData
//		err = sonic.Unmarshal(msg.Frames[0], &cData)
//		if err != nil {
//			log.Printf("客户端数据错误: %#v", err)
//		}
//		log.Printf("客户端数据: %#v", cData)
//		omc := OmcInit(cData.Port)
//		data := omc.SendExpression(cData.Cmd)
//		err = rep.Send(zmq4.NewMsgString(string(data)))
//		if err != nil {
//			log.Printf("发送omc返回数据出现错误: %v", err)
//		}
//		err = omc.Close()
//		if err != nil {
//			return
//		}
//		if err != nil {
//			log.Printf("关闭连接出现错误: %v", err)
//		}
//	}
//}

type clientData struct {
	Cmd  string `json:"cmd,omitempty"`
	Port string `json:"port,omitempty"`
}

func process(conn net.Conn) {
	// 处理完关闭连接
	defer conn.Close()

	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [4096]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		recv := buf[:n]
		fmt.Printf("收到的数据：%v\n", string(recv))
		var cData clientData
		err = sonic.Unmarshal(recv, &cData)
		if err != nil {
			log.Printf("客户端数据错误: %#v", err)
		}
		log.Printf("客户端数据: %#v", cData)
		omc := OmcInit(cData.Port)
		data := omc.SendExpression(cData.Cmd)
		if err != nil {
			log.Printf("发送omc返回数据出现错误: %v", err)
		}
		err = omc.Close()
		if err != nil {
			return
		}
		if err != nil {
			log.Printf("关闭连接出现错误: %v", err)
		}
		// 将接受到的数据返回给客户端
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}
