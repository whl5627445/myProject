package timeConvert

import "strconv"

func SecondToTimeString(simulateSecond float64) string {
	hours := int(simulateSecond / 3600)
	minutes := int(simulateSecond) % 3600 / 60
	seconds := int(simulateSecond) % 60
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
