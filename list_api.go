package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging PagingDTO
		
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		paging.Process()
		
		var data []Category
		
		ldb := db.Where("status in (?)", []string{StatusActive})
		
		if err := ldb.Table(Category{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		if err := ldb.Order("id desc").
			Limit(paging.Limit).
			Offset((paging.Page - 1) * paging.Limit).
			Find(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"data": data, "paging": paging})
	}
}
