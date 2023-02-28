package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
	"updateTool/sftp"
	"updateTool/util"
)

/*
服务器管理Controller
*/

// AddServer 新增服务器
func AddServer(c *gin.Context) {
	var server = model.Server{}
	err := c.BindJSON(&server)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	// 检查参数
	if server.ServerName == "" || server.Host == "" || server.Port == 0 || server.ServerType == 0 {
		response.Fail(c, nil, "参数不完整")
		return
	}

	DB := common.GetDB()

	// 检查host是否添加
	var count int64
	DB.Model(&model.Server{}).Where("host = ?", server.Host).Count(&count)
	if count > 0 {
		response.Fail(c, nil, "该主机地址已存在")
		return
	}

	DB.Create(&server)
	response.Success(c, nil, "请求成功")
}

// DelServer 删除服务器
func DelServer(c *gin.Context) {
	var param = make(map[string]int, 0)
	err := c.BindJSON(&param)
	if err != nil {
		log.Println("参数接收发生错误 -> ", err)
		response.Fail(c, nil, "参数不正确")
		return
	}

	if param["id"] == 0 {
		response.Fail(c, nil, "服务器ID不能为空")
		return
	}

	DB := common.GetDB()
	var count int64
	DB.Model(&model.Server{}).Where("id = ?", param["id"]).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该服务器不存在")
		return
	}
	// 删除项目与服务器关联
	DB.Unscoped().Model(&model.ProjectServerCon{}).Delete("server_id = ?", param["id"])

	// 删除服务器
	DB.Model(&model.Server{}).Delete("id = ?", param["id"])

	response.Success(c, nil, "删除成功")
}

// EditServer 编辑项目
func EditServer(c *gin.Context) {
	var server = model.Server{}
	err := c.BindJSON(&server)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	// 检查参数
	if server.ID == 0 || server.ServerName == "" || server.Host == "" || server.Port == 0 || server.ServerType == 0 {
		response.Fail(c, nil, "参数不完整")
		return
	}

	DB := common.GetDB()
	// 查询是否存在
	var count int64
	DB.Model(&model.Server{}).Where("id = ?", server.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "修改的服务器不存在")
		return
	}
	// 更新
	DB.Updates(&server)
	response.Success(c, nil, "请求成功")
}

// GetServerList 获取所有服务器列表
func GetServerList(c *gin.Context) {
	var server = model.Server{}
	err := c.BindJSON(&server)
	if err != nil {
		log.Println("参数解析失败 -> ", err)
		response.Fail(c, nil, "参数不正确")
		return
	}

	DB := common.GetDB()
	tx := DB.Model(&model.Server{})
	var serverList []model.Server

	if server.ServerName != "" {
		tx.Where("server_name like ?", "%"+server.ServerName+"%")
	}
	if server.ID != 0 {
		tx.Where("id = ?", server.ID)
	}
	if server.Host != "" {
		tx.Where("host like ?", "%"+server.Host+"%")
	}
	if server.ServerType != 0 {
		tx.Where("server_type = ?", server.ServerType)
	}
	tx.Find(&serverList)
	response.Success(c, serverList, "请求成功")
}

// GetServerListByProjectId 根据项目ID获取服务器列表
func GetServerListByProjectId(c *gin.Context) {
	var param = make(map[string]int, 0)
	err := c.BindJSON(&param)
	if err != nil {
		log.Println("参数接收发生错误 -> ", err)
		response.Fail(c, nil, "参数不正确")
		return
	}

	if param["projectId"] == 0 {
		response.Fail(c, nil, "项目ID不能为空")
		return
	}

	// DB := common.GetDB()
	// var serverList []model.Server
	// todo 查询关联表，连表查询

	response.Success(c, nil, "删除成功")
}

// CheckServers 检查server是否可用
func CheckServers() {
	var serverStr string
	servers := util.GetServers()
	for _, item := range servers {
		client, err := sftp.GetSftpClient(item.Username, item.Password, item.Host, item.Port)
		if err != nil {
			panic(item.Host + " 无法连接，请检查该服务器的参数及配置")
		}
		// 创建连接后首先defer进行关闭操作，防止遗忘
		client.Close()
		serverStr += item.Host + " "
	}
	log.Println("所有服务器配置检查通过 [ " + serverStr + "]")
}
