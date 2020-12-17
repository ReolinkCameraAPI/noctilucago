package models

type Camera struct {
	ID      uint64       `gorm:"primary_key"`
	UUID    string       `json:"uuid,omitempty" gorm:"uniqueIndex"`
	Name    string       `json:"name"`
	Model   *CameraModel `json:"model"` // Belongs to a camera model
	ModelID uint64
}

type CameraModel struct {
	ID   uint64 `gorm:"primary_key"`
	UUID string `json:"uuid" gorm:"uniqueIndex"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CameraLocation struct {
	ID           uint64 `gorm:"primary_key"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	AddressLine3 string `json:"addressLine3,omitempty"`
	Province     string `json:"province,omitempty"`
	City         string `json:"city"`
	Country      string `json:"country"`
}
