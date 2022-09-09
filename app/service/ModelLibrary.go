package service

import (
	"fmt"
	"yssim-go/app/DataBaseModel"
	"yssim-go/library/omc"
)

func ModelLibraryInitialization(packageModel []DataBaseModel.YssimModels) {
	omc.OMC.Clear()
	packageALL := omc.OMC.GetPackages()
	var packageMap map[string]bool
	for _, p := range packageALL {
		packageMap[p] = true
	}
	for _, models := range packageModel {
		ok := false
		if packageMap[models.PackageName] == true {
			if models.SysUser == "sys" {
				continue
			} else {
				omc.OMC.DeleteClass(models.PackageName)
			}
		}
		if models.FilePath == "" {
			cmd := fmt.Sprintf("loadModel(%s, {\"%s\"},true,\"\",false)", models.PackageName, models.Version)
			_, ok = omc.OMC.SendExpressionNoParsed(cmd)
		} else {
			ok = omc.OMC.LoadFile(models.FilePath)
		}
		if ok {
			if models.SysUser == "sys" {
				modelsALL := omc.OMC.GetClassNames(models.PackageName, true)
				dataLen := func() int {
					if len(modelsALL) > 500 {
						return 500
					}
					return len(modelsALL)
				}()
				omc.OMC.CacheRefreshSet(true)
				fmt.Println("正在缓存：", models.PackageName, " 的图形数据")
				for p := 0; p < dataLen; p++ {
					GetGraphicsData(modelsALL[p])
				}
			}
			fmt.Printf("初始化模型库： %s  %t \n", models.PackageName, true)
		} else {
			fmt.Println("模型库：" + models.PackageName + "  初始化失败")
		}
		omc.OMC.CacheRefreshSet(false)
	}
	omc.OMC.CacheRefreshSet(false)
	fmt.Println("标准库初始化完成")
}
