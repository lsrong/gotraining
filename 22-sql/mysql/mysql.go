package mysql

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func InitDb() error {
	var err error
	dsn := "root:root123456@tcp(127.0.0.1:3306)/golang"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return nil
}

func Insert() {
	sqlStr := "insert into user(name, password,age) values (?, ?, ?)"
	w := md5.New()
	_, _ = io.WriteString(w, "123456")
	password := fmt.Sprintf("%x", w.Sum(nil))
	// fmt.Println(password)
	ret, err := DB.Exec(sqlStr, "小明", password, 20)
	if err != nil {
		fmt.Printf("Insert failed,err:%v \n", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("failed to get lastinsterid,err:%v \n", err)
		return
	}
	fmt.Printf("Insert success! id is %d \n", id)
}

func QueryRow() {
	sqlStr := "select id,`name` from `user` where id = ? "
	row := DB.QueryRow(sqlStr, 2)
	var user User
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		fmt.Printf("Scan fialed, err %v \n", err)
		return
	}
	fmt.Printf("id=%d,name=%s\n", user.Id, user.Name)

	return
}

func Update() {
	sqlStr := "update user set name=? where id =?"
	ret, err := DB.Exec(sqlStr, "golang", 1)
	if err != nil {
		fmt.Printf("Exec fialed, err:%v \n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Exec fialed, err:%v \n", err)
		return
	}
	fmt.Printf("Update database, affected rows:%d \n", affected)
}

func Delete() {
	sqlStr := "delete from user where id = ?"
	ret, err := DB.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete affected failed, err%v \n", err)
		return
	}

	fmt.Printf("Delete id = 1,affected rows: %d \n", affected)
}

func QueryMulti() {
	sqlStr := "select id,name from user where id > ?"
	rows, err := DB.Query(sqlStr, 2)
	// 必须关闭rows对象
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("query failed, err:%v \n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			fmt.Printf("Scan failed, err:%v \n", err)
			return
		}
		fmt.Printf("User:%#v \n", user)
	}
}
