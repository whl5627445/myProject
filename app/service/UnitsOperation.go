package service

import "yssim-go/library/omc"

func ConvertUnits(s1, s2 string) []string {
	var data []string
	result := omc.OMC.ConvertUnits(s1, s2)
	for _, i := range result {
		data = append(data, i.(string))
	}
	return data
}
