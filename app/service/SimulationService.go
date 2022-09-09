package service

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
)

type SimulateTask struct {
	SRecord DataBaseModel.YssimSimulateRecord
	Package DataBaseModel.YssimModels
}

var SimulateTaskChan = make(chan *SimulateTask, 100)

func openModelica(task *SimulateTask, resultFilePath string) bool {
	res := false
	SimulationPraData := map[string]string{
		"startTime": task.SRecord.StartTime,
		"stopTime":  task.SRecord.StopTime,
		"method":    task.SRecord.Method,
		//"outputFormat": "\"csv\"",  // csv不能使用omc的api读取结果
		//"numberOfIntervals": "500",
		"numberOfIntervals": task.SRecord.NumberOfIntervals,
		"tolerance":         task.SRecord.Tolerance,
	}
	fileOperation.CreateFilePath(resultFilePath)
	pwd, _ := os.Getwd()
	buildModelRes := omc.OMC.BuildModel(task.SRecord.SimulateModelName, pwd+"/"+resultFilePath, SimulationPraData)
	if buildModelRes {
		MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型编译成功"})
		cmd := exec.Command(resultFilePath + "result")
		out, err := cmd.Output()
		simulateResultStr := string(out)
		if err != nil {
			fmt.Println("err: ", err)
			fmt.Println("仿真执行失败")
		}
		if strings.Index(simulateResultStr, "successfully") != -1 {
			res = true
		} else {
			task.SRecord.SimulateStatus = "3"
		}
		task.SRecord.SimulateResultStr = simulateResultStr
	} else {
		task.SRecord.SimulateStatus = "3"
		task.SRecord.SimulateResultStr = "编译失败"
	}
	config.DB.Save(&task.SRecord)
	return res
}

func ModelSimulate(task *SimulateTask) {
	resultFilePath := "public/UserFiles/ModelResult/" + task.SRecord.Username + "/" + strings.ReplaceAll(task.SRecord.SimulateModelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/"
	task.SRecord.SimulateStartTime = time.Now().Local().Format("2006-01-02 15:04:05")
	task.SRecord.SimulateStart = true
	task.SRecord.SimulateStatus = "2"

	config.DB.Save(&task.SRecord)
	if task.Package.FilePath != "" {
		SaveModelCode(task.Package.PackageName, task.Package.FilePath)
	}
	MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型开始编译"})
	sResult := true
	switch task.SRecord.SimulateType {
	case "OM":
		sResult = openModelica(task, resultFilePath)
	}
	if sResult {
		task.SRecord.SimulateModelResultPath = resultFilePath
		task.SRecord.SimulateStatus = "4"
		MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型仿真完成"})
	} else {
		MessageNotice(map[string]string{"message": task.SRecord.SimulateModelName + " 模型仿真失败"})
	}

	task.SRecord.SimulateEndTime = time.Now().Local().Format("2006-01-02 15:04:05")
	task.SRecord.SimulateStart = false
	config.DB.Save(&task.SRecord)
}
