package repository

import (
	"database/sql"

	"goTestApi/models"
)

type ItemStore struct {
	db *sql.DB
}

func NewItemStore(db *sql.DB) *ItemStore {
	return &ItemStore{
		db: db,
	}
}

func (t *ItemStore) GetAll() (*models.Items, error) {
	var m models.Items

	if err := t.db.QueryRow(
		"select id, name, price, description, created_at from items",
	).Scan(&m.ID, &m.Name, &m.Price, &m.Description, &m.CreatedAt); err != nil {
		return nil, err
	}
	return &m, nil
}
