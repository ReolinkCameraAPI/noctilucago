package models

// Camera represents a singular camera which is accessible over the network
type Camera struct {
	ID uint64 `gorm:"primary_key"`
	// an auto generated unique identifier for the camera
	// required: false
	UUID string `json:"uuid,omitempty" gorm:"uniqueIndex"`

	// a custom name given to the camera (a short description)
	// required: true
	Name string `json:"name" binding:"required"`

	// an ip address or domain
	// required: true
	Host string `json:"host" binding:"required"`

	// connection settings for a camera behind a proxy
	// required: false
	Proxy   *Proxy `json:"proxy"`
	ProxyID uint64

	// the cameras' authentication details
	// required: true
	Auth *CameraAuth `json:"auth" binding:"required"`

	// the camera model such as RLC-411WS
	// required: true
	Model   *CameraModel `json:"model"`
	ModelID uint64
}

// CameraAuth contains the camera authentication information, such as it's username and password
type CameraAuth struct {
	ID uint64 `gorm:"primary_key"`

	// username
	// required: true
	Username string `json:"username" binding:"required"`

	// password
	// required: true
	Password string `json:"password"`

	CameraID uint64
}

// CameraModel is the model/type of the camera. Many cameras can have the same model e.g. RLC-411WS.
type CameraModel struct {
	ID uint64 `gorm:"primary_key"`
	// an auto generated unique identifier for the model
	// required: false
	UUID string `json:"uuid" gorm:"uniqueIndex"`
	// the name such as RLC-411WS
	// required: true
	Name string `json:"name"`
}

// CameraLocation is the physical location of the camera.
type CameraLocation struct {
	ID           uint64 `gorm:"primary_key"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	AddressLine3 string `json:"addressLine3,omitempty"`
	Province     string `json:"province,omitempty"`
	City         string `json:"city" binding:"required"`
	Country      string `json:"country" binding:"required"`
}
