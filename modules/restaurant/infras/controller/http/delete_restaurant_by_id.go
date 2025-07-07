package httpcontroller

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/modules/restaurant/service"
)

func (ctl *RestaurantController) DeleteRestaurantByID(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	cmd := restaurantservice.DeleteByIDCommand{ID: id}
	if err = ctl.deleteCmdHdl.Execute(c.Request.Context(), &cmd); err != nil {
		switch {
		case errors.Is(err, restaurantmodel.ErrRestaurantNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, restaurantmodel.ErrRestaurantAlreadyDeleted):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		
		return
	}
	
	c.Status(http.StatusNoContent)
}
