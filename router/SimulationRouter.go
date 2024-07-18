package router

import (
	API "yssim-go/app/v1/api/simulate"
	APIv2 "yssim-go/app/v2/api/simulate"

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
		Models.POST("/record/terminate", API.SimulateTerminateView)

		Models.POST("/experiment/exists", API.ExperimentExistsView)
		Models.POST("/experiment/create", API.ExperimentCreateView)
		Models.POST("/experiment/delete", API.ExperimentDeleteView)
		Models.POST("/experiment/edit", API.ExperimentEditView)
		Models.GET("/experiment/list", API.ExperimentGetView)
		Models.GET("/experiment/parameters", API.ExperimentParametersView)
		Models.POST("/experiment/name/edit", API.ExperimentNameEditView)
		Models.POST("/experiment/compare", API.ExperimentCompareView)

		Models.POST("/snapshot/create", API.CreateSnapshotView)
		Models.POST("/snapshot/delete", API.DeleteSnapshotView)
		Models.POST("/snapshot/edit", API.EditSnapshotView)
		Models.GET("/snapshot/list", API.SnapshotGetListView)

		Models.POST("/calibration/compile", API.CalibrationCompileView)
		Models.POST("/calibration/task/start", API.CalibrationSimulateTaskAddView)
		Models.POST("/calibration/task/stop", API.CalibrationSimulateTaskStopView)
		Models.GET("/calibration/task/status/get", API.GetCalibrationTaskStatusView)

	}

	var ModelsV2 = g.Group("/api/v2/simulation")
	{
		ModelsV2.POST("/simulate", APIv2.ModelSimulateView)
		ModelsV2.GET("/record/delete", APIv2.SimulateResultDeleteView)
		Models.POST("/record/terminate", APIv2.SimulateTerminateView)
		Models.POST("/experiment/delete", APIv2.ExperimentDeleteView)

	}
}
