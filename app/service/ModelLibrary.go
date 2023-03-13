package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/serviceType"
	"yssim-go/config"
	"yssim-go/library/omc"
)

var r = config.R
var redisKey = config.RedisCacheKey

func ModelLibraryInitialization(packageModel []DataBaseModel.YssimModels) {
	setOptions()
	packageModelMap := map[string]DataBaseModel.YssimModels{}
	for _, models := range packageModel {
		packageModelMap[models.PackageName] = models
	}
	packageAll := omc.OMC.GetPackages()
	for _, p := range packageAll {
		version := GetVersion(p)
		if _, ok := packageModelMap[p]; ok && packageModelMap[p].SysUser == "sys" && packageModelMap[p].Version == version {
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
			//cacheStatus, _ := config.R.HGet(context.Background(), "yssim-GraphicsData", "status").Result() // 1是已缓存完成
			//if models.SysUser == "sys" && cacheStatus != "1" {
			//modelCache(models.PackageName, models.SysUser)
			//}
			log.Printf("初始化模型库：%s %s  %s %t \n", models.SysUser, models.PackageName, models.Version, ok)
		} else {
			log.Println("模型库：" + models.PackageName + "  初始化失败")
		}
	}
	//config.R.HSet(context.Background(), "yssim-GraphicsData", map[string]string{"status": "1"})
}

func ModelLibraryInitializationNew(packageModel []DataBaseModel.YssimModels) {
	setOptions()
	ctx := context.Background()
	r.HDel(ctx, config.USERNAME+"-yssim-componentGraphicsData")
	r.HDel(ctx, config.USERNAME+"-yssim-modelGraphicsData")
	packageAll := omc.OMC.GetPackages()
	for _, p := range packageAll {
		DeleteLibrary(p)
	}

	for _, models := range packageModel {
		ok := false
		if models.FilePath == "" {
			//cmd := fmt.Sprintf("loadModel(%s, {\"%s\"},true,\"\",false)", models.PackageName, models.Version)
			ok = omc.OMC.LoadModel(models.PackageName, models.Version)
		} else {
			ok = omc.OMC.LoadFile(models.FilePath)
			_, err := GetLoadPackageConflict(models.PackageName, models.Version, models.FilePath)
			if err != nil {
				deleteModel(models.PackageName)
				ok = false
			}
		}
		if ok {
			//cacheStatus, _ := config.R.HGet(context.Background(), config.USERNAME+"-yssim-GraphicsData", "status").Result() // 1是已缓存完成
			//if models.SysUser == "sys" && cacheStatus != "1" {
			//	modelCache(models.PackageName, models.SysUser)
			//}
			packageCacheKeys := r.HKeys(ctx, models.PackageName+"-"+models.Version+"-GraphicsData").Val()
			packageCacheValues := r.HVals(ctx, models.PackageName+"-"+models.Version+"-GraphicsData").Val()
			NewKeyValues := []string{}
			for i := 0; i < len(packageCacheKeys); i++ {
				NewKeyValues = append(NewKeyValues, packageCacheKeys[i])
				NewKeyValues = append(NewKeyValues, packageCacheValues[i])
			}
			r.HSet(ctx, redisKey, NewKeyValues)
			log.Printf("初始化模型库：%s %s  %s %t \n", models.SysUser, models.PackageName, models.Version, ok)
		} else {
			log.Printf("初始化模型库：%s %s %s 失败 \n", models.SysUser, models.PackageName, models.Version)
		}
	}
	//config.R.HSet(context.Background(), "yssim-GraphicsData", map[string]string{"status": "1"})
}

func setOptions() {
	//commandLineOptions := omc.OMC.GetCommandLineOptions()
	//if strings.Contains(commandLineOptions, "nfAPI") {
	omc.OMC.SetOptions()
	//}
}

func modelCache(packageModel, permissions string) {
	modelsALL := omc.OMC.GetClassNames(packageModel, true)
	omc.OMC.CacheRefreshSet(true)
	for p := 0; p < len(modelsALL); p++ {
		if permissions == "sys" {
			log.Println("正在缓存：", modelsALL[p], " 的图形数据")
			GetGraphicsData(modelsALL[p], permissions)
		}
		//GetGraphicsData(modelsALL[p])

	}
	omc.OMC.CacheRefreshSet(false)
}

func GetLibraryAndVersions() map[string]string {
	// 获取库和版本
	data := map[string]string{}
	loadedLibraries := omc.OMC.GetPackages()
	for _, library := range loadedLibraries {
		libraryVersion := omc.OMC.GetClassInformation(library)[14].(string)
		data[library] = libraryVersion
	}
	return data
}

func GetVersion(packageName string) string {
	libraryVersion := omc.OMC.GetClassInformation(packageName)[14].(string)
	return libraryVersion
}

func GetPackageUses(packageName string) [][]string {
	// 获取包用到的包
	return omc.OMC.GetUses(packageName)
}

func GetLoadedLibraries() []string {
	// 获取已加载库
	return omc.OMC.GetLoadedLibraries()
}

func GetAvailableLibraryVersions(packageName string) []string {
	// 获取库的可用版本
	return omc.OMC.GetAvailableLibraryVersions(packageName)
}

func CheckPackageConflict(packageName, version string) []serviceType.CheckPackageUsesLibrary {
	var data []serviceType.CheckPackageUsesLibrary
	packageVersion := version
	packages := omc.OMC.GetPackages()

	for p := len(packages) - 1; p >= 0; p-- {
		pName := packages[p]
		pVersion := GetVersion(pName)
		switch {
		case pName == packageName && packageVersion == pVersion:
			return nil
		case pName == packageName && packageVersion != pVersion:
			var unloadLibrary serviceType.CheckPackageUsesLibrary
			unloadLibrary.Name = pName
			unloadLibrary.Version = pVersion
			data = append(data, unloadLibrary)
			packages = append(packages[:p], packages[p+1:]...)
			for s := len(packages) - 1; s >= 0; s-- {
				sUsePackageList := GetPackageUses(packages[s])
				for _, nameVersion := range sUsePackageList {
					n := nameVersion[0]
					v := nameVersion[1]
					sName := packages[s]
					sVersion := GetVersion(packages[s])
					if n == pName && v == pVersion {
						var l serviceType.CheckPackageUsesLibrary
						l.Name = sName
						l.Version = sVersion
						data = append(data, l)
						packages = append(packages[:s], packages[s+1:]...)
						p = s
					}
				}
			}
		default:
			packageUses := GetPackageUses(packages[p])
			for _, pUses := range packageUses {
				pUsesName := pUses[0]
				pUsesVersion := pUses[1]
				if pUsesName == packageName && version != pUsesVersion {
					var unloadLibrary serviceType.CheckPackageUsesLibrary
					unloadLibrary.Name = pName
					unloadLibrary.Version = pVersion
					data = append(data, unloadLibrary)
					packages = append(packages[:p], packages[p+1:]...)
				}
			}
		}
	}
	return data
}

func LoadPackage(packageName, version, path string) bool {
	// 加载相应的库与版本
	if path == "" {
		return omc.OMC.LoadModel(packageName, version)
	}
	return omc.OMC.LoadFile(path)
}

func DeleteLibrary(packageName string) bool {
	return omc.OMC.DeleteClass(packageName)
}

func GetLoadPackageConflict(packageName, version, path string) ([]serviceType.CheckPackageUsesLibrary, error) {
	packageNameAndVersion := [][]string{{packageName, version}}
	if path != "" {
		LoadPackage(packageName, version, path)
	}
	uses := GetPackageUses(packageName)
	packageNameAndVersion = append(packageNameAndVersion, uses...)
	if path != "" {
		deleteModel(packageName)
	}
	var unloadPackageNameList []string
	var unloadList []serviceType.CheckPackageUsesLibrary
	for i := 0; i < len(packageNameAndVersion); i++ {
		unloadPackageList := CheckPackageConflict(packageNameAndVersion[i][0], packageNameAndVersion[i][1])
		if len(unloadPackageList) > 0 {
			for _, library := range unloadPackageList {
				unloadPackageNameList = append(unloadPackageNameList, library.Name)
			}
		}
		unloadList = append(unloadList, unloadPackageList...)
	}
	if len(unloadPackageNameList) > 0 {
		errStr := fmt.Sprintf("加载 %s 模型库需要先卸载 %s 模型库", packageName, strings.Join(unloadPackageNameList, ", "))
		return unloadList, errors.New(errStr)
	}
	return nil, nil
}

func LoadAndDeleteLibrary(packageName, version, path, loadOrUnload string) error {
	result := false
	if loadOrUnload == "unload" {
		result = DeleteLibrary(packageName)
	} else {
		result = LoadPackage(packageName, version, path)
	}
	if !result {
		return errors.New(fmt.Sprintf("操作模型库 %s %s 时出错，请联系管理员", packageName, version))
	}
	return nil
}
