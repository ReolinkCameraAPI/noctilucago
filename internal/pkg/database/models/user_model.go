package models

type User struct {
	// swagger:ignore
	ID uint64 `gorm:"primary_key"`
	// UUID is an auto generated unique identifier for the user
	UUID string `json:"uuid,omitempty" gorm:"uniqueIndex"`
	// Username is a unique name used by the user to log in to the system
	Username string `json:"username" gorm:"uniqueIndex"`
	// Password is a secure hash
	// swagger:ignore
	Password string
}
