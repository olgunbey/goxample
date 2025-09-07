package usecase

import (
	"example/internal/models"
	"example/internal/person/dtos"
	"example/internal/person/repository"
)

type IPersonService interface {
	AddPerson(p models.Person) dtos.AddPersonResponseDto
}

type PersonService struct {
	repo repository.IPersonRepository
}

func NewPersonService(r repository.IPersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (s *PersonService) AddPerson(p *dtos.AddPersonRequestDto) *dtos.AddPersonResponseDto {
	addPerson := models.Person{Name: p.Name, Age: p.Age}
	result, err := s.repo.Add(&addPerson)
	if err != nil {
		return &dtos.AddPersonResponseDto{Id: 0, Message: err.Error(), Successfully: false}
	}
	id, err := result.LastInsertId()
	return &dtos.AddPersonResponseDto{Id: id, Message: "Person added successfully", Successfully: true}
}
