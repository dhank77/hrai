package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	var err error
	dsn := "host=localhost user=postgres password=root1234 dbname=hrai port=5432 sslmode=disable TimeZone=Asia/Makassar"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
}
