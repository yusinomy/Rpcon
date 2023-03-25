package main

import (
	cmdpackage "Rpcon/cmd"
	"Rpcon/common"
)

func main() {
	common.Flag()
	cmdpackage.Parse()
}
