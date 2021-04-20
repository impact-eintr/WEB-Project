package respcode

type RespCode int

const (
	CodeSuccess RespCode = 1000 + iota

	CodeInvalidParam

	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword

	CodeNullAuth
	CodeInvalidToken

	CodeServerBusy

	CodeNotFound
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已经存在",
	CodeUserNotExist:    "该用户不存在",
	CodeInvalidPassword: "用户密码错误",

	CodeNullAuth:     "请头中auth为空",
	CodeInvalidToken: "无效的Token",

	CodeServerBusy: "服务器繁忙，通知后端查看日志",

	CodeNotFound: "路由不存在",
}

func (rescode RespCode) Msg() string {
	return codeMsgMap[rescode]
}
