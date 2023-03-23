package stringOperation

import (
	"regexp"
	"strconv"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
)

func NewAnotherName(Username, SimulateModelName, UserspaceId string) string {
	var anotherNameList []string
	var recordList []DataBaseModel.YssimSimulateRecord
	config.DB.Where("username = ? AND simulate_model_name = ? AND userspace_id = ? AND simulate_status = ?", Username, SimulateModelName, UserspaceId, "4").Order("create_time desc").Find(&recordList)
	for i := 0; i < len(recordList); i++ {
		anotherNameList = append(anotherNameList, recordList[i].AnotherName)
	}
	var maxSuffix int
	re := regexp.MustCompile(`\d+$`)
	for _, s := range anotherNameList {
		matches := re.FindStringSubmatch(s)
		if len(matches) > 0 {
			suffix, _ := strconv.Atoi(matches[0])
			if suffix > maxSuffix {
				maxSuffix = suffix
			}
		}
	}
	return "Result" + strconv.Itoa(maxSuffix+1)

}
