package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta", ENV.DB_USER, ENV.DB_PASS, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE, "%2f")

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.

	DB = db
	log.Println("Database connected")
}
