package repository

import (
	"database/sql"
	"example/domain"
)

type IPersonRepository interface {
	Add(p domain.Person) error
}

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{db: database}
}

func (r *PersonRepository) Add(p domain.Person) error {
	_, err := r.db.Exec(`INSERT INTO users VALUES ($1,$2,$3)`, p.Id, p.Name, p.Age)
	return err
}
