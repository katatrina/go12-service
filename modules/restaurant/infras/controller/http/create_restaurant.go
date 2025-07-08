package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/modules/restaurant/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *RestaurantController) CreateRestaurant(c *gin.Context) {
	var requestBodyData restaurantmodel.CreateRestaurantDTO
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	cmd := restaurantservice.CreateCommand{DTO: &requestBodyData}
	
	restaurant, err := ctl.createCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusCreated, gin.H{"data": restaurant})
}
