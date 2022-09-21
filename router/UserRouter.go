package router

import (
	"github.com/gin-gonic/gin"
	UserAPI "yssim-go/app/api/user"
)

func UserRouter(g *gin.Engine) {
	var Models = g.Group("/user")
	{
		Models.GET("/userspace/get", UserAPI.GetUserSpaceView)
		Models.POST("/userspace/create", UserAPI.CreateUserSpaceView)
		Models.POST("/userspace/delete", UserAPI.DeleteUserSpaceView)
		Models.POST("/userspace/login", UserAPI.LoginUserSpaceView)
		Models.GET("/examples", UserAPI.ExamplesView)

		Models.GET("/userspace/recently", UserAPI.GetUserRecentlyOpenedView)
	}
}
