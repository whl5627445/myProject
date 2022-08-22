package router

import (
	"github.com/gin-gonic/gin"
	UserAPI "yssim-go/app/api/user"
)

func UserRouter(e *gin.Engine) {
	var Models = e.Group("/user")
	{
		Models.GET("/userspace/get", UserAPI.GetUserSpaceView)
		Models.POST("/userspace/create", UserAPI.CreateUserSpaceView)
		Models.POST("/userspace/delete", UserAPI.DeleteUserSpaceView)
		Models.POST("/userspace/login", UserAPI.LoginUserSpaceView)
	}
}
