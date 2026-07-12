package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	V1()
}

func V1() {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		panic("db fail")
	}

	defer db.Close()

	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")

	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}

	rows := db.QueryRow("select Name from User limit 1")
	var name string
	if err := rows.Scan(&name); err == nil {
		log.Println(name)
	}
}
