package controllers

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/api/models"
	dbmodels "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/gin-gonic/gin"
)

type Camera struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port,omitempty"`
}

func (ac *ApiController) CameraCreate(c *gin.Context) {
	modelUUID := c.Param("model")

	var camera *models.CameraInput

	if err := c.ShouldBindJSON(&camera); err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Camera data parse error. Ensure the data sent to the server is correct.",
		})
		return
	}

	dbCamera := &dbmodels.Camera{
		Name: camera.Name,
		Host: camera.Host,
		Auth: &dbmodels.CameraAuth{
			Username: camera.Username,
			Password: camera.Password,
		},
	}

	if camera.ProxyUUID != "" {
		dbProxy, err := ac.db.NetworkProxyReadUUID(camera.ProxyUUID)

		if err != nil {
			c.JSON(500, GenericResponse{
				Status:  "error",
				Message: "Database error.",
			})
			return
		}

		dbCamera.ProxyID = dbProxy.ID
	}

	newCamera, err := ac.db.CameraCreate(modelUUID, dbCamera)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error.",
		})
		return
	}

	c.JSON(200, newCamera)
}

func (ac *ApiController) CameraRead(c *gin.Context) {
	cameras, err := ac.db.CameraRead()

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error.",
		})
		return
	}

	c.JSON(200, cameras)
}

func (ac *ApiController) CameraDelete(c *gin.Context) {
	cameraUUID, ok := c.Params.Get("uuid")

	if !ok {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "UUID parameter needs to be set",
		})
		return
	}

	_, err := ac.db.CameraDelete(cameraUUID)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error.",
		})
		return
	}

	c.JSON(200, GenericResponse{
		Status:  "success",
		Message: "camera successfully deleted",
	})
}

func (ac *ApiController) CameraUpdate(c *gin.Context) {
	cameraUUID, ok := c.Params.Get("uuid")

	if !ok {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "UUID parameter needs to be set",
		})
		return
	}

	var camera models.CameraInput

	if err := c.BindJSON(&camera); err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Camera data parse error. Ensure the data sent to the server is correct.",
		})
		return
	}

	dbCamera := &dbmodels.Camera{
		Name: camera.Name,
		Host: camera.Host,
		Auth: &dbmodels.CameraAuth{
			Username: camera.Username,
			Password: camera.Password,
		},
	}

	updatedCamera, err := ac.db.CameraUpdate(cameraUUID, dbCamera)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error.",
		})
		return
	}

	c.JSON(200, updatedCamera)
}

func (ac *ApiController) CameraModelCreate(c *gin.Context) {
	var model models.CameraModelInput
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Camera Model data parse error. Ensure the data sent to the server is correct.",
		})
		return
	}

	dbModel := &dbmodels.CameraModel{
		Name: model.Name,
	}

	dbModel, err := ac.db.CameraModelCreate(dbModel)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Camera Model data parse error. Ensure the data sent to the server is correct.",
		})
		return
	}

	c.JSON(200, dbModel)
}

func (ac *ApiController) CameraModelRead(c *gin.Context) {
	cameraModels, err := ac.db.CameraModelRead()
	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error",
		})
		return
	}

	c.JSON(200, cameraModels)
}
