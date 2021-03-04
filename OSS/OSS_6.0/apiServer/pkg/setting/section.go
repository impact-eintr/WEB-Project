package setting

import (
	"time"
)

type ServerSettings struct {
	RunMode string //运行模式

	RabbitmqAddr string //rabbitmq服务器地址
	EsAddr       string //es服务器地址
	ListenAddr   string //监听地址
	ListenPort   string //监听端口

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (this *Setting) ReadSection(k string, v interface{}) error {
	err := this.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
