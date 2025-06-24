package controller

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/modules/category/internal/service"
)

func (ctl *CategoryController) UpdateCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	var dto model.UpdateCategoryDTO
	
	if err = c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	cmd := service.UpdateByIDCommand{
		ID:  id,
		DTO: &dto,
	}
	err = ctl.updateCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrCategoryNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		case errors.Is(err, model.ErrCategoryAlreadyDeleted):
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": true})
}
