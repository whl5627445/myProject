package main

import (
	"github.com/gin-gonic/gin"
	_ "yssim-go/Init"
	"yssim-go/router"
)

func main() {
	g := gin.Default()
	{
		router.ModelRouter(g)
		router.SimulateRouter(g)
		router.UserRouter(g)

	}
	g.Run(":8912")
}
