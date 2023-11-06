package config

import (
	_ "embed"
	"os"
)

var ModelicaKeywords = map[string]bool{"model": true, "class": true, "connector": true, "block": true, "function": true, "record": true, "expandable connector": true, "der": true, "and": true, "or": true, "not": true, "constant": true, "sum": true, "abs": true, "sign": true, "sqrt": true}

// var ParameterTranslation = map[string]string{
//	"Initialization": "初始化",
//	"General":        "通用设置",
//	"Advanced":       "高级设置",
//	"Attributes":     "属性设置",
//	"Parameters":     "参数",
//	"Modifiers":      "Modifiers",
//	"Dummy":          "Dummy",
//	"Component":      "组件",
//	"Name":           "名称",
//	"comment":        "注释",
// }

var MoldelSimutalionStatus = map[string]string{"1": "仿真排队中", "2": "正在仿真", "3": "仿真失败", "4": "仿真完成", "5": "删除任务", "6": "正在编译"}
var MoldelCompileStatus = map[string]string{"1": "编译排队中", "2": "正在编译", "3": "编译失败", "4": "编译完成"}
var ClassTypeAll = map[string]bool{"model": true, "class": true, "connector": true, "block": true, "function": true, "record": true, "expandable connector": true}

const dymolaConnect = "http://gateway:6535"
const DymolaSimutalionConnect = dymolaConnect + "/dymola"
const OmcFlaskConnect = dymolaConnect + "/omc-python"
const ADDR = "0.0.0.0:"
const CADConnect = dymolaConnect + "/caa"

var USERNAME = os.Getenv("USERNAME")
var PORT = os.Getenv("PORT")
var DEBUG = os.Getenv("DEBUG")
var Solver = map[string]string{"OM": "默认", "DM": "dymola", "JM": "第三方"}

var RedisCacheKey = USERNAME + "-yssim-GraphicsData"
var UserSpaceId = ""
var ModelCodeChan = make(chan string, 100)

//go:embed key/private_key.pem
var PrivateKey []byte

//go:embed key/public_key.pem
var PublicKey []byte
