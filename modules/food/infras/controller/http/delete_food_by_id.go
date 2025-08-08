package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/food/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *FoodController) DeleteFoodByID(c *gin.Context) {
	requester := c.MustGet(datatype.KeyRequester).(datatype.Requester)
	foodID := c.Param("id")
	
	if foodID == "" {
		panic(datatype.ErrBadRequest.WithError("food id is required"))
	}
	
	// Additional validation could be added here to check if the requester
	// has permission to delete this food item
	_ = requester // For now, we allow authenticated users to delete food
	
	cmd := foodservice.DeleteCommand{ID: foodID}
	
	err := ctl.deleteCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "food deleted successfully"})
}