package serviceV2

import (
	"bufio"
	"io"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"strings"
	"yssim-go/library/fileOperation"
)

// 保存资源文件
func SaveResourceFile(fileHeader *multipart.FileHeader, userName, resourceFileId string) (filepath string, ok bool) {
	file, _ := fileHeader.Open()
	data, _ := io.ReadAll(file)

	filePath := "static" + "/resourceLib/" + userName + "/" + resourceFileId + "/" + fileHeader.Filename
	if ok := fileOperation.WriteFileByte(filePath, data); !ok {
		log.Println("保存资源文件时出现错误")
		return "", false
	}
	return filePath, true
}

// 获取资源文件内容
func GetResourceFileContent(path string) string {
	// 读取文件内容
	contentByte, err := os.ReadFile(path)
	if err != nil {
		log.Println("获取资源文件内容出错: ", err)
		return ""
	}

	return string(contentByte)
}

// 解析资源文件内容
func ParseResourceFileContent(path string) map[string]map[string]any {

	file, err := os.Open(path) // 打开文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保在函数结束时关闭文件

	res := map[string]map[string]any{}
	coordName := ""
	column := 0
	var xData []float64
	var yData [][]float64
	var compiledRegexpOnlyHasNum = regexp.MustCompile(`^[\d .]+$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // 逐行扫描
		line := scanner.Text()
		if strings.HasPrefix(line, "double") {
			if coordName != "" {
				res[coordName] = make(map[string]any)
				res[coordName]["xData"] = xData
				res[coordName]["yData"] = yData
			}

			strs := strings.Split(line, " ")
			headerStr := strs[1]
			strs2 := strings.Split(headerStr, "(")
			// 坐标名字
			coordName = strs2[0]

			// 提取列数字
			strs = strings.Split(line, "(")
			rowColumnStr := strs[1]
			strs = strings.Split(rowColumnStr, ")")
			rowColumnStr = strs[0]
			re := regexp.MustCompile("[0-9]+")
			nums := re.FindAllString(rowColumnStr, -1)
			column, _ = strconv.Atoi(nums[1])
			xData = []float64{}
			yData = make([][]float64, int(column-1))
			for i := 0; i < column-1; i++ {
				yData[i] = make([]float64, 0)
			}
		}

		if onlyHasSum := compiledRegexpOnlyHasNum.MatchString(line); onlyHasSum {
			strs := strings.Fields(line)
			xdata, _ := strconv.ParseFloat(strs[0], 64)
			xData = append(xData, xdata)
			for i := 1; i < column; i++ {
				ydata, _ := strconv.ParseFloat(strs[i], 64)
				yData[i-1] = append(yData[i-1], ydata)
			}
		}
	}

	res[coordName] = make(map[string]any)
	res[coordName]["xData"] = xData
	res[coordName]["yData"] = yData

	return res
}
