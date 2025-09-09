package repository

import (
	"context"
	"database/sql"
	"example/internal/models"
	"time"
)

type IPersonRepository interface {
	Add(p *models.Person) (int, error)
	RemoveGetById(id int) (int, error)
}

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{db: database}
}

func (r *PersonRepository) Add(p *models.Person) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var returnId int
	err := r.db.QueryRowContext(ctx, `INSERT INTO users (username,age) VALUES ($1,$2) RETURNING id`, p.Name, p.Age).Scan(&returnId)

	return returnId, err
}

func (r *PersonRepository) RemoveGetById(id int) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var deletedId int
	err := r.db.QueryRowContext(ctx, `DELETE FROM users WHERE id=$1 RETURNING id`, id).Scan(&deletedId)
	return deletedId, err
}
