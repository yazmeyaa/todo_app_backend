package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDBConfig() *gorm.DB {
	fmt.Println("[CONFIG][DATABASE]: Start initiating database")
	const (
		filename = "database.db"
	)
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("[DATABASE]: Cannot initialize database: %s", err.Error()))
	}

	fmt.Println("[CONFIG][DATABASE]: Database file initiated")
	return db
}
