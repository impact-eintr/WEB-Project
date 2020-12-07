package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	test := ""
	op := os.Args[1]
	klen := os.Args[2]
	if op == "S" {
		vlen := os.Args[3]
		k := os.Args[4]
		v := os.Args[5]
		test = op + klen + " " + vlen + " " + k + v
	} else {
		k := os.Args[3]
		test = op + klen + " " + k
	}
	serverAddr := "127.0.0.1:54321"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println([]byte(test))
	_, err = conn.Write([]byte(test))
	if err != nil {
		log.Println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	log.Println("Write to server : ", test)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		log.Println("Write to server failed:", err.Error())
		os.Exit(1)

	}

	println("reply from server=", string(reply))

	conn.Close()
}
