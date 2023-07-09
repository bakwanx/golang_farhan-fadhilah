package config

import (
	"restfull_test/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "golang_db:golang123@tcp(golang-db.cbuoaypgqh0v.us-east-1.rds.amazonaws.com:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Book{})
}
