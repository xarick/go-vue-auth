package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authControllers "github.com/xarick/gin-sso/controllers/auth"
)

func UserRoutes(route *gin.Engine) {

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"page": "404"})
	})

	api := route.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authControllers.Login)
			auth.POST("/register", authControllers.Register)
			auth.POST("/logout", authControllers.Logout)
		}

		// secured := api.Use(middleware.AuthMiddleware())
		// {
		// 	secured.GET("/esbbc/get-tariffs", inhouseControllers.GetTariffs)
		// 	secured.GET("/esbbc/get-tariff-info", inhouseControllers.GetTarifInfo)
		// 	secured.GET("/esbbc/get-tariff-with-amount", inhouseControllers.GetTariffWithAmount)

		// 	secured.GET("/esbca/get-jur-acc", inhouseControllers.GetJurAcc)
		// 	secured.GET("/esbca/acc-info-short", inhouseControllers.AccInfoShort)

		// 	secured.GET("/esbhr/get-bxo-info", inhouseControllers.GetBxoInfo)
		// 	secured.POST("/esbhr/add-sum-operand", inhouseControllers.AddSumOperand)
		// 	secured.GET("/esbhr/check-state-operand", inhouseControllers.CheckStateOperand)

		// 	secured.POST("/esbcrm/send-sms", inhouseControllers.SendSMS)
		// 	secured.GET("/esbcrm/get-sms", inhouseControllers.GetSMS)
		// }
	}
}
