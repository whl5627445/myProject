package router

import (
	API "yssim-go/app/api/file"

	"github.com/gin-gonic/gin"
)

func FileRouter(g *gin.Engine) {
	var Models = g.Group("/file")
	{
		Models.POST("/upload/package", API.UploadModelPackageView)
		Models.POST("/update/package", API.UpdateModelPackageView)
		Models.POST("/create/package", API.CreateModelPackageView)

		Models.POST("/upload/icon", API.UploadModelIconView)

		Models.GET("/package/list", API.GetPackageFileListView)
		Models.POST("/package/get", API.GetPackageFileView)
		Models.GET("/package/encrypt/export", API.GetPackageEncryptView)

		Models.POST("/result/all/get", API.GetResultFileView)
		Models.POST("/result/filter/get", API.GetFilterResultFileView)

		Models.POST("/fmu/export", API.FmuExportModelView)
		Models.POST("/code/save", API.ModelCodeSaveView)
		Models.POST("/upload/var", API.UploadModelVarFileView)

		Models.POST("/package/resources/download/", API.DownloadResourcesFileView)
		Models.POST("/package/resources/get", API.GetPackageResourcesList)
		Models.POST("/package/resources/upload", API.UploadResourcesFileView)
		Models.POST("/package/resources/dir/create", API.CreateResourcesDirView)
		Models.POST("/package/resources/dir/delete", API.DeleteResourcesDirAndFileView)
		Models.POST("/package/resources/images/path/get", API.ResourcesImagesPathGetView)
		Models.GET("/package/resources/images/get", API.ResourcesImagesGetView)

		Models.POST("/package/", API.ResourcesImagesPathGetView)

		Models.POST("/icon/set", API.ModelIconSetView)

	}
}
