package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Readfile(filename string) (string, error) {
	s := ""
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("打开文件失败")
	}
	s = fmt.Sprintf("%s", file)
	return s,err
}
