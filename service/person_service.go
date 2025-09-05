package service

import (
	"example/domain"
	"example/repository"
)

type IPersonService interface {
	AddPerson(p domain.Person) error
}

type PersonService struct {
	repo repository.IPersonRepository
}

func NewPersonService(r repository.IPersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (s *PersonService) AddPerson(p domain.Person) error {
	return s.repo.Add(p)
}
