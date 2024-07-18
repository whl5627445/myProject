package simulate

import (
	"fmt"
	"net/http"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	serviceV2 "yssim-go/app/v2/service"
	"yssim-go/config"

	"github.com/gin-gonic/gin"
)

var DB = config.DB
var userName = config.USERNAME

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
	var res DataType.ResponseData
	userSpaceId := c.GetHeader("space_id")
	token := c.GetHeader("Authorization")
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
		"token":               token,
	}
	replyId, err := serviceV2.GrpcSimulation(itemMap)
	if err != nil {
		fmt.Println("调用(GrpcSimulation)出错：", err)
		res.Err = "仿真出错"
		c.JSON(http.StatusOK, res)
	}

	res.Msg = "仿真任务正在准备，请等待仿真完成"
	res.Data = map[string]string{"id": replyId}
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
	serviceV2.DeleteSimulateTask(resultRecord.TaskId, resultRecord.SimulateModelResultPath)
	config.DB.Delete(&resultRecord)
	DB.Delete(&DataBaseModel.YssimSnapshots{}, "simulate_result_id = ?", recordId) //删除相关的快照
	res.Msg = "删除成功"
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

	if err := serviceV2.TerminateSimulateTask(resultRecord.TaskId); err != nil {
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
		serviceV2.DeleteSimulateTask(resultRecord[i].TaskId, resultRecord[i].SimulateModelResultPath)
		config.DB.Delete(&resultRecord[i])
	}

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}
