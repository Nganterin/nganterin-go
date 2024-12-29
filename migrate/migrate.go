package main

import (
	"nganterin-go/config"
	"nganterin-go/models/database"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		database.AllModels...
	)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
