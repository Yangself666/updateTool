package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
	"strings"
	"time"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
	"updateTool/sftp"
	"updateTool/util"
)

// UploadFile 上传文件接口
func UploadFile(c *gin.Context) {
	remotePath := c.PostForm("remotePath")
	if remotePath == "" {
		response.Fail(c, nil, "上传文件的远程路径不能为空")
		return
	}
	// 如果远程路径不是以/结尾，添加/
	if !strings.HasSuffix(remotePath, "/") {
		remotePath += "/"
	}
	enablePath := util.IsEnablePath(remotePath)
	if !enablePath {
		response.Fail(c, nil, "远程路径不在白名单，请联系管理员添加")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("文件不存在或上传文件时发生错误：", err)
		response.Fail(c, nil, "文件不存在或上传文件时发生错误")
		return
	}
	// 计算处理开始时间
	start := time.Now()

	// 保留原文件名和文件随机名称和相对路径
	// 保存到数据库中
	// 文件生成随机名称uuid
	newUuidName := util.GetUUID()
	// 生成年月格式日期
	formatData := time.Now().Format("2006-01")
	// 检查文件夹是否存在
	dirPath := path.Join("fileHistory", formatData)
	_, err = os.Stat(dirPath)
	if err != nil {
		// 文件夹不存在，进行创建
		os.MkdirAll(dirPath, 0755)
	}
	localFilePath := path.Join(dirPath, newUuidName)
	// 存放在文件夹中
	err = c.SaveUploadedFile(file, localFilePath)
	if err != nil {
		log.Println("上传文件时发生错误：", err)
		response.Fail(c, nil, "上传文件时发生错误")
		return
	}

	// 传输文件到所有服务器
	isZipFile := util.FileIsZip(file.Filename)
	if isZipFile {
		// 这里进行解压缩的上传
		err = sftp.SendZipFileToAllServer(
			localFilePath,
			remotePath)
	} else {
		err = sftp.SendFileToAllServer(
			localFilePath,
			remotePath,
			file.Filename)
	}

	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// 保存关联关系
	DB := common.GetDB()

	history := model.UpdateHistory{
		Model:          gorm.Model{},
		RemotePath:     remotePath,
		LocalPath:      localFilePath,
		FileName:       file.Filename,
		UniqueFileName: newUuidName,
	}

	// 将上传记录保存到数据库
	saveHistory := DB.Create(&history)
	if saveHistory.Error != nil {
		log.Println("保存上传记录时发生错误：", err)
		response.Fail(c, nil, "保存上传记录时发生错误")
		return
	}
	// 计算处理总时间
	elapsed := time.Since(start)
	response.Success(c, nil, fmt.Sprintf("上传文件成功，耗时: %v", elapsed))
}
