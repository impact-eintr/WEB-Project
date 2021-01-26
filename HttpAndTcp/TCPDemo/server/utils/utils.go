package utils

import (
	"TCPDemo/server/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (this *Transfer) PkgRead() (mes common.Message, err error) {
	//读取对端发送的信息长度
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		return
	}

	//解析信息长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//读取对端发送的信息本体
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || (err != nil && err != io.EOF) {
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		return
	}
	return
}

func (this *Transfer) PkgWrite(data []byte) (err error) {

	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		return
	}

	n, err = this.Conn.Write(data)

	if n != len(data) || (err != nil && err != io.EOF) {
		err = fmt.Errorf("写入内容出错%v", err)
		return
	}

	return
}
