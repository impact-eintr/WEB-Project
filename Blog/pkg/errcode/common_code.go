package errcode

const (
	_ int = iota + 10000000
	serverError
	invalidParams
	notFound
	unAuthorizedAuthNotExist
	unAuthorizedTokenError
	unAuthorizedTokenTimeout
	unAuthorizedTokenGenerate
	tooManyRequests
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(serverError, "服务器内部错误")
	InvalidParams             = NewError(invalidParams, "传参有误")
	NotFound                  = NewError(notFound, "找不到")
	UnauthorizedAuthNotExist  = NewError(unAuthorizedAuthNotExist, "鉴权失败，找不到对应的AppKey和appSecret")
	UnauthorizedTokenError    = NewError(unAuthorizedTokenError, "鉴权失败，Tockn错误")
	UnauthorizedTokenTimeout  = NewError(unAuthorizedTokenTimeout, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(unAuthorizedTokenGenerate, "健全失败，token生成失败")
	TooManyRequests           = NewError(tooManyRequests, "请求过多")
)
