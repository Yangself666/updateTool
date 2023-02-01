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
	"updateTool/util"
)

func SendFileToAllServer(localFilePath string, remotePath string, remoteFileName string) error {
	servers := util.GetServers()
	var err error
	for _, info := range servers {
		err = SendFileToServer(
			info.Username,
			info.Password,
			info.Host,
			info.Port,
			localFilePath,
			remotePath,
			remoteFileName)
		if err != nil {
			break
		}
	}
	return err
}

func SendZipFileToAllServer(localFilePath string, remotePath string) error {
	servers := util.GetServers()
	var err error
	for _, info := range servers {
		err := SendZipFileToServer(
			info.Username,
			info.Password,
			info.Host,
			info.Port,
			localFilePath,
			remotePath)
		if err != nil {
			break
		}
	}
	return err
}

// SendZipFileToServer 发送文件到服务器
func SendZipFileToServer(user string, password string, host string, port int, localZipFilePath string, remotePath string) error {
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
	fileInfo, errLocalFilePath := os.Stat(localZipFilePath)
	if errLocalFilePath != nil {
		log.Println("本地文件路径["+localZipFilePath+"]不存在或权限不足", errLocalFilePath)
		return common.Error("本地文件路径[" + localZipFilePath + "]不存在或权限不足")
	}
	if fileInfo.IsDir() {
		log.Println("[" + localZipFilePath + "]文件路径为文件夹，无法上传")
		return common.Error("[" + localZipFilePath + "]文件路径为文件夹，无法上传")
	}

	// 路径检查没有问题，开始解压文件传输
	zipReader, err := zip.OpenReader(localZipFilePath)
	if err != nil {
		log.Println("zip文件读取失败", err)
		return common.Error("zip文件读取失败")
	}
	// 关闭zip包
	defer zipReader.Close()

	// 遍历文件
	for _, file := range zipReader.File {
		uploadZipFile(client, file, remotePath)
	}

	// 计算处理总时间
	elapsed := time.Since(start)
	fmt.Println("上传到"+host+"耗时: ", elapsed)

	// 上传结束
	return nil
}

// SendFileToServer 发送文件到服务器
func SendFileToServer(user string, password string, host string, port int, localFilePath string, remotePath string, remoteFileName string) error {
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
	fileInfo, errLocalFilePath := os.Stat(localFilePath)
	if errLocalFilePath != nil {
		log.Println("本地文件路径["+localFilePath+"]不存在或权限不足", errLocalFilePath)
		return common.Error("本地文件路径[" + localFilePath + "]不存在或权限不足")
	}
	if fileInfo.IsDir() {
		log.Println("[" + localFilePath + "]文件路径为文件夹，无法上传")
		return common.Error("[" + localFilePath + "]文件路径为文件夹，无法上传")
	}

	// 路径检查没有问题，开始文件传输
	err = uploadFile(client, localFilePath, remotePath, remoteFileName)
	if err != nil {
		return err
	}

	// 计算处理总时间
	elapsed := time.Since(start)
	fmt.Println("上传到"+host+"耗时: ", elapsed)

	// 上传结束
	return nil
}

// SendDirectoryToServer 发送文件夹到服务器
func SendDirectoryToServer(user string, password string, host string, port int, localPath string, remotePath string) {
	var (
		client *sftp.Client
		err    error
	)
	// 计算处理开始时间
	start := time.Now()

	client, err = GetSftpClient(user, password, host, port)
	if err != nil {
		log.Fatalln(err)
	}
	// 创建连接后首先defer进行关闭操作，防止遗忘
	defer client.Close()

	// 检查远程文件夹状态
	_, errRemotePath := client.Stat(remotePath)
	if errRemotePath != nil {
		errRemotePath = client.MkdirAll(remotePath)
		if errRemotePath != nil {
			log.Fatalln("远程文件路径[" + remotePath + "]不存在或权限不足")
		}
	}

	// 检查本地文件夹状态
	_, errLocalPath := os.ReadDir(localPath)
	if errLocalPath != nil {
		log.Fatalln("本地文件路径[" + localPath + "]不存在")
	}

	// 路径检查没有问题，开始文件夹传输
	uploadDirectory(client, localPath, remotePath)

	// 计算处理总时间
	elapsed := time.Since(start)
	fmt.Println("上传到"+host+"耗时: ", elapsed)
}

// 上传文件夹
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

// 上传单个文件
func uploadFile(client *sftp.Client, localFilePath string, remotePath string, remoteFileName string) error {
	// 打开本地文件
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Println("["+localFilePath+"] 文件打开失败", err)
		return common.Error("[" + localFilePath + "] 文件打开失败")
	}
	// 提前关闭文件
	defer srcFile.Close()
	if remoteFileName == "" {
		remoteFileName = path.Base(localFilePath)
	}
	dstFile, err := client.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		log.Println("["+path.Join(remotePath, remoteFileName)+"] 文件创建失败：", err)
		return common.Error("[" + path.Join(remotePath, remoteFileName) + "] 文件创建失败")
	}
	defer dstFile.Close()

	file, err := io.ReadAll(srcFile)
	if err != nil {
		log.Println("["+localFilePath+"] 文件读取失败", err)
		return common.Error("[" + localFilePath + "] 文件读取失败")
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
		log.Println("["+remoteFilePath+"] 文件创建失败：", err)
		return common.Error("[" + remoteFilePath + "] 文件创建失败")
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
