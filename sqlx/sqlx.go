package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB //连接池

func initDB() (err error) {
	//数据库信息
	//同户名:密码@tcp(ip:port)/数据库表名字
	Dns := "root:wkp159262@tcp(localhost:3306)/sql_demo"
	db, err = sqlx.Connect("mysql", Dns)
	//处理异常
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxOpenConns(10) //设置数据库连接池最大个数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return
}

type User struct {
	Id   int `id`
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlStr := `   `
	var u User
	err = db.Get(&u, sqlStr) //通过反射机制，对应起来了
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u)

	var userList User[]
	err = db.Get(userList, sqlStr) //查询多个
	if err != nil {
		fmt.Println(err)
		return
	}
}
