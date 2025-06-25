package controller

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/modules/restaurant/service"
)

func (ctl *RestaurantController) GetRestaurantByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := service.GetByIDQuery{ID: id}
	restaurant, err := ctl.getQryHdl.Execute(c.Request.Context(), &query)
	if err != nil {
		if errors.Is(err, model.ErrRestaurantNotFound) || errors.Is(err, model.ErrRestaurantAlreadyDeleted) {
			c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}
