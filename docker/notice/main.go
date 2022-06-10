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

var rdb = redis.NewClient(&redis.Options{
	Addr:        "yssim-redis:6379",
	Username:    "",
	Password:    "",
	DB:          0,
	PoolSize:    0,
	PoolTimeout: 0,
})
var ctx = context.Background()

type Message struct {
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

func main() {
	fmt.Println(rdb.ClientID(ctx))
	engine := gin.Default()
	engine.GET("/notice/", func(context *gin.Context) {
		// 将普通的http GET请求升级为websocket请求
		client, _ := upgrader.Upgrade(context.Writer, context.Request, nil)
		for {
			req := map[string]interface{}{}
			err := client.ReadJSON(&req)
			if err != nil {
				return
			}
			fmt.Println("message:", req)
			res := map[string]interface{}{"status": false, "msg": "", "type": "message"}
			// 每隔两秒给前端推送一句消息“hello, WebSocket”
			username, ok := req["user"]
			if !ok {
				client.WriteMessage(websocket.TextMessage, []byte("not found"))
				continue
			}
			rData, err := rdb.RPop(ctx, username.(string)+"_"+"notification").Result()
			if err != nil {
				continue
			}
			if rData == "" {
				sendData, err := json.Marshal(res)
				err = client.WriteMessage(websocket.TextMessage, sendData)
				if err != nil {
					log.Println("send err: ", err)
				}
				continue
			}
			var message Message
			err = json.Unmarshal([]byte(rData), &message)
			if err != nil {
				fmt.Println("message err: ", err)
				continue
			}
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
	})

	err := engine.Run(":5555")
	if err != nil {
		log.Fatalln(err)
	}
}
