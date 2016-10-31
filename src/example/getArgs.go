package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("use os.Args:")
	fmt.Println(os.Args)
	fmt.Println()
	//使用flag来操作命令行参数，支持的格式如下：
	//-id=1
	//--id=1
	//-id 1
	//--id 1
	fmt.Println("use flag:")
	ok := flag.Bool("ok", false, "is ok")
	id := flag.Int("id", 0, "id")
	port := flag.String("port", ":8080", "http listen port")
	var name string
	flag.StringVar(&name, "name", "liupeng", "your name")
	flag.Parse()
	fmt.Println("ok:", *ok)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)

}
