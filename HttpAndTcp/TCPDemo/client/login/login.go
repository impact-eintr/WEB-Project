package login

import (
	"TCPDemo/client/common"
	//"bufio"
	"encoding/binary"
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"net"
	//"os"
)

//登录封装
func LogIn(user common.User) {
	conn, err := TalkToServer()
	if err != nil {
		log.Println(err)
		return
	}
	//开始登陆
	var mes common.Message
	mes.Type = common.LoginMesType

	//构建登陆信息
	var loginMes common.LoginMes
	loginMes.Uid = user.Uid
	loginMes.Pwd = user.Pwd

	//协议数据部分序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		log.Println(err)
		return
	}
	mes.Data = string(data)

	//协议序列化
	data, err = json.Marshal(mes)
	if err != nil {
		log.Println(err)
		return
	}

	pkgLen := uint32(len(data))
	lenMes := make([]byte, 4096)
	binary.BigEndian.PutUint32(lenMes[:4], pkgLen)
	log.Println(lenMes[:4])

	n, err := conn.Write(lenMes[:4])
	if n != 4 || err != nil {
		log.Println(n, "字节返回", err)
		return
	}

	n, _ = conn.Read(lenMes)
	color.Cyan("%v\n", string(lenMes[:n]))

	//buf := bufio.NewReader(os.Stdin)
	//for {
	//	line, _, _ := buf.ReadLine()
	//	if string(line) == "exit" {
	//		log.Println("退出")
	//		return
	//	}

	//	//发送
	//	_, err := conn.Write(line)
	//	if err != nil {
	//		log.Println(err)
	//	}

	//	//得到回应
	//	res := make([]byte, 4096)
	//	n, err := conn.Read(res)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	log.Printf("%v", string(res[:n]))
	//}
}

func TalkToServer() (net.Conn, error) {

	conn, err := net.Dial("tcp", "127.0.0.1:6066")
	if err != nil {
		defer conn.Close()
		return nil, err
	}

	color.Green("连接成功\n 本地地址%v 远端地址%v\n",
		conn.LocalAddr(),
		conn.RemoteAddr())
	return conn, nil
}
