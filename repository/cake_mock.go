package repository

import (
	"cake-store-api/config"
	"cake-store-api/model"
)

func (r *RepositoryMock) CountCake() (int, error) {
	return 0, nil
}

func (r *RepositoryMock) GetCakeList(params config.M) ([]*model.CakeModel, error) {
	return nil, nil
}

func (r *RepositoryMock) GetCake(cakeID int) (*model.CakeModel, error) {

	var cake *model.CakeModel
	var err error

	args := r.Mock.Called(cakeID)

	if args.Get(0) != nil {
		cake = args.Get(0).(*model.CakeModel)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return cake, err
}

func (r *RepositoryMock) CreateCake(cakeForm *model.CakeForm) (int, error) {
	return 0, nil
}

func (r *RepositoryMock) UpdateCake(cakeID int, cakeForm *model.CakeForm) (int, error) {
	return 0, nil
}

func (r *RepositoryMock) DeleteCake(cakeID int) error {
	return nil
}
