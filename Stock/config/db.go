package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func DatabaseInit() {
	for retries := 5; retries > 0; retries-- {
		database, err = gorm.Open(mysql.Open("root:rabbitmq@tcp(localhost:4433)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

		if err != nil {
			fmt.Println("Error", err.Error())
		} else {
			fmt.Println("Initial DB Success")
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func DB() *gorm.DB {
	return database
}
