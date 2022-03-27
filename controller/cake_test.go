package controller

import (
	"cake-store-api/config"
	"cake-store-api/model"
	"cake-store-api/repository"
	"cake-store-api/service"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repo       = &repository.RepositoryMock{Mock: mock.Mock{}}
	svc        = service.Service{Repo: repo}
	controller = Controller{Svc: svc}
)

func TestGetCakeController(t *testing.T) {

	type Arg struct {
		CakeID string
	}

	type Result struct {
		Status  int
		Message string
		Body    string
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
				CakeID: "2",
			},
			RepoGetCakeArg: repository.GetCakeArg{
				CakeID: 2,
			},
			RepoGetCakeResult: repository.GetCakeResult{
				Cake: &model.CakeModel{
					ID:          2,
					Title:       "Cake A",
					Description: "AAA",
					Rating:      4.8,
					Image:       "https://gallery.com/asdadfs.jpg",
				},
				Err: nil,
			},
			Result: Result{
				Status: http.StatusOK,
			},
		},
		{
			Name: "GetCake Not Found",
			Arg: Arg{
				CakeID: "0",
			},
			RepoGetCakeArg: repository.GetCakeArg{
				CakeID: 0,
			},
			RepoGetCakeResult: repository.GetCakeResult{
				Cake: nil,
				Err:  sql.ErrNoRows,
			},
			Result: Result{
				Status:  http.StatusNotFound,
				Message: "Cake not found",
			},
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			// Mock Repo
			repo.Mock.On("GetCake", tc.RepoGetCakeArg.CakeID).Return(tc.RepoGetCakeResult.Cake, tc.RepoGetCakeResult.Err)

			// Generate expected json body
			expectedResp := new(config.Response)
			if len(tc.Result.Message) > 0 {
				expectedResp.Message = tc.Result.Message
			}
			if tc.RepoGetCakeResult.Cake != nil {
				expectedResp.Data = tc.RepoGetCakeResult.Cake
			}
			expectedJson, _ := json.Marshal(expectedResp)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/cakes/:cake_id")

			// Set Query params
			c.SetParamNames("cake_id")
			c.SetParamValues(tc.Arg.CakeID)

			// Call Method
			if assert.NoError(t, controller.GetCake(c)) {
				// Assertion
				assert.Equal(t, tc.Result.Status, rec.Code)
				assert.Equal(t, string(expectedJson), strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}
