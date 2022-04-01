package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type User struct {
	Username string
	Password string
}

const (
	username = "root"
	password = "wkp159262"
	ip       = "127.0.0.1"
	port     = "3306"
	dbname   = "loginserver"
) //设置连接数据库的密码信息

var db *sql.DB //声明数据库连接池
func main() {
	connectMysql()
	//insertUser(User{
	//	Username: "wangkaipeng",
	//	Password: "demo01",
	//})
}
func connectMysql() {
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, " )", dbname, "?charset=utf-8"}, "")
	db, _ := sql.Open("mysql", path) //driverName, dataSourceName string
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	//验证是否连接成功
	if err := db.Ping(); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("Connection successful")
}
func insertUser(user User) bool { //updata，delete，select操作都一直，区别为Prepare的区别
	tx, err := db.Begin() //开启事物
	if err != nil {
		fmt.Println("err:", err)
		return false
	}
	//准备sql语句
	statement, err1 := tx.Prepare("insert into loginserve (`name`,`password`) value (?,?)")
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	res, err2 := statement.Exec(user.Username, user.Password)
	if err2 != nil {
		fmt.Println("exec err")
		fmt.Println("err2:", err2)
		return false
	}
	tx.Commit() //提交事务
	fmt.Println(res.LastInsertId())
	return true
}
