package service

import (
	"cake-store-api/config"
	"cake-store-api/model"
	"database/sql"
	"errors"
	"log"
	"net/http"
)

func (s *Service) GetCakeList(params config.M) (*config.M, int, error) {

	page := params["page"].(int)
	count := params["count"].(int)

	if page < 1 {
		page = 1
	}
	if count < 1 || count > config.MAX_PAGINATION_COUNT {
		count = config.MAX_PAGINATION_COUNT
	}

	params = config.M{
		"limit":  (page - 1) * count,
		"offset": count,
	}

	cakes, err := s.Repo.GetCakeList(params)
	if err != nil {
		log.Println("ERROR GetCakeList: GetCakeList", err.Error())
		return nil, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
	}

	var totalData int
	if len(cakes) > 0 {
		totalData, err = s.Repo.CountCake()
		if err != nil {
			log.Println("ERROR GetCakeList: CountCake", err.Error())
			return nil, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
		}
	}

	result := &config.M{
		"cakes":      cakes,
		"total_data": totalData,
		"page":       page,
		"count":      count,
	}

	return result, http.StatusOK, nil
}

func (s *Service) GetCake(cakeID int) (*model.CakeModel, int, error) {

	cake, err := s.Repo.GetCake(cakeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New(config.MSG_ERROR_CAKE_NOT_FOUND)
		}
		log.Println("ERROR GetCake", err.Error())
		return nil, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
	}

	return cake, http.StatusOK, nil
}

func (s *Service) CreateCake(cakeForm *model.CakeForm) (*model.CakeModel, int, error) {

	cakeID, err := s.Repo.CreateCake(cakeForm)
	if err != nil {
		log.Println("ERROR CreateCake: CreateCake", err.Error())
		return nil, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
	}

	cakeDetail, err := s.Repo.GetCake(cakeID)
	if err != nil {
		log.Println("ERROR CreateCake: GetCake", err.Error())
		return nil, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
	}

	return cakeDetail, http.StatusCreated, nil
}

func (s *Service) UpdateCake(cakeID int, cakeForm *model.CakeForm) (*model.CakeModel, int, error) {

	_, err := s.Repo.UpdateCake(cakeID, cakeForm)
	if err != nil {
		log.Println("ERROR UpdateCake: UpdateCake", err.Error())
		return nil, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
	}

	cakeDetail, err := s.Repo.GetCake(cakeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New(config.MSG_ERROR_CAKE_NOT_FOUND)
		}
		log.Println("ERROR UpdateCake: GetCake", err.Error())
		return nil, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
	}

	return cakeDetail, http.StatusOK, nil
}

func (s *Service) DeleteCake(cakeID int) (int, int, error) {
	err := s.Repo.DeleteCake(cakeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, http.StatusNotFound, errors.New(config.MSG_ERROR_CAKE_NOT_FOUND)
		}
		log.Println("ERROR DeleteCake", err.Error())
		return cakeID, http.StatusInternalServerError, errors.New(config.MSG_ERROR_DATABASE)
	}

	return cakeID, http.StatusOK, nil
}
