package sshClient

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"os/exec"
)

type SSHClient struct {
	Ip string
	Port int
	Passwd string
	User string
}

func (c *SSHClient)Test() bool {
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Passwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d",c.Ip,c.Port), config)
	if err != nil {
		log.Println("Failed to dial: ", err)
		return false
	}
	defer client.Close()
	return true
}


func (c *SSHClient)ConnectByOSCmd() error {
	cmdStr := fmt.Sprintf("ssh %s@%s",c.User,c.Ip)
	if c.Port != 22 {
		cmdStr = cmdStr + fmt.Sprintf(" -p %d",c.Port)
	}
	cmd := exec.Command(cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Wait()
	return err
}

