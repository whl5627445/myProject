package service

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/grpc/grpcPb"
	"yssim-go/library/fileOperation"
	"yssim-go/library/mapProcessing"
)

type modelVarData struct {
	FinalAttributesStr map[string]interface{} `json:"final_attributes_str"`
}

var DB = config.DB

func GetEnvLibrary(userName, spaceId string) map[string]string {

	// 获取需要加载的系统模型
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
	// 获取需要加载的用户模型
	DB.Where("sys_or_user = ? AND userspace_id = ?", userName, spaceId).Find(&envPackageModel)
	for i := 0; i < len(envPackageModel); i++ {
		loadVersions, ok := libraryAndVersions[envPackageModel[i].PackageName]
		if ok && loadVersions == envPackageModel[i].Version {
			environmentModelData[envPackageModel[i].PackageName] = envPackageModel[i].FilePath
		}
	}
	return environmentModelData

}
func GrpcReadSimulationResult(VarList []string, SimulateModelResultPath string) ([][]float64, bool) {

	SaveFilterResultTest := &grpcPb.ReadSimulationResultRequest{ // 构造请求体
		Vars:       VarList,
		ResultPath: SimulateModelResultPath,
	}
	reply, err := grpcPb.Client.ReadSimulationResult(grpcPb.Ctx, SaveFilterResultTest) // 调用grpc服务
	data := reply.Data
	var replyData [][]float64
	for i := 0; i < len(data); i++ {
		replyData = append(replyData, data[i].Row)
	}

	// 二维数组转置
	rows, cols := len(replyData), len(replyData[0])
	result := make([][]float64, cols)
	for i := range result {
		result[i] = make([]float64, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = replyData[i][j]
		}
	}

	var ok bool
	ok = reply.Ok
	if err != nil {
		fmt.Println("调用grpc服务(SaveFilterResultToCsv)出错：", err)
		return result, false
	}
	return result, ok

}

func GrpcZarrToCsv(SimulateModelResultPath string) bool {
	ZarrToCsvRequestTest := &grpcPb.ZarrToCsvRequest{
		ZarrPath: SimulateModelResultPath + "zarr_res.zarr",
	} // 构造请求体
	ZarrToCsvRes, err := grpcPb.Client.ZarrToCsv(grpcPb.Ctx, ZarrToCsvRequestTest) // 调用grpc服务
	if err != nil {
		fmt.Println("调用grpc服务(ZarrToCsv)出错：", err)
		return false
	}
	return ZarrToCsvRes.Ok
}

func GrpcMatToCsv(SimulateModelResultPath string) bool {
	MatToCsvRequestTest := &grpcPb.MatToCsvRequest{
		MatPath: SimulateModelResultPath + "result_res.mat",
	} // 构造请求体
	MatToCsvRes, err := grpcPb.Client.MatToCsv(grpcPb.Ctx, MatToCsvRequestTest) // 调用grpc服务
	if err != nil {
		fmt.Println("调用grpc服务(MatToCsv)出错：", err)
		return false
	}
	return MatToCsvRes.Ok
}

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

func GrpcGetResult(recordId string, varName string) (*grpcPb.GetResultReply, error) {
	GetResultRequestVar := &grpcPb.GetResultRequest{
		Uuid:     recordId,
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
	//查询数据库中的实验id对应的记录
	var experimentRecord DataBaseModel.YssimExperimentRecord
	DB.Where("id = ? ", itemMap["experiment_id"]).First(&experimentRecord)
	//查询数据库中的模型对应的记录
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
	fileOperation.CreateFilePath(resultFilePath)
	record.SimulateModelResultPath = resultFilePath
	config.DB.Save(&record)
	// 将实验参数写入模型
	if packageModel.SysUser != "sys" {
		//YssimExperimentRecord表的json数据绑定到结构体
		var componentValue modelVarData
		if experimentRecord.ModelVarData.String() != "" {
			err := sonic.Unmarshal(experimentRecord.ModelVarData, &componentValue)
			if err == nil {
				mapAttributesStr := mapProcessing.MapDataConversion(componentValue.FinalAttributesStr)
				//设置组件参数
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

	// 获取需要加载的系统模型
	environmentModelData := GetEnvLibrary(itemMap["username"], itemMap["space_id"])
	// 转为json，保存到数据库
	jsonEnvData, err := sonic.Marshal(environmentModelData)
	if err != nil {
		log.Println("环境依赖包解析错误：", err)
	}
	//SimulateTaskMap[record.ID] = record
	record.SimulateStart = true
	record.EnvModelData = jsonEnvData
	config.DB.Save(&record)
	// 发送仿真请求
	GrpcBuildModelRequest := &grpcPb.SubmitTaskRequest{

		Uuid:              record.ID,
		UserSpaceId:       record.UserspaceId,
		UserName:          record.UserName,
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
	//查询数据库中的实验id对应的记录
	var experimentRecord DataBaseModel.YssimExperimentRecord
	DB.Where("id = ? ", record.ExperimentId).First(&experimentRecord)
	//查询数据库中的模型对应的记录
	var packageModel DataBaseModel.YssimModels
	err := DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", record.PackageId, []string{"sys", record.UserName}, []string{"0", record.UserSpaceId}).First(&packageModel).Error
	if err != nil {
		return "", errors.New("not found")
	}
	// 将实验参数写入模型
	if packageModel.SysUser != "sys" {
		//YssimExperimentRecord表的json数据绑定到结构体
		var componentValue modelVarData
		if experimentRecord.ModelVarData.String() != "" {
			err = sonic.Unmarshal(experimentRecord.ModelVarData, &componentValue)
			if err == nil {
				mapAttributesStr := mapProcessing.MapDataConversion(componentValue.FinalAttributesStr)
				//设置组件参数
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

	// 获取需要加载的系统模型
	environmentModelData := GetEnvLibrary(record.UserName, record.UserSpaceId)
	// 转为json，保存到数据库
	jsonEnvData, err := sonic.Marshal(environmentModelData)
	if err != nil {
		log.Println("环境依赖包解析错误：", err)
	}
	//SimulateTaskMap[record.ID] = record
	//record.SimulateStart = true
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

//
//
//func GrpcPyOmcSimulation(uid, userSpaceId, userName, simulateModelName, resultFilePath string, simulationPraData, envModelData map[string]string) (*grpcPb.PyOmcSimulationReply, error) {
//
//	GrpcBuildModelRequest := &grpcPb.PyOmcSimulationRequest{
//		Uuid:              uid,
//		UserSpaceId:       userSpaceId,
//		UserName:          userName,
//		SimulateModelName: simulateModelName,
//		ResultFilePath:    resultFilePath,
//		SimulationPraData: simulationPraData,
//		EnvModelData:      envModelData,
//	}
//	replyTest, err := grpcPb.Client.PyOmcSimulation(grpcPb.Ctx, GrpcBuildModelRequest)
//	return replyTest, err
//
//}

//
//func GrpcFmuSimulation(ID, userSpaceId, FilePath, SimulateModelName, Username, resultFilePath string, startTime, finalTime, interval, tolerance float64) (*grpcPb.FmuSimulationReply, error) {
//	FmuSimulationRequestTest := &grpcPb.FmuSimulationRequest{
//		Uuid:           ID,
//		UserSpaceId:    userSpaceId,
//		MoPath:         FilePath,
//		ClassName:      SimulateModelName,
//		UserName:       Username,
//		StartTime:      startTime,
//		StopTime:       finalTime,
//		ResPath:        resultFilePath,
//		OutputInterval: interval,
//		Tolerance:      tolerance,
//	} // 构造请求体
//	FmuSimulationRes, err := grpcPb.Client.FmuSimulation(grpcPb.Ctx, FmuSimulationRequestTest) // 调用grpc服务
//	if err != nil {
//		fmt.Println("调用grpc服务(FmuSimulation)出错：", err)
//	}
//	return FmuSimulationRes, err
//
//}
