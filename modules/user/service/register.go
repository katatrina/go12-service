package userservice

import (
	"context"
	"errors"
	"strings"
	"time"
	
	"github.com/google/uuid"
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	"github.com/katatrina/go12-service/shared"
	"github.com/katatrina/go12-service/shared/datatype"
	"golang.org/x/crypto/bcrypt"
)

type RegisterCommand struct {
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
}

func (cmd RegisterCommand) Validate() error {
	cmd.Email = strings.TrimSpace(cmd.Email)
	cmd.Password = strings.TrimSpace(cmd.Password)
	cmd.FirstName = strings.TrimSpace(cmd.FirstName)
	cmd.LastName = strings.TrimSpace(cmd.LastName)
	
	if !shared.CheckValidEmailFormat(cmd.Email) {
		return usermodel.ErrInvalidEmail
	}
	
	if !shared.CheckValidPasswordFormat(cmd.Password) {
		return usermodel.ErrInvalidPassword
	}
	
	if len(cmd.FirstName) < 2 || len(cmd.FirstName) > 50 {
		return usermodel.ErrInvalidFirstName
	}
	
	if len(cmd.LastName) < 2 || len(cmd.LastName) > 50 {
		return usermodel.ErrInvalidLastName
	}
	
	return nil
}

type RegisterRepo interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.User, error)
	Insert(ctx context.Context, user *usermodel.User) error
}

type RegisterCommandHandler struct {
	registerRepo RegisterRepo
}

func NewRegisterCommandHandler(registerRepo RegisterRepo) *RegisterCommandHandler {
	return &RegisterCommandHandler{
		registerRepo: registerRepo,
	}
}

func (h *RegisterCommandHandler) Execute(ctx context.Context, cmd *RegisterCommand) (*usermodel.User, error) {
	// Validate the command
	if err := cmd.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithError(err.Error())
	}
	
	// Check if the email already exists
	existUser, err := h.registerRepo.FindByEmail(ctx, cmd.Email)
	if err != nil && !errors.Is(err, datatype.ErrRecordNotFound) {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	if existUser != nil {
		if existUser.Status == usermodel.StatusDeleted || existUser.Status == usermodel.StatusBanned {
			return nil, datatype.ErrBadRequest.WithError("user is deleted or banned")
		}
		
		return nil, datatype.ErrBadRequest.WithError(usermodel.ErrEmailAlreadyExists.Error())
	}
	
	salt, err := shared.RandomStr(16)
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	userID, err := uuid.NewV7()
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	user := &usermodel.User{
		ID:        userID,
		Email:     cmd.Email,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Password:  string(hashPassword),
		Salt:      salt,
		Type:      usermodel.TypeEmailPassword,
		Role:      usermodel.RoleUser,
		Status:    usermodel.StatusActive,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	
	// Insert the user into the repository
	err = h.registerRepo.Insert(ctx, user)
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	// Return the newly created user
	return user, nil
}
