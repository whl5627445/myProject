package mapProcessing

import "strconv"

// dymola仿真参数处理
func GetMapKeysAndValues(m map[string]string) ([]string, []float64) {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	// flase为0,true为1,数字类型字符串都改为float类型,""舍去字段
	var keys []string
	var values []float64
	for k, v := range m {
		switch v {
		case "flase":
			keys = append(keys, k)
			values = append(values, 0)
		case "true":
			keys = append(keys, k)
			values = append(values, 1)
		case "":
			continue
		default:
			keys = append(keys, k)
			val, _ := strconv.ParseFloat(v, 64)
			values = append(values, val)
		}
	}
	return keys, values
}
