package Init

import (
	"github.com/gin-gonic/gin"
)

func LogInit() {
	gin.ForceConsoleColor()
	//f, _ := os.Create("log/yssim.log")
	//e, _ := os.Create("log/yssimErr.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter = io.MultiWriter(e)
}
