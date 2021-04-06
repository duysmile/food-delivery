package main

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile(".env.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading env %s\n", err)
	}

	dsn := viper.GetString("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(db, err)
}
