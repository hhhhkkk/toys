package main

import (
	"fmt"
)

type Users struct {
	Id   int    `hhhhkkk:"id"`
	Name string `hhhhkkk:"string"`
}

func main() {
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	s := engine.NewSession()
	if _, err := s.Raw("CREATE TABLE IF NOT EXISTS Users(Id integer, Name text);").Exec(); err != nil {
		panic(err)
	}

	insert, err := s.Insert(&Users{
		Name: "hhhhkkk",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(insert)
	// _, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	// _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	// _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	// result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	// count, _ := result.RowsAffected()
	// fmt.Printf("Exec success, %d affected\n", count)
}
