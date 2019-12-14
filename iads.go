package main

import (
	"iads/cmd"
	. "iads/lib/logging"
	"os/user"
)

func main() {
	me, err := user.Current()
	if err != nil {
		FatalPrintln(err.Error())
		return
	}

	if me != nil && me.Name != "root" {
		FatalPrintln("error: Current user must be root. Please run as root or add sudo!\nerror_code: -1")
		return
	}
	cmd.Execute()
}
