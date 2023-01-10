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
