package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		DB: db,
	}
}
func (c *Category) Create(name, description string) (*Category, error) {
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return nil, err
	}

	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}
func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.DB.Query("SELECT * FROM categories;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var categories []Category

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		categories = append(categories, Category{
			ID:          id,
			Name:        name,
			Description: description,
		})
	}

	return categories, nil
}
func (c *Category) FindById(id string) (*Category, error) {
	row := c.DB.QueryRow("SELECT * FROM categories WHERE id = $1", id)

	var categoryId, name, description string
	row.Scan(&id, &name, &description)

	return &Category{
		ID:          categoryId,
		Name:        name,
		Description: description,
	}, nil
}
