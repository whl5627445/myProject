package config

import (
	_ "embed"
	"os"

	"yssim-go/library/net"
)

var ModelicaKeywords = map[string]bool{"model": true, "class": true, "connector": true, "block": true, "function": true, "record": true, "expandable connector": true, "der": true, "and": true, "or": true, "not": true, "constant": true, "sum": true, "abs": true, "sign": true, "sqrt": true}

var Units = [][]string{{"rad", "deg"}, {"rad/s", "deg/s", "Hz", "rpm", "rev/min"}, {"m", "km", "mm"}, {"m2", "mm2", "cm2"}, {"m3", "cm3", "ml", "l"},
	{"s", "ms", "min", "h"}, {"m/s", "km/h", "mm/s"}, {"kg", "g"}, {"kg/m3", "g/cm3", "kg/l"}, {"N", "mN", "kN"}, {"Pa", "kPa", "MPa", "bar", "psi"},
	{"1", "%"}, {"J", "kJ", "kWh", "Wh"}, {"W", "kW", "MW", "mW"}, {"m3/s", "l/min", "l/h", "m3/h"}, {"K", "degC"}, {"1/K", "ppm/K"}, {"A", "mA", "kA"}, {"V", "mV", "kV"},
	{"C", "As", "Ah", "mAh"}, {"F", "µF"},
}
var MoldelSimutalionStatus = map[string]string{"1": "仿真排队中", "2": "正在仿真", "3": "仿真失败", "4": "仿真完成", "5": "删除任务", "6": "正在编译", "7": "仿真终止"}
var MoldelSimutalionStatus_ = map[string]string{"1": "排队中", "2": "仿真中", "3": "结束", "4": "结束", "5": "结束", "6": "编译中", "7": "结束"}
var MoldelCompileStatus = map[string]string{"1": "编译排队中", "2": "正在编译", "3": "编译失败", "4": "编译完成"}
var ClassTypeAll = map[string]bool{"model": true, "class": true, "connector": true, "block": true, "function": true, "record": true, "expandable connector": true}
var SensitiveWords = map[string]bool{"dymola": true, "omc": true, "openmodelica": true} // 需要被替换的敏感词

const gatewayConnect = "http://gateway:6535"

// const DymolaSimutalionConnect = gatewayConnect + "/dymola"
const FmuExportConnect = gatewayConnect + "/fmu-export"
const ADDR = "0.0.0.0:"
const CADConnect = gatewayConnect + "/caa"

var GrpcServerName = os.Getenv("GrpcServerName")
var GrpcPort = os.Getenv("GrpcPort")

var USERNAME = os.Getenv("USERNAME")
var PORT = os.Getenv("PORT")
var DEBUG = os.Getenv("DEBUG")
var Solver = map[string]string{"OM": "默认", "DM": "dymola", "JM": "第三方"}

var UserSpaceId = ""
var ModelCodeChan = make(chan string, 100)

//go:embed key/private_key.pem
var PrivateKey []byte

//go:embed key/public_key.pem
var PublicKey []byte

var NacosIp = os.Getenv("NacosIp")
var NacosPort = os.Getenv("NacosPort")

var ServiceIp = net.GetLocalIp()

var DebugDB = os.Getenv("DebugDB")       // 192.168.121.12:13306
var DebugRedis = os.Getenv("DebugRedis") // 192.168.121.12:6379
var DebugMongo = os.Getenv("DebugMongo") // 192.168.121.12:27017
