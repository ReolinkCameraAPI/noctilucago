package models

type User struct {
	ID uint64 `gorm:"primary_key"`
	// an auto generated unique identifier for the user
	// required: false
	UUID     string `json:"uuid,omitempty" gorm:"uniqueIndex"`
	Username string `json:"username" gorm:"uniqueIndex" binding:"required"`
	Password string `json:"password"`
}
