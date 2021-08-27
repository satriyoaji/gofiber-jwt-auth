package database

import (
	"fiber-jwt-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()  {

	connection, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3307)/fiber_jwt_auth?parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	connection.AutoMigrate(models.User{})
}