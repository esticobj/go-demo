package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:root@/archetype_test")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func main() {

	err := initDB()
	if err != nil {
		fmt.Printf("连接失败,err:%v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	r, err := db.Query("select user_id,nick_name,avatar_url from user")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	type User struct {
		userId    int
		nickName  string
		avatarUrl string
	}
	var rs = make([]User, 0, 10)
	for r.Next() {
		var u User
		err := r.Scan(&u.userId, &u.nickName, &u.avatarUrl)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		rs = append(rs, u)
	}
	fmt.Printf("rs: %v\n", rs)
}
