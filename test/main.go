package main

import (
	"fmt"
	"log"
)

func secondToTimeString(simulateSecond int) string {
	months := simulateSecond / (86400 * 30)
	days := simulateSecond % (86400 * 30) / 86400
	hours := (simulateSecond % 86400) / 3600
	minutes := simulateSecond % 3600 / 60
	seconds := simulateSecond % 60
	if simulateSecond < 0 {
		log.Println("仿真运行时间为负： ", simulateSecond)
		return "1秒"
	}
	if months > 0 {

	}
	if months > 0 {
		return fmt.Sprintf("%d月%d天%d小时%d分%d秒", months, days, hours, minutes, seconds)
	} else if days > 0 {
		return fmt.Sprintf("%d天%d小时%d分%d秒", days, hours, minutes, seconds)
	} else if hours > 0 {
		return fmt.Sprintf("%d小时%d分%d秒", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d分%d秒", minutes, seconds)
	} else {
		return fmt.Sprintf("%d秒", seconds)
	}
}

func main() {
	sec := 8864010
	str := secondToTimeString(sec)
	fmt.Println(str)
}
