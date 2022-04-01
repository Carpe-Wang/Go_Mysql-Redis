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
	//一个一个放入
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3.读取redis数据
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	fmt.Printf("操作ok r1=%v r2=%v\n", r1, r2)

	//批量放入
	_, err = conn.Do("HMSet", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMSet err=", err)
		return
	}
	//一次读取
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}
