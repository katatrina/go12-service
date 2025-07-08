package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/restaurant/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *RestaurantController) ListRestaurants(c *gin.Context) {
	var query restaurantservice.ListQuery
	
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	if err := query.FilterRestaurantDTO.Validate(); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	query.PagingDTO.Process()
	
	restaurants, err := ctl.listQryHdl.Execute(c.Request.Context(), &query)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data":   restaurants,
		"paging": query.PagingDTO,
		"filter": query.FilterRestaurantDTO,
	})
}
