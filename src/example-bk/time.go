package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time.Now():", time.Now())
	fmt.Println("time.Now().Unix():", time.Now().Unix())
}
