package router

import (
	API "yssim-go/app/v2/api/resourceLib"

	"github.com/gin-gonic/gin"
)

func ResourceLibRouter(g *gin.Engine) {
	var Models = g.Group("/api/v2/resource")
	{
		Models.POST("/folder/create", API.CreateResourceFolderView)
		Models.POST("/file/upload", API.UploadResourceFileView)
		Models.GET("/file/content/get", API.GetResourceFileContentView)
		Models.GET("/file/content/parse", API.ParseResourceFileContentView)
		Models.POST("/edit", API.EditResourceInfoView)
		Models.POST("/delete", API.DeleteResourceView)
		Models.GET("/list/get", API.GetResourceListView)
		Models.GET("/search", API.SearchSubListView)
	}
}
