package main

import (
	"github.com/HttpCache/cache"
	"github.com/HttpCache/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
