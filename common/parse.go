package common

import (
	"flag"
	"os"
)

func Parse() {
	if Host == "" && Password == "" && Method == "" {
		flag.Usage()
		os.Exit(0)
	}
	if Host == "" && Password == "" {
		flag.Usage()
		os.Exit(0)
	}

}
