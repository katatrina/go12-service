package mysqlrepository

import (
	"context"
	
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func (repo *FoodRepository) Insert(ctx context.Context, data *foodmodel.Food) error {
	_, span := otel.Tracer("go12-service").Start(ctx, "food-repo-mysql.insert")
	defer span.End()
	
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}