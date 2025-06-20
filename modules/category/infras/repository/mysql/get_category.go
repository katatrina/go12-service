package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"gorm.io/gorm"
)

func (repo *CategoryRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	var category model.Category

	if err := repo.db.First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}
