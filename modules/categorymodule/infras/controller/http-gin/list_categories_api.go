package categoryhttpgin

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/categorymodule/internal/service"
)

func (ctl *CategoryHTTPController) ListCategoriesAPI(c *gin.Context) {
	var dto categoryservice.ListCategoriesDTO
	
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	dto.Process()
	
	categories, err := ctl.catService.ListCategories(c.Request.Context(), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data":   categories,
		"paging": dto.PagingDTO,
		"filter": dto.FilterCategoryDTO,
	})
}
