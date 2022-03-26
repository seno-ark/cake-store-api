package repository

import (
	"cake-store-api/config"
	"cake-store-api/model"
)

func (r *Repo) CountCake() (int, error) {

	var cakeCount int

	err := r.DB.Get(&cakeCount, "SELECT COUNT(id) FROM cakes")
	if err != nil {
		return cakeCount, err
	}

	return cakeCount, nil
}

func (r *Repo) GetCakeList(params config.M) ([]*model.CakeModel, error) {

	var cakes = []*model.CakeModel{}

	var args = []interface{}{
		params["limit"],
		params["offset"],
	}

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes ORDER BY updated_at DESC LIMIT ?, ?"

	err := r.DB.Select(&cakes, query, args...)
	if err != nil {
		return cakes, err
	}

	return cakes, nil
}

func (r *Repo) GetCake(cakeID int) (*model.CakeModel, error) {

	var cake = &model.CakeModel{}

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"

	err := r.DB.Get(cake, query, cakeID)
	if err != nil {
		return cake, err
	}

	return cake, nil
}

func (r *Repo) CreateCake(cakeForm *model.CakeForm) (int, error) {

	query := "INSERT INTO cakes (title, description, rating, image) VALUES (:title, :description, :rating, :image)"

	res, err := r.DB.NamedExec(query,
		map[string]interface{}{
			"title":       cakeForm.Title,
			"description": cakeForm.Description,
			"rating":      cakeForm.Rating,
			"image":       cakeForm.Image,
		})

	if err != nil {
		return 0, err
	}

	cakeID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(cakeID), nil
}

func (r *Repo) UpdateCake(cakeID int, cakeForm *model.CakeForm) (int, error) {

	query := "UPDATE cakes SET title = :title, description = :description, rating = :rating, image = :image, updated_at = now() WHERE id = :id"

	_, err := r.DB.NamedExec(query,
		map[string]interface{}{
			"id":          cakeID,
			"title":       cakeForm.Title,
			"description": cakeForm.Description,
			"rating":      cakeForm.Rating,
			"image":       cakeForm.Image,
		})

	if err != nil {
		return 0, err
	}

	return cakeID, nil
}

func (r *Repo) DeleteCake(cakeID int) error {

	_, err := r.DB.Exec("DELETE FROM cakes WHERE id = ?", cakeID)
	if err != nil {
		return err
	}

	return nil
}