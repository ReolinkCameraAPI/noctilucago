package controllers

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/api/models"
	dbmodels "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/enum"
	"github.com/gin-gonic/gin"
)

func (ac *ApiController) NetworkProxyCreate(c *gin.Context) {
	var proxy models.ProxyInput

	if c.ShouldBindJSON(&proxy) != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Network Proxy information sent to the server is incorrect.",
		})
		return
	}

	data, err := json.Marshal(proxy)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Network Proxy information sent to the server is incorrect.",
		})
		return
	}

	var dbProxy *dbmodels.Proxy

	err = json.Unmarshal(data, &dbProxy)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Network Proxy information sent to the server is incorrect.",
		})
		return
	}

	p, err := ac.db.NetworkProxyCreate(dbProxy)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error.",
		})
		return
	}

	c.JSON(200, p)
}

func (ac *ApiController) NetworkProxyRead(c *gin.Context) {

	proxies, err := ac.db.NetworkProxyRead()

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error",
		})
	}

	c.JSON(200, proxies)
}

func (ac *ApiController) NetworkProxyReadUUID(c *gin.Context) {
	proxyUUID := c.Param("uuid")

	proxy, err := ac.db.NetworkProxyReadUUID(proxyUUID)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error",
		})
	}

	c.JSON(200, proxy)

}

func (ac *ApiController) NetworkProxyUpdate(c *gin.Context) {
	proxyUUID := c.Param("uuid")

	var proxy *models.ProxyInput

	if c.ShouldBindJSON(&proxy) != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Network Proxy information sent to the server is incorrect.",
		})
		return
	}

	data, err := json.Marshal(proxy)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Network Proxy information sent to the server is incorrect",
		})
	}

	var dbProxy *dbmodels.Proxy

	err = json.Unmarshal(data, dbProxy)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Network Proxy information sent to the server is incorrect",
		})
	}

	p, err := ac.db.NetworkProxyUpdate(proxyUUID, dbProxy)

	if err != nil {
		c.JSON(500, GenericResponse{
			Status:  "error",
			Message: "Database error",
		})
	}

	c.JSON(200, p)
}

func (ac *ApiController) NetworkReadProtocol(c *gin.Context) {
	protocols := enum.ProtocolList()
	c.JSON(200, map[string]interface{}{
		"protocols": protocols,
	})
}

func (ac *ApiController) NetworkProxyReadScheme(c *gin.Context) {
	schemes := enum.SchemeList()
	c.JSON(200, map[string]interface{}{
		"schemes": schemes,
	})
}
