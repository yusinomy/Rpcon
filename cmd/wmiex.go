package cmdpackage

import (
	"Rpcon/common"
	"Rpcon/pkg"
	"fmt"
	"strings"
	"time"
)

var flag = false

func WmiExec() (tmperr error) {
	starttime := time.Now().Unix()
	flag, err := pkg.Wmiexec()
	errlog := fmt.Sprintf("[-] WmiExec  %v:%v %v %v %v", common.Host, 445, common.User, common.Password, err)
	errlog = strings.Replace(errlog, "\n", "", -1)
	pkg.LogError(errlog)
	if flag == true {
		var result string
		if common.Domain != "" {
			result = fmt.Sprintf("[+] WmiExec:%v:%v:%v\\%v ", common.Host, common.Port, common.Domain, common.User)
		} else {
			result = fmt.Sprintf("[+] WmiExec:%v:%v:%v ", common.Host, common.Port, common.User)
		}
		if common.Hash != "" {
			result += "hash: " + common.Hash
		} else {
			result += common.Password
		}
		pkg.LogSuccess(result)
		return err
	} else {
		tmperr = err
		if pkg.CheckErrs(err) {
			return err
		}
		if time.Now().Unix()-starttime > (int64(len(common.User)*len(common.Password)) * common.Timeout) {
			return err
		}
	}
	if len(common.Hash) == 32 {
		return
	}
	return tmperr
}
