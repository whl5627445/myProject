package service

import (
	"yssim-go/app/serviceType"
	"yssim-go/library/omc"
)

func GetGraphicsDataNew(modelName string) interface{} {
	m := serviceType.ModelInstance{}
	ok := omc.OMC.ModelInstance(modelName, &m)
	if ok {
		//fmt.Println(m)
		//data, _ := json.Marshal(m)
		return m
	}
	return nil
}
