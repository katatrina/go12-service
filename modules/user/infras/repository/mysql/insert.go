package userrepository

import (
	"context"
	
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	"github.com/pkg/errors"
)

func (repo *UserRepository) Insert(ctx context.Context, user *usermodel.User) error {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
