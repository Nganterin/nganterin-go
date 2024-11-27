package main

import (
	"nganterin-go/config"
	"nganterin-go/models"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Client{},
		&models.Users{},
		&models.Partners{},
		&models.Tokens{},
		&models.Hotels{},
	)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
