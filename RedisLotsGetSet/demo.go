package main

import "fmt"

func main() {
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	//2.向redis写入数据
	//批量放入
	_, err = conn.Do("MSet", "name", "tan", "address", "西安")
	if err != nil {
		fmt.Println("MSet err=", err)
		return
	}
	//一次读取
	r, err := redis.Strings(conn.Do("MGet", "name", "address"))
	if err != nil {
		fmt.Println("MGet err=", err)
		return
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}
