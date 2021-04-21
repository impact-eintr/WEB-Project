package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

const (
	CTXUserID string = "userID"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数 参数校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeServerBusy)

		}
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)

}

/*
之前实现的Token都是Access Token 也就是访问资源接口时需要的Token
接下来引入Refresh Token 通常有效期会比较长 AccessToken 的有效期比较短
- 用户端使用用户名密码认证
- 服务端生成有效期较短的AccessToken(10min) 和有效时间较长的RefreshToken(7day)
- 客户端访问时需要认证的接口时 携带Access Token
- 如果携带Access Token访问需要认证时鉴权失败 则客户端使用RefreshToken向刷新接口申请新的Access Token
- 客户端使用新的Access Token访问需要认证的接口
*/
func LoginHandler(c *gin.Context) {
	// 获取请求参数以及参数校验

	p := new(models.ParamLogin)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}

	// 2. 业务处理
	aToken, rToken, err := logic.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)

		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)

		} else {
			ResponseError(c, CodeServerBusy)

		}
		return
	}

	data := map[string]string{"aToken": aToken, "rToken": rToken}

	// 3. 返回响应
	ResponseSuccess(c, data)

}
