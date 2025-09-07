package repository

import (
	"database/sql"
	"example/internal/models"
)

type IPersonRepository interface {
	Add(p *models.Person) (sql.Result, error)
}

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{db: database}
}

func (r *PersonRepository) Add(p *models.Person) (sql.Result, error) {
	result, err := r.db.Exec(`INSERT INTO users VALUES ($1,$2,$3)`, p.Id, p.Name, p.Age)
	if err != nil {
		return nil, err
	}
	return result, nil
}
