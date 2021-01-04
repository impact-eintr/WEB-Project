package main

import (
	"OSS_1.0/dataServer/Objects"
	"OSS_1.0/dataServer/heartbeat"
	"OSS_1.0/dataServer/locate"
	"log"
	"net/http"
	"os"
)

func main() {

	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/Objects/", Objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
