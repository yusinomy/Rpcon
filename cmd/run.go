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
			Query()
		case "mssql":
			common.Port = 1433
			Mssquery()
		case "oracle":
			common.Port = 1521
			ORquery()
		case "postgresql":
			common.Port = 5432
			PoQuery()
		case "smb":
			common.Port = 445
			Smb()
		case "wmi":
			common.Port = 135
			WmiExec()
		}
	}
	if common.Redis != "" {
		switch common.Redis {
		case "rk":
			common.Port = 6379
			Wkey()
		case "rs":
			common.Port = 6379
			Wshell()
		case "rn":
			common.Port = 6379
			ncshell()
		case "mysl":
			common.Port = 3306
			Mysqlshell()
		case "mycg":
			common.Port = 3306
			Myconfig()
		case "mssg":
			common.Port = 1433
			Msscfg()
		case "cmshell":
			common.Port = 1433
			Cmdshell()
		case "oleshell":
			common.Port = 1433
			Oleshell()
		case "org":
			common.Port = 1521
			ORcfg()
		case "ors":
			common.Port = 1521
			ORacleshell()
		case "pocfg":
			common.Port = 5432
			Pocfg()
		case "poshell":
			common.Port = 5432
			Powriteshell()
		case "pocode":
			common.Port = 5432
			Pocode()
		}
	} else {
		con()
	}
}
