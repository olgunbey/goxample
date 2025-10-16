package repository

import (
	"database/sql"
	"example/internal/common"
	"example/internal/models"
)

type PersonRepository struct {
	db                *sql.DB
	genericRepository *common.BaseRepository[models.Person]
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	genericRepo := common.NewGenericRepository[models.Person](database, "person")
	return &PersonRepository{db: database, genericRepository: genericRepo}
}

func (r *PersonRepository) FindByEmail(email string) (*models.Person, error) {
	query := "SELECT id, name, email FROM person WHERE email = $1"
	row := r.db.QueryRow(query, email)
	var person models.Person
	err := row.Scan(&person.Id, &person.Name, &person.Email)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepository) Add(item *models.Person) (int, error) {
	return r.genericRepository.Add(item)
}
func (r *PersonRepository) RemoveGetById(id int) (int, error) {
	return r.genericRepository.RemoveGetById(id)
}
