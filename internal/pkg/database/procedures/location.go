package procedures

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"

func (db *DB) LocationReadByUUID(uuid string) (*models.Location, error) {

	var location *models.Location

	if err := db.First(&location, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return location, nil
}
