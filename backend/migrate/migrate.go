package main

import (
	"fmt"
	"log"

	"github.com/xarick/gin-sso/initializers"
	"github.com/xarick/gin-sso/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Panic("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("Migration complete")
}
