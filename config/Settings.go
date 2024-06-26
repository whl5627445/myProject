package config

import (
	_ "embed"
	"os"
)

var ModelicaKeywords = map[string]bool{"model": true, "class": true, "connector": true, "block": true, "function": true, "record": true, "expandable connector": true, "der": true, "and": true, "or": true, "not": true, "constant": true, "sum": true, "abs": true, "sign": true, "sqrt": true}

var MoldelSimutalionStatus = map[string]string{"1": "仿真排队中", "2": "正在仿真", "3": "仿真失败", "4": "仿真完成", "5": "删除任务", "6": "正在编译", "7": "仿真终止"}
var ClassTypeAll = map[string]bool{"model": true, "class": true, "connector": true, "block": true, "function": true, "record": true, "expandable connector": true}

const ADDR = "0.0.0.0:"

var USERNAME = os.Getenv("USERNAME")
var PORT = os.Getenv("PORT")
var Solver = map[string]string{"OM": "默认"}

var UserSpaceId = ""
var ModelCodeChan = make(chan string, 100)
