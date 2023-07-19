package middleware

import (
	"net/http"
	"strings"
	"time"
	"yssim-go/library/omc"

	"github.com/gin-gonic/gin"
)

func CheckOMC() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/static") && (!omc.OMCInstance.Start || omc.OMCInstance.Cmd == nil) {
			res := map[string]any{"data": nil, "msg": "", "err": "", "status": 3}
			c.AbortWithStatusJSON(http.StatusOK, res)
			return
		} else {
			UseTime := time.Now().Local()
			omc.OMCInstance.UseTime = &UseTime
			c.Next()
		}
	}
}
