package module

import (
	"TCPDemo/client/common"
	"net"
)

type CurUser struct {
	Conn net.Conn
	common.User
}
