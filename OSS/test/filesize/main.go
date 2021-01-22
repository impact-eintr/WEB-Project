package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(stat.Size())

}
