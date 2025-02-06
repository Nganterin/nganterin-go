package main

import (
	"nganterin-go/models"
	"nganterin-go/pkg/config"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		models.AllModels...,
	)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
