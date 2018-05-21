# easyssh

wrapper for easy use of ssh :

```go

import (
	"fmt"

	. "github.com/xjj1/easyssh"
)

func main() {
	c, err := NewSSH("1.2.3.4", "user", "pass")
	if err != nil {  
		panic(err)
	}

	res, err := c.ExecCmd("ls -al")
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(res)
}

```
