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
	token, err := logic.Login(p)
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

	data := map[string]string{"token": token}

	// 3. 返回响应
	ResponseSuccess(c, data)

}
