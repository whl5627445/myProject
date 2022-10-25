package API

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
)

var DB = config.DB

func UploadModelPackageView(c *gin.Context) {
	/*
		# 上传模型包文件，支持.mo与rar、zip两种压缩格式
	*/
	var res ResponseData
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
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
	nameList := strings.Split(modelFile.Filename, ".")
	if len(nameList) < 2 || (nameList[1] != "rar" && nameList[1] != "mo" && nameList[1] != "zip") {
		res.Msg = "请上传后缀为：mo与rar、zip三种格式的文件"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	removeSuffix := strings.Split(modelFile.Filename, ".")[0]
	saveFilePath := "public/UserFiles/UploadFile/" + username + "/" + removeSuffix + "/" + time.Now().Local().Format("20060102150405") + "/"
	zipPackagePath := saveFilePath + fileName
	packageName, ok := service.PackageFileParse(fileName, saveFilePath, zipPackagePath, file)

	if ok {
		var packageModel DataBaseModel.YssimModels
		DB.Where("sys_or_user = ? AND userspace_id = ? AND package_name = ?", username, userSpaceId, packageName).First(&packageModel)
		if packageModel.PackageName != "" {
			res.Err = packageName + "， 已存在相同名字的包，请重新检查后上传"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
		packagePathNew := "public/UserFiles/UploadFile/" + username + "/" + packageName + "/" + time.Now().Local().Format("20060102150405") + "/" + packageName + ".mo"
		result := service.SaveModelCode(packageName, packagePathNew)
		if result {
			packageRecord := DataBaseModel.YssimModels{
				ID:          uuid.New().String(),
				PackageName: packageName,
				SysUser:     username,
				FilePath:    packagePathNew,
				UserSpaceId: userSpaceId,
			}
			err := DB.Create(&packageRecord).Error
			if err != nil {
				res.Err = "上传失败，请重试"
				res.Status = 2
				service.DeleteLibrary(packageName)
			} else {
				res.Msg = packageName + " 包已上传成功"
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
	service.DeleteLibrary(packageName)
	res.Err = "模型包解析失败,存在语法错误，请检查后重新上传"
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
	var res ResponseData
	var item UpdateModelPackageData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var packageRecord DataBaseModel.YssimModels
	parentName := ""
	modelStr := item.ModelStr
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&packageRecord)
	if packageRecord.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	nameList := strings.Split(item.UpdateName, ".")
	if len(nameList) > 1 {
		parentName = strings.Join(nameList[:len(nameList)-1], ".")
		modelStr = "within " + parentName + ";" + item.ModelStr
	}
	oldCode := service.GetModelCode(packageRecord.PackageName)
	parseResult, ok := service.ParseCodeString(modelStr, packageRecord.PackageName)
	if ok && len(parseResult) > 0 {
		loadResult := service.LoadCodeString(modelStr, packageRecord.PackageName)
		if loadResult {
			saveResult := service.SaveModelToFile(packageRecord.PackageName, packageRecord.FilePath)
			if saveResult {
				res.Data = map[string]string{
					"id":        packageRecord.ID,
					"model_str": modelStr,
					"name":      parseResult,
				}
				res.Msg = "模型保存成功"
				c.JSON(http.StatusOK, res)
				return
			}
		} else {
			service.LoadCodeString(oldCode, packageRecord.PackageName)
		}
	}
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
	var res ResponseData
	var item CreateModelPackageData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	re1, _ := regexp.Compile("^[a-zA-Z_]")
	f := re1.Find([]byte(item.Name))
	createPackageName := item.Name
	createPackageNameALL := item.Name
	if f == nil {
		res.Err = "名称请以子母和下划线开头"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var packageRecord DataBaseModel.YssimModels
	var newPackage = DataBaseModel.YssimModels{
		ID:          uuid.New().String(),
		PackageName: createPackageName,
		SysUser:     username,
		FilePath:    "public/UserFiles/UploadFile/" + username + "/" + createPackageName + "/" + time.Now().Local().Format("20060102150405") + "/" + createPackageName + ".mo",
		UserSpaceId: userSpaceId,
	}
	DB.Where("package_name = ? AND sys_or_user = ? AND userspace_id = ?", item.Name, username, userSpaceId).First(&packageRecord)
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
			if modelChildList[i]["name"] == item.Name {
				res.Err = "名称已存在，请修改后再试。"
				res.Status = 2
				c.JSON(http.StatusOK, res)
				return
			}
		}
		newPackage = insertPackageRecord
	}
	result := service.CreateModelAndPackage(createPackageName, item.Vars.InsertTo, item.Vars.Expand, item.StrType, createPackageNameALL, item.Comment, item.Vars.Partial, item.Vars.Encapsulated, item.Vars.State)
	if result {
		saveResult := service.SaveModelCode(newPackage.PackageName, newPackage.FilePath)
		if saveResult {
			if item.Vars.InsertTo == "" {
				DB.Create(&newPackage)
			}
			res.Msg = "创建成功"
			res.Data = map[string]string{
				"model_str": service.GetModelCode(createPackageName),
				"id":        newPackage.ID,
			}
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
	var res ResponseData
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
	result := service.UploadIcon(modelName, fileBase64Str)
	if result {
		saveResult := service.SaveModelToFile(packageRecord.PackageName, packageRecord.FilePath)
		if saveResult {
			res.Msg = "图标上传成功"
			c.JSON(http.StatusOK, res)
			return
		}
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
	var res ResponseData
	username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	var packageRecord []map[string]interface{}
	DB.Raw("select m.id, m.package_name, m.create_time, m.update_time, s.space_name from yssim_models as m, yssim_user_spaces as s where m.sys_or_user = ? AND m.userspace_id = s.id AND m.deleted_at IS NULL AND s.deleted_at IS NULL;", username).Find(&packageRecord)
	var dataList []map[string]interface{}
	for id, models := range packageRecord {
		data := map[string]interface{}{
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
	   ## return: 包id， 包名， 上传时间， 修改时间
	*/
	//var res ResponseData
	username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	var item PackageFileData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var packageRecord DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ?", item.PackageId, username).First(&packageRecord)
	//c.JSON(http.StatusOK, res)
	c.Header("content-disposition", `attachment; filename=`+packageRecord.PackageName+".mo")
	c.File(packageRecord.FilePath)
}

func GetResultFileView(c *gin.Context) {
	/*
	   # 用户仿真结果文件下载
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item ResultFileData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//recordId := c.Query("record_id")
	var resultRecord DataBaseModel.YssimSimulateRecord
	DB.Where("id = ? AND username = ? AND userspace_id = ? ", item.RecordId, username, userSpaceId).First(&resultRecord)
	//c.Header("content-disposition", `attachment; filename=`+time.Now().Local().Format("20060102150405")+".mat")
	c.Header("content-disposition", `attachment; filename=`+resultRecord.SimulateModelName+".mat")
	//c.Data(http.StatusOK, "application/octet-stream", resDy)
	c.File(resultRecord.SimulateModelResultPath + "result_res.mat")
}

func GetFilterResultFileView(c *gin.Context) {
	/*
	   # 用户筛选仿真结果文件下载
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item FilterResultFileData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var resultRecord DataBaseModel.YssimSimulateRecord
	DB.Where("id = ? AND username = ? AND userspace_id = ? ", item.RecordId, username, userSpaceId).First(&resultRecord)
	newFileName := "public/tmp/" + username + "/" + strings.ReplaceAll(resultRecord.SimulateModelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + ".csv"
	ok := service.FilterSimulationResult(item.VarList, resultRecord.SimulateModelResultPath+"result_res.mat", newFileName)
	if ok {
		c.Header("content-disposition", `attachment; filename=`+time.Now().Local().Format("20060102150405")+".csv")
		c.File(newFileName)
		return
	}
	var res ResponseData
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

	var item FmuExportData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userSpaceId := c.GetHeader("space_id")
	username := c.GetHeader("username")
	token := c.GetHeader("Authorization")
	fileName := ""
	filePath := ""

	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&packageModel)
	if packageModel.FilePath != "" {
		filePath = packageModel.FilePath
		fileName = packageModel.PackageName
	}
	resDy, ok := service.DymolaFmuExport(item.FmuPar, token, username, item.FmuName, item.PackageName, item.ModelName, fileName, filePath)
	if ok {
		c.Header("content-disposition", `attachment; filename=`+item.FmuName+".fmu")
		c.Data(http.StatusOK, "application/octet-stream", resDy)
		return
	}
	var res ResponseData
	res.Err = "下载失败，请稍后再试"
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func ModelCodeSaveView(c *gin.Context) {
	/*
	   # 保存模型所在包的代码到.mo文件
	   ## package_id: 包的id
	   ## package_name： 包的名称
	*/
	var res ResponseData
	var item ModelCodeSaveData
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
	result := service.SaveModelToFile(packageModel.PackageName, packageModel.FilePath)
	if result {
		res.Msg = "模型已保存"
	} else {
		res.Err = "保存模型失败"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}
