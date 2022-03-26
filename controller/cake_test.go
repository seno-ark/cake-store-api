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
	"strconv"
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

	testCases := []struct {
		name       string
		cakeID     int
		mockResult *model.CakeModel
		mockErr    error
		status     int
		message    string
	}{
		{
			name:   "GetCake Success",
			cakeID: 2,
			mockResult: &model.CakeModel{
				ID:          2,
				Title:       "Cake A",
				Description: "AAA",
				Rating:      4.8,
				Image:       "https://gallery.com/asdadfs.jpg",
			},
			mockErr: nil,
			status:  http.StatusOK,
			message: "",
		},
		{
			name:       "GetCake Not Found",
			cakeID:     0,
			mockResult: nil,
			mockErr:    sql.ErrNoRows,
			status:     http.StatusNotFound,
			message:    "Cake not found",
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			repo.Mock.On("GetCake", tc.cakeID).Return(tc.mockResult, tc.mockErr)

			expectedResp := new(config.Response)
			if len(tc.message) > 0 {
				expectedResp.Message = tc.message
			}
			if tc.mockResult != nil {
				expectedResp.Data = tc.mockResult
			}
			expectedJson, _ := json.Marshal(expectedResp)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/cakes/:cake_id")
			c.SetParamNames("cake_id")
			c.SetParamValues(strconv.Itoa(tc.cakeID))

			if assert.NoError(t, controller.GetCake(c)) {
				assert.Equal(t, tc.status, rec.Code)
				assert.Equal(t, string(expectedJson), strings.TrimSpace(rec.Body.String()))
			}
		})

	}
}
