package categoryhttpgin

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
)

func (ctl *CategoryHTTPController) CreateCategory(c *gin.Context) {
	var requestBodyData categorymodel.CreateCategoryDTO
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	cmd := categoryservice.CreateCommand{Dto: &requestBodyData}
	category, err := ctl.createCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"data": category.ID})
}
