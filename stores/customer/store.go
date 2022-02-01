package customer

import (
	"database/sql"
	"log"

	"github.com/prashant/layeredArc/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s store) Create(c models.Customer) (models.Customer, error) {
	_, err := s.db.Exec("INSERT INTO CUSTOMER (Name ,Phone,Address) VALUES (?,?,?)", &c.Name, &c.PhoneNo, &c.Address)
	if err != nil {
		log.Fatal(err)
	}
	return models.Customer{}, nil
}
