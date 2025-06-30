package httpcontroller

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/service"
)

func (ctl *CategoryController) DeleteCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	cmd := categoryservice.DeleteByIDCommand{ID: id}
	err = ctl.deleteCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		if errors.Is(err, categorymodel.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.Status(http.StatusNoContent)
}
