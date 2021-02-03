package process

import (
	"fmt"
)

func SayHello() {
	fmt.Println("Hello World")
}

func SayGoodbye() {
	fmt.Println("Hello,")
	fmt.Println("goodbye")
}

func PrintNames() {
	students := make(map[int]string, 4)
	students[1] = "Jim"
	students[2] = "Bob"
	students[3] = "Tom"
	students[4] = "Sue"
	for _, value := range students {
		fmt.Println(value)
	}
}
