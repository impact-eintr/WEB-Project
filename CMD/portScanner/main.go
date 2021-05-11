package main

import (
	"fmt"
	"net"
)

func worker(port chan int, result chan int) {
	for p := range port {
		address := fmt.Sprintf("220.194.140.188:%d", p)
		fmt.Println(address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}

}

func main() {
	address := make(chan int, 20)
	result := make(chan int)

	open := []int{}
	closed := []int{}

	go func() {
		for i := 0; i < cap(address); i++ {
			go worker(address, result)
		}
	}()

	for i := 8090; i <= 8190; i++ {
		address <- i
		res := <-result
		if res > 0 {
			open = append(open, res)
		} else {
			closed = append(closed, res)
		}

	}

	close(address)
	close(result)

	fmt.Println(open)
}
