package models

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:p4ssw0Rd@tcp(localhost:3306)/db_api_jwt?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Product{})
	database.AutoMigrate(&Transaction{})
	database.AutoMigrate(&User{})

	DB = database
}
