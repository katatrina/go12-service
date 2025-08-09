package restaurantlikeservice

import (
	"context"
	"errors"
	
	restaurantlikemodel "github.com/katatrina/go12-service/modules/restaurantlike/model"
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/google/uuid"
)

type UnlikeRestaurantCommand struct {
	RestaurantID uuid.UUID
	UserID       uuid.UUID
}

type IUnlikeRestaurantRepo interface {
	FindLike(ctx context.Context, restaurantID, userID uuid.UUID) (*restaurantlikemodel.RestaurantLike, error)
	DeleteLike(ctx context.Context, restaurantID, userID uuid.UUID) error
}

type UnlikeRestaurantCommandHandler struct {
	repo IUnlikeRestaurantRepo
}

func NewUnlikeRestaurantCommandHandler(repo IUnlikeRestaurantRepo) *UnlikeRestaurantCommandHandler {
	return &UnlikeRestaurantCommandHandler{repo: repo}
}

func (hdl *UnlikeRestaurantCommandHandler) Execute(ctx context.Context, cmd *UnlikeRestaurantCommand) error {
	// Check if like exists
	_, err := hdl.repo.FindLike(ctx, cmd.RestaurantID, cmd.UserID)
	
	if err != nil {
		if errors.Is(err, datatype.ErrRecordNotFound) {
			return datatype.ErrNotFound.WithError(restaurantlikemodel.ErrRestaurantLikeNotFound.Error())
		}
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	// Delete like
	if err = hdl.repo.DeleteLike(ctx, cmd.RestaurantID, cmd.UserID); err != nil {
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	return nil
}
