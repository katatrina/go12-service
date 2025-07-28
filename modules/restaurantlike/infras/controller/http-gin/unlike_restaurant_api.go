package restaurantlikehttpgin

import (
	"net/http"
	
	restaurantlikeservice "github.com/katatrina/go12-service/modules/restaurantlike/service"
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *RestaurantLikeHTTPController) UnlikeRestaurantAPI(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	
	if err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	requester := c.MustGet(datatype.KeyRequester).(datatype.Requester)
	userId := requester.Subject()
	
	cmd := restaurantlikeservice.UnlikeRestaurantCommand{
		RestaurantId: id,
		UserId:       userId,
	}
	
	if err := ctrl.unlikeHandler.Execute(c.Request.Context(), &cmd); err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{"data": true})
}
