package serviceV2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
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
	var compiledRegexpOnlyHasNum = regexp.MustCompile(`^[\d\s\t.]+$`)

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
			strs := strings.FieldsFunc(line, func(r rune) bool {
				return r == ' ' || r == '\t'
			})
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

func CopyLibFileToResources(packageName, parent string, filePath, fileName string) (bool, string) {
	pType := omc.OMC.IsPackage(packageName)
	if !pType {
		return false, ""
	}
	file, _ := os.Open(filePath)
	fileData, _ := io.ReadAll(file)
	fileSavePath := resourcesDir(packageName, parent)

	currentNodes := GetResourcesList(packageName, parent)

	filenames := strings.Split(fileName, ".")
	filenamesLength := len(filenames)
	preName := ""
	postfix := ""
	for i := 0; i < filenamesLength-1; i++ {
		preName = preName + filenames[i] + "."
	}
	preName = preName[0 : len(preName)-1]
	postfix = filenames[filenamesLength-1]

	regex := regexp.MustCompile(preName + "_copy" + "[0-9]+" + "." + postfix)
	nums := []int{}
	findItself := false
	findCopy := false

	for _, namePair := range currentNodes {
		if namePair["name"] == fileName {
			findItself = true
		}

		if namePair["name"] == preName+"_copy"+"."+postfix {
			findCopy = true
		}
		// 判断字符串是否匹配副本形式正则表达式
		if regex.MatchString(namePair["name"]) {
			strs := strings.Split(namePair["name"], "copy")
			strs = strings.Split(strs[1], ".")
			num, _ := strconv.Atoi(strs[0])
			nums = append(nums, num)
		}
	}

	newName := ""
	if !findItself {
		newName = fileName
	} else if !findCopy {
		newName = fmt.Sprintf("%s%s%s", preName+"_copy", ".", postfix)
	} else {
		// 获取待创建的副本的编号
		num := findFirstCopyNum(nums)
		newName = fmt.Sprintf("%s%d%s%s", preName+"_copy", num, ".", postfix)
	}

	return fileOperation.WriteFileByte(fileSavePath+"/"+newName, fileData), newName
}

func resourcesDir(packageName, parent string) string {
	path := GetSourceFile(packageName)
	pathList := strings.Split(path, "/")
	packagePath := pathList[:len(pathList)-1]
	packagePath = append(packagePath, "Resources")
	if parent != "" {
		packagePath = append(packagePath, parent)
	}
	resourcesPath := strings.Join(packagePath, "/")
	return resourcesPath
}

func GetSourceFile(packageName string) string {
	return omc.OMC.GetSourceFile(packageName)
}

func GetResourcesList(packageName, parent string) []map[string]string {
	resourcesPath := resourcesDir(packageName, parent)
	data, err := fileOperation.GetDirChild(resourcesPath)
	if err != nil {
		log.Println("获取Resources子节点失败，错误： ", err)
		return nil
	}
	return data
}

func findFirstCopyNum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	switch n {
	case 0:
		return 1
	case 1:
		if nums[0] == 0 || nums[0] == 1 {
			return nums[0] + 1
		}
		return 1
	default:
		if nums[0] > 1 {
			return 1
		}
		i, j := 0, 1
		for j < n {
			if nums[i]+1 == nums[j] {
				i, j = i+1, j+1
			} else {
				return nums[i] + 1
			}
		}
		return nums[i] + 1
	}
}
