package controllers

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/api/responses"
	"github.com/gin-gonic/gin"
)

func (ac *ApiController) Version(c *gin.Context) {
	c.JSON(200, responses.VersionResponse{Version: "v0.0.1"})
}
