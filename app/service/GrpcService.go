package service

import (
	"fmt"
	"yssim-go/config"
	"yssim-go/grpc/grpcPb"
)

var DB = config.DB

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

func GrpcFmuSimulation(ID, userSpaceId, FilePath, SimulateModelName, Username, resultFilePath string, startTime, finalTime, interval, tolerance float64) (*grpcPb.FmuSimulationReply, error) {
	FmuSimulationRequestTest := &grpcPb.FmuSimulationRequest{
		Uuid:           ID,
		UserSpaceId:    userSpaceId,
		MoPath:         FilePath,
		ClassName:      SimulateModelName,
		UserName:       Username,
		StartTime:      startTime,
		StopTime:       finalTime,
		ResPath:        resultFilePath,
		OutputInterval: interval,
		Tolerance:      tolerance,
	} // 构造请求体
	FmuSimulationRes, err := grpcPb.Client.FmuSimulation(grpcPb.Ctx, FmuSimulationRequestTest) // 调用grpc服务
	if err != nil {
		fmt.Println("调用grpc服务(FmuSimulation)出错：", err)
	}
	return FmuSimulationRes, err

}

func GrpcPyOmcSimulation(uid, userSpaceId, userName, simulateModelName, resultFilePath string, simulationPraData, envModelData map[string]string) (*grpcPb.PyOmcSimulationReply, error) {

	GrpcBuildModelRequest := &grpcPb.PyOmcSimulationRequest{
		Uuid:              uid,
		UserSpaceId:       userSpaceId,
		UserName:          userName,
		SimulateModelName: simulateModelName,
		ResultFilePath:    resultFilePath,
		SimulationPraData: simulationPraData,
		EnvModelData:      envModelData,
	}
	replyTest, err := grpcPb.Client.PyOmcSimulation(grpcPb.Ctx, GrpcBuildModelRequest)
	return replyTest, err

}

func GrpcPyOmcSimulationProcessOperation(uid, operation string) (*grpcPb.ProcessOperationReply, error) {
	PyOmcSimProcessOperationRequest := &grpcPb.ProcessOperationRequest{
		Uuid:          uid,
		OperationName: operation,
	}
	replyTest, err := grpcPb.Client.ProcessOperation(grpcPb.Ctx, PyOmcSimProcessOperationRequest)
	return replyTest, err
}
