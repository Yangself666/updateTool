package util

import (
	"archive/zip"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"strings"
	"updateTool/model"
)

// GetUUID 生成无横杠的UUID
func GetUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// GetServers 获取配置中国的服务器
func GetServers() []model.ServerInfo {
	var serverList []model.ServerInfo
	mapString := viper.Get("remote-servers").([]interface{})
	for _, item := range mapString {
		serverInfo := item.(map[string]interface{})
		info := model.ServerInfo{
			Host:     serverInfo["host"].(string),
			Port:     serverInfo["port"].(int),
			Username: serverInfo["username"].(string),
			Password: serverInfo["password"].(string),
		}
		serverList = append(serverList, info)
	}
	return serverList
}

// IsEnablePath 检测路径是否可以上传
func IsEnablePath(remotePath string) bool {
	// 获取白名单路径
	slice := viper.GetStringSlice("enable-path")
	result := false
	for _, str := range slice {
		if !strings.HasSuffix(str, "/") {
			str += "/"
		}
		if strings.HasPrefix(remotePath, str) {
			result = true
			break
		}
	}
	return result
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
	for i := 0; i < cap(resultList); i++ {
		result += resultList[i]
		if i != cap(resultList)-1 {
			result += "\n"
		}
	}
	return result
}
