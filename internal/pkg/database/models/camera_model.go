package models

// Camera is represents a singular camera which is accessible over the network
// swagger:model
type Camera struct {
	ID      uint64       `gorm:"primary_key"`
	// an auto generated unique identifier for the camera
	// required: false
	UUID    string       `json:"uuid,omitempty" gorm:"uniqueIndex"`
	// a custom name given to the camera (a short description)
	// required: true
	Name    string       `json:"name"`
	// the camera model such as RLC-411WS
	// required: true
	Model   *CameraModel `json:"model"` // Belongs to a camera model
	ModelID uint64
}

// CameraModel is the model/type of the camera. Many cameras can have the same model e.g. RLC-411WS.
// swagger:model
type CameraModel struct {
	ID   uint64 `gorm:"primary_key"`
	// an auto generated unique identifier for the model
	// required: false
	UUID string `json:"uuid" gorm:"uniqueIndex"`
	// the name such as RLC-411WS
	// required: true
	Name string `json:"name"`
}

// CameraLocation is the physical location of the camera.
// swagger:model
type CameraLocation struct {
	ID           uint64 `gorm:"primary_key"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	AddressLine3 string `json:"addressLine3,omitempty"`
	Province     string `json:"province,omitempty"`
	City         string `json:"city"`
	Country      string `json:"country"`
}
