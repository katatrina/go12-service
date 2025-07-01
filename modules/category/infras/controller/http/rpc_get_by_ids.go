package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RPCGetByIDsRequestDTO struct {
	IDs []uuid.UUID `json:"ids"`
}

func (ctl *CategoryController) RPCGetByIDS(c *gin.Context) {
	var dto RPCGetByIDsRequestDTO
	
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ids := dto.IDs
	
	categories, err := ctl.categoryRPCRepo.FindByIDs(c.Request.Context(), ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": categories})
}
