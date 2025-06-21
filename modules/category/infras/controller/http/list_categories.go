package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/category/internal/service"
)

func (ctl *CategoryController) ListCategories(c *gin.Context) {
	var query service.ListQuery
	
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := query.FilterCategoryDTO.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query.PagingDTO.Process()
	
	categories, err := ctl.listQryHdl.Execute(c.Request.Context(), &query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data":   categories,
		"paging": query.PagingDTO,
		"filter": query.FilterCategoryDTO,
	})
}
