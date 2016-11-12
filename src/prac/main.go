package main

import (
	//"crypto/md5"
	//_ "github.com/go-sql-driver/mysql"
	"fmt"
	//"time"
	"bufio"
	"bytes"
	"io"
	"mylib/common"
	"os"
)

func main() {
	file, err := os.Open("/tmp/test.txt")
	_ = err
	defer file.Close()

	r := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := r.ReadBytes('\n')
		line = bytes.TrimRight(line, "\r\n")
		fmt.Println(string(line))
		if err != nil {
			if err != io.EOF {
				fmt.Printf("error:%d: %s\n", lino, err)
			}
			break
		}
	}

	fmt.Println()
	fmt.Println(common.MinimumInt(10, 25, 20, 1))
}
