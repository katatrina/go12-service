package userrepository

import (
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

type UserRepository struct {
	dbCtx sharedinfras.IDbContext
}

func NewUserRepository(dbCtx sharedinfras.IDbContext) *UserRepository {
	return &UserRepository{
		dbCtx: dbCtx,
	}
}
