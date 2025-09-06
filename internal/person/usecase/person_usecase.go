package usecase

import (
	"example/internal/models"
	"example/internal/person/repository"
)

type IPersonService interface {
	AddPerson(p models.Person) error
}

type PersonService struct {
	repo repository.IPersonRepository
}

func NewPersonService(r repository.IPersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (s *PersonService) AddPerson(p models.Person) error {
	return s.repo.Add(p)
}
