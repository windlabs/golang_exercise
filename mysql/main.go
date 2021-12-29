package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "go_user:go_user@tcp(127.0.0.1:3306)/go_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%v invalid, err:%v.", dsn, err)
		return
	}
	rows, err := db.Query("select * from user limit 1")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(rows)
}
