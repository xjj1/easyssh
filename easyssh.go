package easyssh

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

// EasySSH type alias for ssh.ClientEasySSH
type EasySSH struct {
	*ssh.Client
}

// NewSSH creates new ssh connection
func NewSSH(Name, Username, Password string) (*EasySSH, error) {
	deiviceIP := fmt.Sprintf("%s:22", Name)
	config := &ssh.ClientConfig{
		User: Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", deiviceIP, config)
	if err != nil {
		return nil, err
	}
	return &EasySSH{client}, nil
}

// ExecCmd creates new session, executes one command and returns the result as string
func (c *EasySSH) ExecCmd(cmd string) (string, error) {
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
	}
	return b.String(), nil
}

// ExecCmdErr creates new session, executes one command and returns the result as string. Return immediately if err ! nil.
// Use to execute safely multiple commands, while having only one error check at the end.
func (c *EasySSH) ExecCmdErr(cmd string, err *error) string {
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
	}
	return b.String()

}
