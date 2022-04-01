package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	//2.向redis写入数据
	_, err = conn.Do("lpush", "herList", "no1:宋江", 30, "no2:卢俊义", 28)
	if err != nil {
		fmt.Println("lpush err=", err)
		return
	}
	//List 先进先出
	r0, err := redis.String(conn.Do("rpop", "herList"))
	if err != nil {
		fmt.Println("MGet err=", err)
		return
	}
	fmt.Println(r0)
	r1, err := redis.String(conn.Do("rpop", "herList"))
	if err != nil {
		fmt.Println("MGet err=", err)
		return
	}
	fmt.Println(r1)

}
