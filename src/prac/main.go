package main

import (
	//"crypto/md5"
	//_ "github.com/go-sql-driver/mysql"
	"fmt"
	"mylib/common"
	"os"
)

func main() {
	fmt.Println(common.CommandLineFiles(os.Args[1:]))
}
