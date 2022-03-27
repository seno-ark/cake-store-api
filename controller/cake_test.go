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
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCakeListController(t *testing.T) {
	var (
		repo       = &repository.RepositoryMock{Mock: mock.Mock{}}
		svc        = service.Service{Repo: repo}
		controller = Controller{Svc: svc}
	)

	type Arg struct {
		Page  int
		Count int
	}

	type Result struct {
		Status  int
		Message string
		Body    string
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
				Page:  1,
				Count: 10,
			},
			Result: Result{
				Status: http.StatusOK,
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

			e := echo.New()

			// Set Query params
			q := make(url.Values)
			q.Set("page", strconv.Itoa(tc.Arg.Page))
			q.Set("count", strconv.Itoa(tc.Arg.Count))

			req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/cakes")

			// Mock Repo
			repo.Mock.On("GetCakeList", tc.RepoGetCakeListArg.Params).Return(tc.RepoGetCakeListResult.Cakes, tc.RepoGetCakeListResult.Err)
			repo.Mock.On("CountCake").Return(tc.RepoCountCakeResult.Count, tc.RepoGetCakeListResult.Err)

			// Generate expected json body
			expectedResp := new(config.Response)
			expectedResp.Message = tc.Result.Message
			expectedResp.Data = config.M{
				"page":       tc.Arg.Page,
				"count":      tc.Arg.Count,
				"cakes":      tc.RepoGetCakeListResult.Cakes,
				"total_data": tc.RepoCountCakeResult.Count,
			}
			expectedJson, _ := json.Marshal(expectedResp)

			// Call Method
			if assert.NoError(t, controller.GetCakeList(c)) {
				// Assertion
				assert.Equal(t, tc.Result.Status, rec.Code)
				assert.Equal(t, string(expectedJson), strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}

func TestGetCakeController(t *testing.T) {
	var (
		repo       = &repository.RepositoryMock{Mock: mock.Mock{}}
		svc        = service.Service{Repo: repo}
		controller = Controller{Svc: svc}
	)

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
			Result: Result{
				Status: http.StatusOK,
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
		},
		{
			Name: "GetCake Not Found",
			Arg: Arg{
				CakeID: "0",
			},
			Result: Result{
				Status:  http.StatusNotFound,
				Message: "Cake not found",
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

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

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

			// Set Path params
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

func TestCreateCakeController(t *testing.T) {
	var (
		repo       = &repository.RepositoryMock{Mock: mock.Mock{}}
		svc        = service.Service{Repo: repo}
		controller = Controller{Svc: svc}
	)

	type Arg struct {
		Body string
	}

	type Result struct {
		Status  int
		Message string
		Body    string
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
				Body: `{"title":"BBB","description":"BBB - BBB - BBB","rating":3.4,"image":"https://sdf.sg/fd.png"}`,
			},
			Result: Result{
				Status: http.StatusCreated,
			},
			RepoCreateCakeArg: repository.CreateCakeArg{
				CakeForm: &model.CakeForm{
					Title:       "BBB",
					Description: "BBB - BBB - BBB",
					Rating:      3.4,
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
					Rating:      3.4,
					Image:       "https://sdf.sg/fd.png",
				},
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.Arg.Body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/cakes")

			// Mock Repo
			repo.Mock.On("CreateCake", tc.RepoCreateCakeArg.CakeForm).Return(tc.RepoCreateCakeResult.CakeID, tc.RepoCreateCakeResult.Err)
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

			// Call Method
			if assert.NoError(t, controller.CreateCake(c)) {
				// Assertion
				assert.Equal(t, tc.Result.Status, rec.Code)
				assert.Equal(t, string(expectedJson), strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}

func TestUpdateCakeController(t *testing.T) {
	var (
		repo       = &repository.RepositoryMock{Mock: mock.Mock{}}
		svc        = service.Service{Repo: repo}
		controller = Controller{Svc: svc}
	)

	type Arg struct {
		CakeID string
		Body   string
	}

	type Result struct {
		Status  int
		Message string
		Body    string
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
				CakeID: "3",
				Body:   `{"title":"CCC","description":"CCC - CCC - CCC","rating":4.4,"image":"https://sdf.sg/fd.png"}`,
			},
			Result: Result{
				Status: http.StatusOK,
			},
			RepoUpdateCakeArg: repository.UpdateCakeArg{
				ID: 3,
				CakeForm: &model.CakeForm{
					Title:       "CCC",
					Description: "CCC - CCC - CCC",
					Rating:      4.4,
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
					Rating:      4.4,
					Image:       "https://sdf.sg/fd.png",
				},
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {

			e := echo.New()
			req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(tc.Arg.Body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/cakes/:cake_id")

			// Set Path params
			c.SetParamNames("cake_id")
			c.SetParamValues(tc.Arg.CakeID)

			// Mock Repo
			repo.Mock.On("UpdateCake", tc.RepoUpdateCakeArg.ID, tc.RepoUpdateCakeArg.CakeForm).Return(tc.RepoUpdateCakeResult.CakeID, tc.RepoUpdateCakeResult.Err)
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

			// Call Method
			if assert.NoError(t, controller.UpdateCake(c)) {
				// Assertion
				assert.Equal(t, tc.Result.Status, rec.Code)
				assert.Equal(t, string(expectedJson), strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}

func TestDeleteCakeController(t *testing.T) {
	var (
		repo       = &repository.RepositoryMock{Mock: mock.Mock{}}
		svc        = service.Service{Repo: repo}
		controller = Controller{Svc: svc}
	)

	type Arg struct {
		CakeID string
	}

	type Result struct {
		Status  int
		Message string
		Body    string
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
				CakeID: "2",
			},
			Result: Result{
				Status:  http.StatusOK,
				Message: config.MSG_SUCCESS,
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

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/cakes/:cake_id")

			// Set Path params
			c.SetParamNames("cake_id")
			c.SetParamValues(tc.Arg.CakeID)

			// Mock Repo
			repo.Mock.On("DeleteCake", tc.RepoDeleteCakeArg.CakeID).Return(tc.RepoDeleteCakeResult.Err)

			// Generate expected json body
			expectedResp := new(config.Response)
			if len(tc.Result.Message) > 0 {
				expectedResp.Message = tc.Result.Message
			}
			expectedJson, _ := json.Marshal(expectedResp)

			// Call Method
			if assert.NoError(t, controller.DeleteCake(c)) {
				// Assertion
				assert.Equal(t, tc.Result.Status, rec.Code)
				assert.Equal(t, string(expectedJson), strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}
