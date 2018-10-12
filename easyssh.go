package EasySSH

import (
	//	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

// EasySSH contains ssh.Client and error
type EasySSH struct {
	cl  *ssh.Client
	err error
}

// NewSSH creates new ssh connection
func NewSSH(Name, Username, Password string) (*EasySSH, error) {
	config := &ssh.ClientConfig{
		User: Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", Name+":22", config)
	if err != nil {
		return nil, err
	}
	return &EasySSH{client, nil}, nil
}

// ExecCmd creates new session, executes one command and returns the result as string and sets error in c
func (c *EasySSH) ExecCmd(cmd string) string {
	if c.err != nil {
		return ""
	}
	var session *ssh.Session
	///var b bytes.Buffer
	session, err := c.cl.NewSession()
	if err != nil {
		c.err = fmt.Errorf("Error creating new session %v", err)
		return ""
	}
	defer session.Close()
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // doen't make any difference?
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		c.err = fmt.Errorf("request for pseudo terminal failed: %v", err)
		return ""
	}
	//session.Stdout = &b
	var b []byte
	if b, err = session.CombinedOutput(cmd); err != nil {
		c.err = fmt.Errorf("Error executing %s : %v", cmd, err)
		return ""
	}
	//return b.String()
	return string(b)
}

// GetError returns the error
func (c *EasySSH) GetError() error {
	return c.err
}

// Close closes the connetion
func (c *EasySSH) Close() {
	c.cl.Close()
}
