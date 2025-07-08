package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *RestaurantController) GetRestaurantByID(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err := uuid.Parse(idStr)
	if err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error())) // This should be handled by a middleware
	}
	
	query := restaurantservice.GetByIDQuery{ID: id}
	restaurant, err := ctl.getQryHdl.Execute(c.Request.Context(), &query)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}
