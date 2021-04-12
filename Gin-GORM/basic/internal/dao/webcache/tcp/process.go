package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
)

type result struct {
	v []byte
	e error
}

func reply(conn net.Conn, resultCh chan chan *result) {
	defer conn.Close()
	for {
		c, open := <-resultCh
		if !open {
			return
		}

		r := <-c
		e := sendResponse(r.v, r.e, conn)
		if e != nil {
			log.Println("异常关闭:", e)
			return
		}
	}
}

func (s *Server) process(conn net.Conn) {
	r := bufio.NewReader(conn)
	resultCh := make(chan chan *result, 5)

	defer conn.Close()
	go reply(conn, resultCh)

	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("异常关闭:", err)
			}
			return
		}
		if op == 'S' {
			s.set(resultCh, r)
		} else if op == 'G' {
			s.get(resultCh, r)
		} else if op == 'D' {
			s.get(resultCh, r)
		} else {
			log.Println("非法操作", op)
			return
		}
		if err != nil {
			log.Println("异常关闭:", err)
		}
	}
}

//G3 AAA
func (s *Server) get(ch chan chan *result, r *bufio.Reader) {

	c := make(chan *result)

	ch <- c

	k, e := s.readKey(r) //只读键 确认是否存在
	if e != nil {
		c <- &result{nil, e}
		return
	}
	go func() {
		v, e := s.Get(k) //读键对应的值
		c <- &result{v, e}
	}()

}

//S3 5 AAAaaaaa
func (s *Server) set(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c
	k, v, e := s.readKeyAndValue(r)
	if e != nil {
		c <- &result{nil, s.Set(k, v)}
		return
	}
	go func() {
		c <- &result{nil, s.Set(k, v)}
	}()
}

//D3 AAA
func (s *Server) del(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c
	k, e := s.readKey(r)
	if e != nil {
		c <- &result{nil, s.Del(k)}
		return
	}

	go func() {
		c <- &result{nil, s.Del(k)}
	}()
}
