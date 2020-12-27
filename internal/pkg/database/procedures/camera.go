package procedures

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/dgryski/trifles/uuid"
)

// Create a new models.Camera by passing the models.CameraModel uuid along with it
// The newly created models.Camera will be returned
func (db *DB) CameraCreate(modelUUID string, camera *models.Camera) (*models.Camera, error) {

	cameraModel, err := db.CameraModelReadUUID(modelUUID)

	if err != nil {
		return nil, err
	}

	camera.UUID = uuid.UUIDv4()
	camera.ModelID = cameraModel.ID

	if err := db.Create(&camera).Error; err != nil {
		return nil, err
	}

	return camera, nil
}

// Get all the models.Camera in the database
func (db *DB) CameraRead() ([]*models.Camera, error) {

	var cameras []*models.Camera

	if err := db.Preload("Proxy").Preload("Auth").Find(&cameras).Error; err != nil {
		return nil, err
	}

	return cameras, nil
}

// Delete a models.Camera
func (db *DB) CameraDelete(cameraUUID string) (bool, error) {

	if err := db.Delete(&models.Camera{}, "uuid = ?", cameraUUID).Error; err != nil {
		return false, err
	}

	return true, nil
}

// Update an already existing models.Camera
// The newly updated models.Camera will be returned
func (db *DB) CameraUpdate(cameraUUID string, camera *models.Camera) (*models.Camera, error) {

	var dbCamera *models.Camera

	if err := db.First(&dbCamera, "uuid = ?", cameraUUID).Error; err != nil {
		return nil, err
	}

	camera.ID = dbCamera.ID
	camera.UUID = dbCamera.UUID

	if err := db.Save(&camera).Error; err != nil {
		return nil, err
	}

	return camera, nil
}

// Create a new models.CameraModel
func (db *DB) CameraModelCreate(model *models.CameraModel) (*models.CameraModel, error) {

	model.UUID = uuid.UUIDv4()

	if err := db.Create(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

// Get a models.CameraModel with its uuid
func (db *DB) CameraModelReadUUID(uuid string) (*models.CameraModel, error) {

	var model *models.CameraModel

	if err := db.First(&model, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return model, nil

}

// Get all the models.CameraModel in the database
func (db *DB) CameraModelRead() ([]*models.CameraModel, error) {

	var cameraModels []*models.CameraModel

	if err := db.Find(&cameraModels).Error; err != nil {
		return nil, err
	}

	return cameraModels, nil

}
