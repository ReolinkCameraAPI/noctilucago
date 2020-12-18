package procedures

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
)

func (db *DB) CameraCreate(camera models.Camera) (*models.Camera, error) {

	if err := db.Create(&camera).Error; err != nil {
		return nil, err
	}

	return &camera, nil
}

func (db *DB) CameraRead() ([]*models.Camera, error) {

	var cameras []*models.Camera

	if err := db.Find(&cameras).Error; err != nil {
		return nil, err
	}

	return cameras, nil
}

func (db *DB) CameraModelCreate(model models.CameraModel) (*models.CameraModel, error) {

	if err := db.Create(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (db *DB) CameraModelRead() ([]*models.CameraModel, error) {

	var cameraModels []*models.CameraModel

	if err := db.Find(&cameraModels).Error; err != nil {
		return nil, err
	}

	return cameraModels, nil

}
