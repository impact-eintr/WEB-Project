package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var salt = []byte("impact-eintr")

// MyClaims自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: "username", // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "webconsole",                               // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(salt)
}
