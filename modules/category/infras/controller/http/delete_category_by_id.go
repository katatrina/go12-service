package controller

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/modules/category/internal/service"
)

func (ctl *CategoryController) DeleteCategoryByIDAPI(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	cmd := service.DeleteByIDCommand{ID: id}
	err = ctl.deleteCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		if errors.Is(err, model.ErrCategoryNotFound) || errors.Is(err, model.ErrCategoryDeleted) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "category not found",
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}
