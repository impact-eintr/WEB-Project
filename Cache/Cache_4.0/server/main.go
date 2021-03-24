package main

import (
	"Cache/server/cache"
	"Cache/server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
