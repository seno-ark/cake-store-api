package repository

import (
	"cake-store-api/config"
	"cake-store-api/model"
)

func (r *RepositoryMock) CountCake() (count int, err error) {

	args := r.Mock.Called()

	if args.Get(0) != nil {
		count = args.Get(0).(int)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return
}

func (r *RepositoryMock) GetCakeList(params config.M) (cakes []*model.CakeModel, err error) {

	args := r.Mock.Called(params)

	if args.Get(0) != nil {
		cakes = args.Get(0).([]*model.CakeModel)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return
}

func (r *RepositoryMock) GetCake(cakeID int) (cake *model.CakeModel, err error) {

	args := r.Mock.Called(cakeID)

	if args.Get(0) != nil {
		cake = args.Get(0).(*model.CakeModel)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return
}

func (r *RepositoryMock) CreateCake(cakeForm *model.CakeForm) (cakeID int, err error) {

	args := r.Mock.Called(cakeForm)

	if args.Get(0) != nil {
		cakeID = args.Get(0).(int)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return
}

func (r *RepositoryMock) UpdateCake(id int, cakeForm *model.CakeForm) (cakeID int, err error) {

	args := r.Mock.Called(id, cakeForm)

	if args.Get(0) != nil {
		cakeID = args.Get(0).(int)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return
}

func (r *RepositoryMock) DeleteCake(cakeID int) (err error) {

	args := r.Mock.Called(cakeID)

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return
}
