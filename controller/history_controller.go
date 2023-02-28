package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
	"updateTool/common"
	"updateTool/dto"
	"updateTool/model"
	"updateTool/response"
	"updateTool/sftp_util"
	"updateTool/util"
)

// Rollback 恢复某一备份到生产
func Rollback(c *gin.Context) {
	param := make(map[string]int)
	err := c.BindJSON(&param)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	rollbackId := param["id"]
	if rollbackId == 0 {
		response.Fail(c, nil, "回滚的记录ID不能为空")
		return
	}

	// 计算处理开始时间
	start := time.Now()

	DB := common.GetDB()
	history := model.UpdateHistory{}
	DB.First(&history, rollbackId)
	if history.ID == 0 {
		response.Fail(c, nil, "该历史记录不存在")
		return
	}
	// 查看保存的路径ID
	var projectPath model.ProjectPath
	DB.Model(&model.ProjectPath{}).Where("id = ?", history.PathId).First(&projectPath)

	// 路径是否合格
	flag := false
	// 如果路径存在，并且路径相同（或包含子路径）
	if projectPath.ID != 0 && strings.HasPrefix(history.RemotePath, projectPath.Path) {
		if history.RemotePath == projectPath.Path || projectPath.HasSubPath == 1 {
			// 路径相同，或者包含子路径
			flag = true
		}
	}

	var otherInfo string
	if flag {
		// 路径正确，正常回滚
		otherInfo = fmt.Sprintf("[%v]回滚至历史记录[%v]", time.Now().Format("2006-01-02 15:04:05"), history.ID)
	} else {
		// 路径错误，增加额外提示
		otherInfo = fmt.Sprintf("[%v]回滚至历史记录[%v](检测到记录中的绑定路径由[%v]变更到了[%v]，或包含子路径设置有变化。本次上传的远程路径为[%v])", time.Now().Format("2006-01-02 15:04:05"), history.ID, history.RemotePath, projectPath.Path, history.RemotePath)
	}

	// 开始回滚，上传文件
	// 传输文件到所有服务器
	isZipFile := util.FileIsZip(history.FileName)
	var resultList []map[string]interface{}
	if isZipFile {
		// 这里进行解压缩的上传
		resultList, err = sftp_util.SendZipFileToAllServer(
			int(history.ProjectId),
			history.LocalPath,
			history.RemotePath)
	} else {
		resultList, err = sftp_util.SendFileToAllServer(
			int(history.ProjectId),
			history.LocalPath,
			history.RemotePath,
			history.FileName)
	}

	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	resultBool, resultStr := util.UploadResultHandler(resultList)
	// 计算处理总时间
	elapsed := time.Since(start)
	var userId uint
	user, exists := c.Get("user")
	if exists {
		userId = user.(model.User).ID
	}

	if resultBool {
		// 添加关联关系
		history := model.UpdateHistory{
			UserId:         userId,
			RemotePath:     history.RemotePath,
			LocalPath:      history.LocalPath,
			FileName:       history.FileName,
			UniqueFileName: history.UniqueFileName,
			ProjectId:      history.ProjectId,
			PathId:         history.PathId,
			ServerInfo:     resultStr,
			OtherInfo:      otherInfo,
		}
		// 将上传记录保存到数据库
		saveHistory := DB.Create(&history)
		if saveHistory.Error != nil {
			log.Println("保存回滚记录时发生错误：", err)
			response.Success(c, resultStr, fmt.Sprintf("操作成功，但保存回滚记录时发生错误，总耗时: %v", elapsed))
			return
		}
		response.Success(c, resultStr, fmt.Sprintf("操作成功，总耗时: %v", elapsed))
		return
	}

	response.Fail(c, resultStr, fmt.Sprintf("所有服务器回滚至历史记录[%v]失败，总耗时: %v", history.ID, elapsed))
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
	tx := DB.Model(&model.UpdateHistory{})
	// 查询路径不为空
	if remotePath != "" {
		tx = tx.Where("remote_path like ?", remotePath+"%")
	}
	// 如果文件名称不为空，模糊搜索
	if fileName != "" {
		tx = tx.Where("file_name like ?", "%"+fileName+"%")
	}
	// 如果历史ID不为空
	if id != 0 {
		tx = tx.Where("id = ?", id)
	}
	// 如果项目ID不为空
	if projectId != 0 {
		tx = tx.Where("project_id = ?", projectId)
	}
	// 如果路径ID不为空
	if pathId != 0 {
		tx = tx.Where("path_id = ?", pathId)
	}
	tx.Order("updated_at desc").Find(&histories)
	var resultList = make([]dto.UpdateHistoryDto, 0)
	for _, history := range histories {
		var (
			// 用户信息
			userInfo    model.User
			userInfoDto dto.UserDto
			// 项目信息
			projectInfo model.Project
			// 路径信息
			pathInfo model.ProjectPath
		)
		historyDto := dto.ToUpdateHistoryDto(history)
		if historyDto.UserId != 0 {
			DB.Model(&model.User{}).Where("id = ?", historyDto.UserId).First(&userInfo)
		}
		if historyDto.ProjectId != 0 {
			DB.Model(&model.Project{}).Where("id = ?", historyDto.ProjectId).First(&projectInfo)
		}
		if historyDto.PathId != 0 {
			DB.Model(&model.ProjectPath{}).Where("id = ?", historyDto.PathId).First(&pathInfo)
		}
		userInfoDto = dto.ToUserDto(userInfo)
		if userInfo.ID == 0 {
			userInfoDto.Name = "未知用户"
		}
		if projectInfo.ID == 0 {
			projectInfo.ProjectName = "未知项目"
		}
		if pathInfo.ID == 0 {
			pathInfo.Path = "未知路径"
		}
		historyDto.UserInfo = userInfoDto
		historyDto.ProjectInfo = projectInfo
		historyDto.PathInfo = pathInfo
		resultList = append(resultList, historyDto)
	}
	response.Success(c, resultList, "请求成功")
}
