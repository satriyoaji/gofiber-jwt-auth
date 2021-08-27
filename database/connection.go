package database

import (
	"gofiber-jwt-auth/helper"
	"gofiber-jwt-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()  {

	connection, err := gorm.Open(mysql.Open(helper.GoDotEnvVariable("DB_USERNAME") + ":" + helper.GoDotEnvVariable("DB_PASSWORD") + "@tcp(localhost:" + helper.GoDotEnvVariable("DB_PORT") + ")/" + helper.GoDotEnvVariable("DB_NAME") + "?parseTime=True&loc=Local"), &gorm.Config{})
	helper.OutputPanicError(err)

	DB = connection

	connection.AutoMigrate(models.User{})
}