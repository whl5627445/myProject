package service

import (
	"yssim-go/app/serviceType"
	"yssim-go/library/omc"
)

func GetGraphicsDataNew(modelName string) any {
	m := serviceType.ModelInstance{}
	ok := omc.OMC.ModelInstance(modelName, &m)
	if ok {
		//fmt.Println(m)
		//data, _ := sonic.Marshal(m)
		return m
	}
	return nil
}
