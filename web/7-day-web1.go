package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

/**
  简单的web服务
 */
func sayhelloName(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ",strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}
func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
