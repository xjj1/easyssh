# easyssh
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fxjj1%2Feasyssh.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fxjj1%2Feasyssh?ref=badge_shield)


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


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fxjj1%2Feasyssh.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fxjj1%2Feasyssh?ref=badge_large)