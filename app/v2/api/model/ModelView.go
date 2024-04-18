package API

import (
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
	res.Data = map[string]any{"encryption": packageModel.Encryption, "model": service.GetModelInstance(item.ModelName)}
	c.JSON(http.StatusOK, res)
}
