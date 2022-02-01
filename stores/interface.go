package stores

import "github.com/prashant/layeredArc/models"

type Store interface {
	Create(c models.Customer) (models.Customer, error)
}
