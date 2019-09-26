package main

import (
	"net/http"
	"fmt"
)

/**
  Request：用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息
  Response：服务器需要反馈给客户端的信息
  Conn：用户的每次请求链接
  Handler：处理请求和生成返回信息的处理逻辑
 */

 /**
   调用http.ListenAndServe(":8090",handler)设置监听端口

1、实例化Server

2、调用Server的ListenAndServe().

3、调用net.Listen("tcp",addr)监听端口。

4、启动一个for循环，在循环体中Accept请求

5、对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务 go c.serve()

6、读取每个请求的内容w,err:=c.readRquest()

7、判断handler是否为空，如果没有设置handler，handler就设置为defaultServeMux

8、调用handler的ServeHttp

9、根据request选择handler并且进入到这个handler的ServeHTTP

10、选择handler
  */

type MyMux struct {
}

func (p *MyMux)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/"{
		sayHelloName(w,r)
		return
	}
	http.NotFound(w,r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "hello myroute")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8002", mux)
}
