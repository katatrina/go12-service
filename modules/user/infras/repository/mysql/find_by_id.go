package userrepository

import (
	"context"
	
	"github.com/google/uuid"
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	"github.com/katatrina/go12-service/shared/datatype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	db := repo.dbCtx.GetMainConnection()
	var user usermodel.User
	
	if err := db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, datatype.ErrRecordNotFound
		}
		
		return nil, errors.WithStack(err)
	}
	
	return &user, nil
}
