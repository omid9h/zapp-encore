package userprofile

import (
	"context"

	"encore.app/pkg/errorhandler"
)

type RegisterRequest struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type RegisterResponse struct {
	UserProfile UserProfile `json:"user_profile"`
}

// Register inserts a new userprofile with given user detail
//
//encore:api public method=POST path=/userprofile/register
func (s *Service) Register(ctx context.Context, params RegisterRequest) (RegisterResponse, error) {
	user := UserProfile{
		UserID: params.UserID,
		Email:  params.Email,
	}
	err := s.db.Create(&user).Error
	err = errorhandler.HandleDBError(err)
	return RegisterResponse{UserProfile: user}, err
}
