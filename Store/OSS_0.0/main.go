package main

import "log"
import "net/http"
import "os"
import "OSS_0.0/Objects"

func main() {
	http.HandleFunc("/objects/", Objects.Handler)
	log.Fatalln(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
