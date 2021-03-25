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
				log.Println("异常关闭:", err)
			}
			return
		}
		if op == 'S' {
			err = s.set(conn, r)
		} else if op == 'G' {
			err = s.get(conn, r)
		} else if op == 'D' {
			err = s.get(conn, r)
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
func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	k, e := s.readKey(r) //只读键 确认是否存在
	if e != nil {
		return e
	}

	v, e := s.Get(k)                //读键对应的值
	return sendResponse(v, e, conn) //将值作为响应返回

}

//S3 5 AAAaaaaa
func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	k, v, e := s.readKeyAndValue(r)
	if e != nil {
		return e
	}

	return sendResponse(nil, s.Set(k, v), conn)

}

//D3 AAA
func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	k, e := s.readKey(r)
	if e != nil {
		return e
	}

	return sendResponse(nil, s.Del(k), conn)

}
