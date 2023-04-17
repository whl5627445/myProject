package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/library/omc"
)

var r = config.R
var redisKey = config.RedisCacheKey

func ModelLibraryInitialization(packageModel []DataBaseModel.YssimModels) {
	setOptions()
	packageAll := GetLibraryAndVersions()

	for _, models := range packageModel {
		version, ok := packageAll[models.PackageName]
		switch {
		case ok && version == models.Version && models.SysUser == "sys":
			delete(packageAll, models.PackageName)
			continue
		}
		ok = false
		if models.FilePath == "" {
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
			log.Printf("初始化模型库：%s %s  %s %t \n", models.SysUser, models.PackageName, models.Version, ok)
		} else {
			log.Printf("初始化模型库：%s %s %s 失败 \n", models.SysUser, models.PackageName, models.Version)
		}
		delete(packageAll, models.PackageName)
	}
	for k, _ := range packageAll {
		deleteModel(k)
	}
	lPackage := GetLibraryAndVersions()
	refreshCache(lPackage)

}

func setOptions() {
	omc.OMC.SetOptions()
}

func Clear() {
	omc.OMC.SendExpressionNoParsed("clear()")
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

func GetLoadedLibraries() [][]string {
	// 获取已加载库
	lList := [][]string{}
	loadedLibraries := omc.OMC.GetLoadedLibraries()
	for i := 0; i < len(loadedLibraries); i++ {
		name := loadedLibraries[i]
		version := GetVersion(name)
		lList = append(lList, []string{name, version})
	}
	return lList
}

func GetAvailableLibraryVersions(packageName string) []string {
	// 获取库的可用版本
	return omc.OMC.GetAvailableLibraryVersions(packageName)
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

func GetLoadPackageConflict(packageName, version, path string) ([]map[string]string, error) {
	var unloadList []map[string]string
	whiteList := map[string]bool{"Complex": true, "ModelicaServices": true}
	unloadPackageNameList := checkLibraryInterdependenceNoLoad(packageName, version)
	if len(unloadPackageNameList) == 0 {
		_ = LoadPackage(packageName, version, path)
		unloadPackageNameList = checkLibraryInterdependenceIsLoad(packageName, version)
		if len(unloadPackageNameList) > 0 {
			deleteModel(packageName)
		} else {
			return nil, nil
		}
	}
	errStrNameList := []string{}
	for i := len(unloadPackageNameList) - 1; i >= 0; i-- {
		if !whiteList[unloadPackageNameList[i]] {
			unloadList = append(unloadList, map[string]string{"name": unloadPackageNameList[i]})
			errStrNameList = append(errStrNameList, unloadPackageNameList[i])
		}
	}
	errStr := fmt.Sprintf("加载 %s 模型库需要先卸载 %s 模型库", packageName, strings.Join(errStrNameList, ","))
	return unloadList, errors.New(errStr)

}

func LoadAndDeleteLibrary(packageName, version, path, loadOrUnload string) error {
	result := false
	if loadOrUnload == "unload" {
		result = DeleteLibrary(packageName)
	} else {
		result = LoadPackage(packageName, version, path)
		lPackage := GetLibraryAndVersions()
		refreshCache(lPackage)
	}
	if !result {
		return errors.New(fmt.Sprintf("操作模型库 %s %s 时出错，请联系管理员", packageName, version))
	}
	return nil
}

func checkLibraryInterdependenceNoLoad(packageName, version string) []string {
	var unloadPackageNameList []string
	unloadMap := map[string]bool{}
	LoadPackageList := GetLoadedLibraries()
	for _, l := range LoadPackageList {
		if l[0] == packageName {
			if l[1] == version {
				return nil
			} else {
				unloadMap[l[0]] = true
			}
		}
	}
	unloadMap = getInterdependence(unloadMap, LoadPackageList)
	for k, _ := range unloadMap {
		unloadPackageNameList = append(unloadPackageNameList, k)
	}
	return unloadPackageNameList
}

func checkLibraryInterdependenceIsLoad(packageName, version string) []string {
	var unloadPackageNameList []string
	unloadMap := map[string]bool{}
	LoadPackageList := GetLoadedLibraries()
	uses := GetPackageUses(packageName)
	for i := 0; i < len(LoadPackageList); i++ {
		if LoadPackageList[i][0] == packageName {
			LoadPackageList = append(LoadPackageList[:i], LoadPackageList[i+1:]...)
		}
	}
	// 查看需要被卸载的库用到哪些其他库
	for _, u := range uses { // 循环被卸载库的依赖项有没有被加载
		for _, l := range LoadPackageList {
			//lUses := GetPackageUses(l[0]) // 查看已加载的库的依赖项
			if l[0] == u[0] && l[1] != u[1] { //  查看被加载库的依赖项是否被加载
				unloadMap[l[0]] = true
			}
		}
	}
	unloadMap = getInterdependence(unloadMap, LoadPackageList)
	for k, _ := range unloadMap {
		unloadPackageNameList = append(unloadPackageNameList, k)
	}
	return unloadPackageNameList
}

func getInterdependence(unloadMap map[string]bool, LoadPackageList [][]string) map[string]bool {
	for un, _ := range unloadMap {
		uses := GetPackageUses(un) // 查看需要被卸载的库用到哪些其他库
		for _, u := range uses {   // 循环被卸载库的依赖项有没有被加载
			for _, l := range LoadPackageList {
				lUses := GetPackageUses(l[0]) // 查看已加载的库的依赖项
				if l[0] == u[0] {             //  查看被加载库的依赖项是否被加载
					unloadMap[l[0]] = true
				}
				for _, use := range lUses { //查看已加载的库是否依赖需要被卸载的库
					if use[0] == un {
						unloadMap[l[0]] = true
					}
				}
			}
		}
	}
	return unloadMap
}

func refreshCache(packageAndVersion map[string]string) {
	ctx := context.Background()
	r.Del(ctx, config.USERNAME+"-yssim-componentGraphicsData").Result()
	r.Del(ctx, config.USERNAME+"-yssim-GraphicsData").Result()
	r.Del(ctx, config.USERNAME+"-yssim-modelGraphicsData").Result()
	for k, v := range packageAndVersion {
		packageCacheKeys := r.HKeys(ctx, k+"-"+v+"-GraphicsData").Val()
		packageCacheValues := r.HVals(ctx, k+"-"+v+"-GraphicsData").Val()
		NewKeyValues := []string{}
		for i := 0; i < len(packageCacheKeys); i++ {
			NewKeyValues = append(NewKeyValues, packageCacheKeys[i])
			NewKeyValues = append(NewKeyValues, packageCacheValues[i])
		}
		r.HSet(ctx, redisKey, NewKeyValues)
	}
}
