package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Check if category exists and is not deleted
		var existingCategory Category
		if err := db.Table(Category{}.TableName()).
			Where("id = ? AND status != ?", id, StatusDeleted).
			First(&existingCategory).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		
		var requestBodyData CategoryUpdateDTO
		
		if err := c.ShouldBindJSON(&requestBodyData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Validate status if provided
		if requestBodyData.Status != nil {
			statusValue := *requestBodyData.Status
			if statusValue != 0 && statusValue != 1 && statusValue != 2 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 0 (active), 1 (inactive) or 2 (deleted)"})
				return
			}
			
			// Convert status int to string
			var statusStr string
			switch statusValue {
			case 0:
				statusStr = StatusActive
			case 1:
				statusStr = StatusInactive
			case 2:
				statusStr = StatusDeleted
			}
			
			// Update the status in the database directly
			if err := db.Table(Category{}.TableName()).
				Where("id = ?", id).
				Update("status", statusStr).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		
		// Update other fields
		if err := db.Table(Category{}.TableName()).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"name":        requestBodyData.Name,
				"description": requestBodyData.Description,
			}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		// Get updated category
		var updatedCategory Category
		if err = db.Table(Category{}.TableName()).
			Where("id = ?", id).
			First(&updatedCategory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"data": updatedCategory})
	}
}
