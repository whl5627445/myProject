package router

import (
	UserAPI "yssim-go/app/api/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.Engine) {
	var Models = g.Group("/user")
	{
		Models.GET("/userspace/get", UserAPI.GetUserSpaceView)
		Models.POST("/userspace/create", UserAPI.CreateUserSpaceView)
		Models.POST("/userspace/delete", UserAPI.DeleteUserSpaceView)
		Models.POST("/userspace/login", UserAPI.LoginUserSpaceView)
		Models.GET("/examples", UserAPI.ExamplesView)
		Models.GET("/userspace/recent", UserAPI.GetUserRecentlyOpenedView)
		Models.GET("/settings/get", UserAPI.GetUserSettingsView)
		Models.POST("/settings/set", UserAPI.SetUserSettingsView)
	}
}
