package controllers

import (
	"net/http"
	"payment-manager/models"
	"payment-manager/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Send(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var transaction models.Transaction
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := services.ProcessTransaction(&transaction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Transaction sent"})
	}
}

func Withdraw(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var transaction models.Transaction
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := services.ProcessTransaction(&transaction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Transaction withdrawn"})
	}
}
