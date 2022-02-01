package services

import "github.com/prashant/layeredArc/models"

type Service interface {
	Create(c models.Customer) (models.Customer, error)
}
