package router

import (
	"github.com/gin-gonic/gin"
	SimulationAPI "yssim-go/app/api/simulate"
)

func SimulateRouter(e *gin.Engine) {
	var Models = e.Group("/simulation")
	{
		Models.GET("/options/get", SimulationAPI.GetSimulationOptionsView)
		Models.POST("/options/set", SimulationAPI.SetSimulationOptionsView)
		Models.POST("/state/get", SimulationAPI.GetModelStateView)
		Models.POST("/simulate", SimulationAPI.ModelSimulateView)
		Models.POST("/result", SimulationAPI.SimulateResultView)
		Models.POST("/record/list", SimulationAPI.SimulateResultListView)
		Models.POST("/record/tree", SimulationAPI.SimulateResultTreeView)
		Models.POST("/experiment/create", SimulationAPI.ExperimentCreateView)
		Models.POST("/experiment/delete", SimulationAPI.ExperimentDeleteView)
		Models.POST("/experiment/edit", SimulationAPI.ExperimentEditView)
		Models.POST("/experiment/list", SimulationAPI.ExperimentGetView)
		Models.POST("/fmu/export", SimulationAPI.FmuExportModelView)
		Models.POST("/code/save", SimulationAPI.ModelCodeSaveView)
	}
}
