package stringOperation

import (
	"regexp"
	"strings"
	"unicode"
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

func SanitizeName(input string) string {
	// 将字符串中的中文字符和特殊符号转换为下划线
	// 定义正则表达式，匹配非英文、数字、下划线的字符
	re := regexp.MustCompile(`[^\w]`)

	// 遍历字符串，将中文字符替换为下划线
	result := []rune{}
	for _, r := range input {
		if unicode.Is(unicode.Han, r) { // 检查是否为中文字符
			result = append(result, '_')
		} else {
			result = append(result, r)
		}
	}

	// 将其他非英文、数字、下划线的字符替换为下划线
	sanitizedString := re.ReplaceAllString(string(result), "_")

	return sanitizedString
}
