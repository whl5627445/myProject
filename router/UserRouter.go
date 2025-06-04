package router

import (
	API "yssim-go/app/v1/api/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.Engine) {
	var Models = g.Group("/user")
	{
		Models.GET("/userspace/get", API.GetUserSpaceView)
		Models.GET("/userspace/get/new", API.GetUserSpaceNewView)
		Models.POST("/userspace/create", API.CreateUserSpaceView)
		Models.POST("/userspace/edit", API.EditUserSpaceView)
		Models.POST("/userspace/delete", API.DeleteUserSpaceView)
		Models.GET("/userspace/recent", API.GetUserRecentlyOpenedView)

		Models.POST("/userspace/collect", API.CollectUserSpaceView)

		Models.GET("/examples", API.ExamplesView)

		Models.GET("/settings/get", API.GetUserSettingsView)
		Models.POST("/settings/set", API.SetUserSettingsView)

		Models.POST("/background/upload", API.BackgroundUploadView)

		Models.POST("/service/start", API.StartSMCView)
		Models.POST("/service/stop", API.StopSMCView)
		Models.POST("/service/restart", API.RestartSMCView)
	}
}
