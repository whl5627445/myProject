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
		Models.POST("/space/collect", API.AppSpaceCollectView)

		Models.GET("/page/get", API.GetAppPageView)
		Models.POST("/page/create", API.CreateAppPageView)
		Models.POST("/page/edit", API.EditAppPageView)
		Models.POST("/page/delete", API.DeleteAppPageView)

		Models.GET("/page/components/get", API.GetPageComponentView)
		Models.POST("/page/components/create", API.CreatePageComponentView)
		Models.POST("/page/components/edit", API.EditPageComponentView)
		Models.POST("/page/components/delete", API.DeletePageComponentView)
	}
}