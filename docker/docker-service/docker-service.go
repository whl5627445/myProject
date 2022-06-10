package main

import (
	"context"
	"net/http"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()
var cli, _ = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

func main() {
	r := gin.Default()
	r.POST("/create", func(c *gin.Context) {
		var cd ContainerData
		res := map[string]any{
			"ID": "",
			"result": true,
			"err":    nil,
		}
		_ = c.BindJSON(&cd)
		err := cd.verify(ctx, cli)
		if err != nil {
			res["result"] = false
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		id, err := cd.service(ctx, cli)
		if err != nil {
			res["result"] = false
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		res["ID"] = id
		c.JSON(http.StatusOK, res)

	})
	r.POST("/inquire", func(c *gin.Context) {
		var cd ContainerData
		res := map[string]any{
			"data": make([]interface{}, 0, 1),
			"err":  nil,
		}
		_ = c.BindJSON(&cd)
		data, err := cd.containerList(ctx, cli)
		res["data"] = data
		if err != nil {
			res["err"] = err.Error()
		}
		c.JSON(http.StatusOK, res)
	})
	r.POST("/containerstart", func(c *gin.Context) {
		var cd ContainerData
		res := map[string]any{
			"result": true,
			"err":    nil,
		}
		_ = c.BindJSON(&cd)
		err := cd.containerStart(ctx, cd.ContainerID, cli)
		if err != nil {
			res["result"] = false
			res["err"] = err.Error()
		}
		c.JSON(http.StatusOK, res)
	})
	r.POST("/containerstop", func(c *gin.Context) {
		var cd ContainerData
		res := map[string]any{
			"result": true,
			"err":    nil,
		}
		_ = c.BindJSON(&cd)
		err := cd.containerStop(ctx, cd.ContainerID, cli)
		if err != nil {
			res["result"] = false
			res["err"] = err.Error()
		}
		c.JSON(http.StatusOK, res)
	})
	r.POST("/containerremove", func(c *gin.Context) {
		var cd ContainerData
		res := map[string]any{
			"result": true,
			"err":    nil,
		}
		_ = c.BindJSON(&cd)
		err := cd.containerRemove(ctx, cd.ContainerID, cli)
		if err != nil {
			res["err"] = err.Error()
			res["result"] = false
		}
		c.JSON(http.StatusOK, res)
	})
	r.Run("0.0.0.0:36790") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
