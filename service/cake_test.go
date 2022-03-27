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
			result, status, err := service.GetCake(tc.Arg.CakeID)

			// Assertion
			if tc.Result.Cake != nil {
				assert.Equal(t, tc.Result.Cake.ID, result.ID)
			} else {
				assert.Nil(t, result)
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
