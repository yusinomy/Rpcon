package cmdpackage

import (
	"Rpcon/common"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func Sshcmd() {

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            common.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(common.Password)}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", common.Host, common.Port)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	//执行远程命令
	//var b bytes.Buffer
	//session.Stdout = &b
	combo, err := session.CombinedOutput(common.Code)
	if err != nil {
		log.Fatal("远程执行cmd 失败", err)
	}
	//fmt.Println(b.String())
	fmt.Println(string(combo))

}

//func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
//
//	keyPath, err := homedir.Expand(kPath)
//	if err != nil {
//		log.Fatal("find key's home dir failed", err)
//	}
//	key, err := ioutil.ReadFile(keyPath)
//	if err != nil {
//		log.Fatal("ssh key file read failed", err)
//	}
//	// Create the Signer for this private key.
//	signer, err := ssh.ParsePrivateKey(key)
//	if err != nil {
//		log.Fatal("ssh key signer failed", err)
//	}
//	return ssh.PublicKeys(signer)
//}

//cmds := strings.Split(common.Code, ";")
//for _, cmd := range cmds {
//cmd = strings.TrimSpace(cmd)
//if len(cmd) == 0 {
//continue
//}
