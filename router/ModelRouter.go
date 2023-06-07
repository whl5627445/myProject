package router

import (
	ModelAPI "yssim-go/app/api/model"

	"github.com/gin-gonic/gin"
)

func ModelRouter(g *gin.Engine) {
	var Models = g.Group("/model")
	{
		Models.GET("/root_library/sys", ModelAPI.GetSysRootModelView)
		Models.GET("/root_library/user", ModelAPI.GetUserRootModelView)
		Models.GET("/user/get", ModelAPI.GetUserPackageView)
		Models.GET("/list_library", ModelAPI.GetListModelView)

		Models.POST("/graphics", ModelAPI.GetGraphicsDataView)
		Models.POST("/icon/graphics", ModelAPI.GetGraphicsDataView)
		Models.POST("/icon/graphics/new", ModelAPI.GetIconView)
		Models.GET("/code", ModelAPI.GetModelCodeView)

		Models.GET("/parameters/get", ModelAPI.GetModelParametersView)
		Models.POST("/parameters/set", ModelAPI.SetModelParametersView)
		Models.POST("/parameters/add", ModelAPI.AddModelParametersView)
		Models.POST("/parameters/delete", ModelAPI.DeleteModelParametersView)
		Models.GET("/properties/get", ModelAPI.GetComponentPropertiesView)
		Models.POST("/properties/set", ModelAPI.SetComponentPropertiesView)

		Models.POST("/class/copy", ModelAPI.CopyClassView)

		Models.POST("/package/delete", ModelAPI.DeletePackageAndModelView)
		Models.POST("/package/load", ModelAPI.LoadModelView)
		Models.POST("/package/unload", ModelAPI.UnLoadModelView)
		Models.GET("/package/get/all", ModelAPI.GetPackageAndVersionView)

		Models.GET("/component/name", ModelAPI.GetComponentNameView)
		Models.POST("/component/add", ModelAPI.AddModelComponentView)
		Models.POST("/component/delete", ModelAPI.DeleteModelComponentView)
		Models.POST("/component/update", ModelAPI.UpdateModelComponentView)

		Models.POST("/connection/create", ModelAPI.CreateConnectionAnnotationView)
		Models.POST("/connection/delete", ModelAPI.DeleteConnectionAnnotationView)
		Models.POST("/connection/update", ModelAPI.UpdateConnectionAnnotationView)
		Models.POST("/connection/name", ModelAPI.UpdateConnectionNamesView)

		Models.GET("/exists", ModelAPI.ExistsView)
		Models.GET("/check", ModelAPI.CheckView)
		Models.GET("/components/get", ModelAPI.GetComponentsView)
		Models.GET("/document/get", ModelAPI.GetModelDocumentView)
		Models.POST("/document/set", ModelAPI.SetModelDocumentView)
		Models.POST("/units/convert", ModelAPI.ConvertUnitsView)

		Models.POST("/collection/create", ModelAPI.CreateCollectionModelView)
		Models.GET("/collection/get", ModelAPI.GetCollectionModelView)
		Models.GET("/collection/delete", ModelAPI.DeleteCollectionModelView)
		Models.GET("/search", ModelAPI.SearchModelView)

		Models.POST("/reference/resources", ModelAPI.GetModelResourcesReferenceView)

		Models.POST("/userspace/login", ModelAPI.LoginUserSpaceView)
		Models.POST("/mark", ModelAPI.AppModelMarkView)

	}
	g.POST("/test", ModelAPI.Test)
	g.POST("/test1", ModelAPI.Test1)
}
