package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./view/index.html")
	t.Execute(w, "上传文件")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9090", nil)
}
