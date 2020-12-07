package tcp

import (
	"TCPCache/server/cache"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", ":54321")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go s.process(c)
	}
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("close connection due to error: ", err)
			}
			return
		}
		//S3 5 AAAaaaaa
		if op == 'S' {
			err = s.set(conn, r)
		} else if op == 'G' { //G 3 AAA
			err = s.get(conn, r)
		} else if op == 'D' { //D 3 AAA
			err = s.del(conn, r)
		} else {
			log.Println("close connection due to invalid operation: ", op)
			return
		}

		if err != nil {
			log.Println("close connection due to error:", err)
			return
		}
	}
}

func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	k, err := s.readKey(r)
	if err != nil {
		return err
	}
	v, err := s.Get(k)
	return sendResponse(v, err, conn)
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	k, v, err := s.readKeyAndValue(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Set(k, v), conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	k, err := s.readKey(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Del(k), conn)
}

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		temp := fmt.Sprintf("-%d ", len(errString)) + errString //err
		_, e := conn.Write([]byte(temp))
		return e
	}
	vlen := fmt.Sprintf("%d ", len(value))
	_, e := conn.Write(append([]byte(vlen), value...))
	return e
}

func (s *Server) readKey(r *bufio.Reader) (string, error) {
	klen, err := readLen(r)
	if err != nil {
		return "", err
	}
	k := make([]byte, klen)
	_, err = io.ReadFull(r, k)
	if err != nil {
		return "", err
	}
	return string(k), nil

}

func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	klen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	vlen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	k := make([]byte, klen)
	_, err = io.ReadFull(r, k)
	if err != nil {
		return "", nil, err
	}
	v := make([]byte, vlen)
	_, err = io.ReadFull(r, v)
	if err != nil {
		return "", nil, err
	}
	return string(k), v, nil
}

func readLen(r *bufio.Reader) (int, error) {
	temp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	l, err := strconv.Atoi(strings.TrimSpace(temp))
	if err != nil {
		return 0, err
	}
	return l, nil
}
