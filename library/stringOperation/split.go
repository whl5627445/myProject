package stringOperation

import (
	"strings"
)

func PluralSplit(s string, sep []string) []string {
	strData := s
	for i := 0; i < len(sep); i++ {
		strData = strings.ReplaceAll(strData, sep[i], "0x86")
	}
	return strings.Split(strData, "0x86")
}
