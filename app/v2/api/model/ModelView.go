package APIv2

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	serviceV2 "yssim-go/app/v2/service"
	"yssim-go/config"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

var dbModel = config.DB
var userName = config.USERNAME

func GetInstanceDataView(c *gin.Context) {
	/*
		# 获取模型的画图数据，一次性返回
		## package_id: 模型包的id
		## modelname: 需要查询的模型名称，全称， 例如“Modelica.Blocks.Examples.PID_Controller”
	*/
	var item DataType.ModelGraphicsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	userSpaceId := c.GetHeader("space_id")
	var packageModel DataBaseModel.YssimModels
	packageName := strings.Split(item.ModelName, ".")[0]
	err = dbModel.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", packageName, []string{"sys", userName}, []string{"0", userSpaceId}).First(&packageModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var res DataType.ResponseData
	res.Data = map[string]any{"encryption": packageModel.Encryption, "model": serviceV2.GetModelInstanceData(item.ModelName)}
	c.JSON(http.StatusOK, res)
}

func AddModelComponentView(c *gin.Context) {
	/*
		# 创建模型当中的模型组件
		## package_id： 包id
		## model_name: 需要创建的组件在哪个模型之下，例如在"Filter1"模型中创建组件
		## new_component_name: 新创建的组件名称，例如"abs1"
		## old_component_name: 被创建成组件的模型名称， 例如"Modelica.Blocks.Math.Abs"
		## origin: 原点， 例如"0,0"
		## extent: 范围坐标, 例如["-10,-10", "10,10"]
		## rotation: 旋转角度, 例如"0"，不旋转`
	*/

	userSpaceId := c.GetHeader("space_id")
	var res DataType.ResponseData
	var item DataType.AddComponentData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	var modelPackage DataBaseModel.YssimModels
	dbModel.Where("id = ? AND sys_or_user IN ? AND userspace_id IN ?", item.PackageId, []string{userName, "sys"}, []string{userSpaceId, "0"}).First(&modelPackage)
	if modelPackage.SysUser == "sys" || modelPackage.Encryption {
		res.Err = "该模型不允许此操作"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	data, err := serviceV2.AddComponent(item.ModelName, item.OldComponentName, item.NewComponentName, item.Origin)
	if err != nil {
		res.Err = err.Error()
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = data
	res.Msg = "新增组件成功"
	c.JSON(http.StatusOK, res)
}

func GetAllModelView(c *gin.Context) {
	/*
		# 获取用户空间所有模型信息
		## space_id: 用户空间id
	*/
	var res DataType.ResponseData
	var item DataType.LoginUserSpaceModel
	if err := c.BindJSON(&item); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 从用户空间表获取package数据
	var space DataBaseModel.YssimUserSpace
	dbModel.Model(space).Where("id = ? AND username = ?", item.SpaceId, userName).First(&space)
	if space.ID == "" {
		res.Status = 4
		res.Err = "空间已被删除"
		c.JSON(http.StatusOK, res)
		return
	}

	// 从模型表获取package数据
	var packageModelAll []DataBaseModel.YssimModels
	dbModel.Where("sys_or_user IN ?  AND default_version = ? AND userspace_id IN ?", []string{"sys", userName}, true, []string{"0", item.SpaceId}).Find(&packageModelAll)

	dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", item.SpaceId, userName).UpdateColumn("last_login_time", time.Now().Local().Unix())

	information, _ := space.PackageInformation.MarshalJSON()
	packageInformation := map[string]map[string]string{}
	sonic.Unmarshal(information, &packageInformation)

	// 获取所有模型信息
	data := []map[string]string{}
	if len(packageInformation) == 0 {
		for _, packageModel := range packageModelAll {
			m := map[string]string{
				"id":      packageModel.ID,
				"name":    packageModel.PackageName,
				"version": packageModel.Version,
			}
			data = append(data, m)
		}
	} else {
		for packageName, info := range packageInformation {
			m := map[string]string{
				"id":      "",
				"name":    packageName,
				"version": info["version"],
			}
			data = append(data, m)
		}
	}

	serviceV2.RestartOMC()
	res.Data = data
	res.Msg = "获取模型列表成功"
	c.JSON(http.StatusOK, res)
}

func LoadModelView(c *gin.Context) {
	/*
		# 加载单个模型
		## space_id: 用户空间id
	*/
	userSpaceId := c.GetHeader("space_id")
	username := c.GetHeader("username")
	var res DataType.ResponseData
	var item DataType.LoadingModel
	if err := c.BindJSON(&item); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var packageName, packageVersion, filePath string

	if item.Id == "" {
		// 从用户空间表获取filePath数据
		var space DataBaseModel.YssimUserSpace
		dbModel.Model(space).Where("id = ? AND username = ?", userSpaceId, username).First(&space)
		if space.ID == "" {
			res.Status = 4
			res.Err = "空间已被删除"
			c.JSON(http.StatusOK, res)
			return
		}
		information, _ := space.PackageInformation.MarshalJSON()
		packageInformation := map[string]map[string]string{}
		sonic.Unmarshal(information, &packageInformation)

		for name, info := range packageInformation {
			if name == item.Name {
				packageName = name
				packageVersion = info["version"]
				filePath = info["file"]
				break
			}
		}
	} else {
		// 从模型表获取filePath数据
		var packageModel DataBaseModel.YssimModels
		dbModel.Where("id = ? AND default_version = ? AND userspace_id IN ?", item.Id, true, []string{"0", userSpaceId}).Find(&packageModel)

		packageName = packageModel.PackageName
		packageVersion = packageModel.Version
		filePath = packageModel.FilePath
		if filePath != "" {
			pwd, _ := os.Getwd()
			filePath = pwd + "/" + filePath
		}
	}

	if ok := serviceV2.LibraryInitializationSingle(packageName, packageVersion, filePath); !ok {
		// 加载失败
		res.Status = 4
		c.JSON(http.StatusOK, res)
		return
	}

	dbModel.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, username).UpdateColumn("last_login_time", time.Now().Local().Unix())
	res.Msg = packageName
	c.JSON(http.StatusOK, res)
}

func GetIconView(c *gin.Context) {
	/*
		# 获取模型的图标信息
	*/

	var item DataType.ModelGraphicsData
	if err := c.BindJSON(&item); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res DataType.ResponseData
	data := serviceV2.GetIcon(item.ModelName, "", true)
	res.Data = data
	c.JSON(http.StatusOK, res)
}
