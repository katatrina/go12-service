package gormmysql

import (
	"context"
	
	mediamodel "github.com/katatrina/go12-service/modules/media/model"
	
	"github.com/pkg/errors"
)

func (repo *MediaRepository) Insert(ctx context.Context, data *mediamodel.Media) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
