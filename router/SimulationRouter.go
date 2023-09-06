package router

import (
	API "yssim-go/app/api/simulate"

	"github.com/gin-gonic/gin"
)

func SimulateRouter(g *gin.Engine) {
	var Models = g.Group("/simulation")
	{
		Models.GET("/options/get", API.GetSimulationOptionsView)
		Models.POST("/options/set", API.SetSimulationOptionsView)

		Models.GET("/state/get", API.GetModelStateView)

		Models.POST("/simulate", API.ModelSimulateView)

		Models.POST("/result/singular", API.SimulateResultSingularView)
		Models.POST("/result", API.SimulateResultGraphicsView)

		Models.GET("/record/list", API.SimulateResultListView)
		Models.GET("/record/details", API.SimulateResultDetailsView)
		Models.GET("/record/tree", API.SimulateResultTreeView)
		Models.GET("/record/delete", API.SimulateResultDeleteView)
		Models.POST("/record/rename", API.SimulateResultRenameView)

		Models.POST("/experiment/create", API.ExperimentCreateView)
		Models.POST("/experiment/delete", API.ExperimentDeleteView)
		Models.POST("/experiment/edit", API.ExperimentEditView)
		Models.GET("/experiment/list", API.ExperimentGetView)
		Models.GET("/experiment/parameters", API.ExperimentParametersView)

		Models.POST("/snapshot/create", API.CreateSnapshotView)
		Models.POST("/snapshot/delete", API.DeleteSnapshotView)
		Models.POST("/snapshot/edit", API.EditSnapshotView)
		Models.GET("/snapshot/list", API.SnapshotGetListView)

		Models.POST("/calibration/compile", API.CalibrationCompileView)

	}
}
