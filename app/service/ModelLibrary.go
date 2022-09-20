package service

import (
	"context"
	"fmt"
	"log"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/library/omc"
)

func ModelLibraryInitialization(packageModel []DataBaseModel.YssimModels) {
	packageModelMap := map[string]DataBaseModel.YssimModels{}
	for _, models := range packageModel {
		packageModelMap[models.PackageName] = models
	}
	packageAll := omc.OMC.GetPackages()
	for _, p := range packageAll {
		if _, ok := packageModelMap[p]; ok && packageModelMap[p].SysUser != "sys" {
			DeleteLibrary(p)
		} else {
			delete(packageModelMap, p)
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
			log.Printf("初始化模型库： %s  %t \n", models.PackageName, true)
		} else {
			log.Println("模型库：" + models.PackageName + "  初始化失败")
		}
	}
	config.R.HSet(context.Background(), "yssim-GraphicsData", map[string]string{"status": "1"})
}

func init() {
	omc.OMC.Clear()
}

func DeleteLibrary(deletePackage string) {
	omc.OMC.DeleteClass(deletePackage)
}

func modelCache(packageModel string) {
	modelsALL := omc.OMC.GetClassNames(packageModel, true)
	//dataLen := func() int {
	//	if len(modelsALL) > 500 {
	//		return 500
	//	}
	//	return len(modelsALL)
	//}()

	// ("function","",false,false,false,"/usr/lib/omc/NFModelicaBuiltin.mo",true,1006,1,1015,15,{},false,false,"","",false,"")
	// ("package","OpenModelica internal definitions and scripting functions",false,false,true,"/usr/lib/omc/NFModelicaBuiltin.mo",true,974,1,5503,17,{},false,false,"","text",false,"")
	// ("record","",false,false,false,"/usr/lib/omc/NFModelicaBuiltin.mo",true,1009,3,1010,17,{},true,false,"","",false,"")
	// ("type","Integer,Real,String,enumeration or array of some kind",false,false,false,"/usr/lib/omc/NFModelicaBuiltin.mo",true,1019,3,1020,18,{},false,false,"","",false,"")
	// ("impure function","",false,false,false,"/usr/lib/omc/NFModelicaBuiltin.mo",true,1067,3,1072,21,{},false,false,"","",false,"")
	omc.OMC.CacheRefreshSet(true)
	for p := 0; p < len(modelsALL); p++ {
		e := omc.OMC.GetClassInformation(modelsALL[p])
		if len(e) > 1 && e[0].(string) == "model" {
			log.Println("正在缓存：", modelsALL[p], " 的图形数据")
			GetGraphicsData(modelsALL[p])
		}
	}
	omc.OMC.CacheRefreshSet(false)
}

// 暂时不用，参数接口速度并不慢
func parametersCache(packageModel string) {
	modelsALL := omc.OMC.GetClassNames(packageModel, true)

	omc.OMC.CacheRefreshSet(true)
	for p := 0; p < len(modelsALL); p++ {
		e := omc.OMC.GetElements(modelsALL[p])
		for ee := 0; ee < len(e); ee++ {
			log.Println("正在缓存：", modelsALL[p], " 的参数数据")
			GetGraphicsData(modelsALL[p])
		}

	}
	omc.OMC.CacheRefreshSet(false)
}
