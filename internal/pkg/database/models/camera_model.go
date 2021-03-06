package models

// Camera represents a singular camera which is accessible over the network
type Camera struct {
	// swagger:ignore
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
	// swagger:ignore
	ProxyID uint64

	// the cameras' authentication details
	// required: true
	Auth *CameraAuth `json:"auth" binding:"required"`

	// the camera model such as RLC-411WS
	// required: true
	Model   *CameraModel `json:"model"`
	// swagger:ignore
	ModelID uint64

	// the location
	Location   *Location `json:"location"`
	// swagger:ignore
	LocationID uint64
}

// CameraAuth contains the camera authentication information, such as it's username and password
type CameraAuth struct {
	// swagger:ignore
	ID uint64 `gorm:"primary_key"`

	// username
	// required: true
	Username string `json:"username" binding:"required"`

	// password
	// required: true
	Password string `json:"password"`

	// swagger:ignore
	CameraID uint64
}

// CameraModel is the model/type of the camera. Many cameras can have the same model e.g. RLC-411WS.
type CameraModel struct {
	// swagger:ignore
	ID uint64 `gorm:"primary_key"`
	// an auto generated unique identifier for the model
	// required: false
	UUID string `json:"uuid" gorm:"uniqueIndex"`
	// the name such as RLC-411WS
	// required: true
	Name string `json:"name"`
}
