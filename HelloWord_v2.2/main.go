package main

import (
	"fmt"
	"net/http"
)

//HelloHandler is a http.Handler
type HelloHandler struct{}

//HelloHandler is a http.Handler
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

//传出传入不变
func log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler Called type is %T\n", handler)
		handler.ServeHTTP(w, r)
	})
}

func test(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if 1 > 2 {
			fmt.Printf("回调成功\n")
			handler.ServeHTTP(w, r)
		} else {
			fmt.Printf("回调失败\n")
		}
	})
}

func main() {
	hello := new(HelloHandler)
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//使用默认的多路复用器
	}

	http.Handle("/", test(log(hello)))
	http.Handle("/hello/", log(hello))
	http.Handle("/world", test(log(hello)))
	server.ListenAndServe()
}
