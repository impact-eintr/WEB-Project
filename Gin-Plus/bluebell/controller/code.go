package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota

	CodeInvalidParam

	CodeUserExist
	CodeUserNotExist

	CodeInvalidPassword

	CodeNullAuth
	CodeInvalidAuth

	CodeInvalidToken

	CodeServerBusy

	CodeNotFound
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已经存在",
	CodeUserNotExist:    "该用户不存在",
	CodeInvalidPassword: "用户密码错误",
	CodeNullAuth:        "请头中auth为空",
	CodeInvalidAuth:     "请头中auth格式非法",
	CodeInvalidToken:    "无效的Token",
	CodeServerBusy:      "服务器繁忙，通知后端查看日志",
	CodeNotFound:        "路由不存在",
}

func (rescode ResCode) Msg() string {
	return codeMsgMap[rescode]
}
