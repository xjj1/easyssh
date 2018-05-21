package easySSH

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Device struct {
	Name     string
	Username string
	Password string
}

type easySSH struct {
	*ssh.Client
}

func NewSSH(a *Device) (*easySSH, error) {
	deiviceIP := fmt.Sprintf("%s:22", a.Name)
	config := &ssh.ClientConfig{
		User: a.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(a.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", deiviceIP, config)
	if err != nil {
		return nil, err
	}
	return &easySSH{client}, nil
}

func (c *easySSH) ExecCmd(cmd string) (string, error) {
	var session *ssh.Session
	var b bytes.Buffer
	session, err := c.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: " + err.Error())
		return "", err
	}
	defer session.Close()
	session.Stdout = &b
	if err = session.Run(cmd); err != nil {
		return "", err
	} else {
		return b.String(), nil
	}
}

func (c *easySSH) ExecCmdErr(cmd string, err *error) string {
	if *err != nil {
		return ""
	}
	var session *ssh.Session
	var b bytes.Buffer
	session, *err = c.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: " + (*err).Error())
		return ""
	}
	defer session.Close()
	session.Stdout = &b
	if (*err) = session.Run(cmd); err != nil {
		return ""
	} else {
		return b.String()
	}
}
