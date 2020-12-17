package procedures

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"

func (db *DB) NetworkProxyCreate(proxy models.Proxy) (*models.Proxy, error) {

	if err := db.Create(&proxy).Error; err != nil {
		return nil, err
	}

	return &proxy, nil
}
