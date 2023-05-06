package timeConvert

import (
	"log"
	"strconv"
	"time"
)

func secondToTimeString(simulateSecond int) string {
	hours := simulateSecond / 3600
	minutes := simulateSecond % 3600 / 60
	seconds := simulateSecond % 60
	if simulateSecond < 0 {
		log.Println("仿真运行时间为负： ", simulateSecond)
		return "1秒"
	}
	if hours == 0 {
		if minutes == 0 {
			return strconv.Itoa(seconds) + "秒"
		} else {
			return strconv.Itoa(minutes) + "分钟" + strconv.Itoa(seconds) + "秒"
		}
	} else {
		return strconv.Itoa(hours) + "小时" + strconv.Itoa(minutes) + "分钟" + strconv.Itoa(seconds) + "秒"
	}
}

func UseTimeFormat(startTime, endTime int) string {
	if endTime <= 0 {
		endTime = int(time.Now().Unix())
	}
	return secondToTimeString(endTime - startTime)
}
