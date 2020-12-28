package procedures

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/dgryski/trifles/uuid"
)

func (db *DB) NetworkProxyCreate(proxy *models.Proxy) (*models.Proxy, error) {
	proxy.UUID = uuid.UUIDv4()
	if err := db.Create(&proxy).Error; err != nil {
		return nil, err
	}

	return proxy, nil
}

func (db *DB) NetworkProxyRead() ([]*models.Proxy, error) {
	var proxies []*models.Proxy

	if err := db.Find(&proxies).Error; err != nil {
		return nil, err
	}

	return proxies, nil
}

func (db *DB) NetworkProxyReadUUID(uuid string) (*models.Proxy, error) {

	var proxy *models.Proxy

	if err := db.First(&proxy, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return proxy, nil
}

func (db *DB) NetworkProxyUpdate(uuid string, proxy *models.Proxy) (*models.Proxy, error) {

	var dbProxy *models.Proxy

	if err := db.First(&dbProxy, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	proxy.UUID = dbProxy.UUID
	proxy.ID = dbProxy.ID

	if err := db.Save(&proxy).Error; err != nil {
		return nil, err
	}

	return proxy, nil
}
