// Package controllers
// Holds all the controllers
//	Schemes: http, https
//  Host: 0.0.0.0, localhost
//	BasePath: /v1/api
//  Version 0.0.1
//  License: GPLv3 https://opensource.org/licenses/GPL-3.0
//	Contact: Alano Terblanche<alano@oleaintueri.com> https://oleaintueri.com
//
//	Consumes:
//	- application/json
//  Produces:
//  - application/json
//  Schemes: http, https
//  Deprecated: false
//  Responses:
//	 default: genericResponse
//	 200: An array of cameras
//	 500: genericError
//
//  swagger:meta
package controllers

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/procedures"

type ApiController struct {
	db *procedures.DB
}

// A GenericResponse is used to send back information to the client with a general status,
// could be an error or a success message.
// swagger:response generalResponse
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// A VersionResponse is used to send back information to the client with a general status,
// could be an error or a success message.
// swagger:response versionResponse
type VersionResponse struct {
	Version string `json:"version"`
}

func NewApiController(db *procedures.DB) *ApiController {
	return &ApiController{db}
}
