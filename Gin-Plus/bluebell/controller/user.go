package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数 参数校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
			"err": err.Error(),
		})
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}
