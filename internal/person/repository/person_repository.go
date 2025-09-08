package repository

import (
	"context"
	"database/sql"
	"example/internal/models"
	"time"
)

type IPersonRepository interface {
	Add(p *models.Person) (sql.Result, error)
	RemoveGetById(id int) (int, error)
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

func (r *PersonRepository) RemoveGetById(id int) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var deletedId int

	err := r.db.QueryRowContext(ctx, `DELETE FROM users WHERE id=$1 RETURNING id`, id).Scan(&deletedId)
	if err != nil {
		return 0, err
	}
	return deletedId, nil
}
