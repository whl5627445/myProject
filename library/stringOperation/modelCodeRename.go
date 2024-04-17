package stringOperation

import (
	"regexp"
	"strings"
)

func ModelRenam(modelCode, oldName, NewName string) string {
	// 暂时只支持model class block record package end关键字

	// 正则表达式 修改模型名字
	re := regexp.MustCompile(`\s*model\s+([A-Za-z0-9_]+)`)
	matches := re.FindAllStringSubmatch(modelCode, -1)

	for _, match := range matches {
		if len(match) > 1 && match[1] == oldName {
			modelCode = strings.ReplaceAll(modelCode, match[0], strings.Replace(match[0], oldName, NewName, 1))
		}
	}

	re = regexp.MustCompile(`\s*class\s+([A-Za-z0-9_]+)`)
	matches = re.FindAllStringSubmatch(modelCode, -1)

	for _, match := range matches {
		if len(match) > 1 && match[1] == oldName {
			modelCode = strings.ReplaceAll(modelCode, match[0], strings.Replace(match[0], oldName, NewName, 1))
		}
	}

	re = regexp.MustCompile(`\s*block\s+([A-Za-z0-9_]+)`)
	matches = re.FindAllStringSubmatch(modelCode, -1)

	for _, match := range matches {
		if len(match) > 1 && match[1] == oldName {
			modelCode = strings.ReplaceAll(modelCode, match[0], strings.Replace(match[0], oldName, NewName, 1))
		}
	}

	re = regexp.MustCompile(`\s*record\s+([A-Za-z0-9_]+)`)
	matches = re.FindAllStringSubmatch(modelCode, -1)

	for _, match := range matches {
		if len(match) > 1 && match[1] == oldName {
			modelCode = strings.ReplaceAll(modelCode, match[0], strings.Replace(match[0], oldName, NewName, 1))
		}
	}

	re = regexp.MustCompile(`\s*package\s+([A-Za-z0-9_]+)`)
	matches = re.FindAllStringSubmatch(modelCode, -1)

	for _, match := range matches {
		if len(match) > 1 && match[1] == oldName {
			modelCode = strings.ReplaceAll(modelCode, match[0], strings.Replace(match[0], oldName, NewName, 1))
		}
	}

	re = regexp.MustCompile(`\s*end\s+([A-Za-z0-9_]+)`)
	matches = re.FindAllStringSubmatch(modelCode, -1)

	for _, match := range matches {
		if len(match) > 1 && match[1] == oldName {
			modelCode = strings.ReplaceAll(modelCode, match[0], strings.Replace(match[0], oldName, NewName, 1))
		}
	}
	return modelCode
}
