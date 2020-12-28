package responses

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"

// UserResponse is used when creating or updating a user account
// swagger:response userResponse
type UserResponse struct {
	// in: body
	Body *models.User
}

// UserArrayResponse is used when retrieving multiple user accounts
// swagger:response userArrayResponse
type UserArrayResponse struct {
	// in: body
	Body []*models.User
}
