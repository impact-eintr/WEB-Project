package main

import (
	cache "HTTPCache/cache"
	http "HTTPCache/http"
)

func main() {
	c := cache.New("inmemory")
	http.NewServer(c).Listen()
}
