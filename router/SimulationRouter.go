package router

import (
	"github.com/gin-gonic/gin"
	SimulationAPI "yssim-go/app/api/simulate"
)

func SimulateRouter(g *gin.Engine) {
	var Models = g.Group("/simulation")
	{
		Models.GET("/options/get", SimulationAPI.GetSimulationOptionsView)
		Models.POST("/options/set", SimulationAPI.SetSimulationOptionsView)
		Models.GET("/state/get", SimulationAPI.GetModelStateView)
		Models.POST("/simulate", SimulationAPI.ModelSimulateView)
		Models.POST("/result", SimulationAPI.SimulateResultView)
		Models.GET("/record/list", SimulationAPI.SimulateResultListView)
		Models.GET("/record/tree", SimulationAPI.SimulateResultTreeView)
		Models.POST("/experiment/create", SimulationAPI.ExperimentCreateView)
		Models.POST("/experiment/delete", SimulationAPI.ExperimentDeleteView)
		Models.POST("/experiment/edit", SimulationAPI.ExperimentEditView)
		Models.GET("/experiment/list", SimulationAPI.ExperimentGetView)
		Models.POST("/code/save", SimulationAPI.ModelCodeSaveView)
	}
}