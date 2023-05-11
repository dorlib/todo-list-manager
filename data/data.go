package data

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func OpenDataBase() {
	var err error

	var dsn string

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("DB_HOST")

	// check if in container
	if _, ok := os.LookupEnv("HOSTNAME"); ok {
		host = os.Getenv("DB_CONTAINER_HOST")
	}

	dsn = user + ":" + pass + "@tcp(" + host + ":3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

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
