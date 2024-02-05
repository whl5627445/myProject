package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/grpc/PythonGrpc/grpcPb"
	"yssim-go/library/mapProcessing"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

type modelVarData struct {
	FinalAttributesStr map[string]any `json:"final_attributes_str"`
}
type OutputData struct {
	Name string `json:"data"`
	Unit string `json:"unit"`
}

var DB = config.DB

// GetEnvLibraryAll 获取当前环境下的所有已经加载的包和系统库
// func GetEnvLibraryAll(userName, spaceId string) map[string]string {
//
//	// 获取系统模型
//	environmentModelData := make(map[string]string)
//	var envPackageModel []DataBaseModel.YssimModels
//	DB.Where("sys_or_user =  ? AND userspace_id = ?", "sys", "0").Find(&envPackageModel)
//	libraryAndVersions := GetLibraryAndVersions()
//	for i := 0; i < len(envPackageModel); i++ {
//		p, ok := libraryAndVersions[envPackageModel[i].PackageName]
//		if ok && p == envPackageModel[i].Version {
//			environmentModelData[envPackageModel[i].PackageName] = envPackageModel[i].Version
//		}
//	}
//	// 获取用户模型
//	DB.Where("sys_or_user = ? AND userspace_id = ?", userName, spaceId).Find(&envPackageModel)
//	for i := 0; i < len(envPackageModel); i++ {
//		loadVersions, ok := libraryAndVersions[envPackageModel[i].PackageName]
//		if ok && loadVersions == envPackageModel[i].Version {
//			environmentModelData[envPackageModel[i].PackageName] = envPackageModel[i].FilePath
//		}
//	}
//
//	// 获取加密模型
//	var encryptionPackageModel []DataBaseModel.YssimModels
//	DB.Where("sys_or_user =  ? AND userspace_id = ? AND encryption = ?", userName, spaceId, 1).Find(&encryptionPackageModel)
//	for i := 0; i < len(encryptionPackageModel); i++ {
//		packageVersion, ok := libraryAndVersions[encryptionPackageModel[i].PackageName]
//		if ok && packageVersion == encryptionPackageModel[i].Version {
//			environmentModelData[encryptionPackageModel[i].PackageName] = encryptionPackageModel[i].FilePath
//		}
//	}
//	return environmentModelData
//
// }

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

// func GrpcReadSimulationResult(VarList []string, SimulateModelResultPath string) ([][]float64, bool) {
//
//	SaveFilterResultTest := &grpcPb.ReadSimulationResultRequest{ // 构造请求体
//		Vars:       VarList,
//		ResultPath: SimulateModelResultPath,
//	}
//	reply, err := grpcPb.Client.ReadSimulationResult(grpcPb.Ctx, SaveFilterResultTest) // 调用grpc服务
//	data := reply.Data
//	var replyData [][]float64
//	for i := 0; i < len(data); i++ {
//		replyData = append(replyData, data[i].Row)
//	}
//
//	// 二维数组转置
//	rows, cols := len(replyData), len(replyData[0])
//	result := make([][]float64, cols)
//	for i := range result {
//		result[i] = make([]float64, rows)
//	}
//	for i := 0; i < rows; i++ {
//		for j := 0; j < cols; j++ {
//			result[j][i] = replyData[i][j]
//		}
//	}
//
//	var ok bool
//	ok = reply.Ok
//	if err != nil {
//		fmt.Println("调用grpc服务(SaveFilterResultToCsv)出错：", err)
//		return result, false
//	}
//	return result, ok
//
// }

// func GrpcZarrToCsv(SimulateModelResultPath string) bool {
//	ZarrToCsvRequestTest := &grpcPb.ZarrToCsvRequest{
//		ZarrPath: SimulateModelResultPath + "zarr_res.zarr",
//	} // 构造请求体
//	ZarrToCsvRes, err := grpcPb.Client.ZarrToCsv(grpcPb.Ctx, ZarrToCsvRequestTest) // 调用grpc服务
//	if err != nil {
//		fmt.Println("调用grpc服务(ZarrToCsv)出错：", err)
//		return false
//	}
//	return ZarrToCsvRes.Ok
// }
//
// func GrpcMatToCsv(SimulateModelResultPath string) bool {
//	MatToCsvRequestTest := &grpcPb.MatToCsvRequest{
//		MatPath: SimulateModelResultPath + "result_res.mat",
//	} // 构造请求体
//	MatToCsvRes, err := grpcPb.Client.MatToCsv(grpcPb.Ctx, MatToCsvRequestTest) // 调用grpc服务
//	if err != nil {
//		fmt.Println("调用grpc服务(MatToCsv)出错：", err)
//		return false
//	}
//	return MatToCsvRes.Ok
// }

func GrpcCheckVarExist(path string, dataNameList []string) map[string]bool {
	CheckVarExistRequestTest := &grpcPb.CheckVarExistRequest{
		Path:  path,
		Names: dataNameList,
	}
	replyTest, err2 := grpcPb.Client.CheckVarExist(grpcPb.Ctx, CheckVarExistRequestTest)
	if err2 != nil {
		fmt.Println("调用grpc服务(CheckVarExist)出错：", err2)
	}
	return replyTest.Res
}

func GrpcGetResult(matPath string, varName string) (*grpcPb.GetResultReply, error) {
	GetResultRequestVar := &grpcPb.GetResultRequest{
		Path:     matPath,
		Variable: varName,
	}
	replyVar, err := grpcPb.Client.GetResult(grpcPb.Ctx, GetResultRequestVar)
	if err != nil {
		fmt.Println("调用grpc服务(GrpcGetResult)出错：", err)
	}
	return replyVar, err
}

func GrpcSimulation(itemMap map[string]string) (string, error) {
	SimulateTypeDict := map[string]bool{"OM": true, "JM": true, "DM": true, "FmPy": true}
	if !SimulateTypeDict[itemMap["simulate_type"]] {
		return "", errors.New("不存在的仿真类型")
	}
	// 查询数据库中的实验id对应的记录
	var experimentRecord DataBaseModel.YssimExperimentRecord
	DB.Where("id = ? ", itemMap["experiment_id"]).First(&experimentRecord)
	// 查询数据库中的模型对应的记录
	var packageModel DataBaseModel.YssimModels
	err := DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", itemMap["package_id"], []string{"sys", itemMap["username"]}, []string{"0", itemMap["space_id"]}).First(&packageModel).Error
	if err != nil {
		return "", errors.New("not found")
	}
	// SimulateStatus "1"初始(正在准备)  "2"执行  "3"失败(编译失败or仿真运行失败)  "4"成功结束  "5"关闭(killed)  "6"编译阶段
	record := DataBaseModel.YssimSimulateRecord{
		ID:                uuid.New().String(),
		PackageId:         itemMap["package_id"],
		UserspaceId:       itemMap["space_id"],
		UserName:          itemMap["username"],
		SimulateModelName: itemMap["model_name"],
		SimulateStatus:    "1",
		StartTime:         itemMap["start_time"],
		StopTime:          itemMap["stop_time"],
		Method:            itemMap["method"],
		SimulateType:      itemMap["simulate_type"],
		NumberOfIntervals: itemMap["number_of_intervals"],
		Tolerance:         itemMap["tolerance"],
		ExperimentId:      itemMap["experiment_id"],
		Intervals:         itemMap["interval"],
	}
	err = DB.Create(&record).Error
	if err != nil {
		return "", errors.New("新建数据库记录出现错误")
	}
	// 创建结果文件夹,并存入数据库
	resultFilePath := "static/UserFiles/ModelResult/" + itemMap["username"] + "/" + strings.ReplaceAll(itemMap["model_name"], ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/"
	// fileOperation.CreateFilePath(resultFilePath)
	record.SimulateModelResultPath = resultFilePath
	config.DB.Save(&record)
	// 将实验参数写入模型
	if packageModel.SysUser != "sys" {
		// YssimExperimentRecord表的json数据绑定到结构体
		var componentValue modelVarData
		if experimentRecord.ModelVarData.String() != "" {
			err := sonic.Unmarshal(experimentRecord.ModelVarData, &componentValue)
			if err == nil {
				mapAttributesStr := mapProcessing.MapDataConversion(componentValue.FinalAttributesStr)
				// 设置组件参数
				result := SetComponentModifierValue(experimentRecord.ModelName, mapAttributesStr)
				if result {
					log.Println("重新设置参数-完成。")
				} else {
					log.Println("重新设置参数-失败: ", mapAttributesStr)
				}
			} else {
				log.Println("modelVarData: ", experimentRecord.ModelVarData)
				log.Println("err: ", err)
				log.Println("json2map filed!")
			}
		}
	}
	// 构建仿真参数
	SimulationPraData := map[string]string{
		"startTime": record.StartTime,
		"stopTime":  record.StopTime,
		"method":    record.Method,
		// "outputFormat": "\"csv\"",  // csv不能使用omc的api读取结果
		// "numberOfIntervals": "500",
		"numberOfIntervals": record.NumberOfIntervals,
		"tolerance":         record.Tolerance,
	}

	// 获取依赖模型和系统库
	environmentModelData := GetEnvLibrary(packageModel.PackageName, itemMap["username"], itemMap["space_id"])
	// 转为json，保存到数据库
	jsonEnvData, err := sonic.Marshal(environmentModelData)
	if err != nil {
		log.Println("环境依赖包解析错误：", err)
	}
	// SimulateTaskMap[record.ID] = record
	record.SimulateStart = true
	record.EnvModelData = jsonEnvData
	config.DB.Save(&record)
	// 发送仿真请求
	GrpcBuildModelRequest := &grpcPb.SubmitTaskRequest{

		Uuid:              record.ID,
		UserSpaceId:       record.UserspaceId,
		UserName:          record.UserName,
		SimulatePackageId: record.PackageId,
		SimulateModelName: record.SimulateModelName,
		ResultFilePath:    resultFilePath,
		SimulationPraData: SimulationPraData,
		EnvModelData:      environmentModelData,
		SimulateType:      record.SimulateType, // OM DM
		// dm才会用到的参数
		PackageName:     packageModel.PackageName,
		PackageFilePath: packageModel.FilePath,
		// 任务类型simulate  translate run_result
		TaskType: "simulate",
	}
	_, err = grpcPb.Client.SubmitTask(grpcPb.Ctx, GrpcBuildModelRequest)
	return record.ID, err

}

func GrpcTranslate(record DataBaseModel.AppDataSource) (string, error) {
	// 创建文件夹
	//fileOperation.CreateFilePath(record.CompilePath)
	// 获取仿真参数
	SimulationPra := GetSimulationOptions(record.ModelName)
	record.StartTime = SimulationPra["startTime"]
	record.StopTime = SimulationPra["stopTime"]
	record.Method = SimulationPra["method"]
	record.Tolerance = SimulationPra["tolerance"]
	record.NumberOfIntervals = SimulationPra["numberOfIntervals"]
	// record.CompileType = SimulationPra["simulate_type"]
	err := DB.Save(&record).Error
	if err != nil {
		return "", errors.New("save error")
	}
	// 查询数据库中的实验id对应的记录
	var experimentRecord DataBaseModel.YssimExperimentRecord
	DB.Where("id = ? ", record.ExperimentId).First(&experimentRecord)
	// 查询数据库中的模型对应的记录
	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", record.PackageId, []string{"sys", record.UserName}, []string{"0", record.UserSpaceId}).First(&packageModel).Error
	if err != nil {
		return "", errors.New("not found")
	}
	// 将实验参数写入模型
	if packageModel.SysUser != "sys" {
		// YssimExperimentRecord表的json数据绑定到结构体
		var componentValue modelVarData
		if experimentRecord.ModelVarData.String() != "" {
			err = sonic.Unmarshal(experimentRecord.ModelVarData, &componentValue)
			if err == nil {
				mapAttributesStr := mapProcessing.MapDataConversion(componentValue.FinalAttributesStr)
				// 设置组件参数
				result := SetComponentModifierValue(experimentRecord.ModelName, mapAttributesStr)
				if result {
					log.Println("重新设置参数-完成。")
				} else {
					log.Println("重新设置参数-失败: ", mapAttributesStr)
				}
			} else {
				log.Println("modelVarData: ", experimentRecord.ModelVarData)
				log.Println("err: ", err)
				log.Println("json2map filed!")
			}
		}
	}
	// 构建仿真参数
	SimulationPraData := map[string]string{
		"startTime":         record.StartTime,
		"stopTime":          record.StopTime,
		"method":            record.Method,
		"numberOfIntervals": record.NumberOfIntervals,
		"tolerance":         record.Tolerance,
	}

	// 获取依赖模型和系统库
	environmentModelData := GetEnvLibrary(record.ModelName, record.UserName, record.UserSpaceId)
	// 转为json，保存到数据库
	jsonEnvData, err := sonic.Marshal(environmentModelData)
	if err != nil {
		log.Println("环境依赖包解析错误：", err)
	}
	// SimulateTaskMap[record.ID] = record
	// record.SimulateStart = true
	record.EnvModelData = jsonEnvData
	config.DB.Save(&record)
	// 发送仿真请求
	GrpcBuildModelRequest := &grpcPb.SubmitTaskRequest{

		Uuid:              record.ID,
		UserSpaceId:       record.UserSpaceId,
		UserName:          record.UserName,
		SimulateModelName: record.ModelName,
		ResultFilePath:    record.CompilePath,
		SimulationPraData: SimulationPraData,
		EnvModelData:      environmentModelData,
		SimulateType:      record.CompileType, // OM DM
		// dm才会用到的参数
		PackageName:     packageModel.PackageName,
		PackageFilePath: packageModel.FilePath,
		// 任务类型simulate  translate run_result
		TaskType: "translate",
	}
	_, err = grpcPb.Client.SubmitTask(grpcPb.Ctx, GrpcBuildModelRequest)
	return record.ID, err

}

func GrpcRunResult(appPageId string, singleSimulationInputData map[string]float64) error {
	// OM和DM 多轮仿真
	var appPageRecord DataBaseModel.AppPage
	// 查询appPageId是否存在
	err := DB.Where("id = ? ", appPageId).First(&appPageRecord).Error
	if err != nil {
		return errors.New("多轮仿真服务出错！请联系管理员！")
	}
	// 查询对应的数据源是否存在
	var record DataBaseModel.AppDataSource
	err = DB.Where("id = ? ", appPageRecord.DataSourceId).First(&record).Error
	if err != nil {
		return errors.New("数据源不存在！")
	}
	// 构建仿真参数
	SimulationPraData := map[string]string{
		"startTime":         record.StartTime,
		"stopTime":          record.StopTime,
		"method":            record.Method,
		"numberOfIntervals": record.NumberOfIntervals,
		"tolerance":         record.Tolerance,
	}
	// 查询数据库中的模型表
	// var packageModel DataBaseModel.YssimModels
	// err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", record.PackageId, []string{"sys", record.UserName}, []string{"0", record.UserSpaceId}).First(&packageModel).Error
	// if err != nil {
	//	return errors.New("模型不存在！")
	// }
	// 获取依赖
	var environmentModelData map[string]string
	err = sonic.Unmarshal(record.EnvModelData, &environmentModelData)
	if err != nil {
		return errors.New("多轮仿真服务出错！请联系管理员！")
	}

	inputData := make(map[string][]float64)
	// 获取数据库中的输出名称数组
	var data []OutputData
	err = sonic.Unmarshal(appPageRecord.Output, &data)
	if err != nil {
		return errors.New("多轮仿真服务出错！请联系管理员！")
	}
	var outputNames []string
	for _, d := range data {
		outputNames = append(outputNames, d.Name)
	}
	var singleOrMultiple string
	if singleSimulationInputData != nil {
		// log.Println("单次仿真！")
		singleOrMultiple = "single"
		for key, value := range singleSimulationInputData {
			var newValues []float64
			newValues = append(newValues, value)
			// 单次仿真的inputData map中的value是长度为1的数组
			inputData[key] = newValues
		}
	} else {
		singleOrMultiple = "multiple"
		// log.Println("多轮仿真！")
		// 查询数据库中的模型表
		var componentRecord []DataBaseModel.AppPageComponent
		DB.Where("page_id = ? AND type = ? AND input_name != ''", appPageId, "slider").Find(&componentRecord)
		for i := 0; i < len(componentRecord); i++ {
			// 将[1,0.5,5]转换为[1,1.5,2,2.5,3,3.5,4,4.5,5]
			minVal := componentRecord[i].Min
			step := componentRecord[i].Interval
			// maxVal := componentRecord[i].Max
			// 计算新的数组元素
			var newValues []float64
			if step == 0 {
				break
			}
			newValues = append(newValues, minVal)
			for j := 0; j <= 8; j++ {
				minVal = minVal + step
				newValues = append(newValues, minVal)
			}
			inputData[componentRecord[i].InputName] = newValues
		}
	}

	inputValData := make(map[string]*grpcPb.SubmitTaskRequestInputObj)
	for k, v := range inputData {
		inputValData[k] = &grpcPb.SubmitTaskRequestInputObj{
			InputObjList: v,
		}
	}
	// 构建grpc请求体
	GrpcBuildModelRequest := &grpcPb.SubmitTaskRequest{

		Uuid:              record.ID,
		UserSpaceId:       record.UserSpaceId,
		UserName:          record.UserName,
		SimulateModelName: record.ModelName,
		ResultFilePath:    record.CompilePath,
		SimulationPraData: SimulationPraData,
		EnvModelData:      environmentModelData,
		SimulateType:      record.CompileType, // OM DM
		// dm才会用到的参数
		// PackageName:     packageModel.PackageName,
		PackageFilePath: record.ZipMoPath,
		// 任务类型"simulate " "translate " "run"三种
		TaskType: "run",
		// 多轮仿真用到的参数
		SingleOrMultiple: singleOrMultiple, // 单次仿真为“single”，多次仿真为“multiple”
		PageId:           appPageId,
		MulResultPath:    appPageRecord.MulResultPath,
		OutputValNames:   outputNames,
		InputValData:     inputValData,
	}
	_, err = grpcPb.Client.SubmitTask(grpcPb.Ctx, GrpcBuildModelRequest)
	if err != nil {
		return errors.New("多轮仿真服务出错！请联系管理员！")
	} else {
		return err
	}
}

func DeleteSimulateTask(taskID, simulateType, SimulateModelResultPath string) {

	replyVar, err := GrpcSimulationProcessOperation(taskID, "kill", simulateType)
	if err != nil {
		log.Println("调用grpc服务(GrpcPyOmcSimulationProcessOperation)出错：：", err)
	}
	log.Println(replyVar.Msg)

	err = os.RemoveAll(SimulateModelResultPath)
	if err != nil {
		log.Println(err)
		return
	}

}

func GrpcSimulationProcessOperation(uid, operation, simulateType string) (*grpcPb.ProcessOperationReply, error) {
	PyOmcSimProcessOperationRequest := &grpcPb.ProcessOperationRequest{
		Uuid:          uid,
		OperationName: operation,
		SimulateType:  simulateType,
	}
	replyTest, err := grpcPb.Client.ProcessOperation(grpcPb.Ctx, PyOmcSimProcessOperationRequest)
	return replyTest, err
}

func GrpcCalibrationCompile(data map[string]string, EnvModelData map[string]string) error {

	// 发送编译请求
	GrpcBuildModelRequest := &grpcPb.SubmitTaskRequest{
		Uuid:              data["id"],
		UserSpaceId:       data["user_space_id"],
		UserName:          data["username"],
		PackageName:       data["package_name"],
		TaskType:          "compile",
		SimulateModelName: data["model_name"],
		EnvModelData:      EnvModelData,
		SimulateType:      "calibrationCompile",
		ResultFilePath:    data["result_file_path"] + "/",
	}
	_, err := grpcPb.Client.SubmitTask(grpcPb.Ctx, GrpcBuildModelRequest)
	return err

}

func GrpcCalibrationSimulate(data map[string]string) error {

	// 发送仿真请求
	GrpcBuildModelRequest := &grpcPb.SubmitTaskRequest{
		Uuid:         data["id"],
		TaskType:     "simulate",
		SimulateType: "calibrationSimulate",
	}
	_, err := grpcPb.Client.SubmitTask(grpcPb.Ctx, GrpcBuildModelRequest)
	return err

}

func GrpcFittingCalculation(uid string) (*grpcPb.FittingCalculationReply, error) {
	FittingCalculationRequest := &grpcPb.FittingCalculationRequest{
		Uuid: uid,
	}
	res, err := grpcPb.Client.FittingCalculation(grpcPb.Ctx, FittingCalculationRequest)
	return res, err
}

func DeleteCalculationSimulateTask(taskID string) error {
	replyVar, err := GrpcSimulationProcessOperation(taskID, "kill", "")
	if err != nil {
		log.Println(replyVar.Msg)
		log.Println("调用grpc服务(GrpcPyOmcSimulationProcessOperation)出错：：", err)
		return err
	}
	return nil
}
