package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/modules/restaurant/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *RestaurantController) UpdateRestaurantByID(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err := uuid.Parse(idStr)
	if err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	var dto restaurantmodel.UpdateRestaurantDTO
	if err = c.ShouldBindJSON(&dto); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	cmd := restaurantservice.UpdateByIDCommand{ID: id, DTO: &dto}
	if err = ctl.updateCmdHdl.Execute(c.Request.Context(), &cmd); err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{"data": true})
}
