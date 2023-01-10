package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-sso/initializers"
	"github.com/xarick/gin-sso/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	initializers.DB.Where("id = ?", c.Param("id")).First(&user)

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	userResponse := &models.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleId: user.RoleId,
		Status: user.Status,
	}

	c.JSON(http.StatusOK, &userResponse)
}

func CreateUsers(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUsers(c *gin.Context) {
	var user models.User

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUsers(c *gin.Context) {
	var user models.User

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	initializers.DB.Delete(&user)
	c.JSON(http.StatusOK, &user)
}
