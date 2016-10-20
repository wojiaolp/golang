package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayHello")
	fmt.Fprintf(w, "hello world")
}

func sayFuck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayFuck")
	fmt.Fprintf(w, "go fuck yourself")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/fuck", sayFuck)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
