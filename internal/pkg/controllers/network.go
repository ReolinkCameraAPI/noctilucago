package controllers

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/gin-gonic/gin"
)

func (ac *ApiController) NetworkProxyCreate(c *gin.Context) {
	var proxy models.Proxy
	if c.BindJSON(&proxy) != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Network Proxy information sent to the server is incorrect.",
		})
	}

	p, err := ac.db.NetworkProxyCreate(proxy)

	if err != nil {
		c.JSON(500, "Database error.")
	}

	c.JSON(200, p)
}
