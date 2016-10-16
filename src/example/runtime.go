package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	name := shell1(1)
	fmt.Println(name)
}

func shell1(steps int) string {
	return caller(steps)
}

func caller(steps int) string {
	name := "?"
	if pc, _, _, ok := runtime.Caller(steps); ok {
		name = filepath.Base(runtime.FuncForPC(pc).Name())
	}
	return name
}
