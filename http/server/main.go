package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// str := "首页"
	html, err := ioutil.ReadFile("./html/index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(html)
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(r.URL.Query())
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/test", f2)
	http.ListenAndServe(net.JoinHostPort("127.0.0.1", "8080"), nil)
}
