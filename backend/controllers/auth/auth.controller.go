package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-sso/initializers"
	"github.com/xarick/gin-sso/models"
	"github.com/xarick/gin-sso/services"
)

func Login(c *gin.Context) {
	var payload models.SignInUser

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := services.VerifyPassword(user.Password, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	token, err := services.GenerateJWT(config.JWTSecret, config.JWTExpTime, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var user models.SignUpUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Password != user.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		RoleId:   2,
		Status:   true,
	}

	if err := initializers.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"logout": "page"})
}
