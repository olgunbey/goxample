package person

import (
	"database/sql"
)

type IPersonRepository interface {
	Add(p Person) error
}

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{db: database}
}

func (r *PersonRepository) Add(p Person) error {
	_, err := r.db.Exec(`INSERT INTO users VALUES ($1,$2,$3)`, p.Id, p.Name, p.Age)
	return err
}
