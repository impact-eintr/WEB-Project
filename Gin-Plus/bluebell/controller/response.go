package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 1001, // 程序中的错误码
	"msg": xx,    // 提示信息
	"data": xxx,  // 携带的数据
}

*/

type ResponseData struct {
	Code ResCode
	Msg  interface{}
	Data interface{}
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}
