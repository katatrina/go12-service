package categoryhttpgin

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
)

func (ctl *CategoryHTTPController) CreateCategoryAPI(c *gin.Context) {
	var requestBodyData categorymodel.Category
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	cmd := categoryservice.CreateNewCommand{Dto: requestBodyData}
	id, err := ctl.createNewCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"data": id})
}
