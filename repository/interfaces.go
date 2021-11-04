package repository

import (
	"time"

	"github.com/mdmims/go-echo-crud/models"
)

type ItemsI interface {
	GetAll() ([]models.Items, error)
	GetById(id int) (*models.Items, error)
	Create(i *models.Items) (*models.Items, error)
}

type ServerCache interface {
	Get(key string) (interface{}, bool)
	Set(key string, data interface{}) error
	SetTTL(ttl time.Duration)
	SetCacheSizeLimit(limit int)
	Count() int
}