package service

import (
	"slices"

	"yssim-go/config"
	"yssim-go/library/omc"
)

func ConvertUnits(s1, s2 string) []string {
	data := []string{}
	result := omc.OMC.ConvertUnits(s1, s2)
	switch s1 {
	case "ppm/K":
		data = []string{"true", "1000000.0", "0.0"}
		return data
	case "As":
		data = []string{"true", "1.0", "0.0"}
		return data
	case "Ah":
		data = []string{"true", "3600.0", "0.0"}
		return data
	case "mAh":
		data = []string{"true", "3600.0", "0.0"}
		return data
	case "µF":
		data = []string{"true", "1000000.0", "0.0"}
		return data
	}
	for _, i := range result {
		data = append(data, i.(string))
	}
	return data
}

func unitDictionary(unit string) []string {
	for _, v := range config.Units {
		if slices.Contains(v, unit) {
			return v
		}
	}
	return nil
}
