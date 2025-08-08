package mysqlrepository

import (
	"context"
	stderrors "errors"
	
	"github.com/google/uuid"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/shared/datatype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *FoodRepository) FindByID(ctx context.Context, id uuid.UUID) (*foodmodel.Food, error) {
	db := repo.dbCtx.GetMainConnection()
	
	var food foodmodel.Food
	
	if err := db.WithContext(ctx).First(&food, "id = ?", id).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, datatype.ErrRecordNotFound
		}
		
		return nil, errors.WithStack(err)
	}
	
	return &food, nil
}