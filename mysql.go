package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from user")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var id int
	var username string
	var password string
	for rows.Next() {
		rerr := rows.Scan(&id, &username, &password)
		if rerr == nil {
			fmt.Println("id: ", strconv.Itoa(id), "username: ", username, "password: ", password)
		}
	}
	insert_sql := "insert into user(id,username,password) values(?,?,?)"
	_, e4 := db.Exec(insert_sql, "5", "ccccc", "ddddd")
	fmt.Println(e4)
	db.Close()
}
