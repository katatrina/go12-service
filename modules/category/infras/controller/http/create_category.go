package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/modules/category/internal/service"
)

func (ctl *CategoryController) CreateCategory(c *gin.Context) {
	var requestBodyData model.CreateCategoryDTO
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	cmd := service.CreateCommand{DTO: &requestBodyData}
	category, err := ctl.createCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"data": category})
}
