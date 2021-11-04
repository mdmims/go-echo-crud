package handlers

import (
	"github.com/mdmims/go-echo-crud/repository"
)

type Handler struct {
	itemStore repository.ItemsI
	cache       repository.ServerCache
}

func NewHandler(i repository.ItemsI, c repository.ServerCache) *Handler {
	return &Handler{
		itemStore: i,
		cache: c,
	}
}
