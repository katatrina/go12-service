package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBodyData Category
		
		if err := c.ShouldBindJSON(&requestBodyData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if err := requestBodyData.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		requestBodyData.Id, _ = uuid.NewV7()
		
		if err := db.Create(&requestBodyData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusCreated, gin.H{"data": requestBodyData.Id})
	}
}
