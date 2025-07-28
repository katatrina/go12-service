package gormmysql

import sharedinfras "github.com/katatrina/go12-service/shared/infras"

type RestaurantLikeRepository struct {
	dbCtx sharedinfras.IDbContext
}

func NewRestaurantLikeRepository(dbCtx sharedinfras.IDbContext) *RestaurantLikeRepository {
	return &RestaurantLikeRepository{dbCtx: dbCtx}
}
