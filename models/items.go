package models

import "time"

type Items struct {
	name        string
	price       float64
	description string
	createdAt   time.Time
}
