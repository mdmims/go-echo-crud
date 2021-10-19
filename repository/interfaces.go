package repository

import (
	"github.com/mdmims/go-echo-crud/models"
)

type ItemsI interface {
	GetAll() ([]models.Items, error)
	GetById(id int) (*models.Items, error)
	Create(i *models.Items) (*models.Items, error)
}
