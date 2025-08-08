package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/modules/food/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *FoodController) UpdateFoodByID(c *gin.Context) {
	requester := c.MustGet(datatype.KeyRequester).(datatype.Requester)
	foodID := c.Param("id")
	
	if foodID == "" {
		panic(datatype.ErrBadRequest.WithError("food id is required"))
	}
	
	var requestBodyData foodmodel.UpdateFoodDTO
	
	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	// Additional validation could be added here to check if the requester
	// has permission to update this food item
	_ = requester // For now, we allow authenticated users to update food
	
	cmd := foodservice.UpdateCommand{
		ID:  foodID,
		DTO: &requestBodyData,
	}
	
	food, err := ctl.updateCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{"data": food})
}