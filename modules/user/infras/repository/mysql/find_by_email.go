package userrepository

import (
	"context"
	
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	"github.com/katatrina/go12-service/shared/datatype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var user usermodel.User
	
	if err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, datatype.ErrRecordNotFound
		}
		
		return nil, errors.WithStack(err)
	}
	
	return &user, nil
}
