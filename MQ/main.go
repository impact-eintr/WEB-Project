package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- `impact`
		c <- `eintr`

	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		println(`auth: ` + <-c)
		println(`by: ` + <-c)

	}()
	wg.Wait()

}
