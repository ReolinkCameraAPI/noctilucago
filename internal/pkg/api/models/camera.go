package models

// CameraInput is for creating and updating a camera object
// This is used to keep track and manage cameras
// swagger:model
type CameraInput struct {
	// a custom name given to the camera (a short description)
	// required: true
	// example: "garden fence"
	Name string `json:"name" binding:"required"`

	// Host an ip address or domain
	// required: true
	// example: 192.168.0.1 or example.com
	Host string `json:"host" binding:"required"`

	// Username is used to authenticate with the camera
	// required: true
	Username string `json:"username" binding:"required"`

	// Password is used to authenticate with the camera
	// required: true
	Password string `json:"password"`

	// ProxyUUID is the unique proxy identifier that should be added
	// to this camera
	// required: false
	ProxyUUID string `json:"proxyUUID,omitempty"`
}

// CameraModelInput is for creating and updating a camera model object
// Many Cameras can be grouped under a singular CameraModelInput
// swagger:model
type CameraModelInput struct {
	// Name is a unique grouping for Cameras
	// required: true
	Name string `json:"name" binding:"required"`
}
