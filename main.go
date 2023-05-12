package main

import (
	_ "yssim-go/Init"
	"yssim-go/config"
	"yssim-go/middleware"
	"yssim-go/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//if os.Getenv("DEBUG") == "" {
	//	gin.SetMode(gin.ReleaseMode)
	//}
	g := gin.Default()

	g.Use(middleware.Cors())
	g.Static("/static", "./static")
	{
		router.AppDesignRouter(g)
		router.ModelRouter(g)
		router.SimulateRouter(g)
		router.UserRouter(g)
		router.FileRouter(g)
	}

	g.Run(config.ADDR + config.PORT)

	//var sysPackageModelAll []DataBaseModel.YssimModels
	//config.DB.Where("sys_or_user = ? AND userspace_id = ?", "sys", "0").Find(&sysPackageModelAll)
	//log.Println(len(sysPackageModelAll))
	//time.Sleep(time.Second * 10)
	//for i := 0; i < len(sysPackageModelAll); i++ {
	//	config.RedisCacheKey = sysPackageModelAll[i].PackageName + "-" + sysPackageModelAll[i].Version + "-GraphicsData"
	//	packageList := omc.OMC.GetPackages()
	//	for _, p := range packageList {
	//		if p == sysPackageModelAll[i].PackageName {
	//			omc.OMC.DeleteClass(p)
	//			break
	//		}
	//	}
	//	loadPackage := omc.OMC.LoadModel(sysPackageModelAll[i].PackageName, sysPackageModelAll[i].Version)
	//	fmt.Println("加载模型库： ", sysPackageModelAll[i].PackageName, loadPackage)
	//	modelsALL := omc.OMC.GetClassNames(sysPackageModelAll[i].PackageName, true)
	//	omc.OMC.CacheRefreshSet(true)
	//	for p := 0; p < len(modelsALL); p++ {
	//		log.Println("正在缓存：", modelsALL[p], " 的图形数据")
	//		//ClassRestriction := omc.OMC.GetClassRestriction(modelsALL[p])
	//		//if ClassRestriction == "model" {
	//		service.GetGraphicsData(modelsALL[p], "sys")
	//		//}
	//	}
	//	omc.OMC.CacheRefreshSet(false)
	//}
}
