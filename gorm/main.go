package main

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var do sync.Once

func main() {
	db := GetDB()

	Run(db)
}

func GetDB() *gorm.DB {
	var db *gorm.DB
	do.Do(func() {
		nd, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

		if err != nil {
			panic("connect fail" + err.Error())
		}
		db = nd
	})
	return db
}
