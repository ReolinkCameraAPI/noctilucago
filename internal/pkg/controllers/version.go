package controllers

import "github.com/gin-gonic/gin"

func (ac *ApiController) Version(c *gin.Context) {
	c.JSON(200, VersionResponse{Version: "v0.0.1"})
}
