package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayHello")
	r.ParseForm()
	name := r.Form["name"][0]
	value := r.Form["value"][0]
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: name, Value: value, Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "set cookie name:%v value:%v", name, value)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"][0]
	cookie, _ := r.Cookie(name)

	fmt.Fprintf(w, "name:%v value:%v \n", cookie.Name, cookie.Value)
	/*
		for _, cookie := range r.Cookies() {
			fmt.Fprint(w, cookie)
		}
	*/
}

func sayFuck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayFuck")
	fmt.Fprintf(w, "go fuck yourself")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayHello")
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/fuck", sayFuck)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
