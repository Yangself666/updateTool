package util

import (
	"archive/zip"
	"github.com/google/uuid"
	"strings"
)

// GetUUID 生成无横杠的UUID
func GetUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// FileIsZip 检查文件是否为Zip压缩包
func FileIsZip(fileName string) bool {
	split := strings.Split(fileName, ".")
	length := len(split)
	if length < 2 {
		return false
	}
	return "zip" == split[length-1]
}

// IsMacUseless 检查是否为Mac中的无用文件
func IsMacUseless(zipFile *zip.File) bool {
	if strings.HasPrefix(zipFile.Name, "__MACOSX/") {
		return true
	}
	if zipFile.FileInfo().Name() == ".DS_Store" {
		return true
	}
	return false
}

// SliceToString 切片转换为字符串
func SliceToString(resultList []string) string {
	var result string
	for i := 0; i < len(resultList); i++ {
		result += resultList[i]
		if i != len(resultList)-1 {
			result += "\n"
		}
	}
	return result
}

// UploadResultHandler 上传结果解析
func UploadResultHandler(resultMapList []map[string]interface{}) (bool, string) {
	// 全为false为失败，不记录日志
	// 有true为成功，记录日志
	totalResult := false
	var resultStr string
	for i := 0; i < len(resultMapList); i++ {
		resultMap := resultMapList[i]
		var flagStr string
		if resultMap["result"].(bool) {
			// 有一个成功就成功
			totalResult = true
			flagStr = "【成功】"
		} else {
			flagStr = "【失败】"
		}
		resultStr += flagStr + resultMap["info"].(string)
		if i != len(resultMapList)-1 {
			resultStr += "\n"
		}
	}
	return totalResult, resultStr
}
