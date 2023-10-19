package main

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/DanilaNik/IU5_RIP2023/internal/dsn"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db
	// Migrate the schema
	err = db.AutoMigrate(
		&ds.User{},
		&ds.Item{},
		&ds.Request{},
		&ds.ItemsRequest{},
	)
	if err != nil {
		panic("cant migrate db: " + err.Error())
	}
}
