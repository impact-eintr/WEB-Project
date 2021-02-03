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

var RegisterCh = make(chan string, 2)

//数据中继
func (this *Processor) ServerProcessMess() error {
	loginCh := make(chan string, 2)
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.PkgRead()

		if err != nil {
			if err == io.EOF {
				log.Println(<-loginCh, <-loginCh, "下号")
				return err
			} else {
				log.Println("err:", err)
			}
		}
		this.serverProcessMess(&mes, loginCh)

	}
}
func (this *Processor) serverProcessMess(mes *common.Message, loginCh chan string) (err error) {

	switch mes.Type {
	case common.LoginMesType:
		up := UserProcess{
			Conn:    this.Conn,
			LoginCh: loginCh,
		}
		err = up.ServerProcessLogin(mes)

	case common.RegisterMesType:
		up := UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)

	case common.SmsMesType:
		up := SmsProcess{}
		up.SendGroupMes(mes)

	default:

		log.Println("call test:", mes.Type)
		err = errors.New("消息类型不存在")
		return
	}
	return
}
