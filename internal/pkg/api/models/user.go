package models

// UserInput holds the user information for authentication
// swagger:model
type UserInput struct {
	// Username is the unique identifier of the user
	// required: true
	Username string `json:"username" binding:"required"`

	// Password for authenticating the user
	// required: true
	Password string `json:"password" binding:"required"`
}
