package config

import (
	"fmt"
	"yssim-go/app/DataBaseModel"
	"yssim-go/library/omc"
)

func init() {

	omc.OMC.SendExpressionNoParsed("clearCommandLineOptions()")
	omc.OMC.SendExpressionNoParsed("clear()")
	omc.OMC.SendExpressionNoParsed("clearVariables()")
	omc.OMC.SendExpressionNoParsed("clearProgram()")
	omc.OMC.SendExpressionNoParsed("setCommandLineOptions(\"+ignoreSimulationFlagsAnnotation=false\")")
	omc.OMC.SendExpressionNoParsed("setCommandLineOptions(\"-d=nogen,noevalfunc,newInst,nfAPI\")")
	//omc.OMC.SendExpressionNoParsed("loadModel(Modelica, {\"3.2.3\"},true,\"\",false)")
	//omc.OMC.SendExpressionNoParsed("loadModel(Complex, {\"3.2.3\"},true,\"\",false)")
	//omc.OMC.SendExpressionNoParsed("loadModel(ModelicaServices, {\"3.2.3\"},true,\"\",false)")
	//omc.OMC.SendExpressionNoParsed("loadModel(ModelicaReference, {\"\"},true,\"\",false)")
	omc.OMC.SendExpressionNoParsed("loadModel(Applications, {\"\"},true,\"\",false)")
	var userSpace DataBaseModel.YssimUserSpace
	var packageModel []DataBaseModel.YssimModels
	err := DB.Where("username = ?", USERNAME).Order("last_login_time desc").First(&userSpace).Error
	err = DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", USERNAME}, []string{"0", userSpace.ID}).Find(&packageModel).Error
	if err != nil {
		panic("模型库初始化失败")
	}
	fmt.Println("初始化标准库...")
	for _, models := range packageModel {
		res, ok := omc.OMC.SendExpressionNoParsed(fmt.Sprintf("loadModel(%s, {\"\"},true,\"\",false)", models.PackageName))
		if ok {
			fmt.Printf("初始化模型库： %s  %s", models.PackageName, res)
		} else {
			panic("模型库：" + models.PackageName + "  初始化失败")
		}

	}
	fmt.Println("标准库初始化完成")
}
