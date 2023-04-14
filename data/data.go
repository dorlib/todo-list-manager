package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDataBase() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	err = DB.AutoMigrate(&Task{})
	if err != nil {
		panic("failed to migrate database scheme")
	}
}
