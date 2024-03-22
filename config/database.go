package config

import (
	"fmt"
	"gin-gorm/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "root"
	password = "secret"
	dbName   = "test"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
