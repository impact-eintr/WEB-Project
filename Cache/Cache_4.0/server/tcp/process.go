package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
)

func (s *Server) process(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("异常关闭")
			}
			return
		}
	}
}
