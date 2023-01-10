package routes

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/xarick/gin-sso/controllers/admin"
	"github.com/xarick/gin-sso/middlewares"
)

func AdminRoutes(route *gin.Engine) {

	api := route.Group("/api")
	{
		admin := api.Group("/admin").Use(middlewares.AuthMiddleware())
		{
			admin.GET("/users", adminController.GetUsers)
			admin.GET("/users/:id", adminController.GetUser)
			admin.POST("/users", adminController.CreateUsers)
			admin.PUT("/users/:id", adminController.UpdateUsers)
			admin.DELETE("/users/:id", adminController.DeleteUsers)
		}
	}
}
