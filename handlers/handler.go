package handlers

import (
	"goTestApi/repository"
)

type Handler struct {
	itemStore repository.ItemsI
}

func NewHandler(i repository.ItemsI) *Handler {
	return &Handler{
		itemStore: i,
	}
}
