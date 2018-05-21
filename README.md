# easyssh

wrapper for easy use of ssh :

```go

package main

import (
	"fmt"

	ssh "github.com/xjj1/easyssh"
)

func main() {
	c, err := ssh.NewSSH("1.2.3.4", "user", "pass")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	res := c.ExecCmd("ls -al")
	res += c.ExecCmd("ps axf")
	fmt.Println(res)

	if err = c.GetError(); err != nil {
		fmt.Printf("Got some errors: %v\n", err)
	}
}


```
