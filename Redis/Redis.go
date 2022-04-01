package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Println(err)
	}
}
