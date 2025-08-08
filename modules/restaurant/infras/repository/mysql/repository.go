package restaurantrepository

import (
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

type RestaurantRepository struct {
	dbCtx sharedinfras.IDbContext
}

func NewRestaurantRepository(dbCtx sharedinfras.IDbContext) *RestaurantRepository {
	return &RestaurantRepository{
		dbCtx: dbCtx,
	}
}
