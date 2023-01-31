package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
	"updateTool/sftp"
	"updateTool/util"
)

// Rollback 恢复某一备份到生产
func Rollback(c *gin.Context) {
	rollbackId := c.PostForm("rollbackId")
	if rollbackId == "" {
		response.Fail(c, nil, "回滚的记录ID不能为空")
		return
	}

	// 计算处理开始时间
	start := time.Now()

	DB := common.GetDB()
	history := model.UpdateHistory{}
	DB.First(&history, rollbackId)
	//添加新的记录

	otherInfo := fmt.Sprintf("%v回滚至历史记录[%v]", time.Now().Format("2006-01-02 15:04:05"), history.ID)
	newHistory := model.UpdateHistory{
		RemotePath:     history.RemotePath,
		LocalPath:      history.LocalPath,
		FileName:       history.FileName,
		UniqueFileName: history.UniqueFileName,
		OtherInfo:      otherInfo,
	}
	// 保存回滚记录
	DB.Create(&newHistory)

	// 开始回滚，上传文件
	// 传输文件到所有服务器
	isZipFile := util.FileIsZip(history.FileName)
	if isZipFile {
		// 这里进行解压缩的上传
		sftp.SendZipFileToAllServer(
			history.LocalPath,
			history.RemotePath)
	} else {
		sftp.SendFileToAllServer(
			history.LocalPath,
			history.RemotePath,
			history.FileName)
	}
	// 计算处理总时间
	elapsed := time.Since(start)
	response.Success(c, nil, fmt.Sprintf("回滚至历史记录[%v]成功，耗时: %v", history.ID, elapsed))
}

// GetHistoryByRemotePath 获取某生产地址的更新记录
func GetHistoryByRemotePath(c *gin.Context) {
	remotePath := c.PostForm("remotePath")
	if remotePath == "" {
		response.Fail(c, nil, "远程路径不能为空")
		return
	}
	fileName := c.PostForm("fileName")

	DB := common.GetDB()
	var histories []model.UpdateHistory
	dbContext := DB.Where("remote_path like ?", remotePath+"%")
	//如果文件名称不为空，模糊搜索
	if fileName != "" {
		dbContext = dbContext.Where("file_name like ?", "%"+fileName+"%")
	}
	dbContext.Find(&histories)
	response.Success(c, histories, "请求成功")
}

// GetAllHistory 获取所有更新记录
func GetAllHistory(c *gin.Context) {
	var histories []model.UpdateHistory
	DB := common.GetDB()
	DB.Find(&histories)
	response.Success(c, histories, "请求成功")
}
