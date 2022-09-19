package Init

import (
	"log"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
)

func SimulationService() {
	username := config.USERNAME
	var undoneRecordAll []DataBaseModel.YssimSimulateRecord
	err := config.DB.Where("username = ?", username).Not("simulate_status = 4").Find(&undoneRecordAll).Error
	if err != nil {
		panic(err)
	}
	for _, record := range undoneRecordAll {
		var packageModel DataBaseModel.YssimModels
		config.DB.Where("ID = ?", record.PackageId).First(&packageModel)
		task := service.SimulateTask{
			SRecord: record,
			Package: packageModel,
		}
		service.SimulateTaskChan <- &task
		log.Printf("未完成任务进入排队通道： %s \n", task.SRecord.ID)
	}
	log.Println("未完成任务排队完成，仿真任务线程启动")
	for {
		task := <-service.SimulateTaskChan
		service.ModelSimulate(task)
		log.Printf("仿真任务执行完成： %s \n", task.SRecord.ID)
	}

}
