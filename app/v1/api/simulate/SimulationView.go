package simulate

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"yssim-go/library/timeConvert"

	"github.com/bytedance/sonic"

	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/v1/service"
	"yssim-go/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB = config.DB

func GetSimulationOptionsView(c *gin.Context) {
	/*
		# 仿真参数获取接口
		## model_name: 模型名称，
	*/
	modelName := c.Query("model_name")
	var res DataType.ResponseData
	result := service.GetSimulationOptions(modelName)
	res.Data = result
	c.JSON(http.StatusOK, res)
}

func SetSimulationOptionsView(c *gin.Context) {
	/*
		# 仿真参数设置接口
		## model_name: 模型名称， 全称
		## package_id: 模型所在包的id
		## startTime：仿真开始时间，
		## stopTime：仿真结束时间，
		## tolerance：积分方法使用的容差，
		## numberOfIntervals：间隔数，
		## interval：间隔时间，秒计数
	*/
	var item DataType.SetSimulationOptionsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var res DataType.ResponseData
	result := service.SetSimulationOptions(item.ModelName, item.StartTime, item.StopTime, item.Interval, item.SimulationFlags, item.SimulateType)
	if !result {
		res.Err = "设置失败，请稍后再试"
		res.Status = 2
	}
	res.Msg = "仿真参数设置成功"
	c.JSON(http.StatusOK, res)
}

func GetModelStateView(c *gin.Context) {
	/*
	   ## 1、初始状态, 仿真记录刚刚创建
	   ## 2、仿真进行中
	   ## 3、仿真失败
	   ## 4、仿真完成,也可以看做是未仿真状态
	   ## model_name: 模型名称， 全称
	   ## package_id: 模型所在包的id
	*/

	userSpaceId := c.GetHeader("space_id")
	userName := c.GetHeader("username")
	packageId := c.Query("package_id")
	modelName := c.Query("model_name")
	var modelRecord DataBaseModel.YssimSimulateRecord
	DB.Where("package_id = ? AND username = ? AND simulate_model_name = ? AND simulate_start = ? AND userspace_id = ?", packageId, userName, modelName, true, userSpaceId).First(&modelRecord)

	var res DataType.ResponseData
	stateData := service.GetSimulationState(modelRecord.ID, modelRecord.StartTime, modelRecord.StopTime, modelRecord.Intervals, modelRecord.Percentage)
	stateData["simulate_status_msg"] = config.MoldelSimutalionStatus_[modelRecord.SimulateStatus]

	res.Data = stateData
	c.JSON(http.StatusOK, res)
}

func ModelSimulateView(c *gin.Context) {
	/*
		# 仿真接口，用于模型的仿真计算任务下发到仿真服务
		## package_id: 模型所在包的id,
		## simulate_type: 仿真模型时使用的求解器是哪家的,
		## model_name: 仿真模型的名字,
		## start_time: 仿真参数，仿真的开始时间，单位是整数秒。
		## stop_time: 仿真参数，仿真的结束时间，单位是整数秒。
		## number_of_intervals: 仿真参数， 间隔设置当中的间隔数。 与间隔参数是计算关系，
		## method: 仿真参数， 选择求解方法，默认参数是dassl(Openmodelica使用，dymola使用Dassl)。
	*/

	userSpaceId := c.GetHeader("space_id")
	userName := c.GetHeader("username")
	var item DataType.ModelSimulateData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	itemMap := map[string]string{
		"username":            userName,
		"space_id":            userSpaceId,
		"package_id":          item.PackageId,
		"model_name":          item.ModelName,
		"simulate_type":       item.SimulateType,
		"start_time":          item.StartTime,
		"stop_time":           item.StopTime,
		"tolerance":           item.Tolerance,
		"number_of_intervals": item.NumberOfIntervals,
		"interval":            item.Interval,
		"method":              item.Method,
		"experiment_id":       item.ExperimentId,
	}
	replyId, err := service.GrpcSimulation(itemMap)
	if err != nil {
		fmt.Println("调用(GrpcSimulation)出错：", err)
	}
	var res DataType.ResponseData
	res.Msg = "仿真任务正在准备，请等待仿真完成"
	res.Data = map[string]string{"id": replyId}
	c.JSON(http.StatusOK, res)

}

func SimulateResultGraphicsView(c *gin.Context) {
	/*
		# 仿真结果获取接口， 可一次获取多条, 单个变量
		## variable: 模型变量名字，
		## id: 仿真记录id值，在/simulate/record/list接口获取，
		## s1: 单位转换使用，固定为初始单位
		## s2: 位单位转换使用，需要转换为什么单位
	*/

	var item DataType.ModelSimulateResultData
	userName := c.GetHeader("username")
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res DataType.ResponseData
	var resData []map[string]any

	// 判断记录是否存在，有一条不存在就返回"not found"
	recordIdList := item.RecordId
	var recordList []DataBaseModel.YssimSimulateRecord
	err = DB.Where("id IN ? AND username = ?", recordIdList, userName).Order("").Find(&recordList).Error
	for i := 0; i < len(recordList); i++ {
		if err != nil || recordList[i].SimulateStatus != "4" {
			c.JSON(http.StatusBadRequest, "not found")
			return
		}
	}
	// 判断输入id个数和输出结果长度是否一致!
	if len(recordList) != len(recordIdList) {
		c.JSON(http.StatusBadRequest, "输入id个数和输出结果个数不一致!")
		return
	}
	recordDict := map[string]DataBaseModel.YssimSimulateRecord{}
	for _, record := range recordList {
		recordDict[record.ID] = record
	}
	// 遍历入参数中的id，依次读取结果，每次经过插入到resData
	for i := 0; i < len(recordIdList); i++ {
		var data [][]float64
		var ok bool

		data, ok = service.ReadSimulationResult([]string{item.Variable}, recordDict[recordIdList[i]].SimulateModelResultPath+"result_res.mat")

		unitsData := service.ConvertUnits(item.S2, item.S1)
		if ok {
			ordinate := data[1]
			abscissa := data[0]

			if unitsData[0] == "true" {
				// 单位转换
				scaleFactor, _ := strconv.ParseFloat(unitsData[1], 64)
				offset, _ := strconv.ParseFloat(unitsData[2], 64)
				for p := 0; p < len(ordinate); p++ {
					ordinate[p] = ordinate[p]*scaleFactor + offset
				}
			}
			oneData := map[string]any{
				"id":        recordDict[recordIdList[i]].ID,
				"abscissa":  abscissa,
				"ordinate":  ordinate,
				"startTime": recordDict[recordIdList[i]].StartTime,
				"stopTime":  recordDict[recordIdList[i]].StopTime,
			}
			resData = append(resData, oneData)
			res.Data = resData
		}
	}
	c.JSON(http.StatusOK, res)
}

func SimulateResultSingularView(c *gin.Context) {
	/*
		# 仿真结果获取接口,多条记录，每条记录对应多个不同的变量
		## variable: 模型变量名字，
		## id: 仿真记录id值，在/simulate/record/list接口获取，
		## s1: 单位转换使用，固定为初始单位
		## s2: 位单位转换使用，需要转换为什么单位
	*/

	username := c.GetHeader("username")
	var items []DataType.ModelSimulateResultSingularData
	err := c.BindJSON(&items)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res DataType.ResponseData
	// 遍历获取所有recordId
	var recordIdList []string
	for i := 0; i < len(items); i++ {
		recordIdList = append(recordIdList, items[i].RecordId)
	}
	// 判断记录是否存在，有一条不存在就返回"not found"
	var recordList []DataBaseModel.YssimSimulateRecord
	err = DB.Where("id IN ? AND username = ?", recordIdList, username).Find(&recordList).Error
	//for i := 0; i < len(recordList); i++ {
	//	if err != nil || recordList[i].SimulateStatus != "4" {
	//		c.JSON(http.StatusBadRequest, "not found")
	//		return
	//	}
	//}
	// 构建key为id，val为SimulateModelResultPath的健值对,降低时间复杂度
	recordDict := map[string]DataBaseModel.YssimSimulateRecord{}
	for _, record := range recordList {
		recordDict[record.ID] = record
	}
	// 遍历items，依次获取变量结果
	var resData []map[string]any
	for i := 0; i < len(items); i++ { //遍历items的每条记录，如果与数据库查询结果中的一条能对得上，则读取对应变量结果
		var data [][]float64
		var ok bool
		data, ok = service.ReadSimulationResult([]string{items[i].Variable}, recordDict[items[i].RecordId].SimulateModelResultPath+"result_res.mat")
		unitsData := service.ConvertUnits(items[i].S2, items[i].S1)
		if ok {
			ordinate := data[1]
			abscissa := data[0]
			if unitsData[0] == "true" {
				// 单位转换
				scaleFactor, _ := strconv.ParseFloat(unitsData[1], 64)
				offset, _ := strconv.ParseFloat(unitsData[2], 64)
				for p := 0; p < len(ordinate); p++ {
					ordinate[p] = ordinate[p]*scaleFactor + offset
				}
			}
			oneData := map[string]any{
				"id":        recordDict[items[i].RecordId].ID,
				"variable":  items[i].Variable,
				"abscissa":  abscissa,
				"ordinate":  ordinate,
				"startTime": recordDict[items[i].RecordId].StartTime,
				"stopTime":  recordDict[items[i].RecordId].StopTime,
				"s1":        items[i].S1,
				"s2":        items[i].S2,
			}
			resData = append(resData, oneData)
			res.Data = resData
		}
	}
	c.JSON(http.StatusOK, res)
}

func SimulateResultListView(c *gin.Context) {
	/*
	   # 仿真记录列表获取接口
	   # 模型名为空的时候查所有模型，只有查所有模型的时候才会分页。
	   # 模型名不为空时,不分页，最多返回10条数据。
	   ## return: 返回对应用户的所有仿真记录
	*/

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	modelName := c.Query("model_name")
	packageId := c.Query("package_id")
	pageNumStr := c.Query("page_num") //页码
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	var totle int64 //总条数
	var recordList []DataBaseModel.YssimSimulateRecord
	var experimentList []DataBaseModel.YssimExperimentRecord
	var resData map[string]any
	resData = make(map[string]any)
	var dataList []map[string]any
	if modelName != "" {
		DB.Limit(10).Where("username = ? AND simulate_model_name = ? AND userspace_id = ? AND simulate_status = ?  AND package_id = ?", username, modelName, userSpaceId, "4", packageId).Order("simulate_start_time desc").Find(&recordList)
	} else {
		DB.Where("username = ? AND userspace_id = ?", username, userSpaceId).Find(&recordList).Count(&totle)
		DB.Limit(10).Offset((pageNumInt-1)*10).Where("username = ? AND userspace_id = ?", username, userSpaceId).Order("simulate_start_time desc").Find(&recordList)
	}
	// 获取实验名称
	var experimentIdList []string
	for i := 0; i < len(recordList); i++ {
		experimentIdList = append(experimentIdList, recordList[i].ExperimentId)
	}
	DB.Where("id IN ?", experimentIdList).Find(&experimentList)
	var experimentMap map[string]string
	experimentMap = make(map[string]string)
	for i := 0; i < len(experimentList); i++ {
		experimentMap[experimentList[i].ID] = experimentList[i].ExperimentName
	}

	pageCount := math.Ceil(float64(totle) / 10) //总页数
	for i := 0; i < len(recordList); i++ {
		simulateStartTime := time.Unix(recordList[i].SimulateStartTime, 0)
		simulateStartTimeStr := simulateStartTime.Format("2006-01-02 15:04:05")
		simulateEndTime := time.Unix(recordList[i].SimulateEndTime, 0)
		simulateRunTime := timeConvert.UseTimeFormat(int(simulateStartTime.Unix()), int(simulateEndTime.Unix()))
		if recordList[i].SimulateStartTime == 0 {
			simulateRunTime = "-"
			simulateStartTimeStr = "-"
		}
		experimentName := ""
		if value, exists := experimentMap[recordList[i].ExperimentId]; exists {
			experimentName = value
		}
		data := map[string]any{
			"index":               i + 1,
			"id":                  recordList[i].ID,
			"create_time":         recordList[i].CreatedAt.Format("2006-01-02 15:04:05"),
			"simulate_status":     config.MoldelSimutalionStatus[recordList[i].SimulateStatus],
			"simulate_start_time": simulateStartTimeStr,
			"simulate_end_time":   simulateEndTime.Format("2006-01-02 15:04:05"),
			"simulate_model_name": recordList[i].SimulateModelName,
			"simulate_run_time":   simulateRunTime,
			"another_name":        recordList[i].AnotherName,
			"simulate_percentage": recordList[i].Percentage,
			"experiment_name":     experimentName,
			"pipe_net":            recordList[i].PipeNet,
		}
		dataList = append(dataList, data)
	}
	resData["resultList"] = dataList
	resData["pageCount"] = pageCount
	resData["totle"] = totle
	var res DataType.ResponseData
	res.Data = resData
	c.JSON(http.StatusOK, res)

}

func SimulateResultDetailsView(c *gin.Context) {
	/*
	   # 仿真记录相关实验参数获取接口
	   ## return: 返回对应用户的所有仿真记录
	*/

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	id := c.Query("id")
	var simulateRecord DataBaseModel.YssimSimulateRecord
	DB.Where("id = ? AND username = ? AND userspace_id = ? AND simulate_status = ?", id, username, userSpaceId, "4").First(&simulateRecord)
	var experimentRecord DataBaseModel.YssimExperimentRecord
	DB.Where("id = ? AND username = ? AND userspace_id = ?", simulateRecord.ExperimentId, username, userSpaceId).First(&experimentRecord)
	data := map[string]any{"start_time": "", "stop_time": "", "step_size": "", "tolerance": "", "solver": "", "method": "", "number_intervals": "", "model_var_data": ""}
	data["start_time"] = simulateRecord.StartTime                 // 开始时间
	data["stop_time"] = simulateRecord.StopTime                   // 结束时间
	data["step_size"] = experimentRecord.Interval                 // 步长
	data["tolerance"] = experimentRecord.Tolerance                // 容差
	data["solver"] = config.Solver[experimentRecord.SimulateType] // 求解器
	data["method"] = experimentRecord.Method                      // 方法
	data["number_intervals"] = experimentRecord.NumberOfIntervals // 间隔
	data["model_var_data"] = experimentRecord.ModelVarData        // 模型组件相关参数属性

	var res DataType.ResponseData
	res.Data = data

	c.JSON(http.StatusOK, res)
}

func SimulateResultTreeView(c *gin.Context) {
	/*
	  # 仿真结果树接口， root节点只需要id， 其他子节点需要传变量名字
	    ## id: 仿真记录id
	    ## variable_name: 模型变量名称
	*/

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	recordId := c.Query("record_id")
	parentNode := c.Query("parent_node")
	keyWords := c.Query("key_words")
	var record DataBaseModel.YssimSimulateRecord
	DB.Where("username = ? AND userspace_id = ? AND ID = ? AND simulate_status = ?", username, userSpaceId, recordId, "4").First(&record)
	if record.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var result []map[string]any

	var res DataType.ResponseData
	if record.SimulateModelResultPath != "" && record.SimulateStart == false {
		if record.SimulateType == "FmPy" {
			//FmPy的结果树用的xml是用omc的DumpXMLDAE方法生成的xml，入参record.SimulateModelName用于输出指定模型的xml文件
			result = service.FmpySimulationResultTree(record.SimulateModelName, record.SimulateModelResultPath+"result_init_fmpy.xml", parentNode, keyWords)
		} else if record.SimulateType == "DM" {
			//DM生成的fmu解压后的xml文件
			result = service.DymolaSimulationResultTree(record.SimulateModelResultPath+"result_init.xml", parentNode, keyWords)
		} else {
			//OMC仿真完输出的xml文件
			result = service.SimulationResultTree(record.SimulateModelResultPath, parentNode, keyWords)
		}
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

	// OMC仿真顶层结果树中增加cpuTime数据
	if record.SimulateType == "OM" && parentNode == "" && strings.Contains("cputime", strings.ToLower(keyWords)) {
		result = append(result, service.SetCpuTimeResultTree())
	}

	res.Data = result
	c.JSON(http.StatusOK, res)
}

func SimulateResultDeleteView(c *gin.Context) {
	/*
	   # 2022.11.2 徐庆达修改（新接口）：删除仿真结果在数据库中的记录
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	recordId := c.Query("record_id")
	var resultRecord DataBaseModel.YssimSimulateRecord
	DB.Where("id = ? AND username = ? AND userspace_id = ? ", recordId, username, userSpaceId).First(&resultRecord)
	var res DataType.ResponseData

	resultRecord.SimulateStatus = "5"
	config.DB.Save(&resultRecord)
	service.DeleteSimulateTask(recordId, resultRecord.SimulateType, resultRecord.SimulateModelResultPath)
	config.DB.Delete(&resultRecord)
	DB.Delete(&DataBaseModel.YssimSnapshots{}, "simulate_result_id = ?", recordId) //删除相关的快照
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func SimulateResultRenameView(c *gin.Context) {
	/*
	   # 2023.04.11 徐庆达修改（新接口）：修改仿真结果的another_name（别名)
	*/
	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	var item DataType.RecordRenameData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res DataType.ResponseData
	err = DB.Model(&DataBaseModel.YssimSimulateRecord{}).Where("id = ?", item.RecordId).Update("another_name", item.NewAnotherName).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "修改失败")
		return
	}
	res.Msg = "修改成功"
	c.JSON(http.StatusOK, res)
}

func SimulateTerminateView(c *gin.Context) {
	/*
	   # 2024.04.29 周强修改（新接口）：终止仿真进程
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.SimulateTerminateData
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res DataType.ResponseData
	var resultRecord DataBaseModel.YssimSimulateRecord
	if err := DB.Where("id = ? AND username = ? AND userspace_id = ? ", item.RecordId, username, userSpaceId).First(&resultRecord).Error; err != nil {
		res.Err = "终止仿真失败"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}

	if err := service.TerminateSimulateTask(item.RecordId, resultRecord.SimulateType); err != nil {
		res.Err = "终止仿真失败"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}

	resultRecord.SimulateStatus = "7"
	resultRecord.SimulateStart = false
	if err := config.DB.Save(&resultRecord).Error; err != nil {
		res.Err = "仿真已终止，但仿真状态更新失败"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "仿真已终止"
	c.JSON(http.StatusOK, res)
}

func ExperimentExistsView(c *gin.Context) {
	/*
	   # 判断仿真实验记录是否存在接口，
	   ## package_id: 保存的实验是属于哪个包id
	   ## model_var_data: 模型的变量数据，修改过哪个模型变量，保存到当前数组对象
	   ## simulate_var_data: 模型仿真选项数据
	*/
	var res DataType.ResponseData
	var item DataType.ExperimentExistsData
	err := c.BindJSON(&item)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")

	// 将请求数据的ModelVarData转化为slice
	var requestModelParamList interface{}
	if err = json.Unmarshal([]byte(item.ModelVarData), &requestModelParamList); err != nil {
		res.Err = "实验参数对比失败，请稍后再试"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}

	// 查询当前用户空间的package id所有的experiment记录
	var experimentRecordList []DataBaseModel.YssimExperimentRecord
	DB.Where("package_id = ? AND username =? AND userspace_id =? AND model_name =?",
		item.PackageId, username, userSpaceId, item.ModelName).Find(&experimentRecordList)

	// 遍历每一条数据库记录，把每一条数据库记录和当前请求数据对比
	for _, experimentRecord := range experimentRecordList {
		// 将每一个数据库记录转化为ExperimentExistsData类型数据，方便比较数据
		var dbRecord DataType.ExperimentExistsData
		simulateVarData := map[string]string{}
		simulateVarData["startTime"] = experimentRecord.StartTime
		simulateVarData["stopTime"] = experimentRecord.StopTime
		simulateVarData["tolerance"] = experimentRecord.Tolerance
		simulateVarData["numberOfIntervals"] = experimentRecord.NumberOfIntervals
		simulateVarData["interval"] = experimentRecord.Interval
		simulateVarData["simulate_type"] = experimentRecord.SimulateType
		simulateVarData["method"] = experimentRecord.Method

		dbRecord.SimulateVarData = simulateVarData
		dbRecord.ModelVarData = experimentRecord.ModelVarData

		// 将每一条数据库记录的ModelVarData转化为slice
		var dbRecordModelParamList interface{}
		if err := json.Unmarshal([]byte(dbRecord.ModelVarData), &dbRecordModelParamList); err != nil {
			res.Err = "实验参数对比失败，请稍后再试"
			res.Status = 1
			c.JSON(http.StatusOK, res)
			return
		}

		// 比较请求参数和本条数据库记录参数的SimulateVarData和ModelVarData
		if reflect.DeepEqual(item.SimulateVarData, dbRecord.SimulateVarData) && reflect.DeepEqual(requestModelParamList, dbRecordModelParamList) {
			data := map[string]string{"experiment_id": experimentRecord.ID, "name": experimentRecord.ExperimentName}
			res.Data = data
			c.JSON(http.StatusOK, res)
			return
		}
	}

	c.JSON(http.StatusOK, res)
}

func ExperimentCreateView(c *gin.Context) {
	/*
	   # 仿真实验创建记录接口，
	   ## package_id: 保存的实验是属于哪个包id
	   ## model_name: 实验属于哪个模型，全称，例如"Modelica.Blocks.Examples.PID_Controller"
	   ## model_var_data: 模型的变量数据，修改过哪个模型变量，保存到当前数组对象
	   ## simulate_var_data: 模型仿真选项数据
	   ## experiment_name: 实验名称
	*/
	var res DataType.ResponseData
	var item DataType.ExperimentCreateData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var models DataBaseModel.YssimModels
	err = DB.Where("id = ?", item.PackageId).First(&models).Error
	if models.Encryption {
		res.Err = "该模型不支持创建实验"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 名称判重
	var record DataBaseModel.YssimExperimentRecord
	DB.Where("package_id = ? AND experiment_name = ? AND username =? AND userspace_id =? AND model_name =?", item.PackageId, item.ExperimentName, username, userSpaceId, item.ModelName).First(&record)
	if record.ExperimentName != "" || item.ExperimentName == "实验(默认)" {
		res.Err = "实验记录名称已存在，请更换。"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	experimentRecord := DataBaseModel.YssimExperimentRecord{
		ID:                uuid.New().String(),
		PackageId:         item.PackageId,
		UserspaceId:       userSpaceId,
		UserName:          username,
		ExperimentName:    item.ExperimentName,
		ModelName:         item.ModelName,
		StartTime:         item.SimulateVarData["startTime"],
		StopTime:          item.SimulateVarData["stopTime"],
		Method:            item.SimulateVarData["method"],
		SimulateType:      item.SimulateVarData["simulate_type"],
		NumberOfIntervals: item.SimulateVarData["numberOfIntervals"],
		Tolerance:         item.SimulateVarData["tolerance"],
		Interval:          item.SimulateVarData["interval"],
		ModelVarData:      item.ModelVarData,
		IsFullModelVar:    item.IsFullModelVar,
	}
	err = DB.Create(&experimentRecord).Error
	if err != nil {
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	// 超过5个实验时会自动删除最早创建的实验
	var recordList []DataBaseModel.YssimExperimentRecord
	DB.Where("package_id = ? AND username =? AND userspace_id =? AND model_name =?", item.PackageId, username, userSpaceId, item.ModelName).Order("create_time desc").Find(&recordList)
	if len(recordList) > 5 {
		DB.Delete(&recordList[len(recordList)-1])
	}
	item.SimulateVarData["id"] = experimentRecord.ID
	// item.SimulateVarData["ModelVarData"] = item.ModelVarData
	res.Data = item.SimulateVarData
	res.Msg = "实验记录创建成功"
	c.JSON(http.StatusOK, res)
}

func ExperimentDeleteView(c *gin.Context) {
	/*
	   # 仿真实验记录删除接口，
	   ## id: 实验id，此接口其他值无须传入
	*/

	var res DataType.ResponseData
	var item DataType.ExperimentDeleteData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var record DataBaseModel.YssimExperimentRecord
	DB.Where("username =? AND userspace_id =? AND id =?", username, userSpaceId, item.ExperimentId).First(&record)
	DB.Delete(&record)
	//删除相关的快照
	DB.Delete(&DataBaseModel.YssimSnapshots{}, "experiment_id = ?", item.ExperimentId)

	// 删除相关的仿真记录
	var resultRecord []DataBaseModel.YssimSimulateRecord
	DB.Where("experiment_id = ? ", item.ExperimentId).Find(&resultRecord)
	for i := 0; i < len(resultRecord); i++ {
		resultRecord[i].SimulateStatus = "5"
		config.DB.Save(&resultRecord[i])
		service.DeleteSimulateTask(resultRecord[i].ID, resultRecord[i].SimulateType, resultRecord[i].SimulateModelResultPath)
		config.DB.Delete(&resultRecord[i])
	}

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func ExperimentEditView(c *gin.Context) {
	/*
	   # 仿真实验记录编辑接口，
	   ## experiment_id: 实验id
	   ## model_var_data: 模型的变量数据，修改过哪个模型变量，保存到当前数组对象
	   ## simulate_var_data: 模型仿真选项数据
	   ## experiment_name: 实验名称
	*/

	var res DataType.ResponseData
	var item DataType.ExperimentEditData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var recordName DataBaseModel.YssimExperimentRecord
	DB.Where("id != ? AND username =? AND userspace_id =? AND experiment_name =? AND model_name =? AND package_id =?", item.ExperimentId, username, userSpaceId, item.ExperimentName, item.ModelName, item.PackageId).First(&recordName)
	if recordName.ExperimentName != "" {
		res.Msg = "实验记录名称已存在，请更换。"
		c.JSON(http.StatusOK, res)
		return
	}
	var editRecord DataBaseModel.YssimExperimentRecord
	err = DB.Model(&editRecord).Where("id = ? AND username =? AND userspace_id =?", item.ExperimentId, username, userSpaceId).Updates(DataBaseModel.YssimExperimentRecord{
		ExperimentName:    item.ExperimentName,
		StartTime:         item.SimulateVarData["startTime"],
		StopTime:          item.SimulateVarData["stopTime"],
		Method:            item.SimulateVarData["method"],
		SimulateType:      item.SimulateVarData["simulate_type"],
		NumberOfIntervals: item.SimulateVarData["numberOfIntervals"],
		Tolerance:         item.SimulateVarData["tolerance"],
		Interval:          item.SimulateVarData["interval"],
		ModelVarData:      item.ModelVarData,
	}).Error
	if err != nil {
		res.Err = "更新失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "实验记录已更新"
	c.JSON(http.StatusOK, res)
}

func ExperimentGetView(c *gin.Context) {
	/*
	   # 获取仿真实验列表接口，
	   ## package_id: 获取的是哪个包当中的实验列表
	   ## model_name： 哪个模型当中的实验列表，全称，例如："Modelica.Blocks.Examples.PID_Controller"
	*/

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	packageId := c.Query("package_id")
	modelName := c.Query("model_name")
	var recordList []DataBaseModel.YssimExperimentRecord
	DB.Where("package_id = ? AND userspace_id = ? AND username = ? AND model_name = ?", packageId, userSpaceId, username, modelName).Order("create_time desc").Find(&recordList)

	var dataList []map[string]string
	for _, record := range recordList {
		data := map[string]string{
			"id":                record.ID,
			"experiment_name":   record.ExperimentName,
			"interval":          record.Interval,
			"method":            record.Method,
			"numberOfIntervals": record.NumberOfIntervals,
			"simulate_type":     record.SimulateType,
			"startTime":         record.StartTime,
			"stopTime":          record.StopTime,
			"tolerance":         record.Tolerance,
		}
		dataList = append(dataList, data)
	}
	var res DataType.ResponseData
	res.Data = dataList
	c.JSON(http.StatusOK, res)
}

func ExperimentParametersView(c *gin.Context) {
	/*
	   # 获取仿真实验中的模型参数接口，
	   ## experiment_id: 实验id
	*/
	experimentId := c.Query("experiment_id")
	var record DataBaseModel.YssimExperimentRecord
	DB.Where("id =?", experimentId).First(&record)

	var res DataType.ResponseData
	res.Data = record.ModelVarData
	c.JSON(http.StatusOK, res)
}

func ExperimentNameEditView(c *gin.Context) {
	/*
	   # 仿真实验记录名称编辑接口，
	   ## experiment_id: 实验id
	   ## new_experiment_name: 新的实验名称
	*/

	var res DataType.ResponseData
	var item DataType.ExperimentNameEditData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")

	// 判断用户传入的experiment是否存在
	var recordName DataBaseModel.YssimExperimentRecord
	DB.Where("id =? AND package_id =? AND username =? AND userspace_id =?",
		item.ExperimentId, item.PackageId, username, userSpaceId).First(&recordName)
	if recordName.ID == "" {
		res.Err = "实验记录不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 判断新的experiment name是否重复
	var existingName DataBaseModel.YssimExperimentRecord
	DB.Where("id != ? AND username =? AND userspace_id =? AND experiment_name =? AND package_id =? AND model_name =?",
		item.ExperimentId, username, userSpaceId, item.NewExperimentName, item.PackageId, item.ModelName).First(&existingName)
	if existingName.ExperimentName != "" || item.NewExperimentName == "实验(默认)" {
		res.Err = "实验记录名称已存在，请更换"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 查询experiment对应的simulation record
	var simulateRecord DataBaseModel.YssimSimulateRecord
	DB.Where("experiment_id = ? AND username =? AND userspace_id =?", item.ExperimentId, username, userSpaceId).First(&simulateRecord)

	// 修改名称，开启事务
	tx := DB.Begin()

	var newExperimentRecordName DataBaseModel.YssimExperimentRecord
	var newSimulationRecordName DataBaseModel.YssimSimulateRecord

	// 修改实验名称
	txResult := tx.Model(&newExperimentRecordName).Where("id = ? AND username =? AND userspace_id =?", item.ExperimentId, username, userSpaceId).Updates(DataBaseModel.YssimExperimentRecord{
		ExperimentName: item.NewExperimentName,
	})
	if txResult.Error != nil {
		tx.Rollback()
		res.Err = "实验名称更新失败，请稍后再试"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}

	//若实验存在仿真结果，同步修改仿真结果名称
	if simulateRecord.ID != "" {
		txResult = tx.Model(&newSimulationRecordName).Where("experiment_id =? AND username =? AND userspace_id =?", item.ExperimentId, username, userSpaceId).Updates(DataBaseModel.YssimSimulateRecord{
			AnotherName: item.NewExperimentName + "的结果",
		})
		if txResult.Error != nil {
			tx.Rollback()
			res.Err = "实验名称更新失败，请稍后再试"
			res.Status = 1
			c.JSON(http.StatusOK, res)
			return
		}
	}

	tx.Commit()

	res.Msg = "实验名称已更新"
	c.JSON(http.StatusOK, res)
}

func ExperimentCompareView(c *gin.Context) {
	/*
	   # 对比不同实验记录的参数差异，
	   ## package_id: 实验是属于哪个包id
	   ## experiment_id_list: 要对比的实验id列表
	*/
	var res DataType.ResponseData
	var item DataType.ExperimentCompareData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")

	// 查询数据库中当前请求要对比的所有的experiment记录
	var experimentRecordList []DataBaseModel.YssimExperimentRecord
	if err = DB.Where("package_id = ? AND username =? AND userspace_id =? AND id IN ?",
		item.PackageId, username, userSpaceId, item.ExperimentIdList).Find(&experimentRecordList).Error; err != nil {
		res.Err = "实验参数对比失败，请稍后再试"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}
	if len(experimentRecordList) < 2 {
		res.Err = "请求参数错误，请重试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 获取当前模型的名称
	modelName := experimentRecordList[0].ModelName

	// 存储实验id和实验记录之间的对应关系
	experimentMap := map[string]DataBaseModel.YssimExperimentRecord{}
	for _, experimentRecord := range experimentRecordList {
		experimentMap[experimentRecord.ID] = experimentRecord
	}

	// 遍历数据库实验记录, 将组件参数ModelVarData(json类型)转化为slice
	experimentRecordMap := map[string][]map[string]any{}
	for _, experimentRecord := range experimentRecordList {
		var componentParams []map[string]any
		if err := json.Unmarshal([]byte(experimentRecord.ModelVarData), &componentParams); err != nil {
			res.Err = "实验参数对比失败，请稍后再试"
			res.Status = 1
			c.JSON(http.StatusOK, res)
			return
		}
		experimentRecordMap[experimentRecord.ID] = componentParams
	}

	// 获取所有组件名称和组件类型
	components := make(map[string]map[string]bool)
	componentClasses := make(map[string]string)
	for experimentId := range experimentRecordMap {
		length := len(experimentRecordMap[experimentId])
		for index := 0; index < length; {
			// 如果实验参数中某组件不包含parameters字段，则删除该组件数据，后续再从omc中获取
			parameters, ok := experimentRecordMap[experimentId][index]["parameters"].([]any)
			if !ok {
				experimentRecordMap[experimentId] = append(experimentRecordMap[experimentId][:index], experimentRecordMap[experimentId][index+1:]...)
				length = len(experimentRecordMap[experimentId])
				continue
			}
			if _, ok := components[experimentRecordMap[experimentId][index]["name"].(string)]; !ok {
				components[experimentRecordMap[experimentId][index]["name"].(string)] = map[string]bool{}
			}
			components[experimentRecordMap[experimentId][index]["name"].(string)][experimentId] = true
			componentClasses[experimentRecordMap[experimentId][index]["name"].(string)] = parameters[0].(map[string]any)["extend_name"].(string)
			index += 1
		}
	}

	// 与omc交互补全任一实验中缺失的组件及参数
	for experimentId := range experimentRecordMap {
		for component := range components {
			if _, ok := components[component][experimentId]; !ok {
				var parameterOMC []any
				if component == modelName {
					parameterOMC = service.GetModelParameters(modelName, "", componentClasses[component], "")
				} else {
					parameterOMC = service.GetModelParameters(modelName, component, componentClasses[component], "")
				}
				parameter := map[string]any{"name": component, "parameters": parameterOMC}
				experimentRecordMap[experimentId] = append(experimentRecordMap[experimentId], parameter)
			}
		}
	}

	// 给前端返回的数据结构
	tableColumns := []map[string]string{}
	for _, experimentId := range item.ExperimentIdList {
		data := map[string]string{"key": experimentId, "name": experimentMap[experimentId].ExperimentName}
		tableColumns = append(tableColumns, data)
	}
	tableData := []map[string]any{}

	// 分解出所有参数
	compareData := map[string]map[string]map[string]any{}
	keys := map[string]bool{}
	for experimentId := range experimentRecordMap {
		if _, ok := compareData[experimentId]; !ok {
			compareData[experimentId] = make(map[string]map[string]any)
		}

		for _, component := range experimentRecordMap[experimentId] {
			if _, ok := compareData[experimentId][component["name"].(string)]; !ok {
				compareData[experimentId][component["name"].(string)] = make(map[string]any)
			}
			for _, parameter := range component["parameters"].([]any) {
				compareData[experimentId][component["name"].(string)][parameter.(map[string]any)["name"].(string)] = parameter.(map[string]any)["value"]
				keys[parameter.(map[string]any)["name"].(string)] = true
			}
		}
	}

	// 依次比较所有的参数，找出不同
	for component := range components {
		singleComponentAllDifferentParameters := map[string]any{}
		for key := range keys {
			isDuplicated := false
			for i := 0; i < len(experimentRecordList)-1; i++ {
				if !reflect.DeepEqual(compareData[experimentRecordList[i].ID][component][key], compareData[experimentRecordList[i+1].ID][component][key]) {
					isDuplicated = true
					singleComponentAllDifferentParameters["name"] = component
				}
			}

			if isDuplicated {
				for j := 0; j < len(experimentRecordList); j++ {
					if _, ok := singleComponentAllDifferentParameters[experimentRecordList[j].ID]; !ok {
						singleComponentAllDifferentParameters[experimentRecordList[j].ID] = map[string]any{}
					}
					singleComponentAllDifferentParameters[experimentRecordList[j].ID].(map[string]any)[key] = compareData[experimentRecordList[j].ID][component][key]
				}
			}
		}
		if len(singleComponentAllDifferentParameters) != 0 {
			tableData = append(tableData, singleComponentAllDifferentParameters)
		}
	}

	// 返回数据
	res.Data = map[string]any{
		"tableColumns": tableColumns,
		"tableData":    tableData,
	}

	c.JSON(http.StatusOK, res)
}

func ExperimentCompareNewView(c *gin.Context) {
	/*
	   # 对比不同实验记录的参数差异，
	   ## package_id: 实验是属于哪个包id
	   ## experiment_id_list: 要对比的实验id列表
	*/
	var res DataType.ResponseData
	var item DataType.SimulateRecordCompareData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")

	// 查询数据库中的仿真结果记录
	var simulateRecordList []DataBaseModel.YssimSimulateRecord
	if err = DB.Where("package_id = ? AND username =? AND userspace_id =? AND id IN ?",
		item.PackageId, username, userSpaceId, item.RecordIdList).Find(&simulateRecordList).Error; err != nil {
		res.Err = "实验参数对比失败，请稍后再试"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}
	if len(simulateRecordList) < 2 {
		res.Err = "请求参数错误，请重试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 查询数据库中仿真结果对应的的experiment记录
	experimentIdList := []string{}
	for _, simulateRecordId := range item.RecordIdList {
		for _, simulateRecord := range simulateRecordList {
			if simulateRecordId == simulateRecord.ID {
				experimentIdList = append(experimentIdList, simulateRecord.ExperimentId)
			}
		}
	}
	var experimentRecordList []DataBaseModel.YssimExperimentRecord
	if err = DB.Where("package_id = ? AND username =? AND userspace_id =? AND id IN ?",
		item.PackageId, username, userSpaceId, experimentIdList).Find(&experimentRecordList).Error; err != nil {
		res.Err = "实验参数对比失败，请稍后再试"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}
	if len(experimentRecordList) < 2 {
		res.Err = "请求参数错误，请重试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 存储实验id和实验记录之间的对应关系
	experimentMap := map[string]DataBaseModel.YssimExperimentRecord{}
	for _, experimentRecord := range experimentRecordList {
		experimentMap[experimentRecord.ID] = experimentRecord
	}

	// 存储实验id和仿真结果记录文件存储路径之间的对应关系
	simulateResultMap := map[string]string{}
	for _, simulateRecord := range simulateRecordList {
		simulateResultMap[simulateRecord.ExperimentId] = simulateRecord.SimulateModelResultPath
	}

	// 遍历数据库实验记录, 将组件参数ModelVarData(json类型)转化为slice
	experimentRecordMap := map[string][]map[string]any{}
	for _, experimentRecord := range experimentRecordList {
		var componentParams []map[string]any
		if err := json.Unmarshal([]byte(experimentRecord.ModelVarData), &componentParams); err != nil {
			res.Err = "实验参数对比失败，请稍后再试"
			res.Status = 1
			c.JSON(http.StatusOK, res)
			return
		}
		experimentRecordMap[experimentRecord.ID] = componentParams
	}

	// 获取所有组件名称和组件类型
	components := make(map[string]map[string]bool)
	componentParameters := make(map[string][]string)
	for experimentId := range experimentRecordMap {
		length := len(experimentRecordMap[experimentId])
		for index := 0; index < length; {
			// 如果实验参数中某组件不包含parameters字段，则删除该组件数据，后续再从xml中获取
			name, _ := experimentRecordMap[experimentId][index]["name"].(string)
			parameters, ok := experimentRecordMap[experimentId][index]["parameters"].([]any)
			if !ok {
				experimentRecordMap[experimentId] = append(experimentRecordMap[experimentId][:index], experimentRecordMap[experimentId][index+1:]...)
				length = len(experimentRecordMap[experimentId])
				continue
			}

			// 获取数据库中组件的参数列表
			componentParameters[name] = make([]string, 0)
			for _, k := range parameters {
				key, _ := k.(map[string]any)
				componentParameters[name] = append(componentParameters[name], key["name"].(string))
			}

			//
			if _, ok := components[experimentRecordMap[experimentId][index]["name"].(string)]; !ok {
				components[experimentRecordMap[experimentId][index]["name"].(string)] = map[string]bool{}
			}
			components[experimentRecordMap[experimentId][index]["name"].(string)][experimentId] = true
			index += 1
		}
	}

	// 与仿真结果文件交互获取每一个实验中每个组件的参数数据
	parametersXmlMap := map[string][]map[string]any{}
	for experimentId := range experimentRecordMap {
		for component := range components {
			parameterXml := service.SimulateResultParameters(simulateResultMap[experimentId], component, "")
			parameter := map[string]any{"name": component, "parameters": parameterXml}
			if _, ok := parametersXmlMap[experimentId]; !ok {
				parametersXmlMap[experimentId] = make([]map[string]any, 0)
			}
			parametersXmlMap[experimentId] = append(parametersXmlMap[experimentId], parameter)
		}
	}

	// 给前端返回的数据结构
	tableColumns := []map[string]string{}
	for _, experimentId := range experimentIdList {
		data := map[string]string{"key": experimentId, "name": experimentMap[experimentId].ExperimentName}
		tableColumns = append(tableColumns, data)
	}
	tableData := []map[string]any{}

	// 分解出所有参数
	compareData := map[string]map[string]map[string]any{}
	keys := map[string]bool{}
	for experimentId := range parametersXmlMap {
		if _, ok := compareData[experimentId]; !ok {
			compareData[experimentId] = make(map[string]map[string]any)
		}

		for _, component := range parametersXmlMap[experimentId] {
			if _, ok := compareData[experimentId][component["name"].(string)]; !ok {
				compareData[experimentId][component["name"].(string)] = make(map[string]any)
			}
			for _, parameter := range component["parameters"].([]any) {
				compareData[experimentId][component["name"].(string)][parameter.(map[string]any)["name"].(string)] = parameter.(map[string]any)["value"]
				keys[parameter.(map[string]any)["name"].(string)] = true
			}
		}
	}

	// 依次比较所有的参数，找出不同
	for component := range components {
		singleComponentAllDifferentParameters := map[string]any{}
		for _, key := range componentParameters[component] {
			isDuplicated := false
			for i := 0; i < len(experimentRecordList)-1; i++ {
				if !reflect.DeepEqual(compareData[experimentRecordList[i].ID][component][key], compareData[experimentRecordList[i+1].ID][component][key]) {
					isDuplicated = true
					singleComponentAllDifferentParameters["name"] = component
				}
			}

			if isDuplicated {
				for j := 0; j < len(experimentRecordList); j++ {
					if _, ok := singleComponentAllDifferentParameters[experimentRecordList[j].ID]; !ok {
						singleComponentAllDifferentParameters[experimentRecordList[j].ID] = map[string]any{}
					}
					singleComponentAllDifferentParameters[experimentRecordList[j].ID].(map[string]any)[key] = compareData[experimentRecordList[j].ID][component][key]
				}
			}
		}
		if len(singleComponentAllDifferentParameters) != 0 {
			tableData = append(tableData, singleComponentAllDifferentParameters)
		}
	}

	// 返回数据
	res.Data = map[string]any{
		"tableColumns": tableColumns,
		"tableData":    tableData,
	}

	c.JSON(http.StatusOK, res)
}

func CreateSnapshotView(c *gin.Context) {
	/*
		#xqd#创建视图(快照)接口
	*/
	var res DataType.ResponseData
	var item DataType.SnapshotCreatData
	var snapshotRecord DataBaseModel.YssimSnapshots
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	DB.Where("snapshot_name = ? AND username =? AND space_id =? AND model_name =?", item.SnapshotName, username, userSpaceId, item.ModelName).First(&snapshotRecord)

	if snapshotRecord.SnapshotName != "" {
		res.Msg = "视图名称已存在，请更换。"
		c.JSON(http.StatusOK, res)
		return
	}

	snapshot := DataBaseModel.YssimSnapshots{
		ID:                uuid.New().String(),
		SpaceId:           userSpaceId,
		UserName:          username,
		SnapshotName:      item.SnapshotName,
		ModelName:         item.ModelName,
		PackageId:         item.PackageId,
		ComponentName:     item.ComponentName,
		ModelVarData:      item.ModelVarData,
		ExperimentId:      item.ExperimentId,
		SimulateVarData:   item.SimulateVarData,
		SimulateResultId:  item.SimulateResultId,
		SimulateResultObj: item.SimulateResultObj,
	}
	err = DB.Create(&snapshot).Error
	if err != nil {
		log.Println("DB Create err:", err)
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	//item.SimulateVarData["id"] = experimentRecord.ID
	// item.SimulateVarData["ModelVarData"] = item.ModelVarData
	res.Data = snapshot.ID
	res.Msg = "视图创建成功"
	c.JSON(http.StatusOK, res)

}

func DeleteSnapshotView(c *gin.Context) {
	/*
	   #xqd# 删除视图接口
	*/
	var res DataType.ResponseData
	var item DataType.SnapshotDeleteData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var record DataBaseModel.YssimSnapshots
	DB.Where("id =? AND space_id = ? AND username = ?", item.SnapshotId, userSpaceId, username).First(&record)
	DB.Delete(&record)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func EditSnapshotView(c *gin.Context) {
	/*
	   #xqd# 修改视图接口
	   #grom bug记录：snapshotEditData的字段数和名称必须和数据库模型YssimSnapshots一致。
	*/
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.SnapshotEditData

	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//判断名称是否存在
	var recordName DataBaseModel.YssimSnapshots
	DB.Where("id != ? AND snapshot_name =? AND username =? AND space_id =? AND model_name =? AND package_id=?", item.ID, item.SnapshotName, username, userSpaceId, item.ModelName, item.PackageId).First(&recordName)
	if recordName.SnapshotName != "" {
		res.Msg = "视图记录名称已存在，请更换。"
		c.JSON(http.StatusOK, res)
		return
	}
	//更新数据库
	var editRecord DataBaseModel.YssimSnapshots
	result := DB.Model(&editRecord).Omit("ID", "SpaceId", "UserName", "ModelName", "PackageId").Where("id = ? ", item.ID).Updates(item).Error

	if result != nil {
		fmt.Println("DB Model Updates err", result)
		res.Err = "更新失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "视图记录已更新"
	c.JSON(http.StatusOK, res)

}

func SnapshotGetListView(c *gin.Context) {
	/*
	   #xqd# 获取视图列表接口
	*/

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	modelName := c.Query("model_name")
	packageId := c.Query("package_id")

	var snapshotList []DataBaseModel.YssimSnapshots
	DB.Where("space_id = ? AND username = ? AND model_name = ? AND package_id = ?", userSpaceId, username, modelName, packageId).Find(&snapshotList)
	var dataList []map[string]any
	for _, record := range snapshotList {
		data := make(map[string]any)
		data["id"] = record.ID
		data["snapshot_name"] = record.SnapshotName
		data["updated_time"] = record.UpdatedAt.Format("2006-01-02 15:04:05") // .Format("2006-01-02 15:04:05")
		data["component_name"] = record.ComponentName
		data["modelVar_data"] = record.ModelVarData
		data["experiment_id"] = record.ExperimentId
		data["simulateVar_data"] = record.SimulateVarData
		data["simulateResult_id"] = record.SimulateResultId
		data["simulateResult_obj"] = record.SimulateResultObj
		dataList = append(dataList, data)
	}

	var res DataType.ResponseData
	res.Data = dataList
	c.JSON(http.StatusOK, res)
}

func CalibrationCompileView(c *gin.Context) {
	/*
	   # 模型参数标定功能编译模型，验证模型是否可用
	*/

	var res DataType.ResponseData
	var item DataType.CalibrationCompileData
	userName := c.GetHeader("username")
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, userName, item.UserSpaceId).First(&modelPackage)

	var record DataBaseModel.ParameterCalibrationRecord
	err = DB.Where("id = ? AND package_id = ? AND username = ?", item.ID, item.PackageId, userName).First(&record).Error
	if modelPackage.ID == "" || record.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	record.CompilePath = "static/ParameterCalibration/" + userName + "/" + modelPackage.PackageName + "/" + time.Now().Local().Format("20060102150405")

	record.PackagePath, err = service.CopyPackage(modelPackage.FilePath, record.CompilePath+"/"+item.ModelName)
	packageSource := service.GetPackagesSource()
	packageSource[modelPackage.PackageName] = record.PackagePath
	record.CompileDependencies, _ = sonic.Marshal(packageSource)
	if err != nil {
		log.Println("标定参数预编译copy模型失败：", err)
		res.Err = "系统错误请联系管理员处理"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	DB.Save(&record)
	itemMap := map[string]string{
		"id":               record.ID,
		"user_space_id":    modelPackage.UserSpaceId,
		"username":         userName,
		"package_id":       item.PackageId,
		"package_name":     modelPackage.PackageName,
		"model_name":       record.ModelName,
		"result_file_path": record.CompilePath,
	}
	err = service.GrpcCalibrationCompile(itemMap, packageSource)
	if err != nil {
		fmt.Println("调用(GrpcSimulation)出错：", err)
		res.Err = "编译失败，请联系管理员"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]any{"id": record.ID}
	res.Msg = "请等待编译完成"
	c.JSON(http.StatusOK, res)
}

func CalibrationSimulateTaskAddView(c *gin.Context) {
	/*
	   # 模型参数标定功能仿真接口，生成仿真数据
	*/

	var res DataType.ResponseData
	var item DataType.CalibrationSimulateData
	userName := c.GetHeader("username")
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var record DataBaseModel.ParameterCalibrationRecord
	err = DB.Where("id = ?  AND username = ?", item.ID, userName).First(&record).Error
	if record.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	switch record.SimulateStatus {
	case "1":
		res.Msg = "已经在排队仿真当中，请耐心等待"
		res.Status = 0
		c.JSON(http.StatusOK, res)
		return
	case "2":
		res.Msg = "已经在仿真进行中，请耐心等待"
		res.Status = 0
		c.JSON(http.StatusOK, res)
		return
	}

	itemMap := map[string]string{
		"id": record.ID,
	}
	var percentageMap = map[string]any{"0": 1}
	var conditionParameters []map[string]any
	_ = sonic.Unmarshal(record.ConditionParameters, &conditionParameters)
	percentage, _ := sonic.Marshal(percentageMap)
	err = DB.Model(DataBaseModel.ParameterCalibrationRecord{}).Where("id = ?  AND username = ?", item.ID, userName).Updates(map[string]any{"percentage": percentage, "simulate_status": "1"}).Error
	if err != nil {
		log.Println("标定任务更新数据库出错：", err)
	}
	err = service.GrpcCalibrationSimulate(itemMap)
	if err != nil {
		log.Println("调用(GrpcSimulation)出错：", err)
		res.Err = "仿真失败，请联系管理员"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]any{"id": record.ID}
	res.Msg = "请等待仿真完成"
	c.JSON(http.StatusOK, res)
}

func CalibrationSimulateTaskStopView(c *gin.Context) {
	/*
	   # 模型参数标定功能仿真终止接口，终止仿真任务
	*/

	var res DataType.ResponseData
	var item DataType.CalibrationSimulateData

	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	userName := c.GetHeader("username")
	var record DataBaseModel.ParameterCalibrationRecord
	err = DB.Where("id = ?  AND username = ?", item.ID, userName).First(&record).Error
	if record.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	err = service.DeleteCalculationSimulateTask(record.ID)
	if err != nil {
		fmt.Println("调用(GrpcSimulation)出错：", err)
		res.Err = "仿真失败，请联系管理员"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]any{"id": record.ID}
	res.Msg = "仿真任务已终止"
	c.JSON(http.StatusOK, res)
}

func GetCalibrationTaskStatusView(c *gin.Context) {
	/*
	   # 模型参数标定功能仿真状态获取
	*/

	var res DataType.ResponseData
	userName := c.GetHeader("username")
	recordId := c.Query("id")
	var record DataBaseModel.ParameterCalibrationRecord
	err := DB.Where("id = ?  AND username = ?", recordId, userName).First(&record).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	percentageMap := map[string]any{}
	_ = sonic.Unmarshal(record.Percentage, &percentageMap)
	percentageList := []any{}
	for i := 0; i >= 0; i++ {
		index := strconv.Itoa(i)
		percentage, ok := percentageMap[index]
		if ok {
			percentageList = append(percentageList, percentage)
			delete(percentageMap, index)
		}
		if len(percentageMap) == 0 {
			break
		}
	}
	res.Data = map[string]any{"status": record.SimulateStatus, "percentage": percentageList}
	c.JSON(http.StatusOK, res)
}
