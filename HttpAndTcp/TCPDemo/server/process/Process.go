package process

import (
	"errors"
	"io"
	"log"
	"net"

	"TCPDemo/server/common"
	"TCPDemo/server/utils"
)

type Processor struct {
	Conn net.Conn
}

var LoginCh = make(chan string, 2)
var RegisterCh = make(chan string, 2)

//数据中继
func (this *Processor) ServerProcessMess() error {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.PkgRead()

		if err != nil {
			if err == io.EOF {
				log.Println(<-LoginCh, <-LoginCh, "下号")
			} else {
				log.Println("err:", err)
			}

			return err
		}
		this.serverProcessMess(&mes)

	}
}
func (this *Processor) serverProcessMess(mes *common.Message) (err error) {

	switch mes.Type {
	case common.LoginMesType:
		up := UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case common.RefisterMesType:
		up := UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)

	default:
		err = errors.New("消息类型不存在")
		return
	}
	return
}
