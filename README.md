# easyssh

wrapper for easy use of ssh :


import (
	"fmt"

	. "github.com/xjj1/easyssh"
)

func main() {
	c, err := NewSSH(&Device{"1.2.3.4", "user", "pass"})
	if err != nil {  // always check errors
		panic(err)
	}

	res, err := c.ExecCmd("ls -al")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
