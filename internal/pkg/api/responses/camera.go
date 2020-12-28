package responses

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"

// CameraResponse used when creating and updating a camera
// swagger:response cameraResponse
type CameraResponse struct {
	// in: body
	Body *models.Camera
}

// CameraArrayResponse used when retrieving multiple camera models
// swagger:response cameraArrayResponse
type CameraArrayResponse struct {
	// in: body
	Body []*models.Camera
}

// CameraModelResponse used when creating and updating a camera model
// swagger:response cameraModelResponse
type CameraModelResponse struct {
	// in: body
	Body *models.CameraModel
}

// CameraModelArrayResponse is used when retrieving multiple camera models
// swagger:response cameraModelArrayResponse
type CameraModelArrayResponse struct {
	// in: body
	Body []*models.CameraModel
}
