package userrepository

import (
	"context"

	usermodel "github.com/katatrina/go12-service/modules/user/model"

	"github.com/google/uuid"
)

func (repo *UserRepository) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]usermodel.User, error) {
	db := repo.dbCtx.GetMainConnection()
	
	var users []usermodel.User
	
	if err := db.WithContext(ctx).Where("id IN ? AND status != ?", ids, usermodel.StatusDeleted).Find(&users).Error; err != nil {
		return nil, err
	}
	
	return users, nil
}