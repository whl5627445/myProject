package service

import (
	"fmt"
	"strconv"
	"yssim-go/library/stringOperation"
)

// 根据仿真进度百分比和仿真设置的起始终止时间获取当前仿真状态和时间, 时间精度与intervals相同
func GetSimulationState(id, startTimeStr, stopTimeStr, intervals string, percentage int64) map[string]any {

	states := map[string]string{
		"1": "排队中",
		"2": "编译中",
		"3": "仿真中",
		"4": "结束",
	}

	data := map[string]any{"simulate_record_id": id, "simulate_time": ""}
	if id != "" { // 仿真进行中
		data["status"] = 2
		precision := stringOperation.GetDecimalStrPrecision(intervals)
		if percentage >= 0 && percentage < 10 {
			data["simulate_status_msg"] = states["1"]
		} else if percentage >= 10 && percentage <= 30 {
			data["simulate_status_msg"] = states["2"]
		} else {
			startTime, _ := strconv.ParseFloat(startTimeStr, 64)
			stopTime, _ := strconv.ParseFloat(stopTimeStr, 64)
			elapsedTime := startTime + float64(percentage-30)*(stopTime-startTime)/float64(70)
			data["simulate_status_msg"] = states["3"]
			data["simulate_time"] = fmt.Sprintf("%.*f", precision, elapsedTime)
		}
	} else { // 仿真结束
		data["status"] = 4
		data["simulate_status_msg"] = states["4"]
	}

	return data
}
