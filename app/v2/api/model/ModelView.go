package APIv2

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/v2/service"
	"yssim-go/config"
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
