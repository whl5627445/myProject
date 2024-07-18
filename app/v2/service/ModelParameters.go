package serviceV2

import (
	"yssim-go/library/omc"
)

// SetComponentModifierValue 参数操作
func SetComponentModifierValue(className string, parameterValue map[string]string) bool {
	for k, v := range parameterValue {
		result := omc.OMC.SetElementModifierValue(className, k, v)
		if !result {
			return false
		}
	}
	return true
}
