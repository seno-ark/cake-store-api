package controller

import (
	"cake-store-api/service"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s}
}
