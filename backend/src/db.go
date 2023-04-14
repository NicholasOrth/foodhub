package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func dbInit() {
	dsn := os.Getenv("DB_URL")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Like{})
	if err != nil {
		log.Println(err)
	}
}
