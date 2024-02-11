package userprofile

import (
	"context"
	"time"

	"encore.app/pkg/errorhandler"
	"encore.app/pkg/queryhelper"
)

type ListProfilesRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Email         string    `json:"email"`
	FullName      string    `json:"full_name"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedAtFrom time.Time `json:"created_at_from"`
	CreatedAtTo   time.Time `json:"created_at_to"`
}

type ListProfilesResponse struct {
	UserProfiles []UserProfile `json:"user_profiles"`
}

// ListProfiles inserts a new userprofile with given user detail
//
//encore:api public method=GET path=/userprofile/list
func (s *Service) ListProfiles(ctx context.Context, params ListProfilesRequest) (ListProfilesResponse, error) {
	var users []UserProfile
	user := UserProfile{
		ID:        params.ID,
		UserID:    params.UserID,
		Email:     params.Email,
		FullName:  params.FullName,
		IsDeleted: params.IsDeleted,
	}
	query := s.db.Where(&user)
	query = queryhelper.ApplyRangeFilter(query, "created_at", params.CreatedAtFrom, params.CreatedAtTo)
	err := query.Find(&users).Error
	err = errorhandler.HandleDBError(err)
	return ListProfilesResponse{UserProfiles: users}, err
}
