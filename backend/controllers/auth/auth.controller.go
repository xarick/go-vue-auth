package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-sso/initializers"
	"github.com/xarick/gin-sso/models"
	"github.com/xarick/gin-sso/services"
)

func Login(c *gin.Context) {
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
		RoleId:   1,
		Status:   true,
	}

	initializers.DB.Create(&newUser)
	c.JSON(http.StatusOK, &user)
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
