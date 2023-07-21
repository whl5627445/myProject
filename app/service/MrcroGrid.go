package service

import (
	"context"
	"fmt"
	"time"
	"yssim-go/config"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var timeCur time.Time
var MB = config.MB

func GetAppPowSingleData(names []string) []map[string]any {
	//查数据库的集合
	collection := MB.Database("micro_grid").Collection("result_data")
	// 查找文档
	var doc bson.M
	err := collection.FindOne(context.Background(), bson.M{}).Decode(&doc)
	// 检查错误
	if err != nil {
		fmt.Println("查询文档失败：", err)
	}

	//时间格式的2023年起始时间
	layout := "2006/1/2/15"
	timeStrRefer := "2023/1/1/0"
	timeRefer, _ := time.ParseInLocation(layout, timeStrRefer, time.Local)

	//相较于2023/1/1/0的天数及当天的小时数，即为数据库二维数组的行和列
	dataNum := int((timeCur.Sub(timeRefer)) / time.Hour)
	day := dataNum / 24

	var data []map[string]any
	// 遍历文档获取对应数据
	for _, key := range names {
		value := doc[key]
		temp1, ok := value.(primitive.A)
		if !ok {
			fmt.Println("类型转换失败1")
		}
		temp2, ok := temp1[day].(primitive.A)
		if !ok {
			fmt.Println("类型转换失败2")
		}
		for i := 0; i <= 23; i++ {
			data0 := map[string]any{
				"x":    i,
				"y":    temp2[i],
				"name": key,
			}
			data = append(data, data0)
		}
	}
	return data
}

func GetAppPowDoubleData(names []string) []map[string]any {
	//查数据库的集合
	collection := MB.Database("micro_grid").Collection("result_data")
	// 查找文档
	var doc bson.M
	err := collection.FindOne(context.Background(), bson.M{}).Decode(&doc)
	// 检查错误
	if err != nil {
		fmt.Println("查询文档失败：", err)
	}

	//时间格式的2023年起始时间
	layout := "2006/1/2/15"
	timeStrRefer := "2023/1/1/0"
	timeRefer, _ := time.ParseInLocation(layout, timeStrRefer, time.Local)

	//相较于2023/1/1/0的天数及当天的小时数，即为数据库二维数组的行和列
	dataNum := int((timeCur.Sub(timeRefer)) / time.Hour)
	day := dataNum / 24

	var data []map[string]any
	// "蓄电池充放电功率"及"蓄电池SOC"为双轴
	tempy1, _ := doc[names[0]].(primitive.A)
	tempz1, _ := doc[names[1]].(primitive.A)
	tempy2, _ := tempy1[day].(primitive.A)
	tempz2, _ := tempz1[day].(primitive.A)
	for i := 0; i <= 23; i++ {
		data0 := map[string]any{
			"x": i,
			"y": tempy2[i],
			"z": tempz2[i],
		}
		data = append(data, data0)
	}
	return data
}

func GetAppPowPieChartData(names []string) []map[string]any {
	//查数据库的集合
	collection := MB.Database("micro_grid").Collection("result_data")
	// 查找文档
	var doc bson.M
	err := collection.FindOne(context.Background(), bson.M{}).Decode(&doc)
	// 检查错误
	if err != nil {
		fmt.Println("查询文档失败：", err)
	}

	//时间格式的2023年起始时间
	layout := "2006/1/2/15"
	timeStrRefer := "2023/1/1/0"
	timeRefer, _ := time.ParseInLocation(layout, timeStrRefer, time.Local)

	//相较于2023/1/1/0的天数及当天的小时数，即为数据库二维数组的行和列
	dataNum := int((timeCur.Sub(timeRefer)) / time.Hour)
	day := dataNum / 24
	hour := dataNum % 24

	var data []map[string]any
	// 遍历文档获取对应数据
	for _, name := range names {
		value := doc[name]
		temp1, ok := value.(primitive.A)
		if !ok {
			fmt.Println("类型转换失败1")
		}
		temp2, ok := temp1[day].(primitive.A)
		if !ok {
			fmt.Println("类型转换失败2")
		}

		data0 := map[string]any{
			"s": name,
			"v": temp2[hour],
		}
		data = append(data, data0)
	}
	return data
}

func init() {
	/*
		# 获取电网app数据的时间戳每十秒增加一小时
	*/
	//新开一个线程解决阻塞问题
	go updateCount()
}

func updateCount() {
	for {
		layout := "2006/1/2/15"
		timeStr := "2023/1/1/0"
		timeCur, _ = time.ParseInLocation(layout, timeStr, time.Local)
		for k := 0; k < 365*24-1; k++ {
			time.Sleep(time.Second)
			timeCur = timeCur.Add(time.Hour)
		}
	}
}

func GetAppPowLabelData(name string) map[string]interface{} {
	//查数据库的集合
	collection := MB.Database("micro_grid").Collection("result_data")
	// 查找文档
	var doc bson.M
	err := collection.FindOne(context.Background(), bson.M{}).Decode(&doc)
	// 检查错误
	if err != nil {
		fmt.Println("查询文档失败：", err)
	}

	//时间格式的2023年起始时间
	layout := "2006/1/2/15"
	timeStrRefer := "2023/1/1/0"
	timeRefer, _ := time.ParseInLocation(layout, timeStrRefer, time.Local)

	//相较于2023/1/1/0的天数及当天的小时数，即为数据库二维数组的行和列
	dataNum := int((timeCur.Sub(timeRefer)) / time.Hour)
	day := dataNum / 24
	hour := dataNum % 24

	value := doc[name]
	temp1, ok := value.(primitive.A)
	if !ok {
		fmt.Println("类型转换失败1")
	}
	temp2, ok := temp1[day].(primitive.A)
	if !ok {
		fmt.Println("类型转换失败2")
	}

	data := map[string]interface{}{
		"value": temp2[hour],
	}

	return data
}
