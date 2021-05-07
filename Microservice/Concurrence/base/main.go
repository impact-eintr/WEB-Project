package main

import (
	"fmt"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello goroutine")
}

func main() {
	go HelloWorld()
	time.Sleep(1 * time.Second)
	fmt.Println("Now Exit!")

}
