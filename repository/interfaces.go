package repository

import (
	"goTestApi/models"
)

type ItemsI interface {
	GetAll() (*models.Items, error)
	Create(i *models.Items) (*models.Items, error)
}
