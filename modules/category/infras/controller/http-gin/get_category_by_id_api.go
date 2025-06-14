package categoryhttpgin

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

func (ctl *CategoryHTTPController) GetCategoryByIDAPI(c *gin.Context) {
	// Lấy ID từ URL parameter
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// Gọi service để lấy thông tin category
	category, err := ctl.getDetailQryHdl.Execute(c.Request.Context(), &categoryservice.GetDetailQuery{ID: id})
	if err != nil {
		if errors.Is(err, sharedmodel.ErrRecordNotFound) {
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
	
	// Trả về thông tin category
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
