package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	http "github.com/wangluozhe/chttp"
	"yssim-go/grpc/WorkSpace"
)

func CheckSMC() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.GetHeader("space_id")
		workSpace := WorkSpace.WS.Get(spaceID)
		if !strings.HasPrefix(c.Request.URL.Path, "/static") && (spaceID != "" && workSpace.SMC == nil) {
			res := map[string]any{"data": nil, "msg": "", "err": "", "status": 3}
			c.AbortWithStatusJSON(http.StatusOK, res)
			return
		} else {
			workSpace.LastTime = time.Now()
			c.Next()
		}
	}
}
