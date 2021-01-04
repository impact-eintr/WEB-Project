package main

import (
	"OSS_0.0/Objects"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/objects/", Objects.Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
