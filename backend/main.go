package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-sso/routes"
)

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Panic("Error loading .env file")
	// }
}

func main() {

	route := gin.Default()
	routes.UserRoutes(route)

	// db := models.Connect()
	route.Run()
	// r.Run(os.Getenv("SERVER_RUN_PORT"))
	// r.Run("172.25.0.74:8081")
}
