package router

import (
	API "yssim-go/app/v2/api/pipeNet"

	"github.com/gin-gonic/gin"
)

func PipeNetRouter(g *gin.Engine) {
	var Models = g.Group("/pipenet")
	{
		Models.POST("/mapping_config/upload", API.UploadMappingConfigView)
		Models.POST("/mapping_config/download", API.DownloadMappingConfigView)
		Models.POST("/mapping_config/delete", API.DeleteMappingConfigView)
		Models.POST("/mapping_config/copy", API.CopyMappingConfigView)
		Models.GET("/mapping_config/list/get", API.GetMappingConfigListView)
		Models.POST("/mapping_config/edit", API.EditMappingConfigView)
		Models.GET("/mapping_config/details/get", API.GetMappingConfigDetailsView)
		Models.POST("/mapping_config/details/edit", API.EditMappingConfigDetailsView)
	}
}
