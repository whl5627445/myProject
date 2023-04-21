package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/library/fileOperation"
	"yssim-go/library/mapProcessing"
	"yssim-go/library/stringOperation"

	"github.com/bytedance/sonic"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
)

type SimulateTask struct {
	SRecord          *DataBaseModel.YssimSimulateRecord
	Package          *DataBaseModel.YssimModels
	ExperimentRecord *DataBaseModel.YssimExperimentRecord
	Cmd              *exec.Cmd
}

var SimulateTaskMap = make(map[string]*SimulateTask, 0)

type modelVarData struct {
	FinalAttributesStr map[string]interface{} `json:"final_attributes_str"`
}

var SimulateTaskChan = make(chan *SimulateTask, 1000)

func openModelica(task *SimulateTask, resultFilePath string, SimulationPraData map[string]string) bool {
	SimulateTaskMap[task.SRecord.ID] = task
	pwd, _ := os.Getwd()
	//pwd = "/home/xuqingda/GolandProjects/YssimGoService"
	//resultFilePath = "static/UserFiles/ModelResult/"
	//buildModelRes := omc.OMC.BuildModel(task.SRecord.SimulateModelName, pwd+"/"+resultFilePath, SimulationPraData)
	replyVar, err := GrpcPyOmcSimulation(
		task.SRecord.ID,
		task.SRecord.UserspaceId,
		task.SRecord.UserName,
		task.SRecord.SimulateModelName,
		pwd+"/"+resultFilePath,
		resultFilePath,
		SimulationPraData)
	if err != nil {
		fmt.Println("调用grpc服务(PyOmcSimulation)出错：", err)
		return false
	}
	fmt.Println("仿真提交任务:", replyVar.Msg)
	return replyVar.Msg == "Task submitted successfully."
	//
	//var resultRecord DataBaseModel.YssimSimulateRecord
	//config.DB.Where("id = ? ", task.SRecord.ID).First(&resultRecord)
	//if resultRecord.ID == "" {
	//	log.Println("编译完成-不执行仿真-进程被kill")
	//	return false, errors.New("进程被Kill")
	//}
	//
	//if buildModelRes {
	//	MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型编译成功"})
	//	// 更新状态为“2” 仿真执行中
	//	task.SRecord.SimulateStatus = "2"
	//	config.DB.Save(&task.SRecord)
	//	cmd := exec.Command(resultFilePath + "result")
	//	task.Cmd = cmd
	//	SimulateTaskMap[task.SRecord.ID] = task
	//	combinedOutput, err := cmd.Output()
	//	simulateResultStr := string(combinedOutput)
	//	if err != nil {
	//		simulateResultStr = err.Error()
	//		log.Println("err: ", err.Error())
	//		log.Println("仿真执行失败")
	//		return false, errors.New("进程被Kill")
	//	}
	//
	//	if strings.Contains(simulateResultStr, "successfully") {
	//		res = true
	//	} else {
	//		task.SRecord.SimulateStatus = "3"
	//	}
	//	task.SRecord.SimulateResultStr = simulateResultStr
	//} else {
	//	task.SRecord.SimulateStatus = "3"
	//	task.SRecord.SimulateResultStr = "编译失败"
	//}
	//config.DB.Save(&task.SRecord)
	//return res, nil
}

func dymolaSimulate(task *SimulateTask, resultFilePath string, SimulationPraData map[string]string, simulateFilePath string) (bool, error) {

	SimulateTaskMap[task.SRecord.ID] = task
	service, err := dymolaSimulateService(task, simulateFilePath, SimulationPraData, resultFilePath)
	if err != nil {
		return service, err
	}
	return service, nil
}

func dymolaSimulateService(task *SimulateTask, simulateFilePath string, SimulationPraData map[string]string, resultFilePath string) (bool, error) {
	path := task.Package.PackageName
	packageFileName := task.Package.PackageName + ".mo"
	uploadResult := false
	uploadFilePath := ""
	if task.Package.FilePath != "" {
		req := url.NewRequest()
		params := url.NewParams()
		params.Set("url", task.SRecord.UserName+"/"+path)
		req.Params = params
		req.Timeout = 600 * time.Second
		files := url.NewFiles()
		SaveModelToFile(task.Package.PackageName, simulateFilePath)
		files.SetFile("file", packageFileName, simulateFilePath, "")
		req.Files = files
		uploadFileRes, _ := requests.Post(config.DymolaSimutalionConnect+"/file/upload", req)
		uploadRes, _ := uploadFileRes.Json()
		if uploadRes["code"].(float64) == 200 {
			uploadResult = true
			uploadFilePath = uploadRes["data"].(string)
		} else {
			return false, nil
		}
	}
	if uploadResult || task.Package.FilePath == "" {
		fileName := ""
		if task.Package.FilePath != "" {
			fileName = uploadFilePath
		}
		compileReqData := map[string]interface{}{
			"userName":  task.SRecord.UserName,
			"fileName":  fileName,
			"modelName": task.SRecord.SimulateModelName,
			"taskId":    task.SRecord.ID,
		}
		req := url.NewRequest()
		req.Json = compileReqData
		req.Timeout = 10 * time.Minute
		compileRes, err := requests.Post(config.DymolaSimutalionConnect+"/dymola/translate", req)
		if err != nil {
			log.Println("dymola服务编译错误： ", err)
			return false, nil
		}
		compileResData, err := compileRes.Json()
		log.Println("dymola服务编译结果： ", compileResData)
		if err != nil {
			return false, nil
		}

		if compileResData["code"].(float64) == 200 {
			MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 编译成功，开始仿真"})
			simulateReqData := map[string]interface{}{
				"id":                0,
				"fileName":          fileName,
				"modelName":         task.SRecord.SimulateModelName,
				"userName":          task.SRecord.UserName,
				"startTime":         SimulationPraData["startTime"],
				"stopTime":          SimulationPraData["stopTime"],
				"numberOfIntervals": SimulationPraData["numberOfIntervals"],
				"outputInterval":    0.0,
				"method":            SimulationPraData["method"],
				"tolerance":         SimulationPraData["tolerance"],
				"fixedStepSize":     0.0,
				"resultFile":        "dsres",
				"initialNames":      []string{},
				"initialValues":     []float64{},
				//"initialNames":      initialNames,
				//"initialValues":     initialValues,
				"finalNames": "",
				"taskId":     task.SRecord.ID,
			}
			req.Json = simulateReqData
			simulateRes, _ := requests.Post(config.DymolaSimutalionConnect+"/dymola/simulate", req)
			simulateResData, err := simulateRes.Json()
			simulateResDataCode, ok := simulateResData["code"]
			log.Println("dymola仿真结果：", simulateResData)
			if err != nil {
				return false, nil
			}
			if ok && simulateResDataCode.(float64) == 200 {
				resFileUrl := config.DymolaSimutalionConnect + "/file/download?fileName=" + simulateResData["msg"].(string)
				fmuFileUrl := config.DymolaSimutalionConnect + "/file/download?fileName=" + compileResData["msg"].(string)
				downloadResFileUrl, err := requests.Get(resFileUrl, req)
				if err != nil {
					return false, nil
				}
				fileOperation.WriteFileByte(resultFilePath+"result_res.mat", downloadResFileUrl.Content)
				downloadFmuFileUrl, err := requests.Get(fmuFileUrl, req)
				if err != nil {
					return false, nil
				}
				fileOperation.WriteFileByte(resultFilePath+"dymola_model.fmu.zip", downloadFmuFileUrl.Content)
				err = fileOperation.UnZip(resultFilePath+"dymola_model.fmu.zip", resultFilePath)
				if err != nil {
					return false, nil
				}
				err = os.Rename(resultFilePath+"modelDescription.xml", resultFilePath+"result_init.xml")
				if err != nil {
					return false, nil
				}
				task.SRecord.SimulateResultStr = "DM"
				return true, nil
			}
		}
	}
	return false, nil
}

func dymolaServiceStop(taskID string) error {
	req := url.NewRequest()
	req.Json = map[string]interface{}{"taskId": taskID}
	req.Timeout = 10 * time.Minute
	compileRes, err := requests.Post(config.DymolaSimutalionConnect+"/dymola/stopDymola", req)
	if err != nil {
		log.Println("dymola服务停止错误： ", err)
		return err
	}
	d, err := compileRes.Json()
	if err != nil {
		return err
	}
	log.Println("暂停dymola任务结果：", d)
	if d["code"].(float64) == 200 {
		return nil
	}
	return errors.New("停止任务失败")
}

//func jModelicaSimulate(task *SimulateTask, resultFilePath string, SimulationPraData map[string]string, simulateFilePath string) bool {
//	moFilePath := "/" + simulateFilePath
//	if task.Package.FilePath == "" {
//		moFilePath = "/omlibrary/" + task.Package.PackageName + " " + task.Package.Version
//	}
//	finalTime, _ := strconv.ParseFloat(SimulationPraData["stopTime"], 64)
//	startTime, _ := strconv.ParseFloat(SimulationPraData["startTime"], 64)
//	numberOfIntervals, _ := strconv.Atoi(SimulationPraData["numberOfIntervals"])
//	tolerance, _ := strconv.ParseFloat(SimulationPraData["tolerance"], 64)
//	//if err != nil {
//	//	log.Printf("数据转换失败: %s", err)
//	//	return false
//	//}
//	data := map[string]interface{}{
//		"start_time":       startTime,
//		"final_time":       finalTime,
//		"mo_path":          moFilePath,
//		"result_name":      "result_res.mat",
//		"modelname":        task.SRecord.SimulateModelName,
//		"ncp":              numberOfIntervals,    // 结果间隔
//		"result_file_path": "/" + resultFilePath, // 结果文件名字
//		"tolerance":        tolerance,            // 相对公差
//		"type":             "compile",            // 是编译还是计算， 默认是编译
//		//"initialNames":     initialNames,
//		//"initialValues":    initialValues,
//	}
//	dial, err := net.Dial("tcp", config.JmodelicaConnect)
//	defer func(dial net.Conn) {
//		err = dial.Close()
//		if err != nil {
//			log.Println("关闭连接失败，错误： ", err)
//		}
//	}(dial)
//	if err != nil {
//		log.Printf("连接JModelica服务失败: %s", err)
//		return false
//	}
//	dataJson, _ := sonic.Marshal(data)
//	_, err = dial.Write(dataJson)
//
//	if err != nil {
//		log.Printf("发送编译数据失败: %s", data)
//		log.Printf("错误消息为: %s", err)
//		return false
//	}
//	var compileRes [1024]byte
//	n, err := dial.Read(compileRes[:])
//	if err != nil {
//		log.Printf("接收编译结果数据失败，错误为: %s", err)
//		return false
//	}
//	log.Printf("编译结果: %s", string(compileRes[:n]))
//	log.Printf("编译数据: %s", dataJson)
//	if string(compileRes[:n]) == "ok" {
//		MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 编译成功，开始仿真"})
//		modelName_ := strings.ReplaceAll(task.SRecord.SimulateModelName, ".", "_")
//		data["type"] = "simulate"
//		data["modelname"] = modelName_
//		dataJson, _ = sonic.Marshal(data)
//		dialRes, err := net.Dial("tcp", config.JmodelicaConnect)
//		defer func(dialRes net.Conn) {
//			err = dialRes.Close()
//			if err != nil {
//				log.Println("关闭连接失败，错误：", err)
//			}
//		}(dialRes)
//		_, err = dialRes.Write(dataJson)
//		log.Printf("发送仿真数据: %s", data)
//		if err != nil {
//			log.Printf("发送仿真数据失败: %s", data)
//			log.Printf("错误消息为: %s", err)
//			return false
//		}
//		var simulateRes [4096]byte
//		n, err = dialRes.Read(simulateRes[:])
//		log.Printf("仿真结果: %s", string(simulateRes[:n]))
//		if err != nil {
//			log.Printf("接收仿真结果数据失败，错误为: %s", err)
//			return false
//		}
//		if string(simulateRes[:n]) == "ok" {
//			err = os.Rename(resultFilePath+modelName_+".fmu", resultFilePath+modelName_+".fmu.zip")
//			if err != nil {
//				return false
//			}
//			err = fileOperation.UnZip(resultFilePath+modelName_+".fmu.zip", resultFilePath)
//			if err != nil {
//				return false
//			}
//			err = os.Rename(resultFilePath+"modelDescription.xml", resultFilePath+"result_init.xml")
//			task.SRecord.SimulateResultStr = "JM"
//			return true
//		}
//	}
//	return false
//}

//func fmpySimulate(task *SimulateTask, resultFilePath string, SimulationPraData map[string]string) bool {
//
//	finalTime, _ := strconv.ParseFloat(SimulationPraData["stopTime"], 64)
//	startTime, _ := strconv.ParseFloat(SimulationPraData["startTime"], 64)
//	tolerance, _ := strconv.ParseFloat(SimulationPraData["tolerance"], 64)
//	interval, _ := strconv.ParseFloat(task.SRecord.Intervals, 64)
//
//	FmuSimulationRes, err := GrpcFmuSimulation(task.SRecord.ID, task.SRecord.UserspaceId, task.Package.FilePath, task.SRecord.SimulateModelName,
//		task.SRecord.UserName, resultFilePath, startTime, finalTime, interval, tolerance)
//	if err != nil {
//		fmt.Println("调用grpc服务(FmuSimulation)出错：", err)
//		return false
//	}
//	fmt.Println("仿真提交任务:", FmuSimulationRes.Log)
//	return FmuSimulationRes.Log == "Task submitted successfully."
//
//}

func ModelSimulate(task *SimulateTask) {
	resultFilePath := "static/UserFiles/ModelResult/" + task.SRecord.UserName + "/" + strings.ReplaceAll(task.SRecord.SimulateModelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/"
	_, err := fileOperation.CreateFilePath(resultFilePath)
	if err != nil {
		log.Println("仿真目录创建错误： ", err)
		return
	}
	task.SRecord.SimulateStartTime = time.Now().Unix()
	task.SRecord.SimulateStart = true
	// 模型开始编译 状态“6”
	//task.SRecord.SimulateStatus = "6"
	task.SRecord.SimulateModelResultPath = resultFilePath

	config.DB.Save(&task.SRecord)
	if task.Package.SysUser != "sys" {
		//YssimExperimentRecord表的json数据绑定到结构体
		var componentValue modelVarData
		if task.ExperimentRecord.ModelVarData.String() != "" {
			err = sonic.Unmarshal(task.ExperimentRecord.ModelVarData, &componentValue)
			if err == nil {
				mapAttributesStr := mapProcessing.MapDataConversion(componentValue.FinalAttributesStr)
				//设置组件参数
				result := SetComponentModifierValue(task.ExperimentRecord.ModelName, mapAttributesStr)
				if result {
					log.Println("重新设置参数-完成。")
				} else {
					log.Println("重新设置参数-失败: ", mapAttributesStr)
				}
			} else {
				log.Println("modelVarData: ", task.ExperimentRecord.ModelVarData)
				log.Println("err: ", err)
				log.Println("json2map filed!")
			}
		}
	}
	FilePath := "static/tmp/simulateModelFile/" + task.SRecord.UserName + "/" + time.Now().Local().Format("20060102150405") + "/" + task.Package.PackageName + ".mo"

	MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型开始编译"})
	sResult := true
	SimulationPraData := map[string]string{
		"startTime": task.SRecord.StartTime,
		"stopTime":  task.SRecord.StopTime,
		"method":    task.SRecord.Method,
		// "outputFormat": "\"csv\"",  // csv不能使用omc的api读取结果
		// "numberOfIntervals": "500",
		"numberOfIntervals": task.SRecord.NumberOfIntervals,
		"tolerance":         task.SRecord.Tolerance,
	}
	if task.Package.FilePath != "" && (task.SRecord.SimulateType == "DM" || task.SRecord.SimulateType == "JM") {
		sResult = SaveModelToFile(task.Package.PackageName, FilePath)
		if !sResult {
			task.SRecord.SimulateStatus = "3"
			MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 出现错误，请联系管理员"})
			task.SRecord.SimulateEndTime = time.Now().Unix()
			task.SRecord.SimulateStart = false
			config.DB.Save(&task.SRecord)
			return
		}
	}
	switch task.SRecord.SimulateType {
	case "OM":
		sResult = openModelica(task, resultFilePath, SimulationPraData)
		return
	case "DM":
		sResult, err = dymolaSimulate(task, resultFilePath, SimulationPraData, FilePath)
	//case "JM":
	//	sResult = jModelicaSimulate(task, resultFilePath, SimulationPraData, FilePath)
	case "FmPy":
		sResult = openModelica(task, resultFilePath, SimulationPraData)
		//sResult = fmpySimulate(task, resultFilePath, SimulationPraData)
		//if !sResult {
		//	task.SRecord.SimulateStatus = "3"
		//	MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型仿真失败"})
		//	task.SRecord.SimulateEndTime = time.Now().Unix()
		//	task.SRecord.SimulateStart = false
		//	config.DB.Save(&task.SRecord)
		//}
		//return
	}
	if err != nil {
		// 仿真进程被杀掉后直接退出
		return
	}
	if sResult {
		task.SRecord.SimulateModelResultPath = resultFilePath
		task.SRecord.SimulateStatus = "4"
		task.SRecord.AnotherName = stringOperation.NewAnotherName(task.SRecord.UserName, task.SRecord.SimulateModelName, task.SRecord.UserspaceId)
		MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型仿真完成"})
	} else {
		task.SRecord.SimulateStatus = "3"
		MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型仿真失败"})
	}
	task.SRecord.SimulateEndTime = time.Now().Unix()
	task.SRecord.SimulateStart = false
	config.DB.Save(&task.SRecord)
	basePath := strings.Join(strings.Split(FilePath, "/")[:2], "/")
	err = os.RemoveAll(basePath)
	if err != nil {
		log.Println("basePath err", err)
		return
	}
}

func DeleteSimulateTask(taskID, simulateType string) {
	task, ok := SimulateTaskMap[taskID]
	switch simulateType {
	case "OM":
		if ok {
			replyVar, err := GrpcPyOmcSimulationProcessOperation(taskID, "kill")
			if err != nil {
				log.Println("调用grpc服务(GrpcPyOmcSimulationProcessOperation)出错：：", err)
			}
			log.Println(replyVar.Msg)
		}
	case "DM":
		if ok {
			err := dymolaServiceStop(taskID)
			if err != nil {
				log.Println(err)
			}
		}
	}
	os.RemoveAll(task.SRecord.SimulateModelResultPath)
}
