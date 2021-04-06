package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var env Env
	env = Init()

	db, err := gorm.Open(mysql.Open(env.DBConnectionStr), &gorm.Config{})

	log.Println(db, err)
}
