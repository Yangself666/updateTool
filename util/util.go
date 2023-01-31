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
		if strings.Index(remotePath, str) == 0 {
			result = true
			break
		}
	}
	return result
}

func FileIsZip(fileName string) bool {
	split := strings.Split(fileName, ".")
	length := len(split)
	if length < 2 {
		return false
	}
	return "zip" == split[length-1]
}

func IsMacUseless(zipFile *zip.File) bool {
	if strings.Index(zipFile.Name, "__MACOSX/") == 0 {
		return true
	}
	if zipFile.FileInfo().Name() == ".DS_Store" {
		return true
	}
	return false
}
