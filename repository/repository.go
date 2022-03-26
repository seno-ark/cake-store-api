package repository

import (
	"cake-store-api/config"
	"cake-store-api/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CountCake() (int, error)
	GetCakeList(config.M) ([]*model.CakeModel, error)
	GetCake(int) (*model.CakeModel, error)
	CreateCake(*model.CakeForm) (int, error)
	UpdateCake(int, *model.CakeForm) (int, error)
	DeleteCake(int) error
}

func NewRepository(db *sqlx.DB) Repository {
	return &Repo{db}
}

type Repo struct {
	DB *sqlx.DB
}
