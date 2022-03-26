package model

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type CakeModel struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Rating      float64   `db:"rating" json:"rating"`
	Image       string    `db:"image" json:"image"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CakeForm struct {
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Rating      float64 `db:"rating" json:"rating"`
	Image       string  `db:"image" json:"image"`
}

func (ck *CakeForm) Validate() error {

	if len(ck.Title) == 0 {
		return errors.New("Invalid Title")
	}
	ck.Title = strings.TrimSpace(ck.Title)

	if len(ck.Description) > 0 {
		ck.Description = strings.TrimSpace(ck.Description)
	}

	if ck.Rating < 0 {
		ck.Rating = 0
	}

	if len(ck.Image) > 0 {
		regex, err := regexp.Compile(`(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|jpeg|png)`)
		if err != nil {
			return err
		}

		isValid := regex.MatchString(ck.Image)
		if !isValid {
			return errors.New("Invalid Image URL")
		}
	}

	return nil
}
