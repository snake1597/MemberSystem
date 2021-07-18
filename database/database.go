package database

import (
	"fmt"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

// const (
// 	user     = "root"
// 	password = ""
// 	host     = "127.0.0.1"
// 	port     = "3306"
// 	database = "arthur"
// )

var DB *gorm.DB

func ConnectDB(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("get db failed:", err)
		return
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func CloseDB(db *gorm.DB) {

	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("get db failed:", err)
		return
	}

	sqlDB.Close()
}
