package router

import (
	"github.com/gin-gonic/gin"
	API "yssim-go/app/api/model"
)

func ModelRouter(e *gin.Engine) {
	var Models = e.Group("/simulate_model")
	{
		Models.POST("/graphics_data", API.GetGraphicsDataView)
		Models.GET("/root_library", API.GetRootModelView)
		Models.GET("/list_library", API.GetListModelView)
		Models.GET("/model_code", API.GetModelCodeView)
		Models.GET("/model_parameters/get", API.GetModelParametersView)
		Models.POST("/model_parameters/set", API.SetModelParametersView)
		Models.GET("/component_properties/get", API.GetComponentPropertiesView)
		Models.POST("/component_properties/set", API.SetComponentPropertiesView)
		Models.POST("/class/copy", API.CopyClassView)
		Models.POST("/package/delete", API.DeletePackageAndModelView)
		Models.GET("/component/name", API.GetComponentNameView)
		Models.POST("/component/add", API.AddModelComponentView)
		Models.POST("/component/delete", API.DeleteModelComponentView)
		Models.POST("/component/update", API.UpdateModelComponentView)

	}
	e.POST("/test", API.Test)
}
