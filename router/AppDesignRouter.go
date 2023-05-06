package router

import (
	API "yssim-go/app/api/appDesign"

	"github.com/gin-gonic/gin"
)

func AppDesignRouter(g *gin.Engine) {
	var Models = g.Group("/app")
	{
		Models.POST("/model/mark", API.AppModelMarkView)

		Models.GET("/space/get", API.GetAppSpaceView)
		Models.POST("/space/create", API.CreateAppSpaceView)
		Models.POST("/space/edit", API.EditAppSpaceView)
		Models.POST("/space/delete", API.DeleteAppSpaceView)

		Models.GET("/page/get", API.GetAppPageView)
		Models.POST("/page/create", API.CreateAppPageView)
		Models.POST("/page/edit", API.EditAppPageView)
		Models.POST("/page/delete", API.DeleteAppPageView)

	}
}
