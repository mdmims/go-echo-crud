package models

import "time"

type Items struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Price       float64   `db:"price"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}
