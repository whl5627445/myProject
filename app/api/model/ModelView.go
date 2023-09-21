package API

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/service"
	"yssim-go/library/fileOperation"
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
	//userSpaceId := c.Query("space_id")
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
	data := service.GetIconNew(item.OldComponentName, false)
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
	//检测PackageId，userspace_id是否存在
	var packageModel DataBaseModel.YssimModels
	err = dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel).Error
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
	//hasChild := service.GetModelChild(item.ModelName)
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
	//表中插入记录
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
		//检测模型是否存在，不存在就从表中删除记录
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
	//userSpaceId := c.GetHeader("space_id")

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
	//userSpaceId := c.Query("space_id")
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

func GetPackageAndVersionView(c *gin.Context) {
	/*
		# 获取模型库和版本
	*/

	userSpaceId := c.GetHeader("space_id")
	var packageModel []DataBaseModel.YssimModels
	var ids []string
	dbModel.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", userName}, []string{"0", userSpaceId}).Order("create_time desc").Find(&packageModel)
	dbModel.Model(&DataBaseModel.SystemLibrary{}).Where("username = ? AND encryption = ?", userName, true).Or("encryption = ?", false).Select("id").Scan(&ids)
	var res DataType.ResponseData
	var data []map[string]string
	for i := 0; i < len(packageModel); i++ {
		d := map[string]string{
			"package_id":   packageModel[i].ID,
			"package_name": packageModel[i].PackageName,
			"version":      packageModel[i].Version,
			"update_time":  packageModel[i].UpdatedAt.Format("06-01-02 15:04"),
		}
		for _, id := range ids {
			libraryId := packageModel[i].LibraryId
			if libraryId == id {
				d["sys_user"] = "sys"
				break
			} else {
				d["sys_user"] = packageModel[i].SysUser
			}
		}
		if packageModel[i].SysUser == "sys" || d["sys_user"] == "sys" {
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
	//
	//userSpaceId := c.GetHeader("space_id")
	var item DataType.ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//var packageModel []DataBaseModel.YssimModels
	//dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).Order("create_time desc").Find(&packageModel)
	var res DataType.ResponseData
	data := service.GetIconNew(item.ModelName, true)
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

	//result := service.SetWorkSpaceId(&item.SpaceId)
	//if result {
	//	res.Msg = "初始化完成"
	//	c.JSON(http.StatusOK, res)
	//	return
	//}
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
	//log.Println(data)
	var res DataType.ResponseData
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func AppModelMarkView(c *gin.Context) {
	/*
		# 标记模型为可用数据源,(执行一次仿真)
		开发人： 徐庆达
	*/
	var res DataType.ResponseData

	userSpaceId := c.GetHeader("space_id")
	var item DataType.AppModelMarkData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("导出数据源数据错误： ", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// 检测数据源名称是否重复
	var record DataBaseModel.AppDataSource
	dbModel.Where("package_id = ? AND username = ? AND group_name = ? AND data_source_name = ?", item.PackageId, userName, item.GroupName, item.DataSourceName).First(&record)
	if record.ID != "" {
		res.Err = "名称重复"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	CompilePath := "static/UserFiles/modelDataSource/" + userName + "/" + strings.ReplaceAll(item.ModelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/"
	// 数据源表中创建一条记录
	dataSource := DataBaseModel.AppDataSource{
		ID:             uuid.New().String(),
		UserName:       userName,
		UserSpaceId:    userSpaceId,
		PackageId:      item.PackageId,
		ModelName:      item.ModelName,
		CompilePath:    CompilePath,
		CompileType:    item.CompileType,
		ExperimentId:   item.ExperimentId,
		GroupName:      item.GroupName,
		DataSourceName: item.DataSourceName,
		CompileStatus:  0,
	}
	err = dbModel.Create(&dataSource).Error
	record = dataSource
	if err != nil {
		log.Println("标记数据源时创建数据库记录失败： ", err)
		res.Status = 2
		res.Err = "创建失败"
		c.JSON(http.StatusOK, res)
		return
	}
	_, err = service.GrpcTranslate(record)
	if err != nil {
		log.Println("提交任务失败： ", err)
		res.Status = 2
		res.Err = "创建失败"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "请等待导出完成。"
	c.JSON(http.StatusOK, res)

}

func CADParseView(c *gin.Context) {
	/*
		解析CAD文件
	*/
	var res DataType.ResponseData
	var model DataBaseModel.YssimModels

	dbModel.Where("package_name = ? AND version = ?", "Modelica", "4.0.0").First(&model)
	var item DataType.FilePathData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("参数验证失败： ", err)
		c.JSON(http.StatusBadRequest, "参数错误")
		return
	}
	data := service.GetXmlData(item.FilePath, userName)
	if data == "" {
		res.Err = "文件解析失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	modelName := map[string]any{"straight_tube": map[string]any{"id": model.ID, "model_name": []string{"Modelica.Fluid.Pipes.StaticPipe", "Modelica.Fluid.Pipes.DynamicPipe"}}, "bendable_tube": map[string]any{"id": model.ID, "model_name": []string{"Modelica.Fluid.Fittings.Bends.CurvedBend"}}}
	components := service.CADParseParts(data)

	res.Data = map[string]any{"components": components, "model": modelName}

	c.JSON(http.StatusOK, res)
}

func CADParseXmlView(c *gin.Context) {
	/*
		解析CAD文件
	*/
	var res DataType.ResponseData
	var model DataBaseModel.YssimModels
	file, err := c.FormFile("xmlFile")
	if err != nil {
		res.Err = "文件上传失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	xml := service.HandleXMLUpload(file)
	if xml == "" {
		res.Err = "文件解析失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	dbModel.Where("package_name = ? AND version = ?", "Modelica", "4.0.0").First(&model)
	modelName := map[string]any{"straight_tube": map[string]any{"id": model.ID, "model_name": []string{"Modelica.Fluid.Pipes.StaticPipe", "Modelica.Fluid.Pipes.DynamicPipe"}}, "bendable_tube": map[string]any{"id": model.ID, "model_name": []string{"Modelica.Fluid.Fittings.Bends.CurvedBend"}}}

	components := service.CADParseParts(xml)

	res.Data = map[string]any{"components": components, "model": modelName}

	c.JSON(http.StatusOK, res)
}

func CADFilesUploadView(c *gin.Context) {
	var res DataType.ResponseData

	form, err := c.MultipartForm()
	if err != nil {
		res.Err = "文件上传失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Data = service.CadFilesUpload(form, userName)
	res.Msg = "文件上传成功"
	c.JSON(http.StatusOK, res)
}

func CADMappingModelView(c *gin.Context) {
	/*
		利用前端传回的CAD解析数据进行模型映射
	*/

	var res DataType.ResponseData

	userSpaceId := c.GetHeader("space_id")
	var item DataType.CADMappingModelData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("导出数据源数据错误： ", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var packageModel DataBaseModel.YssimModels
	dbModel.Where("sys_or_user = ?  AND userspace_id = ? AND id = ?", userName, userSpaceId, item.PackageId).First(&packageModel)
	if packageModel.ID == "" {
		res.Err = "建模失败， 请检查当前模型是否为您的自建模型"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if len(item.ModelMapping) != len(item.Information) {
		res.Err = "数据错误， 所选零件数量与CAD解析数据不等"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	for i := 0; i < len(item.ModelMapping); i++ {
		if len(item.ModelMapping[i].ModelName) != len(item.Information[i].ModelInformation) {
			res.Err = "建模失败， 映射模型数据与零件对应组件数量不等"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	}
	for i := 0; i < len(item.ModelMapping); i++ {
		service.CADMappingModel(item.ModelName, item.ModelMapping[i].ModelName, item.Information[i])
	}

	res.Msg = "建模完成"
	c.JSON(http.StatusOK, res)
}

func GetSystemLibraryView(c *gin.Context) {
	var res DataType.ResponseData
	var system []DataBaseModel.SystemLibrary

	spaceId := c.Query("space_id")
	if spaceId == "" {
		res.Err = "参数不能为空"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	subQuery := dbModel.Model(&DataBaseModel.YssimModels{}).Where("sys_or_user = ? AND userspace_id = ? AND library_id IS NOT NULL AND library_id != ?", userName, spaceId, "").Select("library_id")
	err := dbModel.Where("encryption = ? AND id NOT IN (?)", false, subQuery).Or("username = ? AND id NOT IN (?)", userName, subQuery).Order("create_time desc").Find(&system).Error
	if err != nil {
		res.Status = 2
		res.Err = "未查询到系统模型"
		c.JSON(http.StatusOK, res)
		return
	}
	var data []map[string]any
	for i := 0; i < len(system); i++ {
		d := map[string]any{
			"id":              system[i].ID,
			"user_name":       system[i].UserName,
			"package_name":    system[i].PackageName,
			"file_path":       system[i].FilePath,
			"version_control": system[i].VersionControl,
			"version_branch":  system[i].VersionBranch,
			"version_tag":     system[i].VersionTag,
			"encryption":      system[i].Encryption,
		}
		data = append(data, d)
	}
	res.Msg = "查询成功"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func DeleteDependencyLibraryView(c *gin.Context) {
	var res DataType.ResponseData
	var model []DataBaseModel.YssimModels
	var item DataType.DeleteDependencyLibraryData
	e := c.BindJSON(&item)
	if e != nil {
		c.JSON(http.StatusBadRequest, "参数验证失败")
		return
	}
	err := dbModel.Where("id IN (?)", item.Ids).Find(&model).Error
	if err != nil {
		res.Status = 2
		res.Err = "删除失败，查看id是否正确"
		c.JSON(http.StatusOK, res)
		return
	}
	dbModel.Delete(&model)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func CreateDependencyLibraryView(c *gin.Context) {
	var res DataType.ResponseData
	var item DataType.CreateDependencyLibraryData
	var models DataBaseModel.YssimModels
	var system DataBaseModel.SystemLibrary

	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数验证失败")
		return
	}
	dbModel.Where("id = ? AND username = ?", item.ID, item.UserName).First(&system)
	dbModel.Where("sys_or_user = ? AND library_id = ? AND userspace_id = ?", userName, system.ID, item.SpaceId).First(&models)
	if models.ID != "" {
		res.Err = "模型已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var newDependencyLibrary = DataBaseModel.YssimModels{
		ID:             uuid.New().String(),
		LibraryId:      system.ID,
		PackageName:    system.PackageName,
		Version:        system.Version,
		SysUser:        userName,
		FilePath:       system.FilePath,
		UserSpaceId:    item.SpaceId,
		VersionControl: system.VersionControl,
		VersionBranch:  system.VersionBranch,
		VersionTag:     system.VersionTag,
		Encryption:     system.Encryption,
	}
	err = dbModel.Create(&newDependencyLibrary).Error
	if err != nil {
		res.Status = 2
		res.Err = "创建失败"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "创建成功"
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

func GetDependencyLibraryView(c *gin.Context) {

	var res DataType.ResponseData
	var model []DataBaseModel.YssimModels

	userSpaceId := c.Query("space_id")
	if userSpaceId == "" {
		res.Err = "参数不能为空"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	subQuery := dbModel.Model(&DataBaseModel.SystemLibrary{}).Where("username = ? AND encryption = ?", userName, true).Or("encryption = ?", false).Select("id")
	err := dbModel.Where("userspace_id = ? AND sys_or_user != ? AND sys_or_user = ? AND library_id IS NOT NULL AND library_id != ? AND library_id IN (?)", userSpaceId, "sys", userName, "", subQuery).Order("create_time desc").Find(&model).Error
	if err != nil {
		res.Status = 2
		res.Err = "无依赖模型"
		c.JSON(http.StatusOK, res)
		return
	}
	var data []map[string]any
	for i := 0; i < len(model); i++ {
		d := map[string]any{
			"id":              model[i].ID,
			"package_name":    model[i].PackageName,
			"version":         model[i].Version,
			"sys_or_user":     model[i].SysUser,
			"file_path":       model[i].FilePath,
			"version_control": model[i].VersionControl,
			"version_branch":  model[i].VersionBranch,
			"version_tag":     model[i].VersionTag,
			"default_version": model[i].Default,
		}
		data = append(data, d)
	}
	res.Msg = "查询成功"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetAvailableLibrariesView(c *gin.Context) {
	/*
		根据username used 查询可用库列表  0未占用 1占用  查询used为0
	*/
	var res DataType.ResponseData

	var userLibraries []DataBaseModel.UserLibrary
	dbModel.Where("username = ? AND used = ?", userName, false).Find(&userLibraries)
	var data []map[string]any
	for _, library := range userLibraries {
		d := map[string]any{
			"id":             library.ID,
			"package_name":   library.PackageName,
			"version_branch": library.VersionBranch,
		}
		data = append(data, d)
	}
	res.Msg = "查询成功"
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

func GetVersionAvailableLibrariesView(c *gin.Context) {
	/*
		根据sys_or_user  userspace_id   查询可编辑的模型库
	*/
	var res DataType.ResponseData
	sysOrUser := c.GetHeader("username")
	userspaceId := c.Query("space_id")
	var versionModels []DataBaseModel.YssimModels
	var noVersionModels []DataBaseModel.YssimModels
	subQuery := dbModel.Model(&DataBaseModel.SystemLibrary{}).Where("username = ? AND encryption = ?", sysOrUser, true).Or("encryption = ?", false).Select("id")
	dbModel.Where("sys_or_user = ? AND userspace_id = ? AND version_control = ? AND library_id not in(?)", sysOrUser, userspaceId, true, subQuery).Find(&versionModels)
	dbModel.Where("sys_or_user = ? AND userspace_id = ? AND  library_id = ?", sysOrUser, userspaceId, "").Find(&noVersionModels)

	result := make(map[string][]DataType.GetVersionLibraryData)
	result["version"] = []DataType.GetVersionLibraryData{}
	result["noVersion"] = []DataType.GetVersionLibraryData{}
	for _, model := range versionModels {
		versionLibraryData := DataType.GetVersionLibraryData{
			Id:             model.ID,
			PackageName:    model.PackageName,
			LibraryId:      model.LibraryId,
			VersionControl: model.VersionControl,
			VersionBranch:  model.VersionBranch,
			Version:        model.Version,
		}
		result["version"] = append(result["version"], versionLibraryData)
	}
	for _, model := range noVersionModels {
		versionLibraryData := DataType.GetVersionLibraryData{
			Id:             model.ID,
			PackageName:    model.PackageName,
			VersionControl: model.VersionControl,
			VersionBranch:  model.VersionBranch,
		}
		result["noVersion"] = append(result["noVersion"], versionLibraryData)
	}
	res.Data = result
	c.JSON(http.StatusOK, res)
}

func DeleteVersionAvailableLibrariesView(c *gin.Context) {
	/*
		根据sys_or_user  userspace_id  删除可编辑模型库
	*/
	var res DataType.ResponseData
	var username = c.GetHeader("username")
	var deleteLibrary DataType.DeleteVersionLibraryData
	err := c.BindJSON(&deleteLibrary)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if deleteLibrary.VersionControl {
		var userLibrary DataBaseModel.UserLibrary
		dbModel.Where("id = ? AND username = ?", deleteLibrary.LibraryId, username).First(&userLibrary)
		if userLibrary.ID == "" {
			res.Status = 2
			res.Err = "删除失败,未查询到该模型"
			c.JSON(http.StatusOK, res)
			return
		}
		dbModel.Model(&userLibrary).Update("used", false)
	}
	var yssimModel DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user = ? AND userspace_id = ? ", deleteLibrary.Id, username, deleteLibrary.SpaceId).First(&yssimModel)
	if yssimModel.ID == "" {
		res.Status = 2
		res.Err = "删除失败，未查询到该模型"
		c.JSON(http.StatusOK, res)
		return
	}
	dbModel.Delete(&yssimModel)
	packageInformation := service.GetPackageInformation()
	service.DeleteLibrary(yssimModel.PackageName)
	delete(packageInformation, yssimModel.PackageName)
	packageInformationJson, _ := sonic.Marshal(packageInformation)
	dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", deleteLibrary.SpaceId, username).Update("package_information", packageInformationJson)

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)

}

func CreateVersionAvailableLibrariesView(c *gin.Context) {
	/*
			创建可编辑有版本的模型库
			先检查模型是否已加载，如果已加载，就卸载掉
		    加载mo文件，获取版本号，如果获取成功更新数据库，获取失败，返回解析mo文件失败
	*/
	var res DataType.ResponseData

	var username = c.GetHeader("username")

	var createLibrary DataType.CreateVersionLibraryData
	err := c.BindJSON(&createLibrary)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var userLibrary DataBaseModel.UserLibrary
	dbModel.Where("id = ? AND username = ? ", createLibrary.Id, username).First(&userLibrary)
	if userLibrary.ID == "" {
		res.Status = 2
		res.Err = "未查询到该模型,添加失败"
		c.JSON(http.StatusOK, res)
		return
	}
	var yssimModels []DataBaseModel.YssimModels
	dbModel.Where("sys_or_user = ? AND userspace_id = ?", username, createLibrary.SpaceId).Find(&yssimModels)
	for _, model := range yssimModels {
		if model.PackageName == userLibrary.PackageName {
			res.Status = 2
			res.Err = "该空间下已有同名的模型"
			c.JSON(http.StatusOK, res)
			return
		}
	}
	flag := service.ExistClass(userLibrary.PackageName)
	if flag {
		service.DeleteLibrary(userLibrary.PackageName)
	}
	flag = service.LoadPackage(userLibrary.PackageName, "", userLibrary.FilePath)
	if !flag {
		res.Status = 2
		res.Err = "加载模型失败"
		c.JSON(http.StatusOK, res)
		return
	}
	version := service.GetVersion(userLibrary.PackageName)
	var newVersionLibrary = DataBaseModel.YssimModels{
		ID:             uuid.New().String(),
		LibraryId:      userLibrary.ID,
		PackageName:    userLibrary.PackageName,
		Version:        version,
		SysUser:        username,
		FilePath:       userLibrary.FilePath,
		UserSpaceId:    createLibrary.SpaceId,
		VersionControl: userLibrary.VersionControl,
		VersionBranch:  userLibrary.VersionBranch,
		VersionTag:     userLibrary.VersionTag,
	}
	err = dbModel.Create(&newVersionLibrary).Error
	if err != nil {
		res.Status = 2
		res.Err = "创建失败"
		c.JSON(http.StatusOK, res)
		return
	}
	dbModel.Model(&userLibrary).Update("used", true)
	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func InitVersionControlView(c *gin.Context) {
	/*
		将无版本控制的包初始化为有版本控制
	*/
	var res DataType.ResponseData
	var item DataType.InitVersionControlData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// 检查是否已经存在
	var record DataBaseModel.UserLibrary
	dbModel.Where("username = ? AND repository_address = ?", userName, item.RepositoryAddress).First(&record)
	if record.ID != "" {
		res.Err = "该存储库已经存在！"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	//查询无版本控制的包
	var noVersionRecord DataBaseModel.YssimModels
	dbModel.Where("id = ?", item.NoVersionPackageId).First(&noVersionRecord)
	var errorMessage string
	// 克隆到本地
	repositoryPath, _, errorMessage, cloneRes := service.RepositoryClone(item.RepositoryAddress, "", userName)
	if cloneRes {
		//克隆成功
		// 将无版本控制的包添加到有版本控制
		addVersionRes, msg := service.InitVersionControl(noVersionRecord.FilePath, repositoryPath, item.UserName, item.PassWord)
		errorMessage = msg
		if addVersionRes {
			packageName, packagePath, msg_, ok := service.GitPackageFileParse(noVersionRecord.PackageName, repositoryPath)
			errorMessage = msg_
			if ok {
				//分支名称默认是master
				versionBranch := "master"
				// 获取克隆到本地的存储库的tag
				versionTag := service.GetTag(repositoryPath)
				anotherName, _ := service.GetRepositoryName(item.RepositoryAddress)
				versionRecord := DataBaseModel.UserLibrary{
					ID:                uuid.New().String(),
					UserName:          userName,
					PackageName:       packageName,            //package名称，一般称为包名或库的名字
					FilePath:          packagePath,            //package所在路径
					VersionControl:    true,                   //是否有版本控制
					VersionBranch:     versionBranch,          //版本控制分支
					VersionTag:        versionTag,             //版本控制tag
					AnotherName:       anotherName,            // 别名
					RepositoryAddress: item.RepositoryAddress, //存储库地址
				}
				err = dbModel.Create(&versionRecord).Error
				dbModel.Delete(&noVersionRecord)
				res.Msg = "初始化成功!"
				c.JSON(http.StatusOK, res)
				return
			}
		}
	}
	os.RemoveAll(repositoryPath)
	res.Err = errorMessage
	res.Status = 2
	c.JSON(http.StatusOK, res)
}

func RepositoryCloneView(c *gin.Context) {
	/*
		添加存储库接口：根据git地址拉去存储库
	*/
	var res DataType.ResponseData
	var item DataType.RepositoryCloneData

	//userSpaceId := c.GetHeader("space_id")
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 检查是否已经存在
	var record DataBaseModel.UserLibrary
	dbModel.Where("username = ? AND repository_address = ?", userName, item.RepositoryAddress).First(&record)
	if record.ID != "" {
		res.Err = "该存储库已经存在！"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var errorMessage string
	// 克隆到本地
	repositoryPath, repositoryName, errorMessage, cloneRes := service.RepositoryClone(item.RepositoryAddress, item.Branch, userName)

	if cloneRes { //克隆成功
		//分支名称默认是master
		versionBranch := "master"
		if item.Branch != "" {
			versionBranch = item.Branch
		}
		// 获取克隆到本地的存储库的tag
		versionTag := service.GetTag(repositoryPath)

		// 解析包文件
		packageName, packagePath, msg_, ok := service.GitPackageFileParse(repositoryName, repositoryPath)
		errorMessage = msg_

		if ok { // 创建数据库记录
			libraryRecord := DataBaseModel.UserLibrary{
				ID:                uuid.New().String(),
				UserName:          userName,
				PackageName:       packageName,            //package名称，一般称为包名或库的名字
				FilePath:          packagePath,            //package所在路径
				VersionControl:    true,                   //是否有版本控制
				VersionBranch:     versionBranch,          //版本控制分支
				VersionTag:        versionTag,             //版本控制tag
				AnotherName:       item.Name,              // 别名
				RepositoryAddress: item.RepositoryAddress, //存储库地址
				//Version:     packageVersion, //package版本号
				//Used:           bool           			//是否已经被某空间使用

			}
			err = dbModel.Create(&libraryRecord).Error
			res.Msg = "拉取成功"
			c.JSON(http.StatusOK, res)
			return
		}
	}
	res.Err = errorMessage
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func RepositoryDeleteView(c *gin.Context) {
	/*
		删除存储库
	*/
	var res DataType.ResponseData
	var item DataType.RepositoryDeleteData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// 删除UserLibrary数据库记录
	var userLibraryRecord DataBaseModel.UserLibrary
	dbModel.Where("id = ?", item.ID).First(&userLibraryRecord)
	dbModel.Delete(&userLibraryRecord)
	// 删除YssimModels数据库记录
	var yssimModelsRecord []DataBaseModel.YssimModels
	dbModel.Where("library_id = ?", item.ID).Find(&yssimModelsRecord)
	dbModel.Delete(&yssimModelsRecord)
	// 删除包文件
	fileOperation.DeleteProjectPath(userLibraryRecord.FilePath)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func RepositoryGetView(c *gin.Context) {
	/*
		获取存储库列表
	*/
	var res DataType.ResponseData

	// 删除数据库记录
	var records []DataBaseModel.UserLibrary
	dbModel.Where("username = ? ", userName).Order("create_time desc").Find(&records)
	var data []map[string]any
	for i := 0; i < len(records); i++ {
		d := map[string]any{
			"id":           records[i].ID,
			"package_name": records[i].PackageName,
			//"version":            records[i].Version,
			"another_name":       records[i].AnotherName,
			"repository_address": records[i].RepositoryAddress,
			"version_branch":     records[i].VersionBranch,
			//"create_time":        records[i].CreatedAt,
		}
		data = append(data, d)
	}
	res.Msg = "查询成功"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetParameterCalibrationRootView(c *gin.Context) {
	/*
		参数标定的package获取与筛选
	*/
	var res DataType.ResponseData

	userSpaceId := c.Query("space_id")
	if userSpaceId == "" {
		userSpaceId = c.GetHeader("space_id")
	}
	var modelData []map[string]any
	var packageModel []DataBaseModel.YssimModels
	//var space DataBaseModel.YssimUserSpace
	//dbModel.Where("id = ? AND username = ?", userSpaceId, userName).First(&space)
	subQuery := dbModel.Model(&DataBaseModel.SystemLibrary{}).Where("encryption = ?", true).Or("encryption = ? AND username = ?", false, userName).Select("id")
	dbModel.Where("sys_or_user = ? AND userspace_id = ? AND encryption = ? AND library_id = ''", userName, userSpaceId, false).
		Or("sys_or_user = ? AND userspace_id = ? AND library_id NOT IN (?)", userName, userSpaceId, subQuery).Find(&packageModel)
	libraryAndVersions := service.GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		loadVersions, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && loadVersions == packageModel[i].Version {
			t := service.GetModelType(packageModel[i].PackageName)
			data := map[string]any{
				"package_id":      packageModel[i].ID,
				"package_name":    packageModel[i].PackageName,
				"package_version": packageModel[i].Version,
				"model_name":      packageModel[i].PackageName,
				"haschild":        service.GetModelHasChild(packageModel[i].PackageName),
				"type":            t,
			}

			if data != nil {
				modelData = append(modelData, data)
			}
		}
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)
}

func GetParameterCalibrationListView(c *gin.Context) {
	/*
		# 参数标定的package子节点获取与筛选
	*/
	parent := c.Query("parent")
	packageId := c.Query("package_id")
	userSpaceId := c.Query("space_id")
	var packageModel DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", packageId, userName, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "数据错误")
		return
	}
	var res DataType.ResponseData
	modelChildList := service.GetModelChild(parent)
	var modelChildListNew []any
	for i := 0; i < len(modelChildList); i++ {
		if !modelChildList[i].HasChild && modelChildList[i].Type == "package" {
			continue
		}
		modelChildListNew = append(modelChildListNew, modelChildList[i])
	}
	res.Data = modelChildListNew
	c.JSON(http.StatusOK, res)
}

func GetParameterCalibrationRecordView(c *gin.Context) {
	/*
		# 获取某模型的参数标定记录数据，如果没有，则创建
	*/
	packageId := c.Query("package_id")
	userSpaceId := c.Query("space_id")
	modelName := c.Query("model_name")
	var record DataBaseModel.ParameterCalibrationRecord
	simulationOptions := service.GetSimulationOptions(modelName)

	dbModel.Where(DataBaseModel.ParameterCalibrationRecord{
		UserSpaceId: userSpaceId,
		PackageId:   packageId,
		UserName:    userName,
		ModelName:   modelName,
	}).Attrs(DataBaseModel.ParameterCalibrationRecord{
		ID:                uuid.New().String(),
		StartTime:         simulationOptions["startTime"],
		StopTime:          simulationOptions["stopTime"],
		Tolerance:         simulationOptions["tolerance"],
		NumberOfIntervals: simulationOptions["numberOfIntervals"],
		Interval:          simulationOptions["interval"],
		Method:            simulationOptions["method"],
	}).FirstOrCreate(&record)
	var res DataType.ResponseData
	res.Data = map[string]any{
		"id":                    record.ID,
		"start_time":            record.StartTime,
		"stop_time":             record.StopTime,
		"tolerance":             record.Tolerance,
		"number_of_intervals":   record.NumberOfIntervals,
		"interval":              record.Interval,
		"method":                record.Method,
		"compile_status":        record.CompileStatus,
		"actual_data":           record.ActualData,
		"rated_condition":       record.RatedCondition,
		"formula":               record.Formula,
		"associated_parameters": record.AssociatedParameters,
		"condition_parameters":  record.ConditionParameters,
		"formula_string":        record.FormulaString,
		"result_parameters":     record.ResultParameters,
	}
	c.JSON(http.StatusOK, res)
}

func SetActualDataView(c *gin.Context) {
	/*
		# 设置参数标定功能模型的实测参数字段与数据
	*/
	var item DataType.SetActualData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("实测数据错误：", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	actualDataList, _ := sonic.Marshal(&item.ActualDataList)
	dbModel.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ? AND package_id = ? AND username = ?", item.ID, item.PackageId, userName).UpdateColumn("actual_data", actualDataList)
	var res DataType.ResponseData
	c.JSON(http.StatusOK, res)
}

func SetRatedConditionView(c *gin.Context) {
	/*
		# 设置参数标定功能模型的额定工况参数
	*/
	var item DataType.SetRatedConditionData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//var ratedConditionList []any
	ratedConditionList, _ := sonic.Marshal(&item.RatedConditionList)
	dbModel.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ? AND package_id = ? AND username = ?", item.ID, item.PackageId, userName).UpdateColumn("rated_condition", ratedConditionList)
	var res DataType.ResponseData
	c.JSON(http.StatusOK, res)
}

func SetConditionParametersView(c *gin.Context) {
	/*
		# 设置参数标定功能模型的条件参数
	*/
	var item DataType.SetConditionParametersData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	conditionParametersList, _ := sonic.Marshal(&item.ConditionParametersList)
	dbModel.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ? AND package_id = ? AND username = ?", item.ID, item.PackageId, userName).UpdateColumn("condition_parameters", conditionParametersList)
	var res DataType.ResponseData
	c.JSON(http.StatusOK, res)
}

func SetResultParametersView(c *gin.Context) {
	/*
		# 设置参数标定功能模型的结果参数
	*/
	var item DataType.SetResultParametersData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	resultParametersList, _ := sonic.Marshal(&item.ResultParametersList)
	dbModel.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ? AND package_id = ? AND username = ?", item.ID, item.PackageId, userName).UpdateColumn("result_parameters", resultParametersList)
	var res DataType.ResponseData
	c.JSON(http.StatusOK, res)
}

func GetVariableParameterView(c *gin.Context) {
	/*
	  # 获取参数标定功能模型的额定工况参数与条件参数节点
	*/
	recordId := c.Query("id")
	packageId := c.Query("package_id")
	parentNode := c.Query("parent")
	var record DataBaseModel.ParameterCalibrationRecord
	dbModel.Where("id = ? AND package_id = ? AND username = ?", recordId, packageId, userName).First(&record)
	if record.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var result []map[string]any

	var res DataType.ResponseData
	if record.CompileStatus == "4" {
		//OMC仿真完输出的xml文件
		result = service.GetVariableParameter(record.CompilePath+"/result_init.xml", parentNode)
	} else {
		res.Err = "查询失败"
		res.Status = 2
	}

	sortByFirstLetter := func(i, j int) bool {
		// 从每个map中提取指定键的值进行比较
		value1 := fmt.Sprintf("%v", result[i]["variables"])
		value2 := fmt.Sprintf("%v", result[j]["variables"])
		return strings.ToLower(string(value1[0])) < strings.ToLower(string(value2[0]))
	}
	// 使用排序函数对切片进行排序
	sort.Slice(result, sortByFirstLetter)
	res.Data = result
	c.JSON(http.StatusOK, res)
}

func ParameterCalibrationFormulaParserView(c *gin.Context) {
	/*
		# 设置参数标定功能公式解析
	*/
	var item DataType.FormulaParserData
	var res DataType.ResponseData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	formulaData, variableList, coefficientNameList, err := service.GetFormulaList(item.FormulaStr)
	if err != nil {
		res.Err = err.Error()
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	formula, _ := sonic.Marshal(&formulaData)
	coefficientName, _ := sonic.Marshal(&coefficientNameList)
	variables, _ := sonic.Marshal(&variableList)
	dbModel.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ? AND package_id = ? AND username = ?", item.ID, item.PackageId, userName).Updates(
		map[string]any{"formula": formula, "coefficient_name": coefficientName, "variable_list": variables, "formula_string": item.FormulaStr})
	res.Data = map[string]any{"variable": variableList, "formula": formulaData}
	c.JSON(http.StatusOK, res)
}

func SetAssociatedParametersView(c *gin.Context) {
	/*
	  # 设置参数标定功能模型的拟合计算中的关联参数
	*/

	var item DataType.AssociatedParametersData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	parameters, _ := sonic.Marshal(&item.Parameters)
	dbModel.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ? AND package_id = ? AND username = ?", item.ID, item.PackageId, userName).UpdateColumn("associated_parameters", parameters)
	var res DataType.ResponseData
	c.JSON(http.StatusOK, res)
}

func FittingCalculationView(c *gin.Context) {
	/*
		# 进行模型参数标定的拟合计算
	*/
	var item DataType.FittingCalculationData
	var res DataType.ResponseData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var record DataBaseModel.ParameterCalibrationRecord
	dbModel.Where("id = ? AND username = ?", item.ID, userName).First(&record)
	if record.ID == "" {
		res.Err = "未找到记录"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	result, err := service.GrpcFittingCalculation(item.ID)
	if err != nil {
		res.Err = "系统出现错误，请联系管理员"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if result.Status != 0 {
		res.Err = result.Err
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	coefficient, _ := sonic.Marshal(&result.Coefficient)
	dbModel.Model(&record).Update("coefficient", coefficient)
	var formulaList []map[string]string
	var coefficientNameList []string
	var coefficientList []map[string]any
	_ = sonic.Unmarshal(record.Formula, &formulaList)
	_ = sonic.Unmarshal(record.CoefficientName, &coefficientNameList)
	for i := 1; i < len(coefficientNameList); i++ {
		coefficientList = append(coefficientList, map[string]any{"name": coefficientNameList[i], "value": result.Coefficient[i-1]})
	}
	res.Data = map[string]any{"coefficient": coefficientList, "score": result.Score}
	c.JSON(http.StatusOK, res)
}

func SetParameterCalibrationSimulationOptionsView(c *gin.Context) {
	/*
		# 设置某模型的参数标定的仿真求解设置信息
	*/
	var item DataType.SimulationOptionsData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	dbModel.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ? AND package_id = ? AND model_name = ? AND username = ?", item.ID, item.PackageId, item.ModelName, userName).
		Updates(&item.Options)
	var res DataType.ResponseData
	c.JSON(http.StatusOK, res)
}

func GetParameterCalibrationTemplateView(c *gin.Context) {
	/*
		# 获取参数标定模板
	*/
	var res DataType.ResponseData
	recordId := c.Query("id")
	var template []DataBaseModel.ParameterCalibrationTemplate
	if recordId == "" {
		dbModel.Find(&template)
	} else {
		dbModel.Where("id = ?", recordId).First(&template)
	}
	dataList := []map[string]any{}
	for _, calibrationTemplate := range template {
		d := map[string]any{
			"id":   calibrationTemplate.RecordID,
			"name": calibrationTemplate.TemplateName,
		}
		dataList = append(dataList, d)
	}
	res.Data = dataList
	c.JSON(http.StatusOK, res)
}

func CreateParameterCalibrationTemplateView(c *gin.Context) {
	/*
		# 创建参数标定模板
	*/
	var res DataType.ResponseData
	var item DataType.CreateCalibrationTemplateData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var record map[string]any
	var template DataBaseModel.ParameterCalibrationTemplate

	dbModel.Model(&DataBaseModel.ParameterCalibrationRecord{}).Where("id = ?", item.ID).First(&record)
	if record == nil {
		log.Println("未查到模型标定记录，请先创建记录再保存模板")
		c.JSON(http.StatusBadRequest, "")
		return
	}
	template.ID = uuid.New().String()
	template.TemplateName = item.TemplateName
	dbModel.Create(&template)
	dbModel.Model(&DataBaseModel.ParameterCalibrationTemplate{}).Where("record_id = ?", template.ID).Updates(&record)
	res.Msg = "模板创建成功"
	c.JSON(http.StatusOK, res)
}

func DeleteParameterCalibrationTemplateView(c *gin.Context) {
	/*
		# 删除参数标定模板
	*/
	var res DataType.ResponseData
	var item DataType.DeleteCalibrationTemplateData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	dbModel.Delete(DataBaseModel.ParameterCalibrationTemplate{}, item.ID)
	res.Msg = "模板删除成功"
	c.JSON(http.StatusOK, res)
}
