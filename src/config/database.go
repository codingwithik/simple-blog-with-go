package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"log"

	"github.com/codingwithik/simple-blog-backend-with-go/src/models"
	"gorm.io/gorm"
)

var dB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

	dB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database:", err)
	}
	log.Println("? Connected Successfully to the Database")
}

func setDB(db *gorm.DB) {
	dB = db
}

func DB() *gorm.DB {
	return dB
}

func Migrate() error {
	return dB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	//log.Println("Database Migration Completed!")
}
