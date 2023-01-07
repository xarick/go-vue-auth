package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Panic("Error loading .env file")
	// }
}

func main() {

	// db := models.Connect()
	// r := routes.ApiRoutes(db)
	// r.Run(os.Getenv("SERVER_RUN_PORT"))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	r.Run()
	// r.Run("172.25.0.74:8081")
}
