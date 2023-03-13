package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authController "github.com/xarick/gin-sso/controllers/auth"
	"github.com/xarick/gin-sso/middlewares"
)

func UserRoutes(route *gin.Engine) {

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"page": "/"})
	})
	route.Use(middlewares.CORSMiddleware())

	api := route.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
			auth.POST("/logout", authController.Logout)
		}

		user := api.Group("/user").Use(middlewares.AuthMiddleware())
		{
			user.GET("/get-user", authController.GetUser)
		}
	}
}
