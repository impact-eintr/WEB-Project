package main

import (
	"github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/cache"
	"github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
