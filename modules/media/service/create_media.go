package mediaservice

import (
	"context"
	"time"
	
	mediamodel "github.com/katatrina/go12-service/modules/media/model"
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/google/uuid"
)

type CreateCommand struct {
	MediaCreate mediamodel.MediaCreateDTO
}

type ICreateRepo interface {
	Insert(ctx context.Context, data *mediamodel.Media) error
}

type CreateCommandHandler struct {
	mediaRepo ICreateRepo
}

func NewCreateCommandHandler(mediaRepo ICreateRepo) *CreateCommandHandler {
	return &CreateCommandHandler{mediaRepo: mediaRepo}
}

func (hdl *CreateCommandHandler) Execute(ctx context.Context, cmd *CreateCommand) (*uuid.UUID, error) {
	newId, _ := uuid.NewV7()
	now := time.Now().UTC()
	
	media := mediamodel.Media{
		ID:        newId,
		Filename:  cmd.MediaCreate.Filename,
		CloudName: cmd.MediaCreate.CloudName,
		Size:      cmd.MediaCreate.Size,
		Ext:       cmd.MediaCreate.Ext,
		Status:    mediamodel.MediaStatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}
	
	if err := hdl.mediaRepo.Insert(ctx, &media); err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	return &media.ID, nil
}
