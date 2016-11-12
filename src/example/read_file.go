package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file := "/home/liup/test.txt"
	read1(file)
	read2(file)
	read3(file)
	read4(file)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read1(file string) {
	data, err := ioutil.ReadFile(file)
	check(err)
	fmt.Println(string(data))
}

func read2(file string) {
	f, err := os.Open(file)
	check(err)
	data := make([]byte, 1024)
	buf := make([]byte, 1)
	for {
		n, err := f.Read(buf)
		//fmt.Println(n)
		//fmt.Println(string(buf))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}
	fmt.Println(string(data))

}

func read3(file string) {
	f, err := os.Open(file)
	check(err)

	r := bufio.NewReader(f)
	data := make([]byte, 1024)
	buf := make([]byte, 10)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}
	fmt.Println(string(data))

}

func read4(file string) {
	f, err := os.Open(file)
	check(err)
	data, err := ioutil.ReadAll(f)
	_ = err
	fmt.Println(string(data))
}
