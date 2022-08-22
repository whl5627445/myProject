package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yssim-go/app/service"
	"yssim-go/config"
)

var DB = config.DB

func GetSimulationOptionsView(c *gin.Context) {
	/*
		# 仿真参数获取接口
		## model_name: 模型名称，
	*/
	modelName := c.Query("model_name")
	var res ResponseData
	result := service.GetSimulationOptions(modelName)
	res.Data = result
	c.JSON(http.StatusOK, res)
}

func SetSimulationOptionsView(c *gin.Context) {
	/*
		# 仿真参数设置接口
		## model_name: 模型名称， 全称
		## package_id: 模型所在包的id
		## package_name: 模型所在包的名称
		## experiment: 仿真参数，对象类型，包含以下几个变量
		   startTime：仿真开始时间，
		   stopTime：仿真结束时间，
		   tolerance：积分方法使用的容差，
		   numberOfIntervals：间隔数，
		   interval：间隔
	*/
	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var res ResponseData

	c.JSON(http.StatusOK, res)
}
func GetModelStateView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func ModelSimulateView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func SimulateResultView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func SimulateResultListView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func SimulateResultTreeView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func ExperimentCreateView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func ExperimentDeleteView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func ExperimentEditView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func ExperimentGetView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func ModelCodeSaveView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
func FmuExportModelView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/

	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	//var packageModel DataBaseModel.YssimModels
	//err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	var res ResponseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}
