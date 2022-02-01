package customer

import (
	"github.com/prashant/layeredArc/models"
	"github.com/prashant/layeredArc/stores"
)

type service struct {
	store stores.Store
}

func New(store stores.Store) service {
	return service{store: store}
}

func (s service) Create(c models.Customer) (models.Customer, error) {
	return s.store.Create(c)
}
