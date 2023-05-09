package router

import (
	FileAPI "yssim-go/app/api/file"

	"github.com/gin-gonic/gin"
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
		Models.POST("/result/filter/get", FileAPI.GetFilterResultFileView)
		Models.POST("/fmu/export", FileAPI.FmuExportModelView)
		Models.POST("/code/save", FileAPI.ModelCodeSaveView)
		Models.POST("/upload/var", FileAPI.UploadModelVarFileView)

		Models.POST("/package/resources/download/", FileAPI.DownloadResourcesFileView)
		Models.POST("/package/resources/get", FileAPI.GetPackageResourcesList)
		Models.POST("/package/resources/upload", FileAPI.UploadResourcesFileView)
		Models.POST("/package/resources/dir/create", FileAPI.CreateResourcesDirView)
		Models.POST("/package/resources/dir/delete", FileAPI.DeleteResourcesDirAndFileView)
		Models.POST("/package/resources/images/path/get", FileAPI.ResourcesImagesPathGetView)
		Models.GET("/package/resources/images/get", FileAPI.ResourcesImagesGetView)

		Models.POST("/icon/set", FileAPI.ModelIconSetView)

		Models.POST("/background/upload", FileAPI.BackgroundUploadView)
	}
}
