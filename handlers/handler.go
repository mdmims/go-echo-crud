package handlers

import (
	"github.com/mdmims/go-echo-crud/repository"
)

type Handler struct {
	itemStore repository.ItemsI
}

func NewHandler(i repository.ItemsI) *Handler {
	return &Handler{
		itemStore: i,
	}
}
