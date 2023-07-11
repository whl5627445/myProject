package router

import (
	API "yssim-go/app/api/appDesign"

	"github.com/gin-gonic/gin"
)

func AppDesignRouter(g *gin.Engine) {
	var Models = g.Group("/app")
	{
		Models.POST("/model/multiple/simulate", API.MultipleSimulateView)
		Models.GET("/model/multiple/simulate/kill", API.MultipleSimulateKillView)
		Models.POST("/model/release", API.AppReleaseView)
		Models.POST("/model/simulate/result", API.GetAppSimulateResultView)
		Models.POST("/model/release/result", API.GetAppReleaseResultView)
		Models.POST("/model/preview/result", API.GetAppPreviewResultView)
		Models.GET("/model/simulate/details", API.GetModelSimulateDetailsView)
		Models.GET("/model/release/details", API.GetModelReleaseDetailsView)
		Models.GET("/model/state/get", API.GetModelStateView)
		Models.POST("/model/message/state/update", API.ModelStateMessageReadView)

		Models.GET("/space/get", API.GetAppSpaceView)
		Models.POST("/space/create", API.CreateAppSpaceView)
		Models.POST("/space/edit", API.EditAppSpaceView)
		Models.POST("/space/delete", API.DeleteAppSpaceView)
		Models.POST("/space/collect", API.AppSpaceCollectView)

		Models.GET("/page/get", API.GetAppPageView)
		Models.GET("/page/space/get", API.GetAppPageSpaceView)
		Models.POST("/page/create", API.CreateAppPageView)
		Models.POST("/page/edit", API.EditAppPageView)
		Models.POST("/page/delete", API.DeleteAppPageView)

		Models.POST("/page/design/edit", API.EditAppPageDesignView)

		Models.GET("/page/input-output/get", API.GetPageInputOutputView)
		Models.POST("/page/input-output/set", API.SetPageInputOutputView)

		Models.GET("/page/components/get", API.GetPageComponentView)
		Models.POST("/page/components/create", API.CreatePageComponentView)
		Models.POST("/page/components/edit", API.EditPageComponentView)
		Models.POST("/page/components/delete", API.DeletePageComponentView)

		Models.GET("/page/components/input-output/get", API.GetPageComponentInputOutputView)
		Models.POST("/page/components/input-output/set", API.SetPageComponentInputOutputView)

		Models.GET("/datasource/get", API.GetDatasourceView)
		Models.POST("/datasource/delete", API.DatasourceDeleteView)
		Models.POST("/datasource/rename", API.DataSourceRenameView)
		Models.GET("/datasource/group/name/get", API.GetDataSourceGroupView)
		Models.GET("/datasource/input/get", API.GetDatasourceInputView)
		Models.GET("/datasource/output/get", API.GetDatasourceOutputView)

		Models.GET("/page/preview", API.AppPagePreviewView)
		Models.GET("/page/release/access", API.AppPageReleaseAccessView)

		Models.POST("/components/basic/information/set", API.SetComponentBasicInformationView)
		Models.GET("/components/basic/information/get", API.GetComponentBasicInformationView)

		Models.POST("/page/alignment-line/set", API.SetPageAlignmentLineView)
		Models.GET("/page/alignment-line/get", API.GetPageAlignmentLineView)
	}
}
