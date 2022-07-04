package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 这个是校验请求来源
	// 在这里我们不做校验，直接return true
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

var rdb = redis.NewClient(&redis.Options{
	Addr:        "redis:6379",
	Username:    "",
	Password:    "",
	DB:          0,
	PoolSize:    0,
	PoolTimeout: 0,
})
var ctx = context.Background()

func main() {
	fmt.Println(rdb.ClientID(ctx))
	engine := gin.Default()
	engine.GET("/notice/", func(context *gin.Context) {
		// 将普通的http GET请求升级为websocket请求
		client, _ := upgrader.Upgrade(context.Writer, context.Request, nil)
		context.Request.ParseForm()
		username := context.Request.Form.Get("username")
		for {
			req := map[string]interface{}{}
			err := client.ReadJSON(&req)
			if err != nil {
				return
			}
			// fmt.Println("message:", req)
			time.Sleep(time.Millisecond*100)
			rData, _ := rdb.RPop(ctx, username+"_"+"notification").Result()
			// if err != nil {
			// 	fmt.Println("err 0: ",err)
			// 	continue
			// }
			if rData != "" {
				var message Message
				err := json.Unmarshal([]byte(rData), &message)
				if err != nil {
					fmt.Println("message err: ", err)
					continue
				}
				var res = make(map[string]interface{})
				res["msg"] = time.Now().Format("2006-01-02 15:04:05") + ";" + message.Message
				res["status"] = true
				res["type"] = message.Type
				sendData, err := json.Marshal(res)
				if err != nil {
					log.Println("sendData: ", err)
				}
				err = client.WriteMessage(websocket.TextMessage, sendData)
				if err != nil {
					log.Println("send: ", err)
				}
			}
		}
	})

	err := engine.Run(":5555")
	if err != nil {
		log.Fatalln(err)
	}
}

// func main() {
// 	var conn = kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{"yssim-kafka:9092"},
// 		// Topic:     topic,
// 		Partition:      0,
// 		MinBytes:       1,    // 10KB
// 		MaxBytes:       10e6, // 10MB
// 		CommitInterval: 1,
// 		// MaxWait:   1,
// 	})
// 	// 	fmt.Println(rdb.ClientID(ctx))
// 	// 	engine := gin.Default()
// 	// 	engine.GET("/notice/", func(context *gin.Context) {
// 	// 		// 将普通的http GET请求升级为websocket请求
// 	// 		client, _ := upgrader.Upgrade(context.Writer, context.Request, nil)
// 	// 		for {
// 	// 			req := map[string]interface{}{}
// 	// 			err := client.ReadJSON(&req)
// 	// 			if err != nil {
// 	// 				return
// 	// 			}
// 	// 			fmt.Println("message:", req)
// 	// 			res := map[string]interface{}{"status": false, "msg": "", "type": "message"}
// 	// 			username, ok := req["user"]
// 	// 			if !ok {
// 	// 				client.WriteMessage(websocket.TextMessage, []byte("not found"))
// 	// 				continue
// 	// 			}
// 	// 			rData, err := rdb.RPop(ctx, username.(string)+"_"+"notification").Result()
// 	// 			if err != nil {
// 	// 				continue
// 	// 			}
// 	// 			if rData == "" {
// 	// 				sendData, err := json.Marshal(res)
// 	// 				err = client.WriteMessage(websocket.TextMessage, sendData)
// 	// 				if err != nil {
// 	// 					log.Println("send err: ", err)
// 	// 				}
// 	// 				continue
// 	// 			}
// 	// 			var message Message
// 	// 			err = json.Unmarshal([]byte(rData), &message)
// 	// 			if err != nil {
// 	// 				fmt.Println("message err: ", err)
// 	// 				continue
// 	// 			}
// 	// 			res["msg"] = time.Now().Format("2006-01-02 15:04:05") + ";" + message.Message
// 	// 			res["status"] = true
// 	// 			res["type"] = message.Type
// 	// 			sendData, err := json.Marshal(res)
// 	// 			if err != nil {
// 	// 				log.Println("sendData: ", err)
// 	// 			}
// 	// 			err = client.WriteMessage(websocket.TextMessage, sendData)
// 	// 			if err != nil {
// 	// 				log.Println("send: ", err)
// 	// 			}
// 	// 		}
// 	// 	})
// 	//
//
// 	engine := gin.Default()
// 	engine.GET("/notice/", func(ctx *gin.Context) {
// 		// 将普通的http GET请求升级为websocket请求
// 		client, _ := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
// 		ctx.Request.ParseForm()
// 		username := ctx.Request.Form.Get("username")
//
// 		topic := username + "_" + "notification"
// 		fmt.Println("topic: ", topic)
// 		for {
// 			res := map[string]interface{}{"status": false, "msg": "", "type": "message"}
// 			// req := map[string]interface{}{}
// 			// fmt.Println("message:", req)
// 			// client.ReadJSON(&req)
//
// 			for {
// 				MessageData, err := conn.ReadMessage(ctx)
//
// 				if err != nil {
// 					fmt.Println("err 0:", err)
// 					break
// 				}
// 				// fmt.Printf("mes: %s", MessageData)
// 				if MessageData.Value == nil {
// 					sendData, err := json.Marshal(res)
// 					fmt.Println("sendData: ", string(sendData))
//
// 					err = client.WriteMessage(websocket.TextMessage, sendData)
// 					if err != nil {
// 						log.Println("send err: ", err)
// 					}
// 					break
// 				}
// 				var message Message
// 				err = json.Unmarshal(MessageData.Value, &message)
// 				if err != nil {
// 					fmt.Println("message err: ", err)
// 					break
// 				}
// 				fmt.Printf("message: %s", message)
// 				res["msg"] = time.Now().Format("2006-01-02 15:04:05") + ";" + message.Message
// 				res["status"] = true
// 				res["type"] = message.Type
// 				sendData, err := json.Marshal(res)
// 				fmt.Printf("sendData: %s", sendData)
// 				if err != nil {
// 					log.Println("sendData: ", err)
// 					break
// 				}
// 				err = client.WriteMessage(websocket.TextMessage, sendData)
// 				if err != nil {
// 					log.Println("send: ", err)
// 					break
// 				}
// 			}
// 		}
// 	})
// 	err := engine.Run(":5555")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
