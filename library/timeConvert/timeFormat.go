package timeConvert

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func secondToTimeString(simulateSecond, num int) string {
	months := simulateSecond / (86400 * 30)
	days := simulateSecond % (86400 * 30) / 86400
	hours := (simulateSecond % 86400) / 3600
	minutes := simulateSecond % 3600 / 60
	seconds := simulateSecond % 60
	if simulateSecond < 0 {
		log.Println("仿真运行时间为负： ", simulateSecond)
		return "1秒"
	}

	timeStr := ""
	if months > 0 {
		timeStr = fmt.Sprintf("%d月,%d天,%d小时,%d分钟,%d秒", months, days, hours, minutes, seconds)
	} else if days > 0 {
		timeStr = fmt.Sprintf("%d天,%d小时,%d分钟,%d秒", days, hours, minutes, seconds)
	} else if hours > 0 {
		timeStr = fmt.Sprintf("%d小时,%d分钟,%d秒", hours, minutes, seconds)
	} else if minutes > 0 {
		timeStr = fmt.Sprintf("%d分钟,%d秒", minutes, seconds)
	} else {
		timeStr = fmt.Sprintf("%d秒", seconds)
	}
	if num != -1 {
		strList := strings.Split(timeStr, ",")
		if len(strList) >= num {
			return strings.Join(strList[:num], "")
		}
	}
	return strings.ReplaceAll(timeStr, ",", "")
}

func UseTimeFormat(startTime, endTime int) string {
	if endTime <= 0 {
		endTime = int(time.Now().Unix())
	}
	return secondToTimeString(endTime-startTime, -1)
}

func UseTimeFormatNew(startTime, endTime, num int) string {
	if endTime <= 0 {
		endTime = int(time.Now().Unix())
	}
	return secondToTimeString(endTime-startTime, num)
}
