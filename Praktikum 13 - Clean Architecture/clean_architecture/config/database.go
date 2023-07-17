package config

import (
	"clean_architecture/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "golang_db:golang123@tcp(golang-db.cbuoaypgqh0v.us-east-1.rds.amazonaws.com:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // ganti ke db mysql
	if err != nil {
		panic(err.Error())
	}
	migrateDB()
}

func migrateDB() {
	DB.AutoMigrate(
		model.User{},
	)
}
