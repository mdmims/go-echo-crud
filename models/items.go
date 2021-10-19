package models

import "time"

type Items struct {
	ID			int64
	Name        string
	Price       float64
	Description string
	CreatedAt   time.Time
}
