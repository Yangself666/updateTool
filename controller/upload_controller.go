package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"strconv"
	"time"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
	"updateTool/sftp_util"
	"updateTool/util"
)

/*
上传文件Controller
*/

// UploadFile 上传文件接口
func UploadFile(c *gin.Context) {
	projectIdStr := c.PostForm("projectId")
	pathIdStr := c.PostForm("pathId")
	if projectIdStr == "" || pathIdStr == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}
	// 项目ID
	projectId, err := strconv.Atoi(projectIdStr)
	// 路径ID
	pathId, err := strconv.Atoi(pathIdStr)

	if err != nil {
		response.Fail(c, nil, "参数格式不正确")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("文件不存在或上传文件时发生错误：", err)
		response.Fail(c, nil, "文件不存在或上传文件时发生错误")
		return
	}

	DB := common.GetDB()
	// 查询项目是否绑定服务器
	var count int64
	DB.Model(&model.ProjectServerCon{}).Where("project_id = ?", projectId).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该项目未绑定服务器，无法上传")
		return
	}

	// 查询项目ID和路径ID关联
	var projectPath = model.ProjectPath{}
	DB.Model(&model.ProjectPath{}).Where("project_id = ? and id = ?", projectId, pathId).First(&projectPath)

	if projectPath.ID == 0 {
		response.Fail(c, nil, "该项目路径不存在，请联系管理员添加")
		return
	}
	// 上传的远程路径
	remotePath := projectPath.Path
	// 包含子路径
	if projectPath.HasSubPath == 1 {
		subPath := c.PostForm("subPath")
		if subPath != "" {
			// 包含子路径，路径拼接
			remotePath = path.Join(remotePath, subPath)
		}
	}

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

	// 首先保存正在上传的历史记录
	var userId uint
	user, exists := c.Get("user")
	if exists {
		userId = user.(model.User).ID
	}
	// 上传中
	history := model.UpdateHistory{
		UserId:         userId,
		RemotePath:     remotePath,
		LocalPath:      localFilePath,
		FileName:       file.Filename,
		UniqueFileName: newUuidName,
		ProjectId:      uint(projectId),
		PathId:         uint(pathId),
		UpdateStatus:   1, // 上传中
	}
	// 将上传记录保存到数据库
	saveHistory := DB.Create(&history)
	if saveHistory.Error != nil {
		log.Println("保存上传记录时发生错误：", err)
		response.Fail(c, nil, "保存上传记录时发生错误")
		return
	}

	// 开始上传流程
	// 传输文件到所有服务器
	isZipFile := util.FileIsZip(file.Filename)
	if isZipFile {
		// 这里进行解压缩的上传
		go sftp_util.SendZipFileToAllServer(history)
	} else {
		go sftp_util.SendFileToAllServer(history)
	}
	response.Success(c, nil, fmt.Sprintf("文件正在上传中，请在历史记录中查看结果。ID:%v", history.ID))
}
