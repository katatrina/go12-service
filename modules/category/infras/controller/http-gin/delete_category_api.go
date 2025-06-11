package categoryhttpgin

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctl *CategoryHTTPController) DeleteCategoryByIDAPI(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	err = ctl.catService.DeleteCategoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}
