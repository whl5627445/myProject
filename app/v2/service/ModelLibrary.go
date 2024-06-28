package serviceV2

import (
	"log"

	"yssim-go/app/DataBaseModel"
	serviceV1 "yssim-go/app/v1/service"
	"yssim-go/library/omc"
)

// RestartOMC 重启OMC服务
func RestartOMC() {
	serviceV1.StopOMC()
	serviceV1.StartOMC()
}

// LibraryInitializationSingle 加载单个模型
func LibraryInitializationSingle(packageName, packageVersion, filePath string) bool {
	// 根据模型文件加载单个模型
	if filePath != "" {
		ok := omc.OMC.LoadFileNoPwd(filePath)
		log.Printf("初始化模型库：%s %s  %t \n", packageName, packageVersion, ok)
		return ok
	}

	// 根据模型名称和版本加载单个模型
	ok := omc.OMC.LoadModel(packageName, packageVersion)
	log.Printf("初始化模型库：%s %s  %t \n", packageName, packageVersion, ok)
	return ok
}

// GetAllLoadableModel 获取所有可加载的模型信息
func GetAllLoadableModel(packageInformation map[string]map[string]string, packageModelAll []DataBaseModel.YssimModels) []map[string]string {
	data := []map[string]string{}

	if len(packageInformation) == 0 {
		//packageInformation为空，直接取所有的packageModel信息
		for _, packageModel := range packageModelAll {
			m := map[string]string{
				"id":      packageModel.ID,
				"name":    packageModel.PackageName,
				"version": packageModel.Version,
			}
			data = append(data, m)
		}
	} else {
		//packageInformation不为空，以packageModelAll为基准取二者交集
		packageInformationNameMap := map[string]bool{}
		for packageName := range packageInformation {
			packageInformationNameMap[packageName] = true
		}

		for _, packageModel := range packageModelAll {
			if ok := packageInformationNameMap[packageModel.PackageName]; ok {
				m := map[string]string{
					"id":      packageModel.ID,
					"name":    packageModel.PackageName,
					"version": packageModel.Version,
				}
				data = append(data, m)
			}
		}
	}

	return data
}
