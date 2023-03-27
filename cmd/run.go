package cmdpackage

import (
	"Rpcon/common"
)

func Parse() {
	common.Parse()
	if common.Method != "" {
		switch common.Method {
		case "ssh":
			common.Port = 22
			Sshcmd()
		case "redis":
			common.Port = 6379
			Codes()
		case "mysql":
			common.Port = 3306
			mysqlcmd()
		case "wmi":
			common.Port = 135
			WmiExec()
		}
	}

	if common.Redis != "" {
		switch common.Redis {
		case "wk":
			Wkey()
		case "wshell":
			Wshell()
		case "shell":
			ncshell()
		}
	} else {
		con()
	}
}
