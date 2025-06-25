package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/modules/restaurant/service"
)

func (ctl *RestaurantController) CreateRestaurant(c *gin.Context) {
	var requestBodyData model.CreateRestaurantDTO
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd := service.CreateCommand{DTO: &requestBodyData}
	restaurant, err := ctl.createCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": restaurant})
}
