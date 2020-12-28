// Package controllers
// Contains all the api controllers
//
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
//	 500: genericResponse
//
//  swagger:meta
package controllers

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/procedures"

type ApiController struct {
	db *procedures.DB
}

func NewApiController(db *procedures.DB) *ApiController {
	return &ApiController{db}
}
