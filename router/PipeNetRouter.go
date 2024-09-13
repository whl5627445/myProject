package router

import (
	API "yssim-go/app/v2/api/pipeNet"

	"github.com/gin-gonic/gin"
)

func PipeNetRouter(g *gin.Engine) {
	var Models = g.Group("/pipenet")
	{
		Models.POST("/mappingConfig/upload", API.UploadMappingConfigView)
		Models.POST("/mappingConfig/download", API.DownloadMappingConfigView)
		Models.POST("/mappingConfig/delete", API.DeleteMappingConfigView)
		Models.POST("/mappingConfig/copy", API.CopyMappingConfigView)
		Models.GET("/mappingConfig/list/get", API.GetMappingConfigListView)
		Models.POST("/mappingConfig/edit", API.EditMappingConfigView)
		Models.GET("/mappingConfig/details/get", API.GetMappingConfigDetailsView)
		Models.POST("/mappingConfig/details/edit", API.EditMappingConfigDetailsView)
	}
}
