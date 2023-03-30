package cmdpackage

import (
	"Rpcon/common"
	"fmt"
	common2 "github.com/Amzza0x00/go-impacket/pkg/common"
	DCERPCv5 "github.com/Amzza0x00/go-impacket/pkg/dcerpc/v5"
	"github.com/Amzza0x00/go-impacket/pkg/smb/smb2"
	"github.com/Amzza0x00/go-impacket/pkg/util"
	"os"
)

// 1.查找可用共享目录
// 2.上传文件
// 3.打开远程服务
// 4.创建服务并启动

var service string

func init() {

}

func Smb() {
	debug := false
	option := common.ClientOption{
		Host:     common.Host,
		Port:     445,
		Domain:   common.Domain,
		User:     common.User,
		Password: common.Password,
		Hash:     common.Hash,
	}
	session, err := smb2.NewSession(common2.ClientOptions(option), debug)
	if err != nil {
		os.Exit(0)
	}
	defer session.Close()
	if session.IsAuthenticated {
	}
	var serviceName string
	if service == "" {
		serviceName = string(util.Random(4))
	} else {
		serviceName = service
	}
	rpc, _ := DCERPCv5.SMBTransport()
	rpc.Client = *session
	// 创建服务并启动
	servicename, _, _ := rpc.ServiceInstall(serviceName, common.File, common.Path)
	fmt.Printf("[+] Service name is [%s]\n", servicename)
}
