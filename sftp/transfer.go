package sftp

import (
	"archive/zip"
	"fmt"
	"github.com/pkg/sftp"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
	"updateTool/common"
	"updateTool/model"
	"updateTool/util"
)

// SendFileToAllServer 发送文件到所有配置的服务器
// localFilePath	string	本地文件路径
// remotePath		string	远程文件夹路径
// remoteFileName	string	远程文件名
//
// return			error	返回异常
func SendFileToAllServer(projectId int, localFilePath string, remotePath string, remoteFileName string) ([]map[string]interface{}, error) {
	var err error
	// 获取项目关联的服务器
	DB := common.GetDB()
	serverList := make([]model.Server, 0)
	DB.Model(&model.ProjectServerCon{}).Select("servers.*").Joins("left join servers on project_server_cons.server_id = servers.id").Where("project_server_cons.project_id = ?", projectId).Find(&serverList)

	if serverList == nil || len(serverList) <= 0 {
		err = common.Error("该项目未绑定服务器，无法上传")
		return nil, err
	}
	// 返回结果集
	resultList := make([]map[string]interface{}, 0)
	// 循环关联的服务器，进行多协程的传递
	result := make(chan map[string]interface{})
	defer close(result)
	for _, server := range serverList {
		go SendFileToServer(
			server,
			localFilePath,
			remotePath,
			remoteFileName,
			result)
	}
	for i := 0; i < len(serverList); i++ {
		// 收集执行结果
		resultList = append(resultList, <-result)
	}

	return resultList, nil
}

// SendZipFileToAllServer 发送压缩文件到所有配置的服务器
// projectId		int		项目ID
// localFilePath	string	本地文件路径
// remotePath		string	远程文件夹路径
//
// return			error 	返回异常
func SendZipFileToAllServer(projectId int, localFilePath string, remotePath string) ([]map[string]interface{}, error) {
	var err error
	// 获取项目关联的服务器
	DB := common.GetDB()
	serverList := make([]model.Server, 0)
	DB.Model(&model.ProjectServerCon{}).Select("servers.*").Joins("left join servers on project_server_cons.server_id = servers.id").Where("project_server_cons.project_id = ?", projectId).Find(&serverList)

	if serverList == nil || len(serverList) <= 0 {
		err = common.Error("该项目未绑定服务器，无法上传")
		return nil, err
	}
	// 返回结果集
	resultList := make([]map[string]interface{}, 0)
	// 循环关联的服务器，进行多协程的传递
	result := make(chan map[string]interface{})
	defer close(result)
	for _, server := range serverList {
		go SendZipFileToServer(
			server,
			localFilePath,
			remotePath,
			result)
	}
	for i := 0; i < len(serverList); i++ {
		// 收集执行结果
		resultList = append(resultList, <-result)
	}
	return resultList, nil
}

// SendZipFileToServer 发送压缩文件到远程服务器
// server			model.Server	服务器信息
// localZipFilePath	string			本地zip压缩包文件路径
// remotePath 		string			远程文件夹路径
// result			chan string		结果管道
func SendZipFileToServer(server model.Server, localZipFilePath string, remotePath string, result chan map[string]interface{}) {
	var (
		client *sftp.Client
		err    error
	)
	// 计算处理开始时间
	start := time.Now()
	// 结果map
	resultMap := make(map[string]interface{}, 0)
	resultMap["result"] = false

	if server.ServerType == 1 {
		client, err = GetSftpClient(server.Username, server.Password, server.Host, server.Port)
		if err != nil {
			log.Println("["+server.ServerName+"("+server.Host+")]连接失败", err)
			resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]连接失败"
			result <- resultMap
			return
		}
		// 创建连接后首先defer进行关闭操作，防止遗忘
		defer client.Close()
	}

	// 检查远程文件夹状态
	_, errRemotePath := client.Stat(remotePath)
	if errRemotePath != nil {
		errRemotePath = client.MkdirAll(remotePath)
		if errRemotePath != nil {
			log.Println("[" + server.ServerName + "(" + server.Host + ")]远程文件路径[" + remotePath + "]不存在或权限不足")
			resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]远程文件路径[" + remotePath + "]不存在或权限不足"
			result <- resultMap
			return
		}
	}

	// 检查本地文件夹状态
	fileInfo, errLocalFilePath := os.Stat(localZipFilePath)
	if errLocalFilePath != nil {
		log.Println("本地文件路径["+localZipFilePath+"]不存在或权限不足", errLocalFilePath)
		resultMap["info"] = "本地文件路径[" + localZipFilePath + "]不存在或权限不足"
		result <- resultMap
		return
	}
	if fileInfo.IsDir() {
		log.Println("[" + localZipFilePath + "]文件路径为文件夹，无法上传")
		resultMap["info"] = "[" + localZipFilePath + "]文件路径为文件夹，无法上传"
		result <- resultMap
		return
	}

	// 路径检查没有问题，开始解压文件传输
	zipReader, err := zip.OpenReader(localZipFilePath)
	if err != nil {
		log.Println("zip文件读取失败", err)
		resultMap["info"] = "zip文件读取失败，错误信息：" + err.Error()
		result <- resultMap
		return
	}
	// 关闭zip包
	defer zipReader.Close()

	// 遍历文件
	for _, file := range zipReader.File {
		err := uploadZipFile(client, file, remotePath)
		if err != nil {
			log.Println("["+server.ServerName+"("+server.Host+")]上传失败", err)
			resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]上传失败"
			result <- resultMap
			// 上传结束
			return
		}
	}

	// 计算处理总时间
	elapsed := time.Since(start)
	fmt.Println("[" + server.ServerName + "(" + server.Host + ")]上传成功，耗时：" + elapsed.String())
	resultMap["result"] = true
	resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]上传成功，耗时：" + elapsed.String()
	result <- resultMap
	// 上传结束
	return
}

// SendFileToServer 发送文件到远程服务器
// server			model.Server	服务器信息
// localFilePath	string			本地文件路径
// remotePath		string			远程文件夹路径
// remoteFileName	string			远程文件名
// result			chan string		结果管道
func SendFileToServer(server model.Server, localFilePath string, remotePath string, remoteFileName string, result chan map[string]interface{}) {
	var (
		client *sftp.Client
		err    error
	)
	// 计算处理开始时间
	start := time.Now()

	// 结果map
	resultMap := make(map[string]interface{}, 0)
	resultMap["result"] = false

	if server.ServerType == 1 {
		client, err = GetSftpClient(server.Username, server.Password, server.Host, server.Port)
		if err != nil {
			log.Println("["+server.ServerName+"("+server.Host+")]连接失败", err)
			resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]连接失败"
			result <- resultMap
			return
		}
		// 创建连接后首先defer进行关闭操作，防止遗忘
		defer client.Close()
	}

	// 检查远程文件夹状态
	_, errRemotePath := client.Stat(remotePath)
	if errRemotePath != nil {
		errRemotePath = client.MkdirAll(remotePath)
		if errRemotePath != nil {
			log.Println("[" + server.ServerName + "(" + server.Host + ")]远程文件路径[" + remotePath + "]不存在或权限不足")
			resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]远程文件路径[" + remotePath + "]不存在或权限不足"
			result <- resultMap
			return
		}
	}

	// 检查本地文件夹状态
	fileInfo, errLocalFilePath := os.Stat(localFilePath)
	if errLocalFilePath != nil {
		log.Println("本地文件路径["+localFilePath+"]不存在或权限不足", errLocalFilePath)
		resultMap["info"] = "本地文件路径[" + localFilePath + "]不存在或权限不足"
		result <- resultMap
		return
	}
	if fileInfo.IsDir() {
		log.Println("[" + localFilePath + "]文件路径为文件夹，无法上传")
		resultMap["info"] = "[" + localFilePath + "]文件路径为文件夹，无法上传"
		result <- resultMap
		return
	}

	// 路径检查没有问题，开始文件传输
	err = uploadFile(client, localFilePath, remotePath, remoteFileName)
	if err != nil {
		log.Println("[" + server.ServerName + "(" + server.Host + ")]上传失败")
		resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]上传失败"
		result <- resultMap
		// 上传结束
		return
	}

	// 计算处理总时间
	elapsed := time.Since(start)
	log.Println("[" + server.ServerName + "(" + server.Host + ")]上传成功，耗时：" + elapsed.String())
	resultMap["result"] = true
	resultMap["info"] = "[" + server.ServerName + "(" + server.Host + ")]上传成功，耗时：" + elapsed.String()
	result <- resultMap
	// 上传结束
	return
}

// SendDirectoryToServer 发送文件夹到服务器
// user			string	服务器用户名
// password 	string	服务器密码
// host			string	服务器主机地址
// port			int		服务器端口
// localPath	string	本地文件夹路径
// remotePath	string	远程文件夹路径
func SendDirectoryToServer(user string, password string, host string, port int, localPath string, remotePath string) error {
	var (
		client *sftp.Client
		err    error
	)
	// 计算处理开始时间
	start := time.Now()

	client, err = GetSftpClient(user, password, host, port)
	if err != nil {
		log.Println(err)
		return common.Error(host + " 连接失败")
	}
	// 创建连接后首先defer进行关闭操作，防止遗忘
	defer client.Close()

	// 检查远程文件夹状态
	_, errRemotePath := client.Stat(remotePath)
	if errRemotePath != nil {
		errRemotePath = client.MkdirAll(remotePath)
		if errRemotePath != nil {
			log.Println("远程文件路径[" + remotePath + "]不存在或权限不足")
			return common.Error(host + " 远程文件路径[" + remotePath + "]不存在或权限不足")
		}
	}

	// 检查本地文件夹状态
	_, errLocalPath := os.ReadDir(localPath)
	if errLocalPath != nil {
		log.Println("本地文件路径["+localPath+"]不存在或权限不足", localPath)
		return common.Error("本地文件路径[" + localPath + "]不存在或权限不足")
	}

	// 路径检查没有问题，开始文件夹传输
	err = uploadDirectory(client, localPath, remotePath)
	if err != nil {
		return err
	}

	// 计算处理总时间
	elapsed := time.Since(start)
	fmt.Println("上传到"+host+"耗时: ", elapsed)
	return nil
}

// uploadDirectory	上传文件夹
// client		*sftp.Client	服务器连接后的client指针
// localPath	string			本地文件夹路径
// remotePath	string			远程文件夹路径
//
// return 		error			返回异常
func uploadDirectory(client *sftp.Client, localPath string, remotePath string) error {
	localFiles, err := os.ReadDir(localPath)
	if err != nil {
		log.Println("读取本地文件失败", err)
		return common.Error("读取本地文件失败")
	}
	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		if backupDir.IsDir() {
			client.Mkdir(remoteFilePath)
			uploadDirectory(client, localFilePath, remoteFilePath)
		} else {
			uploadFile(client, path.Join(localPath, backupDir.Name()), remotePath, "")
		}
	}
	return nil
}

// uploadFile	上传单个文件
// client			*sftp.Client	服务器连接后的client指针
// localFilePath	string			本地文件路径
// remotePath		string			远程文件夹路径
// remoteFileName	string			远程文件名
//
// return 			error			返回异常
func uploadFile(client *sftp.Client, localFilePath string, remotePath string, remoteFileName string) error {
	// 打开本地文件
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Println("["+localFilePath+"] 文件打开失败，本地文件不存在或权限不足", err)
		return common.Error("[" + localFilePath + "] 本地文件不存在或权限不足")
	}
	// 提前关闭文件
	defer srcFile.Close()
	if remoteFileName == "" {
		remoteFileName = path.Base(localFilePath)
	}
	dstFile, err := client.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		log.Println("["+path.Join(remotePath, remoteFileName)+"] 文件创建失败，远程文件不存在或权限不足：", err)
		return common.Error("[" + path.Join(remotePath, remoteFileName) + "] 文件创建失败，远程文件不存在或权限不足")
	}
	defer dstFile.Close()

	file, err := io.ReadAll(srcFile)
	if err != nil {
		log.Println("["+localFilePath+"] 文件读取失败，本地文件不存在或权限不足", err)
		return common.Error("[" + localFilePath + "] 文件读取失败，本地文件不存在或权限不足")
	}
	// 写入文件
	_, err = dstFile.Write(file)
	if err != nil {
		log.Println("远程文件写入失败")
		return common.Error("远程文件写入失败")
	}

	// 无错误，error 返回nil
	return nil
}

// 上传Zip文件
// uploadZipFile	上传zip压缩文件
// client		*sftp.Client	服务器连接后的client指针
// zipFile		*zip.File		zip包中遍历的File指针
// remotePath	string			远程文件夹路径
//
// return 		error			返回异常
func uploadZipFile(client *sftp.Client, zipFile *zip.File, remotePath string) error {
	zipFileInfo := zipFile.FileInfo()
	// 去除mac压缩包中的无用文件
	if util.IsMacUseless(zipFile) {
		return nil
	}
	// 拼接远程地址绝对路径
	remoteFilePath := path.Join(remotePath, zipFile.Name)
	// 如果是文件夹，直接创建文件夹
	if zipFileInfo.IsDir() {
		client.MkdirAll(remoteFilePath)
		return nil
	}
	// 提取文件文件夹，进行创建
	err := client.MkdirAll(filepath.Dir(remoteFilePath))
	if err != nil {
		log.Println("远程文件夹创建失败", err)
		return common.Error("远程文件夹创建失败")
	}

	// 打开本地文件
	srcFile, err := zipFile.Open()
	if err != nil {
		log.Println("zip文件打开失败", err)
		return common.Error("zip文件打开失败")
	}
	// 提前关闭文件
	defer srcFile.Close()

	// 创建远程文件
	dstFile, err := client.Create(remoteFilePath)
	if err != nil {
		log.Println("["+remoteFilePath+"] 文件创建失败，远程文件不存在或权限不足：", err)
		return common.Error("[" + remoteFilePath + "] 文件创建失败，远程文件不存在或权限不足")
	}
	defer dstFile.Close()

	file, err := io.ReadAll(srcFile)
	if err != nil {
		log.Println("zip文件读取失败", err)
		return common.Error("zip文件读取失败")
	}
	// 写入文件
	_, err = dstFile.Write(file)
	if err != nil {
		log.Println("远程文件写入失败")
		return common.Error("远程文件写入失败")
	}

	// 无错误，error 返回nil
	return nil
}
