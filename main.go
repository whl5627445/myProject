package main

import (
	"github.com/gin-gonic/gin"
	_ "yssim-go/Init"
	"yssim-go/config"
	"yssim-go/middleware"
	"yssim-go/router"
)

func main() {
	//omcPross := Init.StartOMC()
	//defer omcPross.Process.Signal(syscall.SIGQUIT)

	g := gin.Default()

	g.Use(middleware.Cors())
	{
		router.ModelRouter(g)
		router.SimulateRouter(g)
		router.UserRouter(g)
		router.FileRouter(g)
	}

	//port, ok := os.LookupEnv("PORT")
	//if !ok {
	//	port = "8912"
	//}
	//port := "8913"

	g.Run(config.ADDR + config.PORT)
}
