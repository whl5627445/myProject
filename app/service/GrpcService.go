package service

import (
	"fmt"
	"yssim-go/grpc/grpcPb"
)

func GrpcSaveFilterResultToCsv(VarList []string, SimulateModelResultPath string, newFileName string) bool {
	SaveFilterResultTest := &grpcPb.SaveFilterResultToCsvRequest{ // 构造请求体
		Vars:        VarList,
		ResultPath:  SimulateModelResultPath + "zarr_res.zarr",
		NewFileName: newFileName,
	}
	reply, err := grpcPb.Client.SaveFilterResultToCsv(grpcPb.Ctx, SaveFilterResultTest) // 调用grpc服务
	var ok bool
	ok = reply.Ok
	if err != nil {
		fmt.Println("调用grpc服务(SaveFilterResultToCsv)出错：", err)
		return false
	}
	return ok
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

func GrpcFmuSimulation(ID, FilePath, SimulateModelName, Username, resultFilePath string, startTime, finalTime, interval, tolerance float64) (*grpcPb.FmuSimulationReply, error) {
	FmuSimulationRequestTest := &grpcPb.FmuSimulationRequest{
		Uuid:           ID,
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