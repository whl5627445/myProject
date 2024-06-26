package API

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"yssim-go/library/stringOperation"

	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/v1/service"
	"yssim-go/library/omc"

	"github.com/bytedance/sonic"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"yssim-go/config"
)

var dbModel = config.DB
var userName = config.USERNAME

func GetSysRootModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
	*/
	var res DataType.ResponseData
	keywords := c.Query("keywords")

	//
	// userSpaceId := c.Query("space_id")
	userSpaceId := c.GetHeader("space_id")
	var modelData []map[string]any
	var packageModel []DataBaseModel.YssimModels
	subQuery := dbModel.Model(&DataBaseModel.SystemLibrary{}).Where("username = ? AND encryption = ?", userName, true).Or("encryption = ?", false).Select("id")
	dbModel.Where("sys_or_user IN (?) AND userspace_id = ?", "sys", "0").Or("sys_or_user =  ? AND userspace_id = ? AND encryption = ?", userName, userSpaceId, true).
		Or("sys_or_user =  ? AND userspace_id = ? AND library_id IN (?)", userName, userSpaceId, subQuery).Find(&packageModel)
	libraryAndVersions := service.GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		p, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && p == packageModel[i].Version {
			data := map[string]any{}
			if keywords != "" {
				data = service.SearchPackage(packageModel[i], keywords)
			} else {
				data = map[string]any{
					"package_id":      packageModel[i].ID,
					"package_name":    packageModel[i].PackageName,
					"package_version": packageModel[i].Version,
					"model_name":      packageModel[i].PackageName,
					"haschild":        service.GetModelHasChild(packageModel[i].PackageName),
					"image":           "",
					"type":            service.GetModelType(packageModel[i].PackageName),
					"encryption":      packageModel[i].Encryption,
				}
			}
			if data != nil {
				modelData = append(modelData, data)
			}
		}
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)
}

func GetUserRootModelView(c *gin.Context) {
	/*
		# 获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的根节点列表，暂时没有图标信息
	*/

	userSpaceId := c.Query("space_id")
	if userSpaceId == "" {
		userSpaceId = c.GetHeader("space_id")
	}
	keywords := c.Query("keywords")
	var res DataType.ResponseData
	var modelData []map[string]any
	var packageModel []DataBaseModel.YssimModels
	var space DataBaseModel.YssimUserSpace
	dbModel.Where("id = ? AND username = ?", userSpaceId, userName).First(&space)
	subQuery := dbModel.Model(&DataBaseModel.SystemLibrary{}).Where("encryption = ?", false).Or("encryption = ? AND username = ?", true, userName).Select("id")
	dbModel.Where("sys_or_user = ? AND userspace_id = ? AND encryption = ? AND library_id = ''", userName, userSpaceId, false).
		Or("sys_or_user = ? AND userspace_id = ? AND library_id NOT IN (?)", userName, userSpaceId, subQuery).Find(&packageModel)
	libraryAndVersions := service.GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		loadVersions, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && loadVersions == packageModel[i].Version {
			data := map[string]any{}
			if keywords != "" {
				data = service.SearchPackage(packageModel[i], keywords)
			} else {
				data = map[string]any{
					"package_id":      packageModel[i].ID,
					"package_name":    packageModel[i].PackageName,
					"package_version": packageModel[i].Version,
					"model_name":      packageModel[i].PackageName,
					"haschild":        service.GetModelHasChild(packageModel[i].PackageName),
					"image":           "",
					"type":            service.GetModelType(packageModel[i].PackageName),
				}
			}
			if service.GetModelType(packageModel[i].PackageName) == "package" && len(data) > 0 {
				data["haschild"] = true
			}
			if data != nil {
				modelData = append(modelData, data)
			}
		}
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)
}

func GetUserPackageView(c *gin.Context) {
	/*
		# 获取已加载package列表接口
	*/

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var modelData []map[string]any
	var packageModel []DataBaseModel.YssimModels
	dbModel.Where("sys_or_user = ? AND userspace_id = ?", userName, userSpaceId).Find(&packageModel)
	libraryAndVersions := service.GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		loadVersions, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && loadVersions == packageModel[i].Version && service.GetModelType(packageModel[i].PackageName) == "package" {
			data := map[string]any{
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

	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var res DataType.ResponseData
	modelChildList := service.GetModelChild(modelName)
	var modelChildListNew []any
	for i := 0; i < len(modelChildList); i++ {
		modelChildListNew = append(modelChildListNew, modelChildList[i])
	}
	// 如果父节点是包名称的话，追加静态资源管理文件夹节点
	modelType := service.GetModelType(modelName)
	if modelName == packageModel.PackageName && packageModel.SysUser != "sys" && modelType == "package" && !packageModel.Encryption {
		modelChildListNew = append(modelChildListNew, map[string]any{
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
	var item DataType.ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	packageName := strings.Split(item.ModelName, ".")[0]
	err = dbModel.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", packageName, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var res DataType.ResponseData
	var graphicsData any
	if item.ComponentName == "" {
		permissions := userName
		if packageModel.Encryption || packageModel.SysUser == "sys" {
			permissions = "sys"
		}
		graphicsData = service.GetGraphicsData(item.ModelName, permissions)
	} else {
		graphicsData = service.GetComponentGraphicsData(item.ModelName, item.ComponentName)
	}

	res.Data = map[string]any{"encryption": packageModel.Encryption, "graphics": graphicsData}
	c.JSON(http.StatusOK, res)
}

func GetModelCodeView(c *gin.Context) {
	/*
		# 获取模型的源码数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
	*/

	userSpaceId := c.GetHeader("space_id")
	modelName := c.Query("model_name")
	packageId := c.Query("package_id")
	var res DataType.ResponseData
	if modelName == "" || userSpaceId == "" {
		c.JSON(http.StatusOK, res)
		return
	}
	var encryptionPackage DataBaseModel.YssimModels
	dbModel.Where("userspace_id = ?  AND sys_or_user = ? AND id = ?", userSpaceId, userName, packageId).First(&encryptionPackage)
	if encryptionPackage.Encryption {
		c.JSON(http.StatusOK, res)
		return
	}
	modelCode := service.GetModelCode(modelName)
	res.Data = modelCode
	c.JSON(http.StatusOK, res)
}

func ModelRename(c *gin.Context) {
	/*
		# 模型重命名
		## new_name: 新名字
		## model_name: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
		## package_id: 包的id
	*/
	var res DataType.ResponseData
	var item DataType.ModelRenameData
	userSpaceId := c.GetHeader("space_id")
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// 判断名称存不存在
	var packageRecord_ DataBaseModel.YssimModels
	dbModel.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", item.NewName, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageRecord_)
	if packageRecord_.PackageName != "" {
		res.Err = "名称已存在，请修改后再试。"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 判断是不是加密模型
	var packageRecord DataBaseModel.YssimModels
	dbModel.Where("userspace_id = ?  AND sys_or_user = ? AND id = ?", userSpaceId, userName, item.PackageId).First(&packageRecord)
	if packageRecord.Encryption || packageRecord.ID == "" {
		c.JSON(http.StatusOK, res)
		return
	}
	nameList := strings.Split(item.ModelName, ".")

	// 获取模型代码
	modelCode := service.GetModelCode(item.ModelName)

	// 将代码中的名称修改为新的名称
	modelCode = stringOperation.ModelRenam(modelCode, nameList[len(nameList)-1], item.NewName)

	// 保存代码 参考接口/update/package UpdateModelPackageView
	if len(nameList) > 1 {
		parentName := strings.Join(nameList[:len(nameList)-1], ".")
		modelCode = "within " + parentName + ";" + modelCode
	}

	modelPath := packageRecord.FilePath
	if strings.HasSuffix(packageRecord.FilePath, "/package.mo") {
		modelPath = service.GetSourceFile(item.ModelName)
	}
	parseResult, ok := service.ParseCodeString(modelCode, modelPath)
	if ok && len(parseResult) > 0 {
		isExist := service.IsExistPackage(parseResult)
		if isExist && (item.ModelName != parseResult) {
			res.Err = "模型名称重复"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
		loadResult := service.LoadCodeString(modelCode, modelPath)
		if loadResult {
			if parseResult != item.ModelName {
				// 判断是否是子模型
				if !strings.Contains(item.ModelName, ".") {
					dbModel.Model(DataBaseModel.YssimModels{}).Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).Update("package_name", parseResult)
				}
				service.DeleteLibrary(item.ModelName)
			}
			service.ModelSave(parseResult)
			res.Data = map[string]string{
				"id":        packageRecord.ID,
				"model_str": modelCode,
				"name":      parseResult,
			}
			var modelCollectionRecord DataBaseModel.YssimModelsCollection
			dbModel.Where("package_id = ? AND userspace_id = ? AND model_name = ? ", item.PackageId, userSpaceId, item.ModelName).First(&modelCollectionRecord)
			if modelCollectionRecord.ID != "" {
				modelCollectionRecord.ModelName = parseResult
				dbModel.Save(&modelCollectionRecord)
			}
			res.Msg = "模型保存成功"
			c.JSON(http.StatusOK, res)
			return
		}
	}
	res.Err = "语法错误，请重新检查"
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func GetModelResourcesReferenceView(c *gin.Context) {
	/*
		# 获取包级别的静态资源，以Reference的形式返回
		## package_id: 包id
		## parent: 需要查询的节点父级路径
		## path: 被查询节点
	*/

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.PackageResourcesData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if item.PackageId != "" {
		var packageModel DataBaseModel.YssimModels
		dbModel.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
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
		dbModel.Where("sys_or_user = ? AND userspace_id = ?", userName, userSpaceId).Find(&packageModelList)
		var data []map[string]any
		for _, model := range packageModelList {
			information := service.GetClassInformation(model.PackageName)
			if len(information) > 0 && information[0] == "package" {
				data = append(data, map[string]any{"name": model.PackageName + ".Resources", "package_id": model.ID, "type": "dir"})
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

	var res DataType.ResponseData
	properties := make(map[string]any, 0)
	parameters := service.GetModelParameters(modelName, componentName, className, "")
	elements := service.GetElements(modelName, componentName)
	if len(elements) > 0 && componentName != "" {
		dimension := elements[len(elements)-1].(string)
		properties = map[string]any{
			"model_name":     modelName,
			"component_name": componentName,
			"path":           elements[2],
			"dimension":      dimension,
			"annotation":     elements[4],
			"Properties":     []any{elements[6], elements[5], elements[9]},
			"Variability":    elements[10],
			"Inner/Outer":    elements[11],
			"Causality":      elements[12],
		}
	}
	res.Data = map[string]any{"parameters": parameters, "properties": properties}
	c.JSON(http.StatusOK, res)
}

func SetModelParametersView(c *gin.Context) {
	/*
		# 设置模型组件的参数数据，一次性返回
		## package_id: 模型包的id
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## parameter_value: 需要设置的变量和新的值，全称，例如{"PID.k": "200"}， k是模型的组件别名和变量名字的组成， 类似于“别名.变量名”
	*/
	var item DataType.SetComponentModifierValueData
	var res DataType.ResponseData

	userSpaceId := c.GetHeader("space_id")
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := false
	errParameterName := []string{}
	for _, parameter := range item.Parameter {
		if !parameter.IsExtend {
			result = service.SetElementModifierValue(item.ModelName, parameter.ParameterName, parameter.ParameterValue)
		} else {
			result = service.SetExtendsModifierValue(item.ModelName, parameter.ExtendName, parameter.ParameterName, parameter.ParameterValue)
		}
		if !result {
			errParameterName = append(errParameterName, parameter.ParameterName)
		}
	}

	if len(errParameterName) == 0 {
		service.ModelSave(item.ModelName)
		res.Msg = "设置完成"
	} else {
		res.Err = "设置失败: 请检查参数是否正确"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func SetModelParametersUnitView(c *gin.Context) {
	/*
		# 设置模型组件的参数数据，一次性返回
		## package_id: 模型包的id
		## model_name: 需要设置参数的模型名称，全称，例如“ENN.Examples.Scenario1_Status”
		## parameter_value: 需要设置的变量和新的值，全称，例如{"PID.k": "200"}， k是模型的组件别名和变量名字的组成， 类似于“别名.变量名”
	*/
	var item DataType.SetComponentUintData
	var res DataType.ResponseData

	userSpaceId := c.GetHeader("space_id")
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result := false
	errParameterName := []string{}

	for _, parameter := range item.UnitEditorData {
		if !parameter.IsExtend {
			result = service.SetElementModifierUnit(item.ModelName, parameter.ParameterName, parameter.ParameterUnit)
		} else {
			result = service.SetExtendsModifierUnit(item.ModelName, parameter.ExtendName, parameter.ParameterName, parameter.ParameterUnit)
		}
		if !result {
			errParameterName = append(errParameterName, parameter.ParameterName)
		}
	}

	if len(errParameterName) == 0 {
		service.ModelSave(item.ModelName)
		res.Msg = "设置完成"
	} else {
		res.Err = "设置失败: 请检查参数值或单位是否正确-" + strings.Join(errParameterName, ",")
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
	var res DataType.ResponseData

	userSpaceId := c.GetHeader("space_id")
	var item DataType.AddComponentParametersData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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
	var item DataType.DeleteComponentParametersData
	var res DataType.ResponseData

	userSpaceId := c.GetHeader("space_id")
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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
	var res DataType.ResponseData
	result := service.GetElements(modelName, componentsName)
	if len(result) > 0 {
		dimension := result[len(result)-1].(string)
		data := map[string]any{
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
		## comment: 组件注释
		## dimensions: 维数
	*/
	var res DataType.ResponseData
	var item DataType.SetComponentPropertiesData

	userSpaceId := c.GetHeader("space_id")
	err := c.BindJSON(&item)
	if err != nil {
		res.Status = 2
		res.Err = "设置失败"
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var packageModel DataBaseModel.YssimModels
	err = dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil || packageModel.SysUser == "sys" || packageModel.Encryption {
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
		item.Causality,
		item.Comment,
		item.Dimensions)
	if result {
		service.ModelSave(item.ModelName)
		res.Msg = "设置完成"
	} else {
		res.Err = strings.Join(msg, " - ")
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

	userSpaceId := c.GetHeader("space_id")
	var item DataType.CopyClassData
	var res DataType.ResponseData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	t := service.GetModelType(item.CopiedClassName)
	if t == "package" {
		res.Err = "包类型不允许复制，请继承使用"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	matchSpaceName1, _ := regexp.MatchString("^[_a-zA-Z0-9]+$", item.ModelName) // 字母、数字、下划线验证
	matchSpaceName2, _ := regexp.MatchString("^[a-zA-Z_]", item.ModelName)      // 字母、下划线验证
	if !matchSpaceName1 {
		res.Err = "名称只能由字母数字下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if !matchSpaceName2 {
		res.Err = "名称只能由字母下划线开头"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	packageName := item.ModelName
	if item.ParentName != "" {
		packageName = strings.Split(item.ParentName, ".")[0]
	}
	var encryptionPackage DataBaseModel.YssimModels
	dbModel.Where("sys_or_user = ? AND userspace_id = ? AND id = ? AND encryption = ?", userName, userSpaceId, item.FromPackageId, true).Or("sys_or_user = ? AND userspace_id = ? AND id = ? AND encryption = ?", userName, userSpaceId, item.ToPackageId, true).First(&encryptionPackage)
	if encryptionPackage.Encryption {
		res.Msg = "加密库不允许复制与插入模型"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var packageModel DataBaseModel.YssimModels
	dbModel.Where("package_name = ? AND userspace_id = ?", packageName, "0").Or("sys_or_user = ? AND userspace_id = ? AND package_name = ?", userName, userSpaceId, packageName).First(&packageModel)
	if packageModel.SysUser == "sys" && item.ParentName != "" {
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
	filePath := ""
	if item.ParentName != "" {
		packageName = packageModel.PackageName
		filePath = packageModel.FilePath
	} else {
		packageName = item.ModelName
		filePath = "static/UserFiles/UploadFile/" + userName + "/" + time.Now().Local().Format("20060102150405") + "/" + packageName + "/" + item.ModelName + ".mo"
	}

	newModel := DataBaseModel.YssimModels{
		ID:          uuid.New().String(),
		PackageName: packageName,
		SysUser:     userName,
		FilePath:    filePath,
		UserSpaceId: userSpaceId,
	}
	if item.ParentName == "" {
		err = dbModel.Create(&newModel).Error
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
			data["id"] = newModel.ID
			data["model_name"] = item.ModelName
		} else {
			data["id"] = packageModel.ID
			data["model_name"] = item.ParentName + "." + item.ModelName
		}
		res.Data = data
		packageInformation := service.GetPackageInformation()
		packageInformationJson, _ := sonic.Marshal(packageInformation)
		dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, userName).Update("package_information", packageInformationJson)
		if item.ParentName != "" {
			service.SetPackageUses(item.CopiedClassName, item.ParentName)
			service.ModelSave(item.ParentName + "." + item.ModelName)
		} else {
			service.SetPackageUses(item.CopiedClassName, item.ModelName)
			service.ModelSave(item.ModelName)
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

	userSpaceId := c.GetHeader("space_id")
	var item DataType.DeleteClassData
	err := c.BindJSON(&item)
	if err != nil {

		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var res DataType.ResponseData
	var packageModel DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result, msg := service.SaveModel(item.ModelName, "", item.ParentName, "delete", packageModel.FilePath)
	if result {
		res.Msg = msg
		if item.ParentName == "" {
			var simulateRecord []DataBaseModel.YssimSimulateRecord
			dbModel.Where("package_id = ? AND username = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).Find(&simulateRecord)
			dbModel.Delete(&packageModel)
		} else {
			service.ModelSave(item.ParentName)
		}
		var modelCollection []DataBaseModel.YssimModelsCollection
		dbModel.Where("package_id = ? AND model_name = ? AND userspace_id = ?", packageModel.ID, item.ModelName, userSpaceId).Find(&modelCollection)
		dbModel.Delete(&modelCollection)
		packageInformation := service.GetPackageInformation()
		packageInformationJson, _ := sonic.Marshal(packageInformation)
		dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, userName).Update("package_information", packageInformationJson)
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
	var res DataType.ResponseData
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.AddComponentData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	rotation := strconv.Itoa(item.Rotation)
	data := service.GetIconNew(item.OldComponentName, item.NewComponentName, false)
	graphics := data["graphics"].(map[string]any)
	graphics["originDiagram"] = item.Origin
	graphics["original_name"] = item.NewComponentName
	graphics["name"] = item.NewComponentName
	graphics["type"] = "Transformation"
	graphics["ID"] = "0"
	graphics["rotateAngle"] = graphics["rotation"]
	extentDiagram := service.GetModelExtentToString(graphics["coordinate_system"])
	data["graphics"] = graphics
	result, msg := service.AddComponent(item.NewComponentName, item.OldComponentName, item.ModelName, item.Origin, rotation, extentDiagram)
	if !result {
		res.Err = msg
		res.Status = 2
	} else {
		service.SetPackageUses(item.OldComponentName, item.ModelName)
		service.ModelSave(item.ModelName)
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.DeleteComponentData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.UpdateComponentData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)

	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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
	service.ModelSave(item.ModelName)
	res.Msg = "更新组件成功"
	c.JSON(http.StatusOK, res)
}

func BatchUpdateModelComponentView(c *gin.Context) {
	/*
		# 批量更新模型当中的模型组件
		## package_id： 包id
		## model_name: 需要更新的组件在哪个模型当中， 例如"Modelica.Blocks.Examples.PID_Controller"
		## batch_update_data 需要更新的组件参数数组集
	*/

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var data DataType.BatchUpdateComponentData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", data.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)

	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	for _, item := range data.UpdateComponentDatas {
		result := service.UpdateComponent(item.ComponentName, item.ComponentClassName, item.ModelName, item.Origin, item.Rotation, item.Extent)
		if !result {
			res.Err = item.ComponentName + "更新组件失败"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		} else {
			for _, connect := range item.ConnectionList {
				service.UpdateConnection(item.ModelName, connect.ConnectStart, connect.ConnectEnd, connect.Color, connect.LinePoints)
			}
		}
	}

	service.ModelSave(data.ModelName)
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.UpdateConnectionAnnotationData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)

	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.UpdateConnectionNamesData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)

	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.DeleteConnectionData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)

	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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
	res.Msg = "连线已删除"
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.UpdateConnectionAnnotationData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)

	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
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
	var res DataType.ResponseData
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
	var res DataType.ResponseData
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
	var res DataType.ResponseData
	modelName := c.Query("model_name")
	result := service.GetElements(modelName, "")
	var data []map[string]string
	for _, e := range result {
		component := map[string]string{
			"component_model_name":  e.([]any)[2].(string),
			"component_name":        e.([]any)[3].(string),
			"component_description": e.([]any)[4].(string),
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
	var res DataType.ResponseData
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.SetModelDocumentData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var packageModel DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&packageModel)

	if packageModel.SysUser == "sys" || packageModel.Encryption {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
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
	var res DataType.ResponseData
	var item DataType.ConvertUnitsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.ConvertUnits(item.S1, item.S2)
	unitsCompatible, _ := strconv.ParseBool(result[0])
	scaleFactor, _ := strconv.ParseFloat(result[1], 32)
	offset, _ := strconv.ParseFloat(result[2], 32)
	res.Data = map[string]any{
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
	var res DataType.ResponseData

	userSpaceId := c.GetHeader("space_id")
	var item DataType.ModelCollectionData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// 检测PackageId，userspace_id是否存在
	var packageModel DataBaseModel.YssimModels
	err = dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	// omc检测模型是否存在
	result := service.ExistClass(item.ModelName)
	if !result {
		c.JSON(http.StatusBadRequest, "model not found")
		return
	}
	// 检测数据库表中是否存在同名模型
	modelType := service.GetModelType(item.ModelName)
	// hasChild := service.GetModelChild(item.ModelName)
	if modelType == "package" {
		res.Err = "包类型或有子类的模型暂不允许收藏，请选择其他类型"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var modelCollection DataBaseModel.YssimModelsCollection
	dbModel.Where("model_name = ? AND userspace_id = ?", item.ModelName, userSpaceId).First(&modelCollection)
	if modelCollection.ID != "" {
		res.Err = "名称已存在，请修改后再试。"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	// 表中插入记录
	var newCollection = DataBaseModel.YssimModelsCollection{
		ID:          uuid.New().String(),
		PackageId:   item.PackageId,
		UserSpaceId: userSpaceId,
		ModelName:   item.ModelName,
	}
	err = dbModel.Create(&newCollection).Error
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

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var modelData []map[string]any
	var modelCollections []map[string]any
	dbModel.Raw("select mc.id, mc.package_id, m.package_name, mc.model_name, m.version, m.sys_or_user from yssim_models_collections as mc  left join yssim_models m on mc.package_id = m.id where mc.userspace_id = ?  and m.sys_or_user IN (?,\"sys\") and mc.deleted_at is NULL", userSpaceId, userName).Scan(&modelCollections)
	for i := 0; i < len(modelCollections); i++ {
		modelName := modelCollections[i]["model_name"].(string)
		packageName := modelCollections[i]["package_name"].(string)
		version := modelCollections[i]["version"].(string)
		sysOrUser := false
		if modelCollections[i]["sys_or_user"].(string) == "sys" {
			sysOrUser = true
		}
		// 检测模型是否存在，不存在就从表中删除记录
		result := service.ExistClass(modelName)
		if !result {
			go dbModel.Delete(&modelCollections[i])
			continue
		}
		data := map[string]any{
			"id":           modelCollections[i]["id"],
			"package_id":   modelCollections[i]["package_id"],
			"package_name": modelCollections[i]["package_name"],
			"model_name":   modelCollections[i]["model_name"],
			"haschild":     false,
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
	//
	// userSpaceId := c.GetHeader("space_id")

	id := c.Query("id")
	var res DataType.ResponseData
	var modelCollection DataBaseModel.YssimModelsCollection
	dbModel.Where("id = ?", id).First(&modelCollection)
	if modelCollection.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	dbModel.Delete(&modelCollection)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func SearchModelView(c *gin.Context) {
	/*
		# 搜索模型
		## keywords: 需要搜索的关键字
		## parent: 需要搜索的关键字的父节点
	*/

	userSpaceId := c.GetHeader("space_id")
	// userSpaceId := c.Query("space_id")
	keywords := c.Query("keywords")
	parent := c.Query("parent")
	var res DataType.ResponseData
	var data []map[string]any
	var packageModel []DataBaseModel.YssimModels
	libraryAndVersions := service.GetLibraryAndVersions()
	if parent != "" {
		namesList := strings.Split(parent, ".")
		packageName := namesList[0]
		dbModel.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", packageName, []string{"sys", userName}, []string{"0", userSpaceId}).Order("sys_or_user desc").Find(&packageModel)
	} else {
		dbModel.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", userName}, []string{"0", userSpaceId}).Order("sys_or_user desc").Find(&packageModel)
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

func SearchFunctionTypeView(c *gin.Context) {
	/*
		# 搜索类型为function的模型
		## parent: 需要搜索的模型的父节点
	*/
	parent := c.Query("parent")
	var res DataType.ResponseData
	var models []DataBaseModel.YssimModels

	userSpaceId := c.GetHeader("space_id")
	var data []map[string]any
	libraryAndVersions := service.GetLibraryAndVersions()
	if parent == "" {
		dbModel.Where("sys_or_user IN (?) AND userspace_id IN (?)", []string{userName, "sys"}, []string{userSpaceId, "0"}).Find(&models)

		for _, model := range models {
			p, ok := libraryAndVersions[model.PackageName]
			if ok && p == model.Version {
				modelType := service.GetModelType(model.PackageName)
				if modelType == "package" || modelType == "function" || modelType == "record" {
					top := map[string]any{
						"name":       model.PackageName,
						"model_name": model.PackageName,
						"haschild":   modelType == "package",
						"type":       modelType,
						"package_id": model.ID,
					}
					data = append(data, top)
				}
			}
		}
	} else {
		modelNameList := service.SearchFunctionType(parent)
		data = append(data, modelNameList...)
	}

	res.Data = data
	c.JSON(http.StatusOK, res)
}

func LoadModelView(c *gin.Context) {
	/*
		# 加载模型
	*/

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var loadPackage DataType.LoadPackageData
	var packageModel DataBaseModel.YssimModels
	err := c.BindJSON(&loadPackage)
	if err != nil {
		log.Println(err)
	}
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", loadPackage.PackageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel)
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
	dependentLibrary := service.GetPackageUses(packageModel.PackageName)
	for i := 0; i < len(dependentLibrary); i++ {
		var p DataBaseModel.YssimModels
		dbModel.Where("package_name = ? AND version = ? AND sys_or_user = ? AND userspace_id = ?", dependentLibrary[i][0], dependentLibrary[i][1], userName, userSpaceId).First(&p)
		loadPackageMap := service.GetLibraryAndVersions()
		l, ok := loadPackageMap[p.PackageName]
		if p.ID != "" && (!ok || l != p.Version) {
			service.LoadAndDeleteLibrary(p.PackageName, p.Version, p.FilePath, "load")
		}
	}
	if err != nil {
		res.Err = err.Error()
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	packageInformation := service.GetPackageInformation()
	packageInformationJson, _ := sonic.Marshal(packageInformation)
	dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, userName).Update("package_information", packageInformationJson)
	res.Msg = "加载成功"
	c.JSON(http.StatusOK, res)
}

func UnLoadModelView(c *gin.Context) {
	/*
		# 卸载模型
	*/

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var unLoadPackage DataType.UnLoadPackageData
	var packageModel DataBaseModel.YssimModels
	err := c.BindJSON(&unLoadPackage)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", unLoadPackage.PackageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel)
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
	packageInformation := service.GetPackageInformation()
	packageInformationJson, _ := sonic.Marshal(packageInformation)
	dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, userName).Update("package_information", packageInformationJson)
	c.JSON(http.StatusOK, res)
}

func GetIconView(c *gin.Context) {
	/*
		# 获取模型的图标信息
	*/
	//
	// userSpaceId := c.GetHeader("space_id")
	var item DataType.ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// var packageModel []DataBaseModel.YssimModels
	// dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).Order("create_time desc").Find(&packageModel)
	var res DataType.ResponseData
	data := service.GetIconNew(item.ModelName, "", true)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func LoginUserSpaceView(c *gin.Context) {
	/*
		# 进入用户空间
		## space_id: 用户空间id
	*/
	var res DataType.ResponseData
	var item DataType.LoginUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}

	// result := service.SetWorkSpaceId(&item.SpaceId)
	// if result {
	//	res.Msg = "初始化完成"
	//	c.JSON(http.StatusOK, res)
	//	return
	// }
	var space DataBaseModel.YssimUserSpace
	dbModel.Model(space).Where("id = ? AND username = ?", item.SpaceId, userName).First(&space)
	if space.ID == "" {
		res.Status = 4
		res.Err = "空间已被删除"
		c.JSON(http.StatusOK, res)
		return
	}
	var packageModelAll []DataBaseModel.YssimModels
	dbModel.Where("sys_or_user IN ?  AND default_version = ? AND userspace_id IN ?", []string{"sys", userName}, true, []string{"0", item.SpaceId}).Find(&packageModelAll)

	dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", item.SpaceId, userName).UpdateColumn("last_login_time", time.Now().Local().Unix())

	information, _ := space.PackageInformation.MarshalJSON()
	packageInformation := map[string]map[string]string{}
	sonic.Unmarshal(information, &packageInformation)
	service.LibraryInitialization(packageInformation, packageModelAll)

	res.Msg = "初始化完成"
	c.JSON(http.StatusOK, res)
}

func Test(c *gin.Context) {
	/*
		测试omc命令
	*/
	cmd := c.Query("cmd")
	NoParsed := c.Query("NoParsed")
	var data any
	if NoParsed != "" {
		data, _ = omc.OMC.SendExpression(cmd)
	} else {
		d, _ := omc.OMC.SendExpressionNoParsed(cmd)
		data = string(d)
	}
	// log.Println(data)
	var res DataType.ResponseData
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetUMLView(c *gin.Context) {
	/*
		根据路径名获取模型的uml图
	*/
	var res DataType.ResponseData
	var className = c.Query("className")
	var result = map[string]interface{}{}
	modelType := service.GetModelType(className)
	if modelType == "model" || modelType == "block" || modelType == "connector" || modelType == "record" {
		result["dataType"] = "model"
		finalResultData := service.GetModelUMLData(className)
		result["resultData"] = finalResultData
	} else {
		result["dataType"] = "package"
		resultData := service.GetPackageUMLData(className)
		result["resultData"] = resultData
	}
	res.Msg = "获取成功"
	res.Data = result
	c.JSON(http.StatusOK, res)
}

func Test1(c *gin.Context) {
	/*
		测试omc命令
	*/
	cmd := c.Query("cmd")
	// data := service.GetGraphicsDataNew(cmd)
	data, _ := omc.OMC.SendExpression(cmd)
	var res DataType.ResponseData
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetExtendedModelView(c *gin.Context) {
	/*
		# 获取模型继承的父类
	*/

	var res DataType.ResponseData
	var models DataBaseModel.YssimModels

	spaceId := c.GetHeader("space_id")
	modelName := c.Query("model_name")
	if strings.TrimSpace(modelName) == "" {
		res.Err = "参数为空"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	dataList := service.GetExtendedModel(modelName)
	if strings.Contains(modelName, ".") {
		modelName = strings.Split(modelName, ".")[0]
	}
	err := dbModel.Where("package_name = ? AND sys_or_user IN (?) AND userspace_id IN (?)", modelName, []string{userName, "sys"}, []string{spaceId, "0"}).First(&models).Error
	if err != nil {
		res.Err = "未查询到模型"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var data []map[string]interface{}
	if dataList != nil {
		for _, str := range dataList {
			dataList2 := service.GetExtendedModel(str)
			temp := map[string]interface{}{
				"model_name": str,
				"flag":       dataList2 != nil,
				"version":    models.Version,
				"models_id":  models.ID,
				"user_name":  models.SysUser,
				"encryption": models.Encryption,
			}
			if models.UserSpaceId == "0" {
				temp["tag"] = "sys"
			} else {
				temp["tag"] = "user"
			}
			data = append(data, temp)
		}
		res.Data = data
		res.Msg = "查询成功"
	} else {
		res.Err = "此模型没有父类"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}
