package common

const (
	LoginMesType            = "LoginMes"
	LoginResType            = "LoginRes"
	RegisterMesType         = "RegisterMes"
	RegisterResType         = "RegisterRes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

type Message struct {
	Type string `json:"type"` //类型
	Data string `json:"data"` //内容
}

//客户端发送消息
type LoginMes struct {
	Uid   string `json:"uid"`
	Pwd   string `json:"pwd"`
	Uname string `json:"uname"`
}

//服务器返回消息
type LoginRes struct {
	Code int `json:"code"` //返回状态码
	//500 未注册
	//200 登陆成功

	Error   string   `json:"error"` //返回错误消息
	Uid     string   `json:"uid"`   //返回用id
	Uname   string   `json:"uname"`
	UsersId []string `json:"usersid"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterRes struct {
	Code int `json:"code"` //返回状态码
	//400 已占用
	//200 注册成功

	Error string `json:"error"` //返回错误消息
	Uid   string `json:"uid"`   //返回用户id
	Uname string `json:"uname"`
}

type NotifyUserStatusMes struct {
	Uid     string `json:"uid"`
	UStatus Status `json:"ustatus"`
}

type SmsMes struct {
	Content string `json:"content"`
	User
}
