package main

import (
	"OSS_0.0/Objects"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/objects/", Objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
