package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-sso/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			c.JSON(401, gin.H{"error": "Token is empty"})
			c.Abort()
			return
		}

		if len(bearerToken) < 7 || bearerToken[0:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		tokenString := bearerToken[7:]

		err := services.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
