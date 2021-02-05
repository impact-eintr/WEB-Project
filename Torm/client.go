package Torm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

//数据库配置信息
type Settings struct {
	DriverName string

	Host     string
	Database string
	User     string
	Password string

	Options        map[string]string
	MaxOpenConns   int
	MaxIdleConns   int
	LoggongEnabled bool
}

type Client struct {
	db      *sql.DB
	session *Session
}

func NewClient(settings Settings) (c *Client, err error) {
	db
}
