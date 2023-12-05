package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/coba_laravel"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&News{})
	database.AutoMigrate(&Product{})

	DB = database
}
