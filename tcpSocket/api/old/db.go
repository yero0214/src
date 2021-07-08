package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func conn() *gorm.DB {
	dsn := "host=localhost user=postgres password=365365 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	return db
}
