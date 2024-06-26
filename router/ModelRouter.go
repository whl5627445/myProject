package router

import (
	API "yssim-go/app/v1/api/model"
	APIv2 "yssim-go/app/v2/api/model"

	"github.com/gin-gonic/gin"
)

func ModelRouter(g *gin.Engine) {
	var ModelsV1 = g.Group("/model")
	{
		ModelsV1.GET("/root_library/sys", API.GetSysRootModelView)
		ModelsV1.GET("/root_library/user", API.GetUserRootModelView)
		ModelsV1.GET("/user/get", API.GetUserPackageView)
		ModelsV1.GET("/list_library", API.GetListModelView)

		ModelsV1.POST("/graphics", API.GetGraphicsDataView)
		ModelsV1.POST("/icon/graphics", API.GetGraphicsDataView)
		ModelsV1.POST("/icon/graphics/new", API.GetIconView)
		ModelsV1.GET("/code", API.GetModelCodeView)
		ModelsV1.POST("/rename", API.ModelRename)

		ModelsV1.GET("/parameters/get", API.GetModelParametersView)
		ModelsV1.POST("/parameters/set", API.SetModelParametersView)
		ModelsV1.POST("/parameters/unit/set", API.SetModelParametersUnitView)
		ModelsV1.POST("/parameters/add", API.AddModelParametersView)
		ModelsV1.POST("/parameters/delete", API.DeleteModelParametersView)
		ModelsV1.GET("/properties/get", API.GetComponentPropertiesView)
		ModelsV1.POST("/properties/set", API.SetComponentPropertiesView)

		ModelsV1.POST("/class/copy", API.CopyClassView)

		ModelsV1.POST("/package/delete", API.DeletePackageAndModelView)
		ModelsV1.POST("/package/load", API.LoadModelView)
		ModelsV1.POST("/package/unload", API.UnLoadModelView)

		ModelsV1.GET("/component/name", API.GetComponentNameView)
		ModelsV1.POST("/component/add", API.AddModelComponentView)
		ModelsV1.POST("/component/delete", API.DeleteModelComponentView)
		ModelsV1.POST("/component/update", API.UpdateModelComponentView)
		ModelsV1.POST("/component/batch/update", API.BatchUpdateModelComponentView)

		ModelsV1.POST("/connection/create", API.CreateConnectionAnnotationView)
		ModelsV1.POST("/connection/delete", API.DeleteConnectionAnnotationView)
		ModelsV1.POST("/connection/update", API.UpdateConnectionAnnotationView)
		ModelsV1.POST("/connection/name", API.UpdateConnectionNamesView)

		ModelsV1.GET("/exists", API.ExistsView)
		ModelsV1.GET("/check", API.CheckView)
		ModelsV1.GET("/components/get", API.GetComponentsView)
		ModelsV1.GET("/document/get", API.GetModelDocumentView)
		ModelsV1.POST("/document/set", API.SetModelDocumentView)
		ModelsV1.POST("/units/convert", API.ConvertUnitsView)

		ModelsV1.POST("/collection/create", API.CreateCollectionModelView)
		ModelsV1.GET("/collection/get", API.GetCollectionModelView)
		ModelsV1.GET("/collection/delete", API.DeleteCollectionModelView)
		ModelsV1.GET("/search", API.SearchModelView)
		ModelsV1.GET("/function/search", API.SearchFunctionTypeView)

		ModelsV1.POST("/reference/resources", API.GetModelResourcesReferenceView)

		ModelsV1.POST("/userspace/login", API.LoginUserSpaceView)

		ModelsV1.GET("/uml/get", API.GetUMLView)

		ModelsV1.GET("/extend/get", API.GetExtendedModelView)

	}

	var ModelsV2 = g.Group("/api/v2/model")
	{
		ModelsV2.POST("/instance", APIv2.GetInstanceDataView)
		ModelsV2.POST("/component/add", APIv2.AddModelComponentView)
	}
	g.POST("/test", API.Test)
	g.POST("/test1", API.Test1)
}
