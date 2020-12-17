package controllers

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/procedures"

type ApiController struct {
	db *procedures.DB
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewApiController(db *procedures.DB) *ApiController {
	return &ApiController{db}
}
