package main

import (
	"github.com/gin-gonic/gin"
	_ "yssim-go/Init"
	"yssim-go/config"
	"yssim-go/middleware"
	"yssim-go/router"
)

func main() {

	g := gin.Default()
	g.Use(middleware.Cors())

	{
		router.ModelRouter(g)
		router.SimulateRouter(g)
		router.UserRouter(g)
		router.FileRouter(g)
	}

	g.Run(config.ADDR + config.PORT)
}
