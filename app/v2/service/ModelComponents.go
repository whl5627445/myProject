package serviceV2

import (
	"errors"

	modelComponent "yssim-go/library/omc/component"
)

func AddComponent(ModelName, OldComponentName, NewComponentName, Origin string) (map[string]any, error) {
	data := GetIcon(OldComponentName, NewComponentName, false)
	graphics := data["graphics"].(map[string]any)
	graphics["origin"] = Origin
	graphics["name"] = NewComponentName
	extentDiagram := GetModelExtentToString(graphics["coordinateSystem"])
	data["graphics"] = graphics
	result, msg := modelComponent.AddComponent(NewComponentName, OldComponentName, ModelName, Origin, "0", extentDiagram)
	if !result {
		return nil, errors.New(msg)
	} else {
		modelComponent.SetPackageUses(OldComponentName, ModelName)
		modelComponent.ModelSave(ModelName)
	}
	return data, nil
}
