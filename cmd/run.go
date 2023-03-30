package cmdpackage

import (
	"Rpcon/common"
)

func Parse() {
	common.Parse()
	if common.Method != "" {
		switch common.Method {
		case "ssh":
			Sshcmd()
		case "redis":
			Codes()
		case "mysql":
			Query()
		case "mssql":
			Mssquery()
		case "oracle":
			ORquery()
		case "postgresql":
			PoQuery()
		case "smb":
			Smb()
		case "wmi":
			common.Port = 135
			WmiExec()
		}
	}
	if common.Redis != "" {
		switch common.Redis {
		case "rk":
			Wkey()
		case "rs":
			Wshell()
		case "rn":
			ncshell()
		case "ms":
			Mysqlshell()
		case "mg":
			Myconfig()
		case "msg":
			Msscfg()
		case "xcmd":
			Cmdshell()
		case "xol":
			Oleshell()
		case "rg":
			ORcfg()
		case "res":
			ORacleshell()
		case "pog":
			Pocfg()
		case "pos":
			Powriteshell()
		case "poc":
			Pocode()
		}
	} else {
		con()
	}
}
