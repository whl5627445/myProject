package router

import (
	"github.com/gin-gonic/gin"
	ModelAPI "yssim-go/app/api/model"
)

func ModelRouter(e *gin.Engine) {
	var Models = e.Group("/model")
	{
		Models.GET("/root_library", ModelAPI.GetRootModelView)
		Models.GET("/list_library", ModelAPI.GetListModelView)
		Models.POST("/graphics", ModelAPI.GetGraphicsDataView)
		Models.GET("/code", ModelAPI.GetModelCodeView)
		Models.GET("/parameters/get", ModelAPI.GetModelParametersView)
		Models.POST("/parameters/set", ModelAPI.SetModelParametersView)
		Models.GET("/properties/get", ModelAPI.GetComponentPropertiesView)
		Models.POST("/properties/set", ModelAPI.SetComponentPropertiesView)
		Models.POST("/class/copy", ModelAPI.CopyClassView)
		Models.POST("/package/delete", ModelAPI.DeletePackageAndModelView)
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

	}
	e.POST("/test", ModelAPI.Test)
}
