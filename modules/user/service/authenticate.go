package userservice

import (
	"context"
	"errors"
	"strings"
	"time"
	
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	"github.com/katatrina/go12-service/shared"
	"github.com/katatrina/go12-service/shared/datatype"
)

type AuthenticateCommand struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (cmd AuthenticateCommand) Validate() error {
	cmd.Email = strings.TrimSpace(cmd.Email)
	cmd.Password = strings.TrimSpace(cmd.Password)
	
	if !shared.CheckValidEmailFormat(cmd.Email) {
		return usermodel.ErrInvalidEmail
	}
	
	if !shared.CheckValidPasswordFormat(cmd.Password) {
		return usermodel.ErrInvalidPassword
	}
	
	return nil
}

type IAuthenticateRepo interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.User, error)
}

type ITokenIssuer interface {
	IssueToken(userID string) (string, error)
	ExpiresIn() time.Duration
}

type AuthenticateCommandHandler struct {
	userRepo    IAuthenticateRepo
	tokenIssuer ITokenIssuer
}

func NewAuthenticateCommandHandler(userRepo IAuthenticateRepo, tokenIssuer ITokenIssuer) *AuthenticateCommandHandler {
	return &AuthenticateCommandHandler{
		userRepo:    userRepo,
		tokenIssuer: tokenIssuer,
	}
}

type AuthenticateResponse struct {
	Token     string        `json:"token"`
	ExpiresIn time.Duration `json:"expiresIn"`
}

func (h *AuthenticateCommandHandler) Execute(ctx context.Context, cmd *AuthenticateCommand) (*AuthenticateResponse, error) {
	if err := cmd.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithWrap(err)
	}
	
	user, err := h.userRepo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		if errors.Is(err, datatype.ErrRecordNotFound) {
			return nil, datatype.ErrBadRequest.WithError(usermodel.ErrInvalidEmailOrPassword.Error())
		}
		
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	if user.Status == usermodel.StatusDeleted || user.Status == usermodel.StatusBanned {
		return nil, datatype.ErrBadRequest.WithError(usermodel.ErrUserBannedOrDeleted.Error())
	}
	
	if err = shared.CheckPassword(user.Password, cmd.Password, user.Salt); err != nil {
		return nil, datatype.ErrBadRequest.WithError(usermodel.ErrInvalidEmailOrPassword.Error()).WithDebug(err.Error())
	}
	
	token, err := h.tokenIssuer.IssueToken(user.ID.String())
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	return &AuthenticateResponse{
		Token:     token,
		ExpiresIn: h.tokenIssuer.ExpiresIn(),
	}, nil
}
