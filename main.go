package main

import (
	"github.com/gin-gonic/gin"
	"yssim-go/router"
)

func main() {
	g := gin.Default()
	{
		router.ModelRouter(g)
	}
	g.Run(":8912")
}
