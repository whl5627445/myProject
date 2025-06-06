package serviceV2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"yssim-go/grpc/taskManagement"
	"yssim-go/library/fileOperation"
	"yssim-go/library/mapProcessing"

	"go.mongodb.org/mongo-driver/bson"

	"yssim-go/app/DataBaseModel"
	"yssim-go/config"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

type OutputData struct {
	Name string `json:"data"`
	Unit string `json:"unit"`
}

var DB = config.DB
var MB = config.MB

// GetEnvLibraryAll 获取当前环境下的所有已经加载的包和系统库
func GetEnvLibraryAll(userName, spaceId string) map[string]string {

	// 获取系统模型
	environmentModelData := make(map[string]string)
	var envPackageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user =  ? AND userspace_id = ?", "sys", "0").Find(&envPackageModel)
	libraryAndVersions := GetLibraryAndVersions()
	for i := 0; i < len(envPackageModel); i++ {
		p, ok := libraryAndVersions[envPackageModel[i].PackageName]
		if ok && p == envPackageModel[i].Version {
			environmentModelData[envPackageModel[i].PackageName] = envPackageModel[i].Version
		}
	}
	// 获取用户模型
	DB.Where("sys_or_user = ? AND userspace_id = ?", userName, spaceId).Find(&envPackageModel)
	for i := 0; i < len(envPackageModel); i++ {
		loadVersions, ok := libraryAndVersions[envPackageModel[i].PackageName]
		if ok && loadVersions == envPackageModel[i].Version {
			environmentModelData[envPackageModel[i].PackageName] = envPackageModel[i].FilePath
		}
	}

	// 获取加密模型
	var encryptionPackageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user =  ? AND userspace_id = ? AND encryption = ?", userName, spaceId, 1).Find(&encryptionPackageModel)
	for i := 0; i < len(encryptionPackageModel); i++ {
		packageVersion, ok := libraryAndVersions[encryptionPackageModel[i].PackageName]
		if ok && packageVersion == encryptionPackageModel[i].Version {
			environmentModelData[encryptionPackageModel[i].PackageName] = encryptionPackageModel[i].FilePath
		}
	}
	return environmentModelData

}

// GetEnvLibrary 获取已经加载的依赖包和系统库
func GetEnvLibrary(packageName, userName, spaceId string) map[string]string {
	environmentModelData := make(map[string]string)
	var p DataBaseModel.YssimModels
	libraryAndVersions := GetLibraryAndVersions()
	// 获取需要仿真的模型名
	DB.Where("package_name = ? AND sys_or_user = ? AND userspace_id = ?", packageName, userName, spaceId).First(&p)
	if p.ID != "" {
		environmentModelData[p.PackageName] = p.FilePath
	}

	// 获取需要加载的用户模型
	dependentLibrary := GetPackageUses(packageName)
	for i := 0; i < len(dependentLibrary); i++ {
		var usedModel DataBaseModel.YssimModels
		DB.Where("package_name = ? AND version = ? AND sys_or_user = ? AND userspace_id = ?", dependentLibrary[i][0], dependentLibrary[i][1], userName, spaceId).First(&usedModel)
		l, ok := libraryAndVersions[usedModel.PackageName]
		// 数据库你存在且FilePath不为空，并且yssim已经加载。
		if usedModel.ID != "" && ok && l == usedModel.Version && usedModel.FilePath != "" {
			environmentModelData[usedModel.PackageName] = usedModel.FilePath
		}
	}

	// 获取需要加载的系统模型
	var envPackageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user =  ? AND userspace_id = ?", "sys", "0").Find(&envPackageModel)
	for i := 0; i < len(envPackageModel); i++ {
		packageVersion, ok := libraryAndVersions[envPackageModel[i].PackageName]
		if ok && packageVersion == envPackageModel[i].Version {
			environmentModelData[envPackageModel[i].PackageName] = envPackageModel[i].Version
		}
	}

	// 获取需要加载的加密模型
	var encryptionPackageModel []DataBaseModel.YssimModels
	DB.Where("sys_or_user =  ? AND userspace_id = ? AND encryption = ?", userName, spaceId, 1).Find(&encryptionPackageModel)
	for i := 0; i < len(encryptionPackageModel); i++ {
		packageVersion, ok := libraryAndVersions[encryptionPackageModel[i].PackageName]
		if ok && packageVersion == encryptionPackageModel[i].Version {
			environmentModelData[encryptionPackageModel[i].PackageName] = encryptionPackageModel[i].FilePath
		}
	}
	return environmentModelData
}

func GrpcSimulation(itemMap map[string]string) (string, error) {
	SimulateTypeDict := map[string]bool{"OM": true, "JM": true, "DM": true, "FmPy": true}
	if !SimulateTypeDict[itemMap["simulate_type"]] {
		return "", errors.New("不存在的仿真类型")
	}
	// 查询数据库中的实验id对应的记录, 不传实验id 则表示使用默认参数仿真
	anotherName := ""
	var experimentRecord DataBaseModel.YssimExperimentRecord
	DB.Where("id = ? ", itemMap["experiment_id"]).First(&experimentRecord)
	if experimentRecord.ID == "" {
		return "", errors.New("not found")
	} else {
		anotherName = experimentRecord.ExperimentName + "的结果"
	}
	// 查询数据库中的模型对应的记录
	var packageModel DataBaseModel.YssimModels
	err := DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", itemMap["package_id"], []string{"sys", itemMap["username"]}, []string{"0", itemMap["space_id"]}).First(&packageModel).Error
	if err != nil {
		return "", errors.New("not found")
	}

	// 查询是不是管网模型
	var pipeNetModel DataBaseModel.YssimPipeNetCadDownload
	DB.Where("package_id = ? AND model_name = ?", itemMap["package_id"], itemMap["model_name"]).First(&pipeNetModel)
	isPipeNet := false
	if pipeNetModel.ID != "" {
		isPipeNet = true
	}

	// 查询实验id对应的仿真记录
	var simulateRecord DataBaseModel.YssimSimulateRecord
	DB.Where("experiment_id = ? AND package_id = ? AND simulate_model_name = ? AND username = ? AND userspace_id = ?", itemMap["experiment_id"], itemMap["package_id"], itemMap["model_name"], itemMap["username"], itemMap["space_id"]).First(&simulateRecord)
	record := DataBaseModel.YssimSimulateRecord{}

	// 如果没有找到记录，则新建实验记录
	if simulateRecord.ID == "" {
		// SimulateStatus "1"初始(正在准备)  "2"执行  "3"失败(编译失败or仿真运行失败)  "4"成功结束  "5"关闭(killed)  "6"编译阶段
		record = DataBaseModel.YssimSimulateRecord{
			ID:                uuid.New().String(),
			PackageId:         experimentRecord.PackageId,
			UserspaceId:       experimentRecord.UserspaceId,
			UserName:          experimentRecord.UserName,
			SimulateModelName: experimentRecord.ModelName,
			SimulateStatus:    "1",
			StartTime:         experimentRecord.StartTime,
			StopTime:          experimentRecord.StopTime,
			Method:            experimentRecord.Method,
			SimulateType:      experimentRecord.SimulateType,
			NumberOfIntervals: experimentRecord.NumberOfIntervals,
			Tolerance:         experimentRecord.Tolerance,
			ExperimentId:      itemMap["experiment_id"],
			Intervals:         experimentRecord.Interval,
			AnotherName:       anotherName,
			PipeNet:           isPipeNet,
		}
		err = DB.Create(&record).Error
		if err != nil {
			return "", errors.New("新建数据库记录出现错误")
		}
		// 创建结果文件夹,并存入数据库
		resultFilePath := "static/UserFiles/ModelResult/" + itemMap["username"] + "/" + strings.ReplaceAll(itemMap["model_name"], ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/"
		createFilePathRes := fileOperation.CreateFilePath(resultFilePath)
		log.Println(resultFilePath)
		log.Println("创建路径结果", createFilePathRes)

		// 设置文件夹权限为 777
		if err_ := fileOperation.SetPermissions(resultFilePath); err_ != nil {
			log.Println(resultFilePath)
			log.Println("Error setting permissions:", err_)
		}
		record.SimulateModelResultPath = resultFilePath
		config.DB.Save(&record)
	} else {
		// 设置文件夹权限为 777
		if err_ := fileOperation.SetPermissions(simulateRecord.SimulateModelResultPath); err_ != nil {
			log.Println(simulateRecord.SimulateModelResultPath)
			log.Println("Error setting permissions:", err_)
		}
		err_ := fileOperation.ClearDirectory(simulateRecord.SimulateModelResultPath)
		if err_ != nil {
			log.Println(err_)
		}
		simulateRecord.SimulateEndTime = 0
		simulateRecord.SimulateStartTime = 0
		simulateRecord.SimulateStatus = "1"
		// 如果有找到记录，则用老的记录,并更新仿真参数
		// simulateRecord.StartTime = itemMap["start_time"]
		// simulateRecord.StopTime = itemMap["stop_time"]
		// simulateRecord.Method = itemMap["method"]
		// simulateRecord.NumberOfIntervals = itemMap["number_of_intervals"]
		// simulateRecord.Tolerance = itemMap["tolerance"]
		// simulateRecord.SimulateType = itemMap["simulate_type"]

		// 删除mongo中的记录
		if simulateRecord.TaskId != "" {
			coll := MB.Database("SimulationTasks").Collection("task_model")
			filter := bson.D{{"_id", simulateRecord.TaskId}}
			_, err = coll.DeleteOne(context.TODO(), filter)
			if err != nil {
				fmt.Println(err)
			}
		}

		config.DB.Save(&simulateRecord)
		record = simulateRecord
	}

	// 获取依赖模型和系统库
	environmentModelData := GetEnvLibraryAll(itemMap["username"], itemMap["space_id"])
	// 转为json，保存到数据库
	jsonEnvData, err := sonic.Marshal(environmentModelData)
	if err != nil {
		log.Println("环境依赖包解析错误：", err)
	}
	// SimulateTaskMap[record.ID] = record
	record.SimulateStart = true
	record.EnvModelData = jsonEnvData
	record.Percentage = 0
	record.TaskId = uuid.New().String()
	config.DB.Save(&record)

	// 获取实验参数
	params := []*TaskManagement.ParamObj{}
	if packageModel.SysUser != "sys" {
		// 取所有非全量实验的并集参数
		var experimentRecords []DataBaseModel.YssimExperimentRecord
		if err := DB.Where("username = ? AND userspace_id = ? AND package_id = ? AND model_name = ? AND create_time <= ? AND is_full_model_var != ?",
			itemMap["username"], itemMap["space_id"], itemMap["package_id"], itemMap["model_name"], experimentRecord.CreatedAt, 1).Order("create_time").Find(&experimentRecords).Error; err != nil {
			return "", errors.New("query error")
		}

		componentParamsMap := make([][]map[string]any, 0)

		for _, record := range experimentRecords {
			var componentParams []map[string]any
			if err = json.Unmarshal([]byte(record.ModelVarData), &componentParams); err != nil {
				log.Println("json to list filed!")
			}
			componentParamsMap = append(componentParamsMap, componentParams)
		}

		mapAttributesStr := mapProcessing.GetUnionComponentParams(componentParamsMap)

		for component, value := range mapAttributesStr {
			param := &TaskManagement.ParamObj{Key: component, Value: value, IsFile: false}
			params = append(params, param)
		}
	}

	// 发送仿真请求
	GrpcBuildModelRequest := &TaskManagement.TaskAssignmentsRequest{
		Uuid:          record.TaskId,
		Application:   "SimulationModeling",
		ResultAddress: record.SimulateModelResultPath,
		UserName:      record.UserName,
		TaskType:      record.SimulateType,
		FileId:        record.ID,
		Token:         itemMap["token"],
		Params:        params,
	}
	_, err = TaskManagement.CasService.Client.Assignments(context.Background(), GrpcBuildModelRequest)
	return record.ID, err

}

func DeleteSimulateTask(taskID, SimulateModelResultPath string) {

	_, err := GrpcSimulationProcessOperation(taskID)
	if err != nil {
		log.Println("调用grpc服务(GrpcPyOmcSimulationProcessOperation)出错：：", err)
	}

	err = os.RemoveAll(SimulateModelResultPath)
	if err != nil {
		log.Println(err)
		return
	}

}

func TerminateSimulateTask(taskID string) error {

	replyVar, err := GrpcSimulationProcessOperation(taskID)
	if err != nil {
		log.Println("调用grpc服务(GrpcPyOmcSimulationProcessOperation)出错：：", err)
		return err
	}

	log.Println(replyVar.Message)
	return nil
}

func GrpcSimulationProcessOperation(uid string) (*TaskManagement.TerminateTaskResponse, error) {

	PyOmcSimProcessOperationRequest := &TaskManagement.TerminateTaskRequest{
		Uuid: uid,
	}
	replyTest, err := TaskManagement.CasService.Client.TerminateTask(context.Background(), PyOmcSimProcessOperationRequest)
	return replyTest, err

}
