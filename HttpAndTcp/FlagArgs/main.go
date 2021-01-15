package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	num := os.Args
	fmt.Println(num)
	var host string
	var port string
	flag.StringVar(&host, "h", "127.0.0.1", "主机地址")
	flag.StringVar(&port, "p", "8080", "监听端口")
	flag.Parse()

	addr := host + ":" + port
	fmt.Println(addr)
}
