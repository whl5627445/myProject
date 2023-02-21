package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/library/omc"
)

func ModelLibraryInitialization(packageModel []DataBaseModel.YssimModels) {
	setOptions()
	packageModelMap := map[string]DataBaseModel.YssimModels{}
	for _, models := range packageModel {
		packageModelMap[models.PackageName] = models
	}
	packageAll := omc.OMC.GetPackages()
	for _, p := range packageAll {
		if _, ok := packageModelMap[p]; ok && packageModelMap[p].SysUser == "sys" {
			delete(packageModelMap, p)
		} else {
			DeleteLibrary(p)
		}
	}
	for _, models := range packageModelMap {
		ok := false
		if models.FilePath == "" {
			cmd := fmt.Sprintf("loadModel(%s, {\"%s\"},true,\"\",false)", models.PackageName, models.Version)
			_, ok = omc.OMC.SendExpressionNoParsed(cmd)

		} else {
			ok = omc.OMC.LoadFile(models.FilePath)
		}
		if ok {
			cacheStatus, _ := config.R.HGet(context.Background(), "yssim-GraphicsData", "status").Result() // 1是已缓存完成
			if models.SysUser == "sys" && cacheStatus != "1" {
				modelCache(models.PackageName)
			}
			log.Printf("初始化模型库： %s  %t \n", models.PackageName, ok)
		} else {
			log.Println("模型库：" + models.PackageName + "  初始化失败")
		}
	}
	config.R.HSet(context.Background(), "yssim-GraphicsData", map[string]string{"status": "1"})
}

func setOptions() {
	commandLineOptions := omc.OMC.GetCommandLineOptions()
	if strings.Contains(commandLineOptions, "nfAPI") {
		omc.OMC.Clear()
	}
}

func DeleteLibrary(deletePackage string) {
	omc.OMC.DeleteClass(deletePackage)
}

func modelCache(packageModel string) {
	modelsALL := omc.OMC.GetClassNames(packageModel, true)
	omc.OMC.CacheRefreshSet(true)
	for p := 0; p < len(modelsALL); p++ {
		//e := omc.OMC.GetClassInformation(modelsALL[p])
		//if len(e) > 1 && e[0].(string) == "model" {
		//if len(e) > 1 {
		log.Println("正在缓存：", modelsALL[p], " 的图形数据")
		GetGraphicsData(modelsALL[p])
		//} else {
		//	log.Println(modelsALL[p], " 不是model类型，跳过")
		//}
	}
	omc.OMC.CacheRefreshSet(false)
}

// 暂时不用，参数接口速度并不慢
//func parametersCache(packageModel string) {
//	modelsALL := omc.OMC.GetClassNames(packageModel, true)
//
//	omc.OMC.CacheRefreshSet(true)
//	for p := 0; p < len(modelsALL); p++ {
//		e := omc.OMC.GetElements(modelsALL[p])
//		for ee := 0; ee < len(e); ee++ {
//			log.Println("正在缓存：", modelsALL[p], " 的参数数据")
//			GetGraphicsData(modelsALL[p])
//		}
//
//	}
//	omc.OMC.CacheRefreshSet(false)
//}
