package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *RestaurantController) DeleteRestaurantByID(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err := uuid.Parse(idStr)
	if err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	cmd := restaurantservice.DeleteByIDCommand{ID: id}
	if err = ctl.deleteCmdHdl.Execute(c.Request.Context(), &cmd); err != nil {
		panic(err)
	}
	
	c.Status(http.StatusNoContent)
}
