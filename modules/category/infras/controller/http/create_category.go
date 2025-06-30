package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/service"
)

func (ctl *CategoryController) CreateCategory(c *gin.Context) {
	var requestBodyData categorymodel.CreateCategoryDTO
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	cmd := categoryservice.CreateCommand{DTO: &requestBodyData}
	category, err := ctl.createCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"data": category})
}
