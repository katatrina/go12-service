package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/modules/food/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *FoodController) CreateFood(c *gin.Context) {
	requester := c.MustGet(datatype.KeyRequester).(datatype.Requester)
	var requestBodyData foodmodel.CreateFoodDTO
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	// Additional validation could be added here to check if the requester
	// has permission to create food for this restaurant
	_ = requester // For now, we allow authenticated users to create food
	
	cmd := foodservice.CreateCommand{DTO: &requestBodyData}
	
	food, err := ctl.createCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusCreated, gin.H{"data": food})
}