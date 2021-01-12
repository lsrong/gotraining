package main

import (
	"fmt"

	"github.com/learning_golang/22-sql/mysql"
)

func main() {
	err := mysql.InitDb()
	if err != nil {
		fmt.Printf("Mysql init db failed, err:%v", err)
	}
	mysql.Insert()
	mysql.QueryRow()
	mysql.Update()
	mysql.Delete()
	mysql.QueryMulti()

}
