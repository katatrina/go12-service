package categoryhttpgin

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

func (ctl *CategoryHTTPController) CreateCategoryAPI(c *gin.Context) {
	var requestBodyData categorymodel.categorymodel
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// call business logic in service
	if err := ctl.catService.CreateNewCategory(c.Request.Context(), &requestBodyData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"data": requestBodyData.ID})
}
