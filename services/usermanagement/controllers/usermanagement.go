package controllers

// Load required packages
import (
	"context"

	model "github.com/RoomFlow/backend/internal/proto/usermanagement"
)

type UsermanagementServer struct {
}

// LoginUser Attempts to login a user based on the passed on credentials.
func (s *UsermanagementServer) LoginUser(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	// TODO: access the database and check for correctness of credentials
	return nil, nil
}

// RegisterUser Attempts to register a user based on the passed on credentials.
func (s *UsermanagementServer) RegisterUser(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error) {
	// TODO: add the new user to the database.
	return nil, nil
}
