package router

import (
	"github.com/gin-gonic/gin"
	FileAPI "yssim-go/app/api/file"
)

func FileRouter(g *gin.Engine) {
	var Models = g.Group("/file")
	{
		Models.POST("/upload/package", FileAPI.UploadModelPackageView)
		Models.POST("/update/package", FileAPI.UpdateModelPackageView)
		Models.POST("/create/package", FileAPI.CreateModelPackageView)
		Models.POST("/upload/icon", FileAPI.UploadModelIconView)
		Models.GET("/package/list", FileAPI.GetPackageFileListView)
		Models.POST("/package/get", FileAPI.GetPackageFileView)
		Models.POST("/result/all/get", FileAPI.GetResultFileView)
		Models.POST("/result/all/delete", FileAPI.DeleteResultFileView)
		Models.POST("/result/filter/get", FileAPI.GetFilterResultFileView)
		Models.POST("/fmu/export", FileAPI.FmuExportModelView)
		Models.POST("/code/save", FileAPI.ModelCodeSaveView)
	}
}
