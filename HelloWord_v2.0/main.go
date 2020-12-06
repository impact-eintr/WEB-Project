package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

//实现了ServeHTTP接口的结构就是处理器
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}

func main() {
	handler := MyHandler{}
	handler1 := HelloHandler{}
	handler2 := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/", &handler)
	http.Handle("/hello", &handler1)
	http.Handle("/world", &handler2)

	server.ListenAndServe()
}
