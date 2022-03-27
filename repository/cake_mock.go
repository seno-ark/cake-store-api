package repository

import (
	"cake-store-api/config"
	"cake-store-api/model"
)

// CountCake

type CountCakeArg struct{}
type CountCakeResult struct {
	Count int
	Err   error
}

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

// GetCakeList

type GetCakeListArg struct {
	Params config.M
}
type GetCakeListResult struct {
	Cakes []*model.CakeModel
	Err   error
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

// GetCake

type GetCakeArg struct {
	CakeID int
}
type GetCakeResult struct {
	Cake *model.CakeModel
	Err  error
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

// CreateCake

type CreateCakeArg struct {
	CakeForm *model.CakeForm
}
type CreateCakeResult struct {
	CakeID int
	Err    error
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

// UpdateCake

type UpdateCakeArg struct {
	ID       int
	CakeForm *model.CakeForm
}
type UpdateCakeResult struct {
	CakeID int
	Err    error
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

// DeleteCake

type DeleteCakeArg struct {
	CakeID int
}
type DeleteCakeResult struct {
	Err error
}

func (r *RepositoryMock) DeleteCake(cakeID int) (err error) {

	args := r.Mock.Called(cakeID)

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return
}
