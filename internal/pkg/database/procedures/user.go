package procedures

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/dgryski/trifles/uuid"
)

func (db *DB) UserCreate(username string, password string) (*models.User, error) {
	user := &models.User{
		UUID:     uuid.UUIDv4(),
		Username: username,
		Password: password,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (db *DB) UserReadByUsername(username string) (*models.User, error) {

	var user *models.User

	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (db *DB) UserUpdate(userUUID string, username string, password string) (*models.User, error) {

	var user *models.User

	if err := db.First(&user, "uuid = ?", userUUID).Error; err != nil {
		return nil, err
	}

	user.Username = username
	user.Password = password

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (db *DB) UserDelete(userUUID string) (bool, error) {

	if err := db.Delete(&models.User{}, "uuid = ?", userUUID).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (db *DB) UserRead() ([]*models.User, error) {

	var users []*models.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
