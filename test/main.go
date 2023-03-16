package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func openRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		//Addr:     "119.3.155.11:6379",
		Addr:     "124.70.211.127:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

var R = openRedis()

func main() {
	//ctx := context.Background()
	//a := R.HGetAll(ctx, "yssim-GraphicsData")
	//b := R.HVals(ctx, "yssim-GraphicsData")
	//c := R.HKeys(ctx, "yssim-GraphicsData")
	//d := R.HLen(ctx, "yssim-GraphicsData")
	//fmt.Println(len(a.Val()))
	//R.HSet()
	//key := c.Val()
	//value := b.Val()
	//s := time.Now().UnixNano() / 1e6
	//NewKeyValues := []string{}
	//for i := 0; i < len(key); i++ {
	//	NewKeyValues = append(NewKeyValues, key[i])
	//	NewKeyValues = append(NewKeyValues, value[i])
	//}
	//fmt.Println(time.Now().UnixNano()/1e6 - s)
	rdb1 := redis.NewClient(&redis.Options{
		//Addr:     "119.3.155.11:6379",
		Addr:     "124.70.211.127:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb2 := redis.NewClient(&redis.Options{
		//Addr:     "119.3.155.11:6379",
		Addr:     "119.3.155.11:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	keyList := []string{"Buildings-8.0.0-GraphicsData", "Buildings-9.1.0-GraphicsData", "Modelica-3.2.3-GraphicsData", "Modelica-4.0.0-GraphicsData"}
	for _, key := range keyList {
		packageCacheKeys := rdb1.HKeys(ctx, key).Val()
		packageCacheValues := rdb1.HVals(ctx, key).Val()
		NewKeyValues := []string{}
		for i := 0; i < len(packageCacheKeys); i++ {
			NewKeyValues = append(NewKeyValues, packageCacheKeys[i])
			NewKeyValues = append(NewKeyValues, packageCacheValues[i])
		}
		rdb2.HSet(ctx, key, NewKeyValues)
	}
	//fmt.Println(len(b.Val()))
	//fmt.Println(len(c.Val()))
	//fmt.Println(d.Val())
}

//{
//	{"Boolean", "tableOnFile", "= true, if table is defined on file or in function usertab", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Real", "table", "Table matrix (time = first column; e.g., table=[0, 0; 1, 1; 2, 4])", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{:,:}"},
//	{"String", "tableName", "Table name on file or in function usertab (see docu)", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"String", "fileName", "File where matrix is stored", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Boolean", "verboseRead", "= true, if info message that file is loading is to be printed", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Integer", "columns", "Columns of table to be interpolated", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{:}"},
//	{"Modelica.Blocks.Types.Smoothness", "smoothness", "Smoothness of table interpolation", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Modelica.Blocks.Types.Extrapolation", "extrapolation", "Extrapolation of data outside the definition range", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Modelica.Units.SI.Time", "timeScale", "Time scale of first table column", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Real", "offset", "Offsets of output signals", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{:}"},
//	{"Modelica.Units.SI.Time", "startTime", "Output = offset for time < startTime", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Modelica.Units.SI.Time", "shiftTime", "Shift time of first table column", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Modelica.Blocks.Types.TimeEvents", "timeEvents", "Time event handling of table interpolation", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Boolean", "verboseExtrapolation", "= true, if warning messages are to be printed if time is outside the table definition range", "public", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Modelica.Units.SI.Time", "t_min", "Minimum abscissa value defined in table", "public", "true", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Modelica.Units.SI.Time", "t_max", "Maximum abscissa value defined in table", "public", "true", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Real", "t_minScaled", "Minimum (scaled) abscissa value defined in table", "public", "true", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Real", "t_maxScaled", "Maximum (scaled) abscissa value defined in table", "public", "true", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Real", "p_offset", "Offsets of output signals", "protected", "true", "false", "false", "false", "parameter", "none", "unspecified", "{nout}"},
//	{"Modelica.Blocks.Types.ExternalCombiTimeTable", "tableID", "External table object", "protected", "false", "false", "false", "false", "parameter", "none", "unspecified", "{}"},
//	{"Modelica.Units.SI.Time", "nextTimeEvent", "Next time event instant", "protected", "false", "false", "false", "false", "discrete", "none", "unspecified", "{}"},
//	{"Real", "nextTimeEventScaled", "Next scaled time event instant", "protected", "false", "false", "false", "false", "discrete", "none", "unspecified", "{}"},
//	{"Real", "timeScaled", "Scaled time", "protected", "false", "false", "false", "false", "unspecified", "none", "unspecified", "{}"}
//}
//
//{
//	{Dialog("General","Table data definition",true,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data definition",not tableOnFile,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data definition",tableOnFile,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data definition",tableOnFile,false,false,"Text files (*.txt);;MATLAB MAT-files (*.mat)","Open file in which table is present",-,-,"",false)},
//	{Dialog("General","Table data definition",tableOnFile,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data interpretation",true,false,false,-,-,-,-,"modelica://Modelica/Resources/Images/Blocks/Sources/CombiTimeTable.png",false)},
//	{Dialog("General","Table data interpretation",true,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data interpretation",true,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data interpretation",true,false,false,-,-,-,-,"",false), Evaluate=true},
//	{Dialog("General","Table data interpretation",true,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data interpretation",true,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data interpretation",true,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data interpretation",smoothness == Modelica.Blocks.Types.Smoothness.LinearSegments,false,false,-,-,-,-,"",false)},
//	{Dialog("General","Table data interpretation",extrapolation == Modelica.Blocks.Types.Extrapolation.LastTwoPoints or extrapolation == Modelica.Blocks.Types.Extrapolation.HoldLastPoint,false,false,-,-,-,-,"",false)}
//	,{},{},{},{},{},{},{},{},{}
//}
