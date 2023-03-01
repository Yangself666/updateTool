package sftp_util

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"time"
)

// GetSftpClient 获取sftp客户端链接
func GetSftpClient(user string, password string, host string, port int) (*sftp.Client, error) {
	// 初始化属性
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // ssh.FixedHostKey(hostKey),
	}
	// connect to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	// create sftp_util client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

// GetSftpClientByKey 获取sftp客户端链接（使用privateKey）
func GetSftpClientByKey(user string, privateKey string, host string, port int) (*sftp.Client, error) {
	// 初始化属性
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	auth = append(auth, ssh.PublicKeys(signer))
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // ssh.FixedHostKey(hostKey),
	}
	// connect to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	// create sftp_util client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}
