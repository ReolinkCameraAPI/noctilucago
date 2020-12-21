package service

import (
	"fmt"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
	"github.com/imkira/go-observer"
)

type CameraService struct {
	cameras  map[string]*reolinkapi.Camera
	property observer.Property
}

func NewCameraService() *CameraService {
	cameras := make(map[string]*reolinkapi.Camera)

	property := observer.NewProperty(cameras)

	return &CameraService{cameras: cameras, property: property}
}

func (c *CameraService) Add(camera models.Camera, uuid string) error {
	cam, err := reolinkapi.NewCamera(camera.Auth.Username, camera.Auth.Password, camera.Host,
		reolinkapi.WithDeferLogin(true))
	if err != nil {
		return err
	}

	c.cameras[uuid] = cam

	return nil
}

func (c *CameraService) Remove(uuid string) error {
	if _, ok := c.cameras[uuid]; !ok {
		return fmt.Errorf("camera object does not exist in camera service pool")
	}
	delete(c.cameras, uuid)
	return nil
}
