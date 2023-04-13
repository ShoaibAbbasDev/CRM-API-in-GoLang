package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnToDb() {
	

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_URL := DB_USER+":"+DB_PASSWORD+"@tcp(127.0.0.1:3306)/"+DB_NAME+"?charset=utf8mb4&parseTime=True&loc=Local"


	DB, err = gorm.Open(mysql.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("unabled to connect db")
	}

}
