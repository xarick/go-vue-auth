package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"login": "page"})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"register": "page"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"logout": "page"})
}
