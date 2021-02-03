package main

import "fmt"

const (
	_ int = iota + 100
	a
	b
)

func main() {
	fmt.Println(a)
}
