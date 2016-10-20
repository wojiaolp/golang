package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./view/index.html")
	t.Execute(w, "上传文件")
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:".r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("./view/index.html")
		t.Execute(w, token)

	} else {

	}
	fmt.Fprintf(w, "upload")
}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9090", nil)
}
