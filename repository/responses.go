package repository

import (
	"time"

	"goTestApi/models"
)

// itemsResponse represents items data as JSON
type itemsResponse struct {
	ID          int64     `json:"ID"`
	Name        string    `json:"Name"`
	Price       float64   `json:"Price"`
	Description string    `json:"Description"`
	CreatedAt   time.Time `json:"CreatedAt"`
}

// NewitemsResponse creates itemsResponse object
func NewitemsResponse(i *models.Items) *itemsResponse {
	m := new(itemsResponse)
	m.ID = i.ID
	m.Name = i.Name
	m.Price = i.Price
	m.Description = i.Description
	m.CreatedAt = i.CreatedAt
	return m
}
