package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yssim-go/Init"
	"yssim-go/library/omc"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func CheckOMC() gin.HandlerFunc {
	return func(context *gin.Context) {
		checkOMC := omc.OMC.IsPackage("Modelica")
		if !checkOMC {
			omc.OMC = omc.OmcInit()
			Init.ModelLibraryInit()
		}
		context.Next()
	}
}
