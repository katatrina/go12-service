package usergrpcctl

import (
	"context"
	"errors"
	"log"

	"github.com/katatrina/go12-service/gen/proto/user"
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	sharedcomponent "github.com/katatrina/go12-service/shared/component"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByIDs(ctx context.Context, ids []uuid.UUID) ([]usermodel.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
}

type UserGrpcServer struct {
	user.UnimplementedUserServer
	repo      UserRepository
	jwtComp   *sharedcomponent.JWTComp
}

func NewUserGrpcServer(repo UserRepository, jwtComp *sharedcomponent.JWTComp) *UserGrpcServer {
	return &UserGrpcServer{
		repo:    repo,
		jwtComp: jwtComp,
	}
}

func (s *UserGrpcServer) IntrospectToken(
	ctx context.Context,
	req *user.IntrospectTokenRequest,
) (*user.IntrospectTokenResp, error) {
	log.Println("IntrospectToken by gRPC")

	if req.Token == "" {
		return nil, errors.New("token is required")
	}

	// Parse and validate JWT token
	userID, err := s.jwtComp.Introspect(req.Token)
	if err != nil {
		log.Printf("Invalid token: %v", err)
		return nil, errors.New("invalid or expired token")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	// Find user in database
	userEntity, err := s.repo.FindByID(ctx, userUUID)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return nil, errors.New("user not found")
	}

	if userEntity.Status == usermodel.StatusDeleted || userEntity.Status == usermodel.StatusBanned {
		return nil, errors.New("user account is inactive")
	}

	// Convert to gRPC response
	userDTO := &user.UserDTO{
		Id:        userEntity.ID.String(),
		Email:     userEntity.Email,
		FirstName: userEntity.FirstName,
		LastName:  userEntity.LastName,
		Role:      string(userEntity.Role),
		Status:    string(userEntity.Status),
		CreatedAt: userEntity.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: userEntity.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return &user.IntrospectTokenResp{Data: userDTO}, nil
}

func (s *UserGrpcServer) GetUsersByIDs(
	ctx context.Context,
	req *user.GetUserIDsRequest,
) (*user.UserIDsResp, error) {
	log.Println("GetUsersByIDs by gRPC")

	uuidIds := make([]uuid.UUID, 0, len(req.Ids))

	for _, id := range req.Ids {
		parsedId, err := uuid.Parse(id)
		if err != nil {
			log.Printf("Invalid UUID: %s, error: %v", id, err)
			continue
		}
		uuidIds = append(uuidIds, parsedId)
	}

	users, err := s.repo.FindByIDs(ctx, uuidIds)
	if err != nil {
		log.Printf("Error finding users: %v", err)
		return nil, err
	}

	result := make([]*user.UserDTO, 0, len(users))

	for _, u := range users {
		if u.Status == usermodel.StatusDeleted {
			continue // Skip deleted users
		}

		result = append(result, &user.UserDTO{
			Id:        u.ID.String(),
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Role:      string(u.Role),
			Status:    string(u.Status),
			CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt: u.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &user.UserIDsResp{Data: result}, nil
}