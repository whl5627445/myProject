package stringOperation

import (
	"regexp"
	"strconv"
)

func NewName(nameList []string, matchingRule string, namePrefix string) string {

	var maxNumber int
	re := regexp.MustCompile(matchingRule)

	for _, exp := range nameList {
		matches := re.FindStringSubmatch(exp)
		if len(matches) > 1 {
			suffix, err := strconv.Atoi(matches[1])
			if err == nil && suffix > maxNumber {
				maxNumber = suffix
			}
		}
	}
	return namePrefix + " " + strconv.Itoa(maxNumber+1)

}
