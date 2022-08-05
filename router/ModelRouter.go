package router

import (
	"github.com/gin-gonic/gin"
	API "yssim-go/app/api/model"
)

func ModelRouter(e *gin.Engine) {
	var Models = e.Group("/simulate_model")
	{
		Models.POST("/get_graphics_data", API.GetGraphicsDataView)
		Models.GET("/list_root_library", API.GetRootModelView)
		Models.GET("/list_library", API.GetListModelView)
		Models.GET("/get_model_code", API.GetModelCodeView)
		Models.GET("/get_model_parameters", API.GetModelParametersView)
		Models.POST("/set_model_parameters", API.SetModelParametersView)
		Models.GET("/get_component_properties", API.GetComponentPropertiesView)
		Models.POST("/set_component_properties", API.SetComponentPropertiesView)
		Models.POST("/copy_class", API.CopyClassView)
	}
}
