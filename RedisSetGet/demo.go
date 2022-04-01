package main

import (
	"fmt"
	"github.com/go-redis/redis"
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
	_, err = conn.Do("Set", "name", "tan")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3.读取redis数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	fmt.Println("操作ok", r)

}
