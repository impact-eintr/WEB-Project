package redis

import (
	"bluebell/setting"
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() (err error) {
	addr := fmt.Sprintf("%s:%d", setting.Conf.RedisConfig.Host, setting.Conf.RedisConfig.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: setting.Conf.RedisConfig.Password,
		DB:       setting.Conf.RedisConfig.DB,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
