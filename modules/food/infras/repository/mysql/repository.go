package mysqlrepository

import (
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

type FoodRepository struct {
	dbCtx sharedinfras.IDbContext
}

func NewFoodRepository(dbCtx sharedinfras.IDbContext) *FoodRepository {
	return &FoodRepository{
		dbCtx: dbCtx,
	}
}