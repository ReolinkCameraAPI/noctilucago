package procedures

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/enum"
)

func (db *DB) NetworkProxyCreate(proxy models.Proxy) (*models.Proxy, error) {

	if err := db.Create(&proxy).Error; err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (db *DB) NetworkReadProtocol() []string {
	return enum.ProtocolList()
}

func (db *DB) NetworkProxyReadScheme() []string {
	return enum.SchemeList()
}
