package controller

import (
	"cake-store-api/config"
	"cake-store-api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (c *Controller) GetCakeList(ctx echo.Context) error {
	response := new(config.Response)

	pageParam := ctx.QueryParam("page")
	countParam := ctx.QueryParam("count")

	var page, count int
	var err error

	if len(pageParam) > 0 {
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, "page")
			return ctx.JSON(http.StatusBadRequest, response)
		}
	}
	if len(countParam) > 0 {
		count, err = strconv.Atoi(countParam)
		if err != nil {
			response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, "count")
			return ctx.JSON(http.StatusBadRequest, response)
		}
	}

	params := config.M{
		"page":  page,
		"count": count,
	}

	result, status, err := c.Svc.GetCakeList(params)

	if err != nil {
		response.Message = err.Error()
	} else {
		response.Data = result
	}

	return ctx.JSON(status, response)
}

func (c *Controller) GetCake(ctx echo.Context) error {
	response := new(config.Response)

	cakeIDParam := ctx.Param("cake_id")
	cakeID, err := strconv.Atoi(cakeIDParam)
	if err != nil {
		response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, "cake_id")
		return ctx.JSON(http.StatusBadRequest, response)
	}

	result, status, err := c.Svc.GetCake(cakeID)

	if err != nil {
		response.Message = err.Error()
	} else {
		response.Data = result
	}

	return ctx.JSON(status, response)
}

func (c *Controller) CreateCake(ctx echo.Context) error {
	response := new(config.Response)

	cakeForm := new(model.CakeForm)
	if err := ctx.Bind(cakeForm); err != nil {
		response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, "json")
		return ctx.JSON(http.StatusBadRequest, response)
	}

	err := cakeForm.Validate()
	if err != nil {
		response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, err.Error())
		return ctx.JSON(http.StatusBadRequest, response)
	}

	result, status, err := c.Svc.CreateCake(cakeForm)

	if err != nil {
		response.Message = err.Error()
	} else {
		response.Data = result
	}

	return ctx.JSON(status, response)
}

func (c *Controller) UpdateCake(ctx echo.Context) error {
	response := new(config.Response)

	cakeIDParam := ctx.Param("cake_id")
	cakeID, err := strconv.Atoi(cakeIDParam)
	if err != nil {
		response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, "cake_id")
		return ctx.JSON(http.StatusBadRequest, response)
	}

	cakeForm := new(model.CakeForm)
	if err := ctx.Bind(cakeForm); err != nil {
		response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, "json")
		return ctx.JSON(http.StatusBadRequest, response)
	}

	err = cakeForm.Validate()
	if err != nil {
		response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, err.Error())
		return ctx.JSON(http.StatusBadRequest, response)
	}

	result, status, err := c.Svc.UpdateCake(cakeID, cakeForm)

	if err != nil {
		response.Message = err.Error()
	} else {
		response.Data = result
	}

	return ctx.JSON(status, response)
}

func (c *Controller) DeleteCake(ctx echo.Context) error {
	response := new(config.Response)

	cakeIDParam := ctx.Param("cake_id")
	cakeID, err := strconv.Atoi(cakeIDParam)
	if err != nil {
		response.Message = fmt.Sprintf("%s: %s", config.MSG_ERROR_INVALID_DATA, "cake_id")
		return ctx.JSON(http.StatusBadRequest, response)
	}

	_, status, err := c.Svc.DeleteCake(cakeID)

	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = config.MSG_SUCCESS
	}

	return ctx.JSON(status, response)
}
