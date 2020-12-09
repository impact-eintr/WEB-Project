package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	test := ""
	op := os.Args[1]
	k := os.Args[2]
	klen := strconv.Itoa(len(k))

	if op == "S" {
		v := os.Args[3]
		vlen := strconv.Itoa(len(v))
		test = op + klen + " " + vlen + " " + k + v

	} else if op == "G" || op == "D" {
		test = op + klen + " " + k
	} else {
		log.Println("Usage: ")
		os.Exit(1)
	}
	serverAddr := "127.0.0.1:54321"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)

	}
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

	fmt.Printf("reply from server:\n%v\n", string(reply))

}
