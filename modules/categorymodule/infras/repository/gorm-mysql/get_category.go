package categorygormmysql

import (
	"context"
	"errors"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/categorymodule/internal/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
	"gorm.io/gorm"
)

func (repo *CategoryRepository) FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error) {
	var category categorymodel.Category
	
	if err := repo.db.First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sharedmodel.ErrRecordNotFound
		}
		return nil, err
	}
	
	return &category, nil
}
