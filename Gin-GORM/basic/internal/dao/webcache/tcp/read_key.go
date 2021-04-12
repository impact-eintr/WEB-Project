package tcp

import (
	"bufio"
	"io"
)

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
	klen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}

	vlen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}

	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", nil, e

	}
	v := make([]byte, vlen)
	_, e = io.ReadFull(r, v)
	if e != nil {
		return "", nil, e

	}
	return string(k), v, nil
}
