package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//ip := "1.2.3.4"
	ip := "10.10.10.10"

	db, err := sql.Open("mysql", fmt.Sprintf("user:password@tcp(%s:port)/db", ip))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connected")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}
