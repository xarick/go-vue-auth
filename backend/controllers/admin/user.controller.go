package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-sso/initializers"
	"github.com/xarick/gin-sso/models"
)

func GetUsers(c *gin.Context) {
	// users := []models.User{}
	// initializers.DB.Find(&users)

	usersResp := []models.UserResponse{}
	initializers.DB.Model(&models.User{}).Find(&usersResp)

	c.JSON(http.StatusOK, &usersResp)
}

func GetUser(c *gin.Context) {
	var user models.User
	// initializers.DB.Where("id = ?", c.Param("id")).First(&user)

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

	var count int64
	initializers.DB.Model(&user).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error adding information"})
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

	// initializers.DB.Save(&user)
	// c.JSON(http.StatusOK, &user)

	if err := initializers.DB.Model(&user).Updates(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred while updating"})
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

func DeleteUsers(c *gin.Context) {
	var user models.User

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	initializers.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"success": "User deleted"})
}
