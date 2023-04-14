package API

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/library/omc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"yssim-go/config"
)

var DB = config.DB

func GetSysRootModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
	*/
	var res responseData
	var modelData []map[string]interface{}
	var packageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user =  ? AND userspace_id = ?", "sys", "0").Find(&packageModel)
	libraryAndVersions := service.GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		p, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && p == packageModel[i].Version {
			data := map[string]interface{}{
				"package_id":      packageModel[i].ID,
				"package_name":    packageModel[i].PackageName,
				"package_version": packageModel[i].Version,
				"model_name":      packageModel[i].PackageName,
				"haschild":        service.GetModelHasChild(packageModel[i].PackageName),
				"image":           service.GetIcon(packageModel[i].PackageName, packageModel[i].PackageName, packageModel[i].Version),
				"type":            service.GetModelType(packageModel[i].PackageName),
			}
			modelData = append(modelData, data)
		}
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)

}

func GetUserRootModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
	*/
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var res responseData
	var modelData []map[string]interface{}
	var packageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user = ? AND userspace_id = ?", userName, userSpaceId).Find(&packageModel)
	libraryAndVersions := service.GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		loadVersions, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && loadVersions == packageModel[i].Version {
			data := map[string]interface{}{
				"package_id":      packageModel[i].ID,
				"package_name":    packageModel[i].PackageName,
				"package_version": packageModel[i].Version,
				"model_name":      packageModel[i].PackageName,
				"haschild":        service.GetModelHasChild(packageModel[i].PackageName),
				"image":           service.GetIcon(packageModel[i].PackageName, packageModel[i].PackageName, packageModel[i].Version),
				"type":            service.GetModelType(packageModel[i].PackageName),
			}
			if service.GetModelType(packageModel[i].PackageName) == "package" {
				data["haschild"] = true
			}
			modelData = append(modelData, data)
		}

	}

	res.Data = modelData
	c.JSON(http.StatusOK, res)

}

func GetUserPackageView(c *gin.Context) {
	/*
		# 获取已加载package列表接口
	*/
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var res responseData
	var modelData []map[string]interface{}
	var packageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user = ? AND userspace_id = ?", userName, userSpaceId).Find(&packageModel)
	libraryAndVersions := service.GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		loadVersions, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && loadVersions == packageModel[i].Version && service.GetModelType(packageModel[i].PackageName) == "package" {
			data := map[string]interface{}{
				"package_id":   packageModel[i].ID,
				"package_name": packageModel[i].PackageName,
			}
			modelData = append(modelData, data)
		}
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)

}

func GetListModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的子节点节点列表(需用传入父节点名称，返回子节点列表)，暂时没有图标信息
		## package_id: 模型包的id
		## modelname: 模型的父节点名称
	*/
	modelName := c.Query("model_name")
	packageId := c.Query("package_id")
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var res responseData
	modelChildList := service.GetModelChild(modelName)
	var modelChildListNew []map[string]interface{}
	for i := 0; i < len(modelChildList); i++ {
		modelChildList[i]["image"] = service.GetIcon(modelName+"."+modelChildList[i]["name"].(string), packageModel.PackageName, packageModel.Version)
		modelChildListNew = append(modelChildListNew, modelChildList[i])
	}
	// 如果父节点是包名称的话，追加静态资源管理文件夹节点
	nameType := service.GetModelType(modelName)
	if modelName == packageModel.PackageName && packageModel.SysUser != "sys" && nameType == "package" {
		modelChildListNew = append(modelChildListNew, map[string]interface{}{
			"name":     "Resources",
			"haschild": true,
			"type":     "static",
		})
	}
	res.Data = modelChildListNew
	c.JSON(http.StatusOK, res)
}

func GetGraphicsDataView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## component_name: 模型的组件名称，用于获取单个组件时传入
	*/
	var item modelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	packageName := strings.Split(item.ModelName, ".")[0]
	err = DB.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", packageName, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var res responseData
	var graphicsData [][]map[string]interface{}
	if item.ComponentName == "" {
		graphicsData = service.GetGraphicsData(item.ModelName, packageModel.SysUser)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}
	res.Data = graphicsData
	c.JSON(http.StatusOK, res)
}

func GetModelCodeView(c *gin.Context) {
	/*
		# 获取模型的源码数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
	*/
	modelName := c.Query("model_name")
	var res responseData
	modelCode := service.GetModelFileCode(modelName)
	res.Data = modelCode
	c.JSON(http.StatusOK, res)
}

func GetModelResourcesReferenceView(c *gin.Context) {
	/*
		# 获取包级别的静态资源，以Reference的形式返回
		## package_id: 包id
		## parent: 需要查询的节点父级路径
		## path: 被查询节点
	*/
	var item packageResourcesData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var res responseData

	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	if item.PackageId != "" {
		var packageModel DataBaseModel.YssimModels
		DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
		if packageModel.ID == "" {
			log.Println(err)
			c.JSON(http.StatusBadRequest, "not found")
			return
		}

		data := service.GetResourcesList(packageModel.PackageName, item.Parent)
		for _, d := range data {
			basePath := ""
			pathList := []string{}
			if item.Parent != "" {
				pathList = append(pathList, item.Parent)
			}
			if d["type"] == "file" {
				pathList = append(pathList, d["name"])
				basePath = "Modelica.Utilities.Files.loadResource(\"modelica://" + packageModel.PackageName + "/Resources/" + strings.Join(pathList, "/") + "\")"
			}
			d["path"] = basePath

		}
		res.Data = data
	} else {
		var packageModelList []DataBaseModel.YssimModels
		DB.Where("sys_or_user = ? AND userspace_id = ?", userName, userSpaceId).Find(&packageModelList)
		var data []map[string]interface{}
		for _, model := range packageModelList {
			information := service.GetClassInformation(model.PackageName)
			if len(information) > 0 && information[0] == "package" {
				data = append(data, map[string]interface{}{"name": model.PackageName + ".Resources", "package_id": model.ID, "type": "dir"})
			}
		}
		res.Data = data
	}

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
	componentName := c.Query("component_name")
	className := c.Query("class_name")

	var res responseData
	properties := make(map[string]interface{}, 0)
	if modelName == "" || componentName == "" || className == "" {
		componentName = modelName
		className = modelName
	}
	parameters := service.GetModelParameters(modelName, componentName, className)
	elements := service.GetElements(modelName, componentName)
	if len(elements) > 0 && componentName != "" {
		dimension := elements[len(elements)-1].(string)
		properties = map[string]interface{}{
			"model_name":     modelName,
			"component_name": componentName,
			"path":           elements[2],
			"dimension":      dimension,
			"annotation":     elements[4],
			"Properties":     []interface{}{elements[6], elements[5], elements[9]},
			"Variability":    elements[10],
			"Inner/Outer":    elements[11],
			"Causality":      elements[12],
		}
	}
	res.Data = map[string]interface{}{"parameters": parameters, "properties": properties}
	c.JSON(http.StatusOK, res)
}

func SetModelParametersView(c *gin.Context) {
	/*
		# 设置模型组件的参数数据，一次性返回
		## package_id: 模型包的id
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## parameter_value: 需要设置的变量和新的值，全称，例如{"PID.k": "200"}， k是模型的组件别名和变量名字的组成， 类似于“别名.变量名”
	*/
	var item setComponentModifierValueData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var res responseData
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.ID == "" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := service.SetComponentModifierValue(item.ModelName, item.ParameterValue)
	if result {
		service.ModelSave(item.ModelName)
		res.Msg = "设置完成"
	} else {
		res.Err = "设置失败: 请检查参数是否正确"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func AddModelParametersView(c *gin.Context) {
	/*
		# 创建模型的全局参数
		## package_id: 模型包的id
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## parameter_name: 参数名称
	*/
	var item addComponentParametersData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var res responseData
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.ID == "" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	_, err = service.AddComponentParameters(item.ParameterName, item.VarType, item.ModelName)
	if err != nil {
		res.Err = err.Error()
		res.Status = 2
	} else {
		service.ModelSave(item.ModelName)
		res.Msg = "设置完成"
	}

	c.JSON(http.StatusOK, res)
}

func DeleteModelParametersView(c *gin.Context) {
	/*
		# 删除模型的全局参数
		## package_id: 模型包的id
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## parameter_name: 参数名称
	*/
	var item deleteComponentParametersData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var res responseData
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.ID == "" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result, err := service.DeleteComponentParameters(item.ParameterName, item.ModelName)
	if result {
		service.ModelSave(item.ModelName)
		res.Msg = "参数已删除"
	} else {
		res.Err = err.Error()
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
	var res responseData
	result := service.GetElements(modelName, componentsName)
	if len(result) > 0 {
		dimension := result[len(result)-1].(string)
		data := map[string]interface{}{
			"model_name":     modelName,
			"component_name": componentsName,
			"path":           result[2],
			"dimension":      "[" + dimension[1:len(dimension)-1] + "]",
			"annotation":     result[4],
			"Properties":     []string{result[6].(string), result[5].(string), result[9].(string)},
			"Variability":    result[10],
			"Inner/Outer":    result[11],
			"Causality":      result[12],
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
	var res responseData
	var item setComponentPropertiesData
	err := c.BindJSON(&item)
	if err != nil {
		res.Status = 2
		res.Err = "设置失败"
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil || packageModel.SysUser == "sys" {
		res.Status = 2
		res.Err = "设置失败"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result, msg := service.SetComponentProperties(
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
		service.ModelSave(item.ModelName)
		res.Msg = "设置完成"
	} else {
		res.Err = msg
		res.Status = 2
	}
	data := make([]string, 0, 1)
	res.Data = data
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
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item copyClassData
	var res responseData
	err := c.BindJSON(&item)
	// if err != nil || item.PackageId == "" && item.ParentName != "" {
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if strings.Contains(item.ModelName, ".") {

		res.Err = "模型名称不能包含: \".\""
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	packageName := item.ModelName
	if item.ParentName != "" {
		packageName = strings.Split(item.ParentName, ".")[0]
	}
	filePath := ""

	var packageModel DataBaseModel.YssimModels
	DB.Where("package_name = ? AND userspace_id = ?", packageName, "0").Or("sys_or_user = ? AND userspace_id = ? AND package_name = ?", userName, userSpaceId, packageName).First(&packageModel)
	if packageModel.SysUser == "sys" {
		res.Msg = "标准库不允许插入模型"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if packageModel.PackageName == item.ModelName {
		res.Msg = "模型名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if item.ParentName != "" {
		packageName = packageModel.PackageName
		filePath = packageModel.FilePath
	} else {
		packageName = item.ModelName
		filePath = "public/UserFiles/UploadFile/" + userName + "/" + packageName + "/" + time.Now().Local().Format("20060102150405") + "/" + item.ModelName + ".mo"
	}
	model := DataBaseModel.YssimModels{
		ID:          uuid.New().String(),
		PackageName: packageName,
		SysUser:     userName,
		FilePath:    filePath,
		UserSpaceId: userSpaceId,
	}
	if item.ParentName == "" {
		err = DB.Create(&model).Error
		if err != nil {
			log.Println("复制模型失败 err：", err)
			res.Msg = "复制模型失败"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	}
	result, msg := service.SaveModel(item.ModelName, item.CopiedClassName, item.ParentName, "copy", filePath)
	if result {

		res.Msg = msg
		data := map[string]string{}
		if item.ParentName == "" {

			data["id"] = model.ID
			data["model_name"] = item.ModelName
		} else {
			data["id"] = packageModel.ID
			data["model_name"] = item.ParentName + "." + item.ModelName
		}
		service.ModelSave(item.ModelName)
		res.Data = data
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
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item deleteClassData
	err := c.BindJSON(&item)
	if err != nil {

		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var res responseData
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result, msg := service.SaveModel(item.ModelName, "", item.ParentName, "delete", packageModel.FilePath)
	if result {
		service.ModelSave(item.ParentName)
		res.Msg = msg
		if item.ParentName == "" {
			var simulateRecord []DataBaseModel.YssimSimulateRecord
			DB.Where("package_id = ? AND username = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).Find(&simulateRecord)
			DB.Delete(&packageModel)
		}
		var modelCollection []DataBaseModel.YssimModelsCollection
		DB.Where("package_id = ? AND model_name = ? AND userspace_id = ?", packageModel.ID, item.ModelName, userSpaceId).Find(&modelCollection)
		DB.Delete(&modelCollection)
		//删除对应的实验记录 暂且搁置
		//var experimentRecord []DataBaseModel.YssimExperimentRecord
		//DB.Where("username =? AND userspace_id =? AND model_name =?", username, userSpaceId, item.ModelName).Find(&experimentRecord)
		//DB.Delete(&experimentRecord)
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
	className := c.Query("class_name")
	name := service.GetComponentName(modelName, className)
	var res responseData
	res.Data = []string{name}
	c.JSON(http.StatusOK, res)
}

func AddModelComponentView(c *gin.Context) {
	/*
		# 创建模型当中的模型组件
		## package_id： 包id
		## model_name: 需要创建的组件在哪个模型之下，例如在"Filter1"模型中创建组件
		## new_component_name: 新创建的组件名称，例如"abs1"
		## old_component_name: 被创建成组件的模型名称， 例如"Modelica.Blocks.Math.Abs"
		## origin: 原点， 例如"0,0"
		## extent: 范围坐标, 例如["-10,-10", "10,10"]
		## rotation: 旋转角度, 例如"0"，不旋转`
	*/
	var res responseData
	var item addComponentData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	rotation := strconv.Itoa(item.Rotation)
	result, msg := service.AddComponent(item.NewComponentName, item.OldComponentName, item.ModelName, item.Origin, rotation, item.Extent)
	data := service.GetIconNew(item.OldComponentName)
	if !result {
		res.Err = msg
		res.Status = 2
	} else {
		s := time.Now().UnixNano() / 1e6
		service.ModelSave(item.ModelName)
		log.Println("time", time.Now().UnixNano()/1e6-s)
		res.Data = data
		res.Msg = "新增组件成功"
	}
	c.JSON(http.StatusOK, res)
}

func DeleteModelComponentView(c *gin.Context) {
	/*
		# 删除模型当中的模型组件
		## package_id： 包id
		## component_list：需要删除的数据数组(delete_type：删除类型，component_name：删除的组件名字，model_name：模型名称，connect_start：连线类型起点，connect_end：终点)
	*/
	var res responseData
	var item deleteComponentData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	DB.Where("id = ?", item.PackageId).First(&modelPackage)
	if modelPackage.SysUser == "sys" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := true
	modelName := ""
	for _, component := range item.ComponentList {
		modelName = component.ModelName
		switch component.DeleteType {
		case "component":
			result = service.DeleteComponent(component.ComponentName, component.ModelName)
		case "connector":
			result = service.DeleteConnection(component.ModelName, component.ConnectStart, component.ConnectEnd)
		}
		if !result {
			res.Err = "删除组件失败"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	}
	service.ModelSave(modelName)

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func UpdateModelComponentView(c *gin.Context) {
	/*
		# 更新模型当中的模型组件
		## package_id： 包id
		## component_name: 需要更新的组件名称，例如"limPID"，
		## component_class_name: 需要更新的组件原本模型名称，例如"Modelica.Blocks.Continuous.LimPID"
		## model_name: 需要更新的组件在哪个模型当中， 例如"Modelica.Blocks.Examples.PID_Controller"
		## origin: 原点， 例如"0,0"
		## extent: 范围坐标, 例如["-10,-10", "10,10"]
		## rotation: 旋转角度, 例如"0"，不旋转
	*/
	var res responseData
	var item updateComponentData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := service.UpdateComponent(item.ComponentName, item.ComponentClassName, item.ModelName, item.Origin, item.Rotation, item.Extent)
	if !result {
		res.Err = "更新组件失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	} else {
		for _, connect := range item.ConnectionList {
			service.UpdateConnection(item.ModelName, connect.ConnectStart, connect.ConnectEnd, connect.Color, connect.LinePoints)
		}
	}
	s := time.Now().UnixNano() / 1e6
	service.ModelSave(item.ModelName)
	log.Println("time", time.Now().UnixNano()/1e6-s)
	res.Msg = "更新组件成功"
	c.JSON(http.StatusOK, res)
}

func CreateConnectionAnnotationView(c *gin.Context) {
	/*
		# 创建模型当中的组件连线
		## package_id： 包id
		## model_name：在哪个模型创建，模型全称
		## connect_start：连线起点， 输出点， 例如"sum1.y"
		## connect_end：连线终点， 输入点， 例如"ChebyshevI.u"
		## color：连线颜色， 例如"0,0,127"
		## line_points：连线拐点坐标，包含起始点坐标，从起点开始到终点 例如["213,-38","-163.25,-38","-163.25,-4"]
	*/
	var res responseData
	var item updateConnectionAnnotationData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := service.AddConnection(item.ModelName, item.ConnectStart, item.ConnectEnd, item.Color, item.LinePoints)
	if !result {
		res.Err = "连接组件失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	service.ModelSave(item.ModelName)
	res.Msg = "连接组件成功"
	c.JSON(http.StatusOK, res)
}

func UpdateConnectionNamesView(c *gin.Context) {
	/*
		# 更新模型当中的组件连线
		## package_id： 包id
		## model_name：在哪个模型修改，模型全称
		## from_name：连线起点， 输出点， 例如"sum1.y"
		## to_name：连线终点， 输入点， 例如"sum2.y"
		## from_name_new：连线起点， 输出点， 例如"sum1new.y"
		## to_name_new：连线终点， 输入点， 例如"sum2new.y"
	*/
	var res responseData
	var item updateConnectionNamesData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := service.UpdateConnectionNames(item.ModelName, item.FromName, item.ToName, item.FromNameNew, item.ToNameNew)
	if !result {
		res.Err = "更新连线失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	service.ModelSave(item.ModelName)
	res.Msg = "更新连线成功"
	c.JSON(http.StatusOK, res)
}

func DeleteConnectionAnnotationView(c *gin.Context) {
	/*
		# 删除模型当中的删除组件连线
		## package_id： 包id
		## model_name_all： 在哪个模型当中删除连线
		## connect_start： 连线起始位置
		## connect_end： 连线终止位置
	*/
	var res responseData
	var item deleteConnectionData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := service.DeleteConnection(item.ModelName, item.ConnectStart, item.ConnectEnd)
	if !result {
		res.Err = "删除连线失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	service.ModelSave(item.ModelName)
	res.Msg = "删除连线成功"
	c.JSON(http.StatusOK, res)
}

func UpdateConnectionAnnotationView(c *gin.Context) {
	/*
		# 更新模型当中的组件连线
		## package_id： 包id
		## model_name_all：在哪个模型中更新，模型全称
		## connect_start：连线起点， 输出点， 例如"sum1.y"
		## connect_end：连线终点， 输入点， 例如"ChebyshevI.u"
		## color：连线颜色， 例如"0,0,127"
		## line_points：连线拐点坐标，包含起始点坐标，从起点开始到终点 例如["213,-38","-163.25,-38","-163.25,-4"]
	*/
	var res responseData
	var item updateConnectionAnnotationData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.ModelName, ".")
	packageName := nameList[0]
	var modelPackage DataBaseModel.YssimModels
	DB.Where("package_name = ?", packageName).First(&modelPackage)
	if modelPackage.SysUser == "sys" {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := service.UpdateConnection(item.ModelName, item.ConnectStart, item.ConnectEnd, item.Color, item.LinePoints)
	if !result {
		res.Err = "更新连线失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	service.ModelSave(item.ModelName)
	res.Msg = "更新连线成功"
	c.JSON(http.StatusOK, res)
}

func ExistsView(c *gin.Context) {
	/*
		# 检查模型是否存在
		## package_id： 包id
		## model_name：模型全称
	*/
	var res responseData
	modelName := c.Query("model_name")
	result := service.ExistClass(modelName)
	res.Data = []bool{result}
	c.JSON(http.StatusOK, res)
}

func CheckView(c *gin.Context) {
	/*
		# 检查模型是否合规
		## package_id： 包id
		## model_name：模型全称
	*/
	var res responseData
	modelName := c.Query("model_name")
	dataList := service.CheckModel(modelName)
	for _, mes := range dataList {
		_ = service.MessageNotice(mes)
	}
	res.Msg = "模型检查完成"
	c.JSON(http.StatusOK, res)
}

func GetComponentsView(c *gin.Context) {
	/*
		# 获取模型的全部组件数据，一次性返回
		##  model_name: 需要查询属性数据的模型名称，全称，例如“Modelica.Blocks.Examples.PID_Controller”
		##  package_id: 所属package的id值，例如“1”
	*/
	var res responseData
	modelName := c.Query("model_name")
	result := service.GetElements(modelName, "")
	var data []map[string]string
	for _, e := range result {
		component := map[string]string{
			"component_model_name":  e.([]interface{})[2].(string),
			"component_name":        e.([]interface{})[3].(string),
			"component_description": e.([]interface{})[4].(string),
		}
		data = append(data, component)
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetModelDocumentView(c *gin.Context) {
	/*
		# 获取模型的文档数据
		##  model_name: 需要查询文档的模型名称，全称，例如“Modelica.Blocks.Examples.PID_Controller”
		##  package_id: 所属package的id值，例如“1”
	*/
	var res responseData
	modelName := c.Query("model_name")
	result := service.GetModelDocument(modelName)
	res.Data = map[string]string{
		"document":   result[0],
		"revisions":  result[1],
		"model_name": modelName,
	}
	c.JSON(http.StatusOK, res)
}

func SetModelDocumentView(c *gin.Context) {
	/*
		# 设置模型的文档数据
		##  model_name: 需要查询文档的模型名称，全称，例如“Modelica.Blocks.Examples.PID_Controller”
		##  document: 文档内容
		##  package_id: 所属package的id值，例如“1”
	*/
	var res responseData
	var item setModelDocumentData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	DB.Where("sys_or_user = ? AND userspace_id = ?", username, userSpaceId).First(&packageModel)

	result := service.SetModelDocument(item.ModelName, item.Document, item.Revisions)
	if !result {
		res.Status = 2
		res.Err = "修改失败"
		c.JSON(http.StatusOK, res)
		return
	}
	service.ModelSave(item.ModelName)
	res.Msg = "修改成功"
	c.JSON(http.StatusOK, res)
}

func ConvertUnitsView(c *gin.Context) {
	/*
		# 转换参数单位
		##  s1: 转换后的单位, new单位
		##  s2: 需要转换的单位， old单位
	*/
	var res responseData
	var item convertUnitsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.ConvertUnits(item.S1, item.S2)
	unitsCompatible, _ := strconv.ParseBool(result[0])
	scaleFactor, _ := strconv.ParseFloat(result[1], 32)
	offset, _ := strconv.ParseFloat(result[2], 32)
	res.Data = map[string]interface{}{
		"units_compatible": unitsCompatible,
		"scale_factor":     scaleFactor,
		"offset":           offset,
	}
	c.JSON(http.StatusOK, res)
}

func CreateCollectionModelView(c *gin.Context) {
	/*
		# 新增收藏模型
		## package_id: 模型包的id
		## modelname: 需要增加的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”

	*/
	var res responseData
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item modelCollectionData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//检测PackageId，userspace_id是否存在
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	//omc检测模型是否存在
	result := service.ExistClass(item.ModelName)
	if !result {
		c.JSON(http.StatusBadRequest, "model not found")
		return
	}
	//检测数据库表中是否存在同名模型
	modelType := service.GetModelType(item.ModelName)
	if modelType == "package" {
		res.Err = "包类型暂不允许收藏，请选择其他类型"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var modelCollection DataBaseModel.YssimModelsCollection
	DB.Where("model_name = ? AND userspace_id = ?", item.ModelName, userSpaceId).First(&modelCollection)
	if modelCollection.ID != "" {
		res.Err = "名称已存在，请修改后再试。"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	//表中插入记录
	var newCollection = DataBaseModel.YssimModelsCollection{
		ID:          uuid.New().String(),
		PackageId:   item.PackageId,
		UserSpaceId: userSpaceId,
		ModelName:   item.ModelName,
	}
	err = DB.Create(&newCollection).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Creation failed")
		return
	}

	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)

}

func GetCollectionModelView(c *gin.Context) {
	/*
		# 获取收藏模型列表
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var res responseData
	var modelData []map[string]interface{}
	var modelCollections []map[string]interface{}
	DB.Raw("select mc.id, mc.package_id, m.package_name, mc.model_name, m.version, m.sys_or_user from yssim_models_collections as mc  left join yssim_models m on mc.package_id = m.id where mc.userspace_id = ?  and m.sys_or_user IN (?,\"sys\") and mc.deleted_at is NULL", userSpaceId, username).Scan(&modelCollections)
	for i := 0; i < len(modelCollections); i++ {
		modelName := modelCollections[i]["model_name"].(string)
		packageName := modelCollections[i]["package_name"].(string)
		version := modelCollections[i]["version"].(string)
		sysOrUser := false
		if modelCollections[i]["sys_or_user"].(string) == "sys" {
			sysOrUser = true
		}
		//检测模型是否存在，不存在就从表中删除记录
		result := service.ExistClass(modelName)
		if !result {
			go DB.Delete(&modelCollections[i])
			continue
		}
		data := map[string]interface{}{
			"id":           modelCollections[i]["id"],
			"package_id":   modelCollections[i]["package_id"],
			"package_name": modelCollections[i]["package_name"],
			"model_name":   modelCollections[i]["model_name"],
			"haschild":     service.GetModelHasChild(modelName),
			"image":        service.GetIcon(modelName, packageName, version),
			"type":         service.GetModelType(modelName),
			"sys_or_user":  sysOrUser,
		}

		modelData = append(modelData, data)
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)
}

func DeleteCollectionModelView(c *gin.Context) {
	/*
		# 删除收藏的模型
		## id： 需要删除的收藏模型id
	*/
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")

	id := c.Query("id")
	var res responseData
	var modelCollection DataBaseModel.YssimModelsCollection
	DB.Where("id = ?", id).First(&modelCollection)
	if modelCollection.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	DB.Delete(&modelCollection)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func SearchModelView(c *gin.Context) {
	/*
		# 搜索模型
		## keywords: 需要搜索的关键字
		## parent: 需要搜索的关键字的父节点
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	keywords := c.Query("keywords")
	parent := c.Query("parent")
	var res responseData
	var data []map[string]interface{}
	var packageModel []DataBaseModel.YssimModels
	libraryAndVersions := service.GetLibraryAndVersions()
	if parent != "" {
		namesList := strings.Split(parent, ".")
		packageName := namesList[0]
		DB.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", packageName, []string{"sys", username}, []string{"0", userSpaceId}).Order("sys_or_user desc").Find(&packageModel)
	} else {
		DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", userSpaceId}).Order("sys_or_user desc").Find(&packageModel)
	}
	for i := 0; i < len(packageModel); i++ {
		if libraryAndVersions[packageModel[i].PackageName] == packageModel[i].Version {
			modelNameList := service.SearchModel(packageModel[i], keywords, parent)
			data = append(data, modelNameList...)
		}
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func LoadModelView(c *gin.Context) {
	/*
		# 加载模型
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var res responseData
	var loadPackage loadPackageData
	var packageModel DataBaseModel.YssimModels
	err := c.BindJSON(&loadPackage)
	if err != nil {
		log.Println(err)
	}
	DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", loadPackage.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel)
	if err != nil || packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if len(loadPackage.LoadPackageConflict) == 0 {
		conflict, err := service.GetLoadPackageConflict(packageModel.PackageName, packageModel.Version, packageModel.FilePath)
		if len(conflict) > 0 && err != nil {
			res.Data = conflict
			res.Msg = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
	} else {
		for _, conflict := range loadPackage.LoadPackageConflict {
			err = service.LoadAndDeleteLibrary(conflict.Name, conflict.Version, "", "unload")

		}
	}
	err = service.LoadAndDeleteLibrary(packageModel.PackageName, packageModel.Version, packageModel.FilePath, "load")
	if err != nil {
		res.Err = err.Error()
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "加载成功"
	c.JSON(http.StatusOK, res)
}

func UnLoadModelView(c *gin.Context) {
	/*
		# 卸载模型
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var res responseData
	var unLoadPackage unLoadPackageData
	var packageModel DataBaseModel.YssimModels
	err := c.BindJSON(&unLoadPackage)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", unLoadPackage.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel)
	if err != nil || packageModel.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	err = service.LoadAndDeleteLibrary(packageModel.PackageName, packageModel.Version, "", "unload")
	if err != nil {
		res.Err = err.Error()
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "卸载成功"
	c.JSON(http.StatusOK, res)
}

func GetPackageAndVersionView(c *gin.Context) {
	/*
		# 获取模型库和版本
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", userSpaceId}).Order("create_time desc").Find(&packageModel)
	var res responseData
	var data []map[string]string
	for i := 0; i < len(packageModel); i++ {
		d := map[string]string{
			"package_id":   packageModel[i].ID,
			"package_name": packageModel[i].PackageName,
			"version":      packageModel[i].Version,
			"sys_user":     packageModel[i].SysUser,
			"update_time":  packageModel[i].UpdatedAt.Format("06-01-02 15:04"),
		}
		if packageModel[i].SysUser == "sys" {
			d["update_time"] = "-"
		}
		data = append(data, d)
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetIconView(c *gin.Context) {
	/*
		# 获取模型的图标信息
	*/
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	var item modelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//var packageModel []DataBaseModel.YssimModels
	//DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).Order("create_time desc").Find(&packageModel)
	var res responseData
	data := service.GetIconNew(item.ModelName)
	res.Data = data
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
	//log.Println(data)
	var res responseData
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func Test1(c *gin.Context) {
	/*
		测试omc命令
	*/
	cmd := c.Query("cmd")
	data := service.GetGraphicsDataNew(cmd)
	//_ = omc.OMC.GetModelInstance(cmd)
	var res responseData
	res.Data = data
	c.JSON(http.StatusOK, res)
}
