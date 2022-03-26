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

func TestGetCake(t *testing.T) {

	testCases := []struct {
		name    string
		arg     int
		result  *model.CakeModel
		status  int
		err     error
		mockErr error
	}{
		{
			name: "GetCake Success",
			arg:  2,
			result: &model.CakeModel{
				ID: 2,
			},
			status:  http.StatusOK,
			err:     nil,
			mockErr: nil,
		},
		{
			name:    "GetCake NotFound",
			arg:     0,
			result:  nil,
			status:  http.StatusNotFound,
			err:     errors.New(config.MSG_ERROR_CAKE_NOT_FOUND),
			mockErr: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			repo.Mock.On("GetCake", tc.arg).Return(tc.result, tc.mockErr)

			result, status, err := service.GetCake(tc.arg)

			if tc.result != nil {
				assert.Equal(t, tc.result.ID, result.ID)
			} else {
				assert.Nil(t, result)
			}

			assert.Equal(t, tc.status, status)

			if tc.err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
