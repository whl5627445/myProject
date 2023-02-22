package API

import (
	"math"
	"net/http"
	"strconv"
	"time"
	"yssim-go/library/timeConvert"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
)

var DB = config.DB

func GetSimulationOptionsView(c *gin.Context) {
	/*
		# 仿真参数获取接口
		## model_name: 模型名称，
	*/
	modelName := c.Query("model_name")
	var res responseData
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
	var item setSimulationOptionsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var res responseData
	result := service.SetSimulationOptions(item.ModelName, item.StartTime, item.StopTime, item.Tolerance, item.Interval)
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

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	packageId := c.Query("package_id")
	modelName := c.Query("model_name")
	var modelRecord DataBaseModel.YssimSimulateRecord
	DB.Where("package_id = ? AND username = ? AND simulate_model_name = ? AND simulate_start = ? AND userspace_id = ?", packageId, username, modelName, true, userSpaceId).First(&modelRecord)
	var res responseData
	if modelRecord.ID != "" {
		res.Data = 2
	} else {
		res.Data = 4
	}
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
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item modelSimulateData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	SimulateTypeDict := map[string]bool{"OM": true, "JM": true, "DM": true, "FmPy": true}
	if !SimulateTypeDict[item.SimulateType] {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//查询数据库中的实验id对应的记录
	var experimentRecord DataBaseModel.YssimExperimentRecord
	DB.Where("id = ? ", item.ExperimentId).First(&experimentRecord)

	var packageModel DataBaseModel.YssimModels
	err = DB.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	record := DataBaseModel.YssimSimulateRecord{
		ID:                uuid.New().String(),
		PackageId:         item.PackageId,
		UserspaceId:       userSpaceId,
		Username:          username,
		SimulateModelName: item.ModelName,
		SimulateStatus:    "1",
		StartTime:         item.StartTime,
		StopTime:          item.StopTime,
		Method:            item.Method,
		SimulateType:      item.SimulateType,
		NumberOfIntervals: item.NumberOfIntervals,
		Tolerance:         item.Tolerance,
		ExperimentId:      item.ExperimentId,
	}
	err = DB.Create(&record).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, "出现错误")
		return
	}
	SData := service.SimulateTask{
		SRecord:          record,
		Package:          packageModel,
		ExperimentRecord: experimentRecord,
	}
	service.SimulateTaskChan <- &SData
	var res responseData
	res.Msg = "仿真任务正在准备，请等待仿真完成"
	res.Data = map[string]string{"id": record.ID}
	c.JSON(http.StatusOK, res)
}

func SimulateResultGraphicsView(c *gin.Context) {
	/*
		# 仿真结果获取接口， 可一次获取多条
		## variable: 模型变量名字，
		## id: 仿真记录id值，在/simulate/record/list接口获取，
		## s1: 单位转换使用，固定为初始单位
		## s2: 位单位转换使用，需要转换为什么单位
	*/

	username := c.GetHeader("username")

	var item modelSimulateResultData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res responseData
	var resData []map[string]interface{}

	// 判断记录是否存在，有一条不存在就返回"not found"
	recordIdList := item.RecordId
	var recordList []DataBaseModel.YssimSimulateRecord
	err = DB.Where("id IN ? AND username = ?", recordIdList, username).Order("").Find(&recordList).Error
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
		if recordDict[recordIdList[i]].SimulateType == "FmPy" {
			data, ok = service.ReadSimulationResultFromGrpc(recordIdList[i], item.Variable)
		} else {
			data, ok = service.ReadSimulationResult([]string{item.Variable}, recordDict[recordIdList[i]].SimulateModelResultPath+"result_res.mat")
		}
		unitsData := service.ConvertUnits(item.S2, item.S1)
		if ok {
			ordinate := data[1]
			abscissa := data[0]
			if unitsData[0] == "true" {
				scaleFactor, _ := strconv.ParseFloat(unitsData[1], 64)
				offset, _ := strconv.ParseFloat(unitsData[2], 64)
				if len(ordinate) > 500 {
					step := len(ordinate) / 500
					o := []float64{}
					a := []float64{}
					for s := 0; s < len(ordinate); s++ {
						index := s * step
						if index >= len(ordinate) {
							break
						}
						o = append(o, data[1][index])
						a = append(a, data[0][index])
					}
					if len(ordinate)%500 != 0 {
						o = append(o, data[1][len(ordinate)-1])
						a = append(a, data[0][len(ordinate)-1])
					}
					ordinate = o
					abscissa = a
				}
				for p := 0; p < len(ordinate); p++ {
					ordinate[p] = ordinate[p]*scaleFactor + offset
				}
			}
			oneData := map[string]interface{}{
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
		# 仿真结果获取接口,单数
		## variable: 模型变量名字，
		## id: 仿真记录id值，在/simulate/record/list接口获取，
		## s1: 单位转换使用，固定为初始单位
		## s2: 位单位转换使用，需要转换为什么单位
	*/

	username := c.GetHeader("username")
	var item modelSimulateResultSingularData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res responseData
	// 判断记录是否存在，有一条不存在就返回"not found"
	recordIdList := item.RecordId
	var record DataBaseModel.YssimSimulateRecord
	err = DB.Where("id = ? AND username = ?", recordIdList, username).First(&record).Error
	if err != nil || record.SimulateStatus != "4" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	data, ok := service.ReadSimulationResult([]string{item.Variable}, record.SimulateModelResultPath+"result_res.mat")
	unitsData := service.ConvertUnits(item.S2, item.S1)
	if ok {
		scaleFactor, _ := strconv.ParseFloat(unitsData[1], 64)
		offset, _ := strconv.ParseFloat(unitsData[2], 64)
		ordinate := data[1]
		if unitsData[0] == "true" {
			for i := 0; i < len(ordinate); i++ {
				ordinate[i] = ordinate[i]*scaleFactor + offset
			}
		}
		res.Data = map[string]interface{}{
			"ordinate":  ordinate,
			"startTime": record.StartTime,
			"stopTime":  record.StopTime,
		}
	} else {
		res.Err = "结果不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func SimulateResultListView(c *gin.Context) {
	/*
	   # 仿真记录列表获取接口
	   # 模型名为空的时候查所有模型，只有查所有模型的时候才会分页。
	   ## return: 返回对应用户的所有仿真记录
	*/

	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	modelName := c.Query("model_name")
	pageNumStr := c.Query("page_num") //页码
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	var totle int64 //总条数
	var recordList []DataBaseModel.YssimSimulateRecord
	var resData map[string]interface{}
	resData = make(map[string]interface{})
	var dataList []map[string]interface{}
	if modelName != "" {
		DB.Limit(10).Where("username = ? AND simulate_model_name = ? AND userspace_id = ? AND simulate_status = ?", username, modelName, userSpaceId, "4").Order("create_time desc").Find(&recordList)
	} else {
		DB.Where("username = ? AND userspace_id = ?", username, userSpaceId).Find(&recordList).Count(&totle)
		DB.Limit(10).Offset((pageNumInt-1)*10).Where("username = ? AND userspace_id = ?", username, userSpaceId).Order("create_time desc").Find(&recordList)
	}
	pageCount := math.Ceil(float64(totle) / 10) //总页数
	for i := 0; i < len(recordList); i++ {
		simulateStartTime := time.Unix(recordList[i].SimulateStartTime, 0)
		simulateEndTime := time.Unix(recordList[i].SimulateEndTime, 0)

		data := map[string]interface{}{
			"index":               i + 1,
			"id":                  recordList[i].ID,
			"create_time":         recordList[i].CreatedAt.Format("2006-01-02 15:04:05"),
			"simulate_status":     config.MoldelSimutalionStatus[recordList[i].SimulateStatus],
			"simulate_start_time": simulateStartTime.Format("2006-01-02 15:04:05"),
			"simulate_end_time":   simulateEndTime.Format("2006-01-02 15:04:05"),
			"simulate_model_name": recordList[i].SimulateModelName,
			"simulate_run_time":   timeConvert.UseTimeFormat(int(simulateStartTime.Unix()), int(simulateEndTime.Unix())),
		}
		dataList = append(dataList, data)
	}
	resData["resultList"] = dataList
	resData["pageCount"] = pageCount
	resData["totle"] = totle
	var res responseData
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
	data := map[string]interface{}{"start_time": "", "stop_time": "", "step_size": "", "tolerance": "", "solver": "", "method": "", "number_intervals": "", "model_var_data": ""}
	data["start_time"] = simulateRecord.StartTime                 // 开始时间
	data["stop_time"] = simulateRecord.StopTime                   // 结束时间
	data["step_size"] = experimentRecord.Interval                 // 步长
	data["tolerance"] = experimentRecord.Tolerance                // 容差
	data["solver"] = config.Solver[experimentRecord.SimulateType] // 求解器
	data["method"] = experimentRecord.Method                      // 方法
	data["number_intervals"] = experimentRecord.NumberOfIntervals // 间隔
	data["model_var_data"] = experimentRecord.ModelVarData        // 模型组件相关参数属性

	var res responseData
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

	var res responseData
	if record.SimulateModelResultPath != "" && record.SimulateStart == false {
		result := service.SimulationResultTree(record.SimulateModelResultPath+"result_init.xml", parentNode, keyWords)
		res.Data = result
	} else {
		res.Err = "查询失败"
		res.Status = 2
	}
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
	var res responseData
	DB.Delete(&resultRecord)
	res.Msg = "删除成功"
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
	var res responseData
	var item experimentCreateData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var record DataBaseModel.YssimExperimentRecord
	DB.Where("package_id = ? AND experiment_name = ? AND username =? AND userspace_id =? AND model_name =?", item.PackageId, item.ExperimentName, username, userSpaceId, item.ModelName).First(&record)
	if record.ExperimentName != "" {
		res.Msg = "实验记录名称已存在，请更换。"
		c.JSON(http.StatusOK, res)
		return
	}
	experimentRecord := DataBaseModel.YssimExperimentRecord{
		ID:                uuid.New().String(),
		PackageId:         item.PackageId,
		UserspaceId:       userSpaceId,
		Username:          username,
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
	}
	err = DB.Create(&experimentRecord).Error
	if err != nil {
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
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

	var res responseData
	var item experimentDeleteData
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

	var res responseData
	var item experimentEditData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var recordName DataBaseModel.YssimExperimentRecord
	DB.Where("id != ? AND username =? AND userspace_id =? AND experiment_name =? ", item.ExperimentId, username, userSpaceId, item.ExperimentName).First(&recordName)
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
	DB.Where("package_id = ? AND userspace_id = ? AND username = ? AND model_name = ?", packageId, userSpaceId, username, modelName).Find(&recordList)

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
	var res responseData
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

	var res responseData
	res.Data = record.ModelVarData
	c.JSON(http.StatusOK, res)
}
