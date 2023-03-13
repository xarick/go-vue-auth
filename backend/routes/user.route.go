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

		// auth := api.Use(middleware.AuthMiddleware())
		// {
		// 	auth.GET("/esbbc/get-tariffs", inhouseControllers.GetTariffs)
		// 	auth.GET("/esbbc/get-tariff-info", inhouseControllers.GetTarifInfo)
		// 	auth.GET("/esbbc/get-tariff-with-amount", inhouseControllers.GetTariffWithAmount)

		// 	auth.GET("/esbca/get-jur-acc", inhouseControllers.GetJurAcc)
		// 	auth.GET("/esbca/acc-info-short", inhouseControllers.AccInfoShort)

		// 	auth.GET("/esbhr/get-bxo-info", inhouseControllers.GetBxoInfo)
		// 	auth.POST("/esbhr/add-sum-operand", inhouseControllers.AddSumOperand)
		// 	auth.GET("/esbhr/check-state-operand", inhouseControllers.CheckStateOperand)

		// 	auth.POST("/esbcrm/send-sms", inhouseControllers.SendSMS)
		// 	auth.GET("/esbcrm/get-sms", inhouseControllers.GetSMS)
		// }
	}
}
