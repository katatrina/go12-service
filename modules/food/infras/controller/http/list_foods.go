package httpcontroller

import (
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/modules/food/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctl *FoodController) ListFoods(c *gin.Context) {
	var filterParams foodmodel.FoodFilterDTO
	
	if err := c.ShouldBindQuery(&filterParams); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	// Parse pagination parameters
	page := 1
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	
	limit := 20
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}
	
	listDTO := &foodmodel.FoodListDTO{
		RestaurantID: filterParams.RestaurantID,
		CategoryID:   filterParams.CategoryID,
		MinPrice:     filterParams.MinPrice,
		MaxPrice:     filterParams.MaxPrice,
		Search:       filterParams.Search,
		Page:         page,
		Limit:        limit,
	}
	
	cmd := foodservice.ListCommand{DTO: listDTO}
	
	result, err := ctl.listCmdHdl.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{"data": result})
}