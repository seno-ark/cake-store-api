package repository

import (
	"cake-store-api/config"
	"cake-store-api/model"
)

func (m *MySql) CountCake() (int, error) {

	var cakeCount int

	err := m.DB.Get(&cakeCount, "SELECT COUNT(id) FROM cakes")

	return cakeCount, err
}

func (m *MySql) GetCakeList(params config.M) ([]*model.CakeModel, error) {

	var cakes = []*model.CakeModel{}

	var args = []interface{}{
		params["limit"],
		params["offset"],
	}

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes ORDER BY updated_at DESC LIMIT ?, ?"

	err := m.DB.Select(&cakes, query, args...)

	return cakes, err
}

func (m *MySql) GetCake(cakeID int) (*model.CakeModel, error) {

	var cake = &model.CakeModel{}

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"

	err := m.DB.Get(cake, query, cakeID)

	return cake, err
}

func (m *MySql) CreateCake(cakeForm *model.CakeForm) (int, error) {

	query := "INSERT INTO cakes (title, description, rating, image) VALUES (:title, :description, :rating, :image)"

	res, err := m.DB.NamedExec(query,
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

func (m *MySql) UpdateCake(cakeID int, cakeForm *model.CakeForm) (int, error) {

	query := "UPDATE cakes SET title = :title, description = :description, rating = :rating, image = :image, updated_at = now() WHERE id = :id"

	_, err := m.DB.NamedExec(query,
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

func (m *MySql) DeleteCake(cakeID int) error {

	_, err := m.DB.Exec("DELETE FROM cakes WHERE id = ?", cakeID)

	return err
}
