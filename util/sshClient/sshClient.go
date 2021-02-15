package sshClient

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	cmdStr := fmt.Sprintf("%s@%s",c.User,c.Ip)
	if c.Port != 22 {
		cmdStr = cmdStr + fmt.Sprintf(" -p %d",c.Port)
	} else {
		cmdStr = cmdStr + fmt.Sprintf(" -p %d",22)
	}
	cmd := exec.Command("ssh",cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	return err
}


func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
			}
			break
		}
	}

	if hostKey == nil {
		return nil, errors.New(fmt.Sprintf("no hostkey for %s", host))
	}
	return hostKey, nil
}