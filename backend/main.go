package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-sso/initializers"
	"github.com/xarick/gin-sso/routes"
)

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Panic("Error loading .env file")
	// }
}

func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Panic("Could not load environment variables", err)
	}

	route := gin.Default()
	routes.UserRoutes(route)

	initializers.ConnectDB(&config)

	route.Run(config.SerRunPort)
}
