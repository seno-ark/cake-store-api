package service

import (
	"cake-store-api/repository"
)

type Service struct {
	Repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{r}
}
