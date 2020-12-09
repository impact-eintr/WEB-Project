package main

import (
	"fmt"
	"log"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	Len := r.ContentLength
	body := make([]byte, Len)
	r.Body.Read(body)
	log.Println(string(body))
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", body)
	server.ListenAndServe()
}
