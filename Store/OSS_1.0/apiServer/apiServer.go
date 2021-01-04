package main

import (
	"OSS_1.0/apiServer/heartbeat"
	"OSS_1.0/apiServer/locate"
	"OSS_1.0/apiServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
