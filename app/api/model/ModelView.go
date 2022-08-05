package API

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	DataBaseModel "yssim-go/app/DataBaseModel/Model"
	"yssim-go/app/service"

	"yssim-go/config"
)

var DB = config.DB

func GetRootModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var modelData []map[string]interface{}
	var packageModel []DataBaseModel.YssimModels
	err := DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", userSpaceId}).Find(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	for i := 0; i < len(packageModel); i++ {
		data := map[string]interface{}{
			"packeage_id":  packageModel[i].ID,
			"package_name": packageModel[i].PackageName,
			"sys_or_user":  packageModel[i].SysUser,
			"haschild":     service.GetModelHasChild(packageModel[i].PackageName),
			"image":        service.GetIcon(packageModel[i].PackageName),
		}
		if packageModel[i].SysUser != "sys" {
			data["sys_or_user"] = "user"
		}
		if data["haschild"] == true {
			modelData = append(modelData, data)
		}
	}
	c.JSON(http.StatusOK, modelData)

}

func GetListModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的子节点节点列表(需用传入父节点名称，返回子节点列表)，暂时没有图标信息
		## package_id: 模型包的id
		## modelname: 模型的父节点名称
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	packageId := c.Query("package_id")
	modelName := c.Query("model_name")
	var req ResponseData
	var packageModel DataBaseModel.YssimModels
	err := DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	modelChildList := service.GetModelChild(modelName)
	for i := 0; i < len(modelChildList); i++ {
		modelChildList[i]["image"] = service.GetIcon(modelName + "." + modelChildList[i]["model_name"].(string))
	}
	req.Data = modelChildList
	c.JSON(http.StatusOK, req)
}

func GetGraphicsDataView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	c.JSON(http.StatusOK, graphicsData)
}

func GetModelCodeView(c *gin.Context) {
	/*
		# 获取模型的源码数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
	*/

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	packageId := c.Query("package_id")
	modelName := c.Query("model_name")
	var req ResponseData
	var packageModel DataBaseModel.YssimModels
	err := DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	modelCode := service.GetModelCode(modelName)
	req.Data = []string{modelCode}
	c.JSON(http.StatusOK, req)
}

func GetModelParametersView(c *gin.Context) {
	/*
		# 获取模型组件的参数数据，一次性返回, 注意，如果是获取整个模型的顶层参数， 只传入模型名称即可， 组件别名和组件名称都不需要传入
		## model_name: 需要查询的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## components_name: 需要查询模型的组件名称，全称， 例如“Modelica.Blocks.Continuous.LimPID“
		## name: 需要查询的组件别名，全称，“PID”
		## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	packageId := c.Query("package_id")
	modelName := c.Query("model_name")
	name := c.Query("name")
	componentsName := c.Query("components_name")
	var req ResponseData
	var packageModel DataBaseModel.YssimModels
	err := DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	data := service.GetModelParameters(modelName, name, componentsName, packageModel.PackageName)
	req.Data = data
	c.JSON(http.StatusOK, req)
}

func SetModelParametersView(c *gin.Context) {
	/*
		# 设置模型组件的参数数据，一次性返回
		## package_id: 模型包的id
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## parameter_value: 需要设置的变量和新的值，全称，例如{"PID.k": "200"}， k是模型的组件别名和变量名字的组成， 类似于“别名.变量名”
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item SetComponentModifierValueModel
	err := c.BindJSON(&item)
	packageId := item.PackageId
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var req ResponseData
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.SetComponentModifierValue(item.ModelName, item.ParameterValue)
	if result {
		req.Msg = "设置完成"
	} else {
		req.Err = "设置失败: 请检查参数是否正确"
		req.Status = 2
	}

	c.JSON(http.StatusOK, req)
}

func GetComponentPropertiesView(c *gin.Context) {
	/*
		# 获取模型组件的属性数据，一次性返回
		## package_id: 模型包的id
		## class_name: 需要查询属性数据的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## component_name: 需要查询的组件别名，全称，“PID”
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	packageId := c.Query("package_id")
	modelName := c.Query("model_name")
	componentsName := c.Query("component_name")
	var req ResponseData
	var packageModel DataBaseModel.YssimModels
	err := DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.GetComponents(modelName, componentsName)
	data := map[string]interface{}{
		"model_name":     modelName,
		"component_name": componentsName,
		"path":           result[0],
		"dimension":      result[len(result)-1],
		"annotation":     result[2],
		"Properties":     []string{result[4].(string), result[3].(string), result[7].(string)},
		"Variability":    result[8],
		"Inner/Outer":    result[9],
		"Causality":      result[10],
	}
	req.Data = data
	c.JSON(http.StatusOK, req)
}

func SetComponentPropertiesView(c *gin.Context) {
	/*
		# 设置模型组件的属性数据，一次性返回
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## old_component_name: 需要设置的组件名，全称，“PID”
		## new_component_name: 需要设置的组件新名称，全称，“PID”
		## final: "true" or "false",
		## protected: "true" or "false",
		## replaceable: "true" or "false",
		## variabilty: "unspecified" or  "parameter" or "discrete" or "constant"
		## inner: "true" or "false",
		## outer: "true" or "false",
		## causality: "output" or "input"
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item SetComponentPropertiesModel
	err := c.BindJSON(&item)
	packageId := item.PackageId
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var req ResponseData
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.SetComponentProperties(item.ModelName, item.NewComponentName, item.OldComponentName, item.Final, item.Protected, item.Replaceable, item.Variability, item.Inner, item.Outer, item.Causality)
	if result {
		req.Msg = "设置完成"
	} else {
		req.Err = "设置失败: 请检查参数是否正确"
		req.Status = 2
	}

	c.JSON(http.StatusOK, req)
}

func CopyClassView(c *gin.Context) {
	/*
		# 复制模型
		## parent_name: 需要复制到哪个父节点之下，例如“ENN.Examples”
		## package_name: 被复制的模型在哪个包之下，例如“ENN”
		## class_name: 复制之后的模型名称，例如“Scenario1_Status_test”
		## copied_class_name: 被复制的模型全称，例如“ENN.Examples.Scenario1_Status”
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item CopyClassModel
	err := c.BindJSON(&item)
	packageId := item.PackageId
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var req ResponseData
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", packageId, username, userSpaceId).First(&packageModel).Error
	//if err != nil {
	//	fmt.Println(err)
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}

	c.JSON(http.StatusOK, req)
}
