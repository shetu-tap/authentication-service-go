package db

import (
	"authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupPostgres() *gorm.DB {
	DSN_STRING := "host=localhost database=authentication-go user=postgres password=shetu2153 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(DSN_STRING), &gorm.Config{})
	if err != nil {
		panic("Can not connect to database")
	}
	db.AutoMigrate(&models.User{})
	return db
}
