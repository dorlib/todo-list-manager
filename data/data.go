package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func OpenDataBase() {
	var err error

	var dsn string

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	dsn = user + ":" + pass + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	DB := DB.Session(&gorm.Session{CreateBatchSize: 1000})

	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	err = DB.AutoMigrate(&Task{}, &Date{}, &User{}, &Group{})
	if err != nil {
		panic("failed to migrate database scheme")
	}
}

// BeforeSave is a gorm hook in order to initiate the deadline field.
func (t *Task) BeforeSave(tx *gorm.DB) error {
	t.Deadline = fmt.Sprintf("%s/%s/%s", t.DeadlineDate.DeadlineDay, t.DeadlineDate.DeadlineMonth, t.DeadlineDate.DeadlineYear)

	return nil
}
