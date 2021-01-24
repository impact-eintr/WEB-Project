package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	redis "github.com/go-redis/redis/v8"
	redisgo "github.com/gomodule/redigo/redis"
	"io"
	"log"
	"net"
)

var ctx = context.Background()

func redisTest(addr string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB

	})

	err := rdb.HSet(ctx, "User01", "name", "yixingwei").Err()
	if err != nil {
		panic(err)
	}

	err = rdb.HSet(ctx, "User01", "age", 22).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.HGet(ctx, "User01", "age").Int()
	if err != nil {
		panic(err)

	}
	fmt.Printf("key:%v\n", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")

	} else if err != nil {
		panic(err)

	} else {
		fmt.Println("key2", val2)

	}
}

var pool *redisgo.Pool

func PoolTest(addr string) {
	//连接池
	pool = &redisgo.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redisgo.Conn, error) {
			return redisgo.Dial("tcp", addr)
		},
	}
	defer pool.Close()

	conn := pool.Get()
	res, err := conn.Do("hget", "User01", "name")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%v\n", string(res.([]byte)))

}

type Status int

const (
	ONLINE  Status = 1
	OFFLINE Status = 0
)

type User struct {
	Uid    string
	Name   string
	Status Status
}

func main() {

	addr := "127.0.0.1:6379"

	PoolTest(addr)

	color.Green("服务器开始监听...")

	listener, err := net.Listen("tcp", "127.0.0.1:6066")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func(conn net.Conn) {
			color.Yellow("来自%v的访问\n", conn.RemoteAddr())
			defer conn.Close()
			for {
				res := make([]byte, 4096)
				n, err := conn.Read(res)
				if err == io.EOF {
					log.Println("客户端退出")
					break
				}

				log.Println(string(res[:n]))

				response := "发送成功\n"
				conn.Write([]byte(response))
			}
		}(conn)
	}

}
