package restaurantlikeservice

import (
	"context"
	"errors"
	"log"
	"time"
	
	restaurantlikemodel "github.com/katatrina/go12-service/modules/restaurantlike/model"
	"github.com/katatrina/go12-service/shared"
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/google/uuid"
)

type LikeRestaurantCommand struct {
	RestaurantId uuid.UUID
	UserId       uuid.UUID
}

type ILikeRestaurantRepo interface {
	FindLike(ctx context.Context, restaurantId, userId uuid.UUID) (*restaurantlikemodel.RestaurantLike, error)
	CreateLike(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type IEvtPublisher interface {
	Publish(ctx context.Context, topic string, evt *datatype.AppEvent) error
}

type LikeRestaurantCommandHandler struct {
	repo         ILikeRestaurantRepo
	evtPublisher IEvtPublisher
}

func NewLikeRestaurantCommandHandler(repo ILikeRestaurantRepo, evtPublisher IEvtPublisher) *LikeRestaurantCommandHandler {
	return &LikeRestaurantCommandHandler{repo: repo, evtPublisher: evtPublisher}
}

func (hdl *LikeRestaurantCommandHandler) Execute(ctx context.Context, cmd *LikeRestaurantCommand) error {
	// Check if already liked
	_, err := hdl.repo.FindLike(ctx, cmd.RestaurantId, cmd.UserId)
	
	if err == nil {
		return datatype.ErrConflict.WithError(restaurantlikemodel.ErrRestaurantLikeExists.Error())
	}
	
	if !errors.Is(err, datatype.ErrRecordNotFound) {
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	// Create like
	now := time.Now().UTC()
	like := restaurantlikemodel.RestaurantLike{
		RestaurantID: cmd.RestaurantId,
		UserID:       cmd.UserId,
		CreatedAt:    &now,
	}
	
	if err = hdl.repo.CreateLike(ctx, &like); err != nil {
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	go func() {
		defer shared.Recover()
		
		evt := datatype.NewAppEvent(
			datatype.WithTopic(datatype.EvtUserLikedRestaurant),
			datatype.WithData(like.ToData()),
		)
		
		if err = hdl.evtPublisher.Publish(ctx, evt.Topic, evt); err != nil {
			log.Println("Failed to publish event", err)
		}
	}()
	
	return nil
}
