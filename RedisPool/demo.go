package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数，0没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "golang")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)
}
