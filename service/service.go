package service

import (
	"cake-store-api/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{r}
}
