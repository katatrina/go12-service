package userservice

import (
	"context"
	"errors"
	
	"github.com/google/uuid"
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	sharedcomponent "github.com/katatrina/go12-service/shared/component"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IntrospectCommand struct {
	Token string `json:"token"`
}

func (c *IntrospectCommand) Validate() error {
	if c.Token == "" {
		return errors.New("token is required")
	}
	
	return nil
}

type IItrospectRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
}

type IntrospectCommandHandler struct {
	jwtComp *sharedcomponent.JWTComp
	repo    IItrospectRepository
}

func NewIntrospectCommandHandler(jwtComp *sharedcomponent.JWTComp, repo IItrospectRepository) *IntrospectCommandHandler {
	return &IntrospectCommandHandler{
		jwtComp: jwtComp,
		repo:    repo,
	}
}

func (h *IntrospectCommandHandler) Execute(ctx context.Context, cmd *IntrospectCommand) (*usermodel.User, error) {
	if err := cmd.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithWrap(err)
	}
	
	userID, err := h.jwtComp.Introspect(cmd.Token)
	if err != nil {
		return nil, datatype.ErrUnauthorized.WithWrap(err)
	}
	
	user, err := h.repo.FindByID(ctx, uuid.MustParse(userID))
	if err != nil {
		return nil, datatype.ErrUnauthorized.WithWrap(err)
	}
	
	if user.Status == usermodel.StatusDeleted || user.Status == usermodel.StatusBanned {
		return nil, datatype.ErrUnauthorized.WithDebug("user has been deleted or banned")
	}
	
	return user, nil
}

type IntrospectCmdHdlWrapper struct {
	hdl *IntrospectCommandHandler
}

func NewIntrospectCmdHdlWrapper(hdl *IntrospectCommandHandler) *IntrospectCmdHdlWrapper {
	return &IntrospectCmdHdlWrapper{hdl: hdl}
}

type dataRequester struct {
	UserID    uuid.UUID `json:"id"`
	RoleValue string    `json:"role"`
}

func (r dataRequester) Subject() uuid.UUID {
	return r.UserID
}

func (r dataRequester) GetRole() string {
	return r.RoleValue
}

func (w *IntrospectCmdHdlWrapper) Introspect(token string) (datatype.Requester, error) {
	user, err := w.hdl.Execute(context.Background(), &IntrospectCommand{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	
	return user, nil
}
