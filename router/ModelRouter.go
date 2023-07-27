package router

import (
	API "yssim-go/app/api/model"

	"github.com/gin-gonic/gin"
)

func ModelRouter(g *gin.Engine) {
	var Models = g.Group("/model")
	{
		Models.GET("/root_library/sys", API.GetSysRootModelView)
		Models.GET("/root_library/user", API.GetUserRootModelView)
		Models.GET("/user/get", API.GetUserPackageView)
		Models.GET("/list_library", API.GetListModelView)

		Models.POST("/graphics", API.GetGraphicsDataView)
		Models.POST("/icon/graphics", API.GetGraphicsDataView)
		Models.POST("/icon/graphics/new", API.GetIconView)
		Models.GET("/code", API.GetModelCodeView)

		Models.GET("/parameters/get", API.GetModelParametersView)
		Models.POST("/parameters/set", API.SetModelParametersView)
		Models.POST("/parameters/add", API.AddModelParametersView)
		Models.POST("/parameters/delete", API.DeleteModelParametersView)
		Models.GET("/properties/get", API.GetComponentPropertiesView)
		Models.POST("/properties/set", API.SetComponentPropertiesView)

		Models.POST("/class/copy", API.CopyClassView)

		Models.POST("/package/delete", API.DeletePackageAndModelView)
		Models.POST("/package/load", API.LoadModelView)
		Models.POST("/package/unload", API.UnLoadModelView)
		Models.GET("/package/get/all", API.GetPackageAndVersionView)

		Models.GET("/component/name", API.GetComponentNameView)
		Models.POST("/component/add", API.AddModelComponentView)
		Models.POST("/component/delete", API.DeleteModelComponentView)
		Models.POST("/component/update", API.UpdateModelComponentView)

		Models.POST("/connection/create", API.CreateConnectionAnnotationView)
		Models.POST("/connection/delete", API.DeleteConnectionAnnotationView)
		Models.POST("/connection/update", API.UpdateConnectionAnnotationView)
		Models.POST("/connection/name", API.UpdateConnectionNamesView)

		Models.GET("/exists", API.ExistsView)
		Models.GET("/check", API.CheckView)
		Models.GET("/components/get", API.GetComponentsView)
		Models.GET("/document/get", API.GetModelDocumentView)
		Models.POST("/document/set", API.SetModelDocumentView)
		Models.POST("/units/convert", API.ConvertUnitsView)

		Models.POST("/collection/create", API.CreateCollectionModelView)
		Models.GET("/collection/get", API.GetCollectionModelView)
		Models.GET("/collection/delete", API.DeleteCollectionModelView)
		Models.GET("/search", API.SearchModelView)
		Models.GET("/function/search", API.SearchFunctionTypeView)

		Models.POST("/reference/resources", API.GetModelResourcesReferenceView)

		Models.POST("/userspace/login", API.LoginUserSpaceView)
		Models.POST("/mark", API.AppModelMarkView)

		Models.POST("/CAD/parse", API.CADParseView)
		Models.POST("/CAD/mapping", API.CADMappingModelView)

		Models.GET("/library/available/get", API.GetAvailableLibrariesView)
		Models.GET("/library/noVersion/get", API.GetNoVersionAvailableLibrariesView)
		Models.POST("/library/noVersion/delete", API.DeleteNoVersionAvailableLibrariesView)
		Models.GET("/library/version/get", API.GetVersionAvailableLibrariesView)

		Models.GET("/extend/get", API.GetExtendedModelView)

		Models.GET("/library/system/get", API.GetSystemLibraryView)
		Models.POST("/library/dependency/delete", API.DeleteDependencyLibraryView)
		Models.POST("/library/dependency/create", API.CreateDependencyLibraryView)
		Models.GET("/library/dependency/get", API.GetDependencyLibraryView)
	}
	g.POST("/test", API.Test)
	g.POST("/test1", API.Test1)
}
