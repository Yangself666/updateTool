package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
	"updateTool/sftp_util"
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

	// 如果远程路径不是以/结尾，添加/
	if !strings.HasSuffix(history.RemotePath, "/") {
		history.RemotePath += "/"
	}

	enablePath := util.IsEnablePath(history.RemotePath)
	if !enablePath {
		response.Fail(c, nil, "远程路径不在白名单，无法回滚，请联系管理员添加")
		return
	}

	// 开始回滚，上传文件
	// 传输文件到所有服务器
	isZipFile := util.FileIsZip(history.FileName)
	var err error
	var resultList []map[string]interface{}
	if isZipFile {
		// 这里进行解压缩的上传
		resultList, err = sftp_util.SendZipFileToAllServer(
			0,
			history.LocalPath,
			history.RemotePath)
	} else {
		resultList, err = sftp_util.SendFileToAllServer(
			0,
			history.LocalPath,
			history.RemotePath,
			history.FileName)
	}
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// 添加新的记录

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

	// 计算处理总时间
	elapsed := time.Since(start)
	response.Success(c, resultList, fmt.Sprintf("回滚至历史记录[%v]成功，耗时: %v", history.ID, elapsed))
}

// GetHistory 获取某生产地址的更新记录
func GetHistory(c *gin.Context) {
	var updateHistory = model.UpdateHistory{}
	err := c.BindJSON(&updateHistory)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	remotePath := updateHistory.RemotePath
	fileName := updateHistory.FileName
	id := updateHistory.ID
	projectId := updateHistory.ProjectId
	pathId := updateHistory.PathId

	DB := common.GetDB()
	var histories []model.UpdateHistory
	// 查询路径不为空
	if remotePath != "" {
		DB = DB.Where("remote_path like ?", remotePath+"%")
	}
	// 如果文件名称不为空，模糊搜索
	if fileName != "" {
		DB = DB.Where("file_name like ?", "%"+fileName+"%")
	}
	// 如果历史ID不为空
	if id != 0 {
		DB = DB.Where("id = ?", id)
	}
	// 如果项目ID不为空
	if projectId != 0 {
		DB = DB.Where("project_id = ?", projectId)
	}
	// 如果路径ID不为空
	if pathId != 0 {
		DB = DB.Where("path_id = ?", pathId)
	}
	DB.Order("updated_at desc").Find(&histories)
	response.Success(c, histories, "请求成功")
}
