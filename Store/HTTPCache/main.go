package main

import "github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/cache"
import "github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/http"

func main() {
	c := New("inmemory")
	NewServer(c).Listen()
}
