package httpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/food/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *FoodController) GetFoodByID(c *gin.Context) {
	foodID := c.Param("id")
	
	if foodID == "" {
		panic(datatype.ErrBadRequest.WithError("food id is required"))
	}
	
	cmd := foodservice.GetByIDCommand{ID: foodID}
	
	foodResponse, err := ctl.getCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{"data": foodResponse})
}