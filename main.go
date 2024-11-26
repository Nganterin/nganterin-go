package main

import (
	"log"
	"nganterin-go/config"
	"nganterin-go/routers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitEnvCheck()

	port := os.Getenv("PORT")
	environment := os.Getenv("ENVIRONMENT")

	r := gin.New()
	r.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	api := r.Group("/v1")
	routers.CompRouter(api)

	if environment == "production" {
		host := "0.0.0.0"
		server := host + ":" + port
		err := r.Run(server)
		if err != nil {
			log.Fatal("Error starting the server: ", err)
		}
	} else if environment == "development" {
		host := "localhost"
		server := host + ":" + port
		err := r.Run(server)
		if err != nil {
			log.Fatal("Error starting the server: ", err)
		}
	} else {
		log.Fatal("ENV ERROR: {ENVIRONMENT} UNKNOWN")
	}

	log.Println("Server started on port :" + port)
}
