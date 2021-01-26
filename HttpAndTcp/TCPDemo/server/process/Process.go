package process

import (
	"errors"
	"io"
	"log"
	"net"
	"time"

	"TCPDemo/server/common"
	"TCPDemo/server/utils"
)

type Processor struct {
	Conn net.Conn
}

//数据中继
func (this *Processor) ServerProcessMess() error {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.PkgRead()
		if err != nil && err != io.EOF {
			log.Println("err:", err)
			return err
		}
		this.serverProcessMess(&mes)
		time.Sleep(3 * time.Second)
	}
}
func (this *Processor) serverProcessMess(mes *common.Message) (err error) {

	switch mes.Type {
	case common.LoginMesType:
		up := UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case common.LogupRegType:
		//err = up.ServerProcessRegister()

	default:
		err = errors.New("消息类型不存在")
		return
	}
	return
}
