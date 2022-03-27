package service

import (
	"cake-store-api/config"
	"cake-store-api/model"
	"cake-store-api/repository"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repo    = &repository.RepositoryMock{Mock: mock.Mock{}}
	service = Service{Repo: repo}
)

func TestGetCakeListService(t *testing.T) {

	type Arg struct {
		Params config.M
	}

	type Result struct {
		Result *config.M
		Status int
		Err    error
	}

	testCases := []struct {
		Name                  string
		Arg                   Arg
		Result                Result
		RepoGetCakeListArg    repository.GetCakeListArg
		RepoGetCakeListResult repository.GetCakeListResult
		RepoCountCakeArg      repository.CountCakeArg
		RepoCountCakeResult   repository.CountCakeResult
	}{
		{
			Name: "GetCakeList Success",
			Arg: Arg{
				Params: config.M{
					"page":  1,
					"count": 10,
				},
			},
			Result: Result{
				Result: &config.M{
					"cakes": []*model.CakeModel{
						{
							ID: 2,
						},
					},
					"total_data": 1,
					"page":       1,
					"count":      10,
				},
				Status: http.StatusOK,
				Err:    nil,
			},
			RepoGetCakeListArg: repository.GetCakeListArg{
				Params: config.M{
					"limit":  0,
					"offset": 10,
				},
			},
			RepoGetCakeListResult: repository.GetCakeListResult{
				Cakes: []*model.CakeModel{
					{
						ID: 2,
					},
				},
				Err: nil,
			},
			RepoCountCakeArg: repository.CountCakeArg{},
			RepoCountCakeResult: repository.CountCakeResult{
				Count: 1,
				Err:   nil,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {

			// Mock Repo
			repo.Mock.On("GetCakeList", tc.RepoGetCakeListArg.Params).Return(tc.RepoGetCakeListResult.Cakes, tc.RepoGetCakeListResult.Err)
			repo.Mock.On("CountCake").Return(tc.RepoCountCakeResult.Count, tc.RepoGetCakeListResult.Err)

			// Call Method
			result, status, err := service.GetCakeList(tc.Arg.Params)

			assert.Equal(t, tc.Result.Result, result)

			assert.Equal(t, tc.Result.Status, status)

			if tc.Result.Err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetCakeService(t *testing.T) {

	type Arg struct {
		CakeID int
	}

	type Result struct {
		Cake   *model.CakeModel
		Status int
		Err    error
	}

	testCases := []struct {
		Name              string
		Arg               Arg
		Result            Result
		RepoGetCakeArg    repository.GetCakeArg
		RepoGetCakeResult repository.GetCakeResult
	}{
		{
			Name: "GetCake Success",
			Arg: Arg{
				CakeID: 2,
			},
			Result: Result{
				Cake: &model.CakeModel{
					ID: 2,
				},
				Status: http.StatusOK,
				Err:    nil,
			},
			RepoGetCakeArg: repository.GetCakeArg{
				CakeID: 2,
			},
			RepoGetCakeResult: repository.GetCakeResult{
				Cake: &model.CakeModel{
					ID: 2,
				},
				Err: nil,
			},
		},
		{
			Name: "GetCake Not Found",
			Arg: Arg{
				CakeID: 0,
			},
			Result: Result{
				Cake:   nil,
				Status: http.StatusNotFound,
				Err:    errors.New(config.MSG_ERROR_CAKE_NOT_FOUND),
			},
			RepoGetCakeArg: repository.GetCakeArg{
				CakeID: 0,
			},
			RepoGetCakeResult: repository.GetCakeResult{
				Cake: nil,
				Err:  sql.ErrNoRows,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {

			// Mock Repo
			repo.Mock.On("GetCake", tc.RepoGetCakeArg.CakeID).Return(tc.RepoGetCakeResult.Cake, tc.RepoGetCakeResult.Err)

			// Call Method
			cake, status, err := service.GetCake(tc.Arg.CakeID)

			// Assertion
			if tc.Result.Cake != nil {
				assert.Equal(t, tc.Result.Cake.ID, cake.ID)
			} else {
				assert.Nil(t, cake)
			}

			assert.Equal(t, tc.Result.Status, status)

			if tc.Result.Err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCreateCakeService(t *testing.T) {

	type Arg struct {
		CakeForm *model.CakeForm
	}

	type Result struct {
		Cake   *model.CakeModel
		Status int
		Err    error
	}

	testCases := []struct {
		Name                 string
		Arg                  Arg
		Result               Result
		RepoCreateCakeArg    repository.CreateCakeArg
		RepoCreateCakeResult repository.CreateCakeResult
		RepoGetCakeArg       repository.GetCakeArg
		RepoGetCakeResult    repository.GetCakeResult
	}{
		{
			Name: "CreateCake Success",
			Arg: Arg{
				CakeForm: &model.CakeForm{
					Title:       "BBB",
					Description: "BBB - BBB - BBB",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
			},
			Result: Result{
				Cake: &model.CakeModel{
					ID:          3,
					Title:       "BBB",
					Description: "BBB - BBB - BBB",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
				Status: http.StatusCreated,
				Err:    nil,
			},
			RepoCreateCakeArg: repository.CreateCakeArg{
				CakeForm: &model.CakeForm{
					Title:       "BBB",
					Description: "BBB - BBB - BBB",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
			},
			RepoCreateCakeResult: repository.CreateCakeResult{
				CakeID: 3,
				Err:    nil,
			},
			RepoGetCakeArg: repository.GetCakeArg{
				CakeID: 3,
			},
			RepoGetCakeResult: repository.GetCakeResult{
				Cake: &model.CakeModel{
					ID:          3,
					Title:       "BBB",
					Description: "BBB - BBB - BBB",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {

			// Mock Repo
			repo.Mock.On("CreateCake", tc.RepoCreateCakeArg.CakeForm).Return(tc.RepoCreateCakeResult.CakeID, tc.RepoCreateCakeResult.Err)
			repo.Mock.On("GetCake", tc.RepoGetCakeArg.CakeID).Return(tc.RepoGetCakeResult.Cake, tc.RepoGetCakeResult.Err)

			// Call Method
			cake, status, err := service.CreateCake(tc.Arg.CakeForm)

			// Assertion
			if tc.Result.Cake != nil {
				assert.Equal(t, tc.Result.Cake.ID, cake.ID)
			} else {
				assert.Nil(t, cake)
			}

			assert.Equal(t, tc.Result.Status, status)

			if tc.Result.Err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdateCakeService(t *testing.T) {

	type Arg struct {
		CakeID   int
		CakeForm *model.CakeForm
	}

	type Result struct {
		Cake   *model.CakeModel
		Status int
		Err    error
	}

	testCases := []struct {
		Name                 string
		Arg                  Arg
		Result               Result
		RepoUpdateCakeArg    repository.UpdateCakeArg
		RepoUpdateCakeResult repository.UpdateCakeResult
		RepoGetCakeArg       repository.GetCakeArg
		RepoGetCakeResult    repository.GetCakeResult
	}{
		{
			Name: "UpdateCake Success",
			Arg: Arg{
				CakeID: 3,
				CakeForm: &model.CakeForm{
					Title:       "CCC",
					Description: "CCC - CCC - CCC",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
			},
			Result: Result{
				Cake: &model.CakeModel{
					ID:          3,
					Title:       "CCC",
					Description: "CCC - CCC - CCC",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
				Status: http.StatusOK,
				Err:    nil,
			},
			RepoUpdateCakeArg: repository.UpdateCakeArg{
				ID: 3,
				CakeForm: &model.CakeForm{
					Title:       "CCC",
					Description: "CCC - CCC - CCC",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
			},
			RepoUpdateCakeResult: repository.UpdateCakeResult{
				CakeID: 3,
				Err:    nil,
			},
			RepoGetCakeArg: repository.GetCakeArg{
				CakeID: 3,
			},
			RepoGetCakeResult: repository.GetCakeResult{
				Cake: &model.CakeModel{
					ID:          3,
					Title:       "CCC",
					Description: "CCC - CCC - CCC",
					Rating:      1.8,
					Image:       "https://sdf.sg/fd.png",
				},
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {

			// Mock Repo
			repo.Mock.On("UpdateCake", tc.RepoUpdateCakeArg.ID, tc.RepoUpdateCakeArg.CakeForm).Return(tc.RepoUpdateCakeResult.CakeID, tc.RepoUpdateCakeResult.Err)
			repo.Mock.On("GetCake", tc.RepoGetCakeArg.CakeID).Return(tc.RepoGetCakeResult.Cake, tc.RepoGetCakeResult.Err)

			// Call Method
			cake, status, err := service.UpdateCake(tc.Arg.CakeID, tc.Arg.CakeForm)

			// Assertion
			if tc.Result.Cake != nil {
				assert.Equal(t, tc.Result.Cake.ID, cake.ID)
			} else {
				assert.Nil(t, cake)
			}

			assert.Equal(t, tc.Result.Status, status)

			if tc.Result.Err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDeleteCakeService(t *testing.T) {

	type Arg struct {
		CakeID int
	}

	type Result struct {
		CakeID int
		Status int
		Err    error
	}

	testCases := []struct {
		Name                 string
		Arg                  Arg
		Result               Result
		RepoDeleteCakeArg    repository.DeleteCakeArg
		RepoDeleteCakeResult repository.DeleteCakeResult
	}{
		{
			Name: "DeleteCake Success",
			Arg: Arg{
				CakeID: 2,
			},
			Result: Result{
				CakeID: 2,
				Status: http.StatusOK,
				Err:    nil,
			},
			RepoDeleteCakeArg: repository.DeleteCakeArg{
				CakeID: 2,
			},
			RepoDeleteCakeResult: repository.DeleteCakeResult{
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {

			// Mock Repo
			repo.Mock.On("DeleteCake", tc.RepoDeleteCakeArg.CakeID).Return(tc.RepoDeleteCakeResult.Err)

			// Call Method
			cakeID, status, err := service.DeleteCake(tc.Arg.CakeID)

			// Assertion
			assert.Equal(t, tc.Result.CakeID, cakeID)
			assert.Equal(t, tc.Result.Status, status)
			assert.Nil(t, err)

		})
	}
}
