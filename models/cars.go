package models

import (
	"database/sql"
	"strings"
)

type Car struct {
	ID    int
	Brand string
	Model string
}

type CarService struct {
	DB *sql.DB
}

func (c *CarService) create(brand, model string) (*Car, error) {
	brand = strings.ToLower(brand)
	model = strings.ToLower(model)

	car := Car{
		Brand: brand,
		Model: model,
	}

	row := c.DB.QueryRow(`INSERT INTO cars (brand, model) VALUES ($1, $2) RETURNING id`, brand, model)
	err := row.Scan(&car.ID)
	return &car, err
}
