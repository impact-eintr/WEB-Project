package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func readLen(r *bufio.Reader) (int, error) {
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}

	l, err := strconv.Atoi(strings.TrimSpace(tmp))
	if err != nil {
		return 0, err
	}
	return l, nil
}

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d ", len(errString)) + errString
		_, err = conn.Write([]byte(tmp))
		if err != nil {
			return err
		}
	}

	vlen := fmt.Sprintf("%d ", len(value))
	resp := append([]byte(vlen), value...)
	_, err = conn.Write(resp)
	return err
}
