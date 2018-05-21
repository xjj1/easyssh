# easyssh

wrapper for easy use of ssh :

c := easyssh.newSSH(easyssh.Device{1.2.3.4,user,password})

output, err := c.ExecCmd("ls -la")

