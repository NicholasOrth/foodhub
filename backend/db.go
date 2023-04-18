package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func dbInit(useSqlite bool) {
	dsn := os.Getenv("DB_URL")
	var err error

	if useSqlite {
		db, err = gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Like{}, &Follow{})
	if err != nil {
		log.Println(err)
	}
}
