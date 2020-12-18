package controllers

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/gin-gonic/gin"
)

type Camera struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port,omitempty"`
}

func (ac *ApiController) CameraCreate(c *gin.Context) {

	var camera models.Camera

	if err := c.BindJSON(&camera); err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Camera data parse error. Ensure the data sent to the server is correct.",
		})
	}

	c.JSON(200, camera)
}

func (ac *ApiController) CameraRead(c *gin.Context) {
	cameras, err := ac.db.CameraRead()

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error.",
		})
	}

	c.JSON(200, cameras)
}

func (ac *ApiController) CameraModelCreate(c *gin.Context) {
	var model models.CameraModel
	if err := c.BindJSON(&model); err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Camera Model data parse error. Ensure the data sent to the server is correct.",
		})
	}

	c.JSON(200, model)
}

func (ac *ApiController) CameraModelRead(c *gin.Context) {
	cameraModels, err := ac.db.CameraModelRead()
	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error",
		})
	}

	c.JSON(200, cameraModels)
}
