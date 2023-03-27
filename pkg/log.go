package pkg

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var Num int64
var End int64
var Results = make(chan *string)
var Start = true
var LogSucTime int64
var LogErrTime int64
var WaitTime int64
var Silent bool
var LogWG sync.WaitGroup

func LogSuccess(result string) {
	LogWG.Add(1)
	LogSucTime = time.Now().Unix()
	Results <- &result
}

func LogError(errinfo interface{}) {
	if WaitTime == 0 {
		fmt.Printf(" %v/%v %v \n", End, Num, errinfo)
		os.Exit(0)
	} else if (time.Now().Unix()-LogSucTime) > WaitTime && (time.Now().Unix()-LogErrTime) > WaitTime {
		fmt.Printf(" %v/%v %v \n", End, Num, errinfo)
		LogErrTime = time.Now().Unix()
		os.Exit(0)
	}
}

func CheckErrs(err error) bool {
	if err == nil {
		return false
	}
	errs := []string{
		"closed by the remote host", "too many connections",
		"i/o timeout", "EOF", "A connection attempt failed",
		"established connection failed", "connection attempt failed",
		"Unable to read", "is not allowed to connect to this",
		"no pg_hba.conf entry",
		"No connection could be made",
		"invalid packet size",
		"bad connection",
	}
	for _, key := range errs {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower(key)) {
			return true
		}
	}
	return false
}
