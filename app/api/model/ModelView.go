package API

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/library/omc"

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
	_ = DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", userSpaceId}).Find(&packageModel).Error

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
		//if data["haschild"] == true {
		modelData = append(modelData, data)
		//}
	}
	c.JSON(http.StatusOK, modelData)

}

func GetListModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的子节点节点列表(需用传入父节点名称，返回子节点列表)，暂时没有图标信息
		## package_id: 模型包的id
		## modelname: 模型的父节点名称
	*/
	modelName := c.Query("model_name")
	var res ResponseData
	modelChildList := service.GetModelChild(modelName)
	for i := 0; i < len(modelChildList); i++ {
		modelChildList[i]["image"] = service.GetIcon(modelName + "." + modelChildList[i]["model_name"].(string))
	}
	res.Data = modelChildList
	c.JSON(http.StatusOK, res)
}

func GetGraphicsDataView(c *gin.Context) {
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
	modelName := c.Query("model_name")
	var res ResponseData
	modelCode := service.GetModelCode(modelName)
	res.Data = []string{modelCode}
	c.JSON(http.StatusOK, res)
}

func GetModelParametersView(c *gin.Context) {
	/*
		# 获取模型组件的参数数据，一次性返回, 注意，如果是获取整个模型的顶层参数， 只传入模型名称即可， 组件别名和组件名称都不需要传入
		## model_name: 需要查询的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## components_name: 需要查询模型的组件名称，全称， 例如“Modelica.Blocks.Continuous.LimPID“
		## name: 需要查询的组件别名，全称，“PID”
		## sys_user: 模型是系统模型还是用户模型，系统模型固定是“sys”, 用户模型固定是“user”
	*/
	modelName := c.Query("model_name")
	name := c.Query("name")
	componentsName := c.Query("components_name")
	var res ResponseData
	data := service.GetModelParameters(modelName, name, componentsName)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func SetModelParametersView(c *gin.Context) {
	/*
		# 设置模型组件的参数数据，一次性返回
		## package_id: 模型包的id
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## parameter_value: 需要设置的变量和新的值，全称，例如{"PID.k": "200"}， k是模型的组件别名和变量名字的组成， 类似于“别名.变量名”
	*/
	var item SetComponentModifierValueModel
	err := c.BindJSON(&item)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var res ResponseData
	result := service.SetComponentModifierValue(item.ModelName, item.ParameterValue)
	if result {
		res.Msg = "设置完成"
	} else {
		res.Err = "设置失败: 请检查参数是否正确"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func GetComponentPropertiesView(c *gin.Context) {
	/*
		# 获取模型组件的属性数据，一次性返回
		## package_id: 模型包的id
		## class_name: 需要查询属性数据的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## component_name: 需要查询的组件别名，全称，“PID”
	*/
	modelName := c.Query("model_name")
	componentsName := c.Query("component_name")
	var res ResponseData
	result := service.GetComponents(modelName, componentsName)
	if len(result) > 0 {
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
		res.Data = data
	}

	c.JSON(http.StatusOK, res)
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
	var res ResponseData
	var item SetComponentPropertiesModel
	err := c.BindJSON(&item)
	if err != nil {
		res.Status = 2
		res.Err = "设置失败"
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil || packageModel.SysUser == "sys" {
		res.Status = 2
		res.Err = "设置失败"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result := service.SetComponentProperties(
		item.ModelName,
		item.NewComponentName,
		item.OldComponentName,
		strconv.FormatBool(item.Final),
		strconv.FormatBool(item.Protected),
		strconv.FormatBool(item.Replaceable),
		item.Variability,
		strconv.FormatBool(item.Inner),
		strconv.FormatBool(item.Outer),
		item.Causality)
	if result {
		_ = service.SaveModelCode(packageModel.PackageName, packageModel.FilePath)
		res.Msg = "设置完成"
	} else {
		res.Err = "设置失败: 请检查参数是否正确"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
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
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	packageName := ""
	filePath := ""
	var res ResponseData
	var packageModel DataBaseModel.YssimModels
	//packageId := item.PackageId
	//if packageId == "" {
	//	err = DB.Where("package_name = ? AND sys_or_user = ? AND userspace_id = ?", item.ClassName, username, userSpaceId).First(&packageModel).Error
	//	if packageModel.ID != "" {
	//		res.Msg = "模型名称已存在"
	//		res.Status = 2
	//		c.JSON(http.StatusOK, res)
	//		return
	//	}
	//}
	if item.ParentName != "" {
		packageName = packageModel.PackageName
		filePath = packageModel.FilePath
	} else {
		packageName = item.ClassName
		filePath = "public/UserFiles/UploadFile/" + username + "/" + time.Now().Local().Format("20060102150405") + "/" + item.ClassName + ".mo"
	}

	result, msg := service.SaveModel(item.ClassName, item.CopiedClassName, item.ParentName, packageName, "copy", filePath)
	if result {
		res.Msg = msg
		if item.ParentName == "" {
			model := DataBaseModel.YssimModels{
				ID:          uuid.New().String(),
				PackageName: packageName,
				//ModelName:   packageName,
				SysUser:     username,
				FilePath:    filePath,
				UserSpaceId: userSpaceId,
			}
			err := DB.Create(&model).Error
			if err != nil {
				fmt.Println("复制模型失败")
			}
		}
	} else {
		res.Msg = msg
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func DeletePackageAndModelView(c *gin.Context) {
	/*
		# 删除模型包或包中的模型
		## parent_name: 需要删除的模型在哪个父节点之下，例如“Modelica.Blocks.Examples”
		## package_name: 被删除的模型在哪个包之下，例如“Modelica”，如果删除的是包，则就是包的名字，
		## class_name: 被删除的的模型名称，例如“Filter”
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DeleteClassModel
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var res ResponseData
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&packageModel).Error
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result, msg := service.SaveModel(item.ClassName, "", item.ParentName, packageModel.PackageName, "delete", packageModel.FilePath)
	if result {
		res.Msg = msg
		if item.ParentName == "" {
			DB.Delete(&packageModel)
		}
	} else {
		res.Msg = msg
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func GetComponentNameView(c *gin.Context) {
	/*
		# 获取模型当中的模型组件的名字
		## package_id： 包id
		## model_name: 需要创建的组件在哪个模型之下，例如在"NN.Examples.Scenario1_Status"模型中创建组件
		## component_name: 被创建成组件的模型名称， 例如"Modelica.Blocks.Math.Abs"
	*/
	modelName := c.Query("model_name")
	componentName := c.Query("component_name")
	name := service.GetComponentName(modelName, componentName)
	var res ResponseData
	res.Data = []string{name}
	c.JSON(http.StatusOK, res)
}

func AddModelComponentView(c *gin.Context) {
	/*
		# 创建模型当中的模型组件
		## package_id： 包id
		## class_name: 需要创建的组件在哪个模型之下，例如在"Filter1"模型中创建组件
		## new_component_name: 新创建的组件名称，例如"abs1"
		## old_component_name: 被创建成组件的模型名称， 例如"Modelica.Blocks.Math.Abs"
		## origin: 原点， 例如"0,0"
		## extent: 范围坐标, 例如["-10,-10", "10,10"]
		## rotation: 旋转角度, 例如"0"，不旋转`
	*/
	var res ResponseData
	var item AddComponentModel
	err := c.BindJSON(&item)
	if err != nil {
		res.Status = 2
		res.Err = "创建失败"
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, res)
}

func Test(c *gin.Context) {
	/*
		测试omc命令
	*/
	cmd := c.Query("cmd")
	NoParsed := c.Query("NoParsed")
	var data interface{}
	if NoParsed != "" {
		data, _ = omc.OMC.SendExpression(cmd)
	} else {
		d, _ := omc.OMC.SendExpressionNoParsed(cmd)
		data = string(d)
	}
	fmt.Println(data)
	var res ResponseData
	res.Data = data
	c.JSON(http.StatusOK, res)
}
