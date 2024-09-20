package router

import (
	API "yssim-go/app/v2/api/pipeNet"

	"github.com/gin-gonic/gin"
)

func PipeNetRouter(g *gin.Engine) {
	var Models = g.Group("/pipenet/mapping_config")
	{
		Models.POST("/upload", API.UploadMappingConfigView)
		Models.POST("/download", API.DownloadMappingConfigView)
		Models.POST("/delete", API.DeleteMappingConfigView)
		Models.POST("/copy", API.CopyMappingConfigView)
		Models.GET("/list/get", API.GetMappingConfigListView)
		Models.POST("/edit", API.EditMappingConfigView)
		Models.GET("/details/get", API.GetMappingConfigDetailsView)
		Models.POST("/details/edit", API.EditMappingConfigDetailsView)

		Models.GET("/model/instance_mapping/get", API.GetInstanceMappingView)
		Models.GET("/model/instance_mapping/log/get", API.GetInstanceMappingLogView)
	}

	var ModelsV2 = g.Group("/pipenet/info")
	{
		ModelsV2.POST("/upload", API.UploadInfoFileView)
		ModelsV2.POST("/download", API.DownloadInfoFileView)
		ModelsV2.POST("/delete", API.DeleteInfoFileView)
		ModelsV2.GET("/list/get", API.GetInfoFileListView)
		ModelsV2.GET("/get", API.GetInfoView)

	}

	var ModelsV3 = g.Group("/pipenet/model")
	{
		ModelsV3.GET("/instance_mapping/get", API.GetInstanceMappingView)
		ModelsV3.GET("/instance_mapping/log/get", API.GetInstanceMappingLogView)
	}
}
