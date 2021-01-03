package main

import (
	"OSS_1.0/Objects"
	"OSS_1.0/heartbeat"
	"OSS_1.0/locate"
	"log"
	"net/http"
	"os"
)

func main() {

	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/Objects/", Objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
