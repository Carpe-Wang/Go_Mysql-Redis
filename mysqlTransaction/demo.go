package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB //连接池

func initDB() (err error) {
	//数据库信息
	//同户名:密码@tcp(ip:port)/数据库表名字
	Dns := "root:wkp159262@tcp(localhost:3306)/sql_demo"
	db, err = sql.Open("mysql", Dns) //driverName连接驱动设定，注意这个方法校验的是输入的String格式是否正常
	//处理异常
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping() //尝试和数据库连接
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxOpenConns(10) //设置数据库连接池最大个数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return
}

type User struct {
	id   int `id`
	name string
	age  int
}

func Transaction() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tx.Prepare()
	//简单的事务，将id为1的User的年龄给id为2两岁
	sqlStr := `updata sql_demo set age=age-2 where id = 1
`
	sqlStr1 := `updata sql_demo set age=age+2 where id =2`
	//tx.Exec(sqlStr,sqlStr1)可以这样写，但是为了防止有错误，我们可以有效回滚，我们采用
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败")
			return
		}
		fmt.Println("开始回滚")
	}
	_, err = tx.Exec(sqlStr)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败")
			return
		}
		fmt.Println("开始回滚")
	}
	err = tx.Commit() //提交事务
	if err != nil {
		fmt.Println("事务提交失败")
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("成功连接mysql")
	}
	Transaction()
}
