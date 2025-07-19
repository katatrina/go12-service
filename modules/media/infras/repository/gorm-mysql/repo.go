package gormmysql

import sharedinfras "github.com/katatrina/go12-service/shared/infras"

type MediaRepository struct {
	dbCtx sharedinfras.IDbContext
}

func NewMediaRepository(dbCtx sharedinfras.IDbContext) *MediaRepository {
	return &MediaRepository{dbCtx: dbCtx}
}
