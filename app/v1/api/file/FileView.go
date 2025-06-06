package file

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/v1/service"
	"yssim-go/config"

	"github.com/bytedance/sonic"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB = config.DB

func UploadModelPackageView(c *gin.Context) {
	/*
		# 上传模型包文件，支持.mo与rar、zip两种压缩格式
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")

	// 检查用户名和用户空间id是否存在
	var userspaceRecord DataBaseModel.YssimUserSpace
	DB.Where("username = ? AND id = ?", userName, userSpaceId).First(&userspaceRecord)
	if userspaceRecord.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	modelFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	file, err := modelFile.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	fileName := modelFile.Filename

	if !strings.HasSuffix(fileName, ".rar") && !strings.HasSuffix(fileName, ".mo") && !strings.HasSuffix(fileName, ".zip") && !strings.HasSuffix(fileName, ".sk") {
		res.Msg = "请上传后缀为：mo与rar、zip、sk四种格式的文件"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	saveFilePathBase := "static/UserFiles/UploadFile/" + userName + "/" + time.Now().Local().Format("20060102150405")
	packageName, packagePath, msg, ok := service.PackageFileParse(fileName, saveFilePathBase, file)
	if ok {
		var packageModel DataBaseModel.YssimModels
		DB.Where("sys_or_user IN ? AND userspace_id IN ? AND package_name = ?", []string{"sys", userName}, []string{"0", userSpaceId}, packageName).First(&packageModel)
		if packageModel.PackageName != "" {
			service.DeletePackageFile(saveFilePathBase)
			res.Err = packageName + "， 已存在相同名字的包，请重新检查后上传"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
		var encryptionPackage DataBaseModel.SystemLibrary

		if strings.HasSuffix(fileName, ".sk") {
			DB.Where("username = ? AND package_name = ? AND encryption = ?", userName, packageName, true).First(&encryptionPackage)
			if encryptionPackage.ID != "" {
				encryptionPackage.Version = service.GetVersion(packageName)
				encryptionPackage.FilePath = packagePath
				DB.Save(&encryptionPackage)
				res.Msg = "加密模型库 \"" + packageName + "\" 已更新"
			} else {
				encryptionPackage.ID = uuid.New().String()
				encryptionPackage.UserName = userName
				encryptionPackage.PackageName = packageName
				encryptionPackage.Version = service.GetVersion(packageName)
				encryptionPackage.FilePath = packagePath
				encryptionPackage.Encryption = true
				DB.Create(&encryptionPackage)
			}
		}
		packageRecord := DataBaseModel.YssimModels{
			ID:          uuid.New().String(),
			PackageName: packageName,
			SysUser:     userName,
			FilePath:    packagePath,
			UserSpaceId: userSpaceId,
			Version:     service.GetVersion(packageName),
			Encryption:  encryptionPackage.Encryption,
			LibraryId:   encryptionPackage.ID,
		}
		err = DB.Create(&packageRecord).Error
		if err != nil {
			res.Err = "上传失败，请重试"
			res.Status = 2
			service.DeleteLibrary(packageName)
		} else {
			conflict, err := service.GetLoadPackageConflict(packageName, packageRecord.Version, packagePath+"/package.mo")
			if len(conflict) > 0 && err != nil {
				service.DeleteLibrary(packageName)
				data := map[string]any{}
				data["package_id"] = packageRecord.ID
				data["conflict"] = conflict
				res.Data = data
				res.Msg = err.Error()
				c.JSON(http.StatusOK, res)
				return
			}
			packageInformation := service.GetPackageInformation()
			packageInformationJson, _ := sonic.Marshal(packageInformation)
			DB.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, userName).Update("package_information", packageInformationJson)
			res.Msg = "上传成功"
		}
		c.JSON(http.StatusOK, res)
		return
	}
	service.DeleteLibrary(packageName)
	res.Err = msg
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func UpdateModelPackageView(c *gin.Context) {
	/*
		# 更新模型源码
		## model_str: 源码数据，字符串形式
		## package_name: 要更新内容的模型或包名，必须是全名
		## package_id: 包的id
	*/
	var res DataType.ResponseData
	var item DataType.UpdateModelPackageData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageRecord DataBaseModel.YssimModels
	// parentName := ""
	modelStr := item.ModelStr
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&packageRecord)
	if packageRecord.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.ModelName, ".")
	if len(nameList) > 1 {
		parentName := strings.Join(nameList[:len(nameList)-1], ".")
		modelStr = "within " + parentName + ";" + item.ModelStr
	}
	// oldCode := service.GetModelCode(item.ModelName)
	modelPath := packageRecord.FilePath
	if strings.HasSuffix(packageRecord.FilePath, "/package.mo") {
		modelPath = service.GetSourceFile(item.ModelName)
	}

	parseResult, ok := service.ParseCodeString(modelStr, modelPath)
	if ok && len(parseResult) > 0 {
		isExist := service.IsExistPackage(parseResult)
		if isExist && (item.ModelName != parseResult) {
			res.Err = "模型名称重复"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
		loadResult := service.LoadCodeString(modelStr, modelPath)
		if loadResult {
			if parseResult != item.ModelName {
				//如果是管网模型，则更新管网数据表
				DB.Model(DataBaseModel.YssimPipeNetCadDownload{}).Where("package_id = ? AND model_name = ?", item.PackageId, item.ModelName).Update("model_name", parseResult)
				//更改实验记录的模型名
				DB.Model(DataBaseModel.YssimExperimentRecord{}).Where("package_id = ? AND userspace_id = ? AND username = ? AND model_name = ?", item.PackageId, userSpaceId, username, item.ModelName).Update("model_name", parseResult)
				//更改仿真历史记录的模型名
				DB.Model(DataBaseModel.YssimSimulateRecord{}).Where("username = ? AND simulate_model_name = ? AND userspace_id = ?  AND package_id = ?", username, item.ModelName, userSpaceId, item.PackageId).Update("simulate_model_name", parseResult)

				// 判断是否是子模型
				if !strings.Contains(item.ModelName, ".") {
					DB.Model(DataBaseModel.YssimModels{}).Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).Update("package_name", parseResult)

					packageInformation := service.GetPackageInformation()
					packageInformationJson, _ := sonic.Marshal(packageInformation)
					DB.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, username).Update("package_information", packageInformationJson)
				}
				service.DeleteLibrary(item.ModelName)
			}
			service.ModelSave(parseResult)
			res.Data = map[string]string{
				"id":        packageRecord.ID,
				"model_str": modelStr,
				"name":      parseResult,
			}
			res.Msg = "模型保存成功"
			c.JSON(http.StatusOK, res)
			return
		}
	}
	// else {
	// 	service.LoadCodeString(oldCode, packageRecord.PackageName)
	// }
	res.Err = "语法错误，请重新检查"
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func CreateModelPackageView(c *gin.Context) {
	/*
		# 用户创建模型和模型包mo文件接口
		## package_name: 包或模型的名字
		## type: 要创建的类型
		## model_str: str = 模型源码字符串
		## package_id: 包的id
		## vars: {
		##     "expand": "", 扩展
		##     "comment": "", 注释
		##     "insert_to": "", 父节点， 要插入哪个节点下
		##     "partial": False,  部分的
		##     "encapsulated": False, 封装
		##     "state": False 状态
		##     }
	*/
	var res DataType.ResponseData
	var item DataType.CreateModelPackageData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	matchSpaceName1, _ := regexp.MatchString("^[_a-zA-Z0-9]+$", item.Name) // 字母、数字、下划线验证
	matchSpaceName2, _ := regexp.MatchString("^[a-zA-Z_]", item.Name)      // 字母、下划线验证
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
	createPackageName := item.Name
	createPackageNameALL := item.Name

	var packageRecord DataBaseModel.YssimModels
	var newPackage = DataBaseModel.YssimModels{
		ID:          uuid.New().String(),
		PackageName: createPackageName,
		SysUser:     username,
		FilePath:    "static/UserFiles/UploadFile/" + username + "/" + time.Now().Local().Format("20060102150405") + "/" + createPackageName + "/" + createPackageName + ".mo",
		UserSpaceId: userSpaceId,
	}
	DB.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", item.Name, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageRecord)
	if packageRecord.PackageName != "" && item.Vars.InsertTo == "" {
		res.Err = "名称已存在，请修改后再试。"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if item.Vars.InsertTo != "" {
		var insertPackageRecord DataBaseModel.YssimModels
		DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&insertPackageRecord)
		createPackageNameALL = item.Vars.InsertTo + "." + item.Name
		modelChildList := service.GetModelChild(item.Vars.InsertTo)
		for i := 0; i < len(modelChildList); i++ {
			if modelChildList[i].Name == item.Name {
				res.Err = "名称已存在，请修改后再试。"
				res.Status = 2
				c.JSON(http.StatusOK, res)
				return
			}
		}
		newPackage = insertPackageRecord
	} else {
		if err = DB.Create(&newPackage).Error; err != nil {
			res.Err = "创建失败，请稍后再试"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	}
	result := service.CreateModelAndPackage(createPackageName, item.Vars.InsertTo, item.Vars.Expand, item.StrType, createPackageNameALL, item.Comment, item.Vars.Partial, item.Vars.Encapsulated, item.Vars.State)
	if result {
		saveResult := service.SaveModelCode(createPackageNameALL, newPackage.FilePath)
		if saveResult {
			res.Msg = "创建成功"
			if item.Vars.InsertTo == "" {
				res.Data = map[string]string{
					"model_name": newPackage.PackageName,
					// "model_str": service.GetModelCode(createPackageName),
					"id": newPackage.ID,
				}
			} else {
				res.Data = map[string]string{
					"model_name": item.Vars.InsertTo + "." + item.Name,
					// "model_str": service.GetModelCode(createPackageName),
					"id": newPackage.ID,
				}
			}
			packageInformation := service.GetPackageInformation()
			packageInformationJson, _ := sonic.Marshal(packageInformation)
			DB.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, username).Update("package_information", packageInformationJson)
		} else {
			DB.Delete(&newPackage)
			res.Err = "创建失败，请稍后再试"
			res.Status = 2
		}
	} else {
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func UploadModelIconView(c *gin.Context) {
	/*
		# 用户上传模型图标接口
		## file: 文件数据，bytes形式的文件流
		## model_name: 模型名称
		## package_id: 包id
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	modelName := c.PostForm("model_name")
	packageId := c.PostForm("package_id")
	var packageRecord DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ? ", packageId, username, userSpaceId).First(&packageRecord)
	if packageRecord.ID == "" {
		res.Err = "暂只支持更新用户区图标"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	iconFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	iconFileName := iconFile.Filename
	iconFileNameList := strings.Split(iconFileName, ".")
	iconFileNameSuffix := iconFileNameList[len(iconFileNameList)-1]
	iconType := map[string]bool{"jpg": true, "png": true, "jpeg": true, "svg": true}
	if !iconType[iconFileNameSuffix] {
		res.Err = "暂只支持jpg,png,jpeg,svg格式的图片"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	file, _ := iconFile.Open()
	iconData, _ := io.ReadAll(file)
	fileBase64Str := base64.StdEncoding.EncodeToString(iconData)
	result := service.SetIcon(modelName, fileBase64Str)
	if result {
		service.ModelSave(modelName)
		res.Msg = "图标上传成功"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = "上传失败，请重新上传"
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func GetPackageFileListView(c *gin.Context) {
	/*
	   # 用户获取mo文件信息接口， 可以进行下载
	   ## return: 包id， 包名， 上传时间， 修改时间
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageRecord []map[string]any
	DB.Raw("select m.id, m.package_name, m.create_time, m.update_time, s.space_name from yssim_models as m, yssim_user_spaces as s where m.sys_or_user = ? AND m.userspace_id = s.id AND m.deleted_at IS NULL AND s.deleted_at IS NULL AND s.id = ? ORDER BY create_time desc;", username, userSpaceId).Find(&packageRecord)
	var dataList []map[string]any
	for id, models := range packageRecord {
		data := map[string]any{
			"id":           id,
			"package_id":   models["id"],
			"space_name":   models["space_name"],
			"package_name": models["package_name"],
			"create_time":  models["create_time"].(time.Time).Format("2006-01-02 15:04:05"),
			"update_time":  models["update_time"].(time.Time).Format("2006-01-02 15:04:05"),
		}
		dataList = append(dataList, data)
	}
	res.Data = dataList
	c.JSON(http.StatusOK, res)

}

func GetPackageFileView(c *gin.Context) {
	/*
	   # 用户mo文件下载
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	var item DataType.PackageFileData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var packageRecord DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ?", item.PackageId, username).First(&packageRecord)

	filePath, err := service.ZipPackageStream(packageRecord.PackageName, packageRecord.FilePath)
	if packageRecord.Encryption {
		filePath, err = service.ZipPackageEncrypt(packageRecord.PackageName, filePath)
	}
	if err != nil {
		res.Err = "导出模型库失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusInternalServerError, res)
	}
	res.Data = map[string]string{"url": filePath}
	c.JSON(http.StatusOK, res)
}

func GetPackageEncryptView(c *gin.Context) {
	/*
	   # 用户mo文件加密后下载
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	packageId := c.Query("package_id")
	spaceId := c.Query("space_id")
	var packageRecord DataBaseModel.YssimModels
	DB.Where("id = ? AND userspace_id IN ? AND sys_or_user IN ?", packageId, []string{spaceId, "0"}, []string{username, "sys"}).First(&packageRecord)
	if packageRecord.ID == "" {
		res.Err = "导出模型库失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusBadRequest, res)
	}
	packagePath := service.GetSourceFile(packageRecord.PackageName)
	filePath, err := service.ZipPackageStream(packageRecord.PackageName, packagePath)
	filePath, err = service.ZipPackageEncrypt(packageRecord.PackageName, filePath)
	if err != nil {
		res.Err = "导出模型库失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusInternalServerError, res)
	}
	res.Data = map[string]string{"url": filePath}
	c.JSON(http.StatusOK, res)
}

func GetResultFileView(c *gin.Context) {
	/*
	   # 用户仿真结果文件下载
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.ResultFileData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// recordId := c.Query("record_id")
	var resultRecord DataBaseModel.YssimSimulateRecord
	DB.Where("id = ? AND username = ? AND userspace_id = ? ", item.RecordId, username, userSpaceId).First(&resultRecord)
	// ok := service.GrpcMatToCsv(resultRecord.SimulateModelResultPath)
	if resultRecord.SimulateModelResultPath != "" && resultRecord.SimulateStatus == "4" {
		res.Data = map[string]string{"url": resultRecord.SimulateModelResultPath + "result_res.mat"}
	} else {
		res.Err = "下载失败"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func GetFilterResultFileView(c *gin.Context) {
	/*
	   # 用户筛选仿真结果文件下载
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	// userSpaceId := c.GetHeader("space_id")
	var items []DataType.FilterResultFileData
	err := c.BindJSON(&items)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 遍历获取所有recordId,用于查询数据库
	itemsMap := map[string][]string{}
	var recordIdList []string
	for i := 0; i < len(items); i++ {
		recordIdList = append(recordIdList, items[i].RecordId)
		itemsMap[items[i].RecordId] = items[i].VarList
	}
	// 判断记录是否存在，有一条不存在就返回"not found"
	var recordList []DataBaseModel.YssimSimulateRecord
	err = DB.Where("id IN ? AND username = ?", recordIdList, username).Find(&recordList).Error
	for i := 0; i < len(recordList); i++ {
		if err != nil || recordList[i].SimulateStatus != "4" {
			c.JSON(http.StatusBadRequest, "not found")
			return
		}
	}
	// 判断输入的id数和数据库查询的id数是否一致
	if len(recordList) != len(recordIdList) {
		c.JSON(http.StatusBadRequest, "判断输入的id数和数据库查询的id数是否一致")
		return
	}
	// 构建key为id，val为SimulateModelResultPath的健值对,降低时间复杂度
	recordDict := map[string]DataBaseModel.YssimSimulateRecord{}
	for _, record := range recordList {
		recordDict[record.ID] = record
	}
	var ok bool
	// newFileName保存为csv的文件名
	newFileName := "static/tmp/" + username + "/" + strings.ReplaceAll(recordList[0].SimulateModelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + ".csv"

	ok = service.FilterSimulationResult(itemsMap, recordDict, newFileName)
	if ok {
		// c.Header("content-disposition", `attachment; filename=`+recordList[0].SimulateModelName+".csv")
		// c.File(newFileName)
		res.Data = map[string]string{"url": newFileName}
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = "下载失败，请稍后再试"
	res.Status = 2
	c.JSON(http.StatusOK, res)
}

func FmuExportModelView(c *gin.Context) {
	/*
	   # 导出模型的fmu文件
	   ## package_id: 包的id
	   ## package_name： 包的名称
	   ## model_name： 模型全名
	   ## fmu_name： fmu文件的名字
	   ## fmu_par： fmu导出的参数
	   ## download_local： 是否下载到本地
	*/
	var res DataType.ResponseData
	var item DataType.FmuExportData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userSpaceId := c.GetHeader("space_id")
	username := c.GetHeader("username")

	filePath := ""

	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&packageModel)
	if packageModel.FilePath != "" {
		filePath = packageModel.FilePath
	}
	// 获取依赖
	envLibrary := service.GetEnvLibraryAll(username, userSpaceId)
	newFileName := ""
	ok := false
	errTips := ""
	if item.Solver == "default" {
		newFileName, ok, errTips = service.OmcFmuExportWithLibrary(item.FmuPar, envLibrary, username, item.FmuName, item.PackageName, item.ModelName, filePath)

	} else {
		newFileName, ok, errTips = service.DymolaFmuExportWithLibrary(item.FmuPar, envLibrary, username, item.FmuName, item.ModelName)
	}
	if ok {

		res.Data = map[string]string{"url": newFileName}
		c.JSON(http.StatusOK, res)
		return
	}
	// 编译问题提示 "下载失败，请查看日志"
	// 系统问题提示 "下载失败，请稍后再试"
	res.Err = errTips
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func ModelCodeSaveView(c *gin.Context) {
	/*
	   # 保存模型所在包的代码到.mo文件
	   ## package_id: 包的id
	   ## package_name： 包的名称
	*/
	var res DataType.ResponseData
	var item DataType.ModelCodeSaveData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&packageModel)
	if packageModel.FilePath == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.SaveModelSource(packageModel.PackageName, packageModel.FilePath)
	if result {
		res.Msg = "模型已保存"
	} else {
		res.Err = "保存模型失败"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func UploadModelVarFileView(c *gin.Context) {
	/*
		# 用户上传参数文件的接口
		## file: 文件数据，bytes形式的文件流
		## model_name: 模型名称
		## package_id: 包id
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	packageId := c.PostForm("package_id")
	userSpaceId := c.GetHeader("space_id")
	modelName := c.PostForm("model_name")
	componentName := c.PostForm("component_name")
	varFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	fileName := varFile.Filename

	if !strings.HasSuffix(fileName, ".txt") && !strings.HasSuffix(fileName, ".mos") && !strings.HasSuffix(fileName, ".csv") {
		res.Msg = "请上传后缀为：txt、mos、csv三种格式的文件"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", packageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel)
	if packageModel.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.UploadResourcesFile(packageModel.PackageName, modelName+"/"+componentName, varFile)
	if result {
		res.Msg = "文件上传成功"
		res.Data = []string{"Modelica.Utilities.Files.loadResource(\"modelica://" + packageModel.PackageName + "/Resources/" + modelName + "/" + componentName + "/" + varFile.Filename + "\")"}
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = "上传失败，请重新上传"
	res.Status = 2
	c.JSON(http.StatusOK, res)
}

func GetPackageResourcesList(c *gin.Context) {
	/*
		# 模型库的静态资源文件夹下资源查找
		## package_id: 包id
		## parent: 需要查询的节点父级路径
		## path: 被查询节点
	*/
	var item DataType.PackageResourcesData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel)
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
			basePath = "modelica://" + packageModel.PackageName + "/Resources/" + strings.Join(pathList, "/")
		}
		d["path"] = basePath
	}
	var res DataType.ResponseData
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func UploadResourcesFileView(c *gin.Context) {
	/*
		# 用户上传静态资源文件接口
		## file: 文件数据，bytes形式的文件流
		## model_name: 模型名称
		## parent: 保存文件的父节点
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	packageId := c.PostForm("package_id")
	parent := c.PostForm("parent")
	userSpaceId := c.GetHeader("space_id")
	varFile, err := c.FormFile("file")
	if varFile.Size > 1500000 {
		res.Err = "上传文件过大，请上传小于1.5M的文件"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if !strings.HasSuffix(varFile.Filename, ".txt") && !strings.HasSuffix(varFile.Filename, ".csv") && !strings.HasSuffix(varFile.Filename, ".png") &&
		!strings.HasSuffix(varFile.Filename, ".mos") {
		res.Err = "暂时只支持*.txt、*.csv、*.png、*.mos格式文件上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", packageId, userName, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.UploadResourcesFile(packageModel.PackageName, parent, varFile)
	if result {
		res.Msg = "文件上传成功"
		pathList := []string{}
		if parent != "" {
			pathList = append(pathList, parent)
		}
		pathList = append(pathList, varFile.Filename)
		data := map[string]string{
			"path": "modelica://" + packageModel.PackageName + "/Resources/" + strings.Join(pathList, "/"),
			"name": varFile.Filename,
		}
		res.Data = data
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = "上传失败，请重新上传"
	res.Status = 2
	c.JSON(http.StatusOK, res)
}

func CreateResourcesDirView(c *gin.Context) {
	/*
		# 静态资源文件夹创建子级文件夹接口
		## package_id: 包id
		## parent: 创建文件夹的父节点
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.PackageResourcesData
	var packageModel DataBaseModel.YssimModels
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if item.Path == "" {
		res.Status = 2
		res.Err = "文件夹名称不能为空"
		c.JSON(http.StatusOK, res)
		return
	}
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
	list := service.GetResourcesList(packageModel.PackageName, item.Parent)
	for _, name := range list {
		if name["name"] == item.Path {
			res.Status = 2
			res.Err = "文件夹名称重复"
			c.JSON(http.StatusOK, res)
			return
		}
	}
	if packageModel.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.CreateResourcesDir(packageModel.PackageName, item.Path, item.Parent)
	if result {
		res.Msg = "文件夹创建成功"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = err.Error()
	res.Status = 2
	c.JSON(http.StatusOK, res)
}

func DeleteResourcesDirAndFileView(c *gin.Context) {
	/*
		# 静态资源文件夹删除子级文件夹与文件接口
		## package_id: 包id
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.PackageResourcesData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	result := service.DeleteResourcesDirAndFile(packageModel.PackageName, item.Parent)
	if result {
		res.Msg = "删除成功"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = err.Error()
	res.Status = 2
	c.JSON(http.StatusOK, res)
}

func DownloadResourcesFileView(c *gin.Context) {
	/*
		# 静态资源文件的下载
		## package_id: 包id
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.PackageResourcesData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	filePath := ""
	if item.Parent != "" {
		filePath = item.Parent + "/" + item.Path
	} else {
		filePath = item.Path
	}
	filePathAll := service.GetResourcesDir(packageModel.PackageName, filePath)
	// c.Header("content-disposition", `attachment;filename=`+url.QueryEscape(item.Path))
	// c.Data(http.StatusOK, "application/octet-stream", fileData)
	res.Data = map[string]any{
		"url": filePathAll,
	}
	c.JSON(http.StatusOK, res)

}

func ResourcesImagesPathGetView(c *gin.Context) {
	/*
		# 静态资源文件png图片的路径获取
		## package_id: 包id
		## keyWord: 筛选关键字
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.ResourcesImagesPathData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	data := service.GetResourcesImagesPath(packageModel.PackageName, item.KeyWord)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func ResourcesImagesGetView(c *gin.Context) {
	/*
		# 静态资源文件png图片的获取
		## path: 文件相对路径
	*/
	// var item getResourcesImagesData
	path := c.Query("path")
	// err := c.BindJSON(&item)
	// if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	// }
	fileData := service.GetResourcesImages(path)
	c.Header("content-disposition", `attachment;filename=`+url.QueryEscape("image.png"))
	c.Data(http.StatusOK, "application/octet-stream", fileData)
}

func ModelIconSetView(c *gin.Context) {
	/*
		# 用户设置模型图标接口
		## path: 文件相对路径
		## model_name: 模型名称
		## package_id: 包id
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.SetResourcesImagesIconData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var packageRecord DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ? ", item.PackageId, username, userSpaceId).First(&packageRecord)
	if packageRecord.ID == "" {
		res.Err = "暂只支持更新用户区图标"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// iconFileNameList := strings.Split(item.Path, ".")
	// iconFileNameSuffix := iconFileNameList[len(iconFileNameList)-1]
	// iconType := map[string]bool{"png": true}
	// file := service.GetResourcesImages(item.Path)
	// if !iconType[iconFileNameSuffix] || len(file) == 0 {
	//	res.Err = ""
	//	res.Status = 2
	//	c.JSON(http.StatusBadRequest, res)
	//	return
	// }
	file := service.GetResourcesImages(item.Path)
	if !strings.HasSuffix(item.Path, ".png") || len(file) == 0 {
		res.Err = ""
		res.Status = 2
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result := service.SetIconPath(item.ModelName, file)
	if result {
		service.ModelSave(packageRecord.PackageName)
		res.Msg = "图标设置成功"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = "设置失败，请重新上传"
	res.Status = 2
	c.JSON(http.StatusOK, res)

}
