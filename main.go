package main

import (
	_ "yssim-go/Init"
	"yssim-go/config"
	"yssim-go/middleware"
	"yssim-go/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//if os.Getenv("DEBUG") == "" {
	//	gin.SetMode(gin.ReleaseMode)
	//}
	g := gin.Default()
	g.Use(middleware.Cors())
	g.Static("/static", "./static")
	{
		router.UserRouter(g)
		router.AppDesignRouter(g)

		g.Use(middleware.CheckOMC())
		router.ModelRouter(g)
		router.SimulateRouter(g)
		router.FileRouter(g)
		router.ResourceLibRouter(g)
		router.PipeNetRouter(g)
	}
	g.Run(config.ADDR + config.PORT)
}
