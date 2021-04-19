package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已经存在",
	CodeUserNotExist:    "该用户不存在",
	CodeInvalidPassword: "用户密码错误",
	CodeServerBusy:      "服务器繁忙，通知后端查看日志",
}

func (rescode ResCode) Msg() string {
	return codeMsgMap[rescode]
}
