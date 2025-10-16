package usecase

import (
	"database/sql"
	"example/internal/models"
	"example/internal/person/dtos"
	"example/internal/person/repository"
)

type IPersonService interface {
	AddPerson(p *dtos.AddPersonRequestDto) dtos.AddPersonResponseDto
	RemovePersonGetById(id int) *dtos.RemovePersonGetByIdResponseDto
}

type PersonService struct {
	repo *repository.PersonRepository
}

func NewPersonService(r *repository.PersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (s *PersonService) AddPerson(p *dtos.AddPersonRequestDto) *dtos.AddPersonResponseDto {
	addPerson := models.Person{Name: p.Name, Age: p.Age}

	result, err := s.repo.Add(&addPerson)
	if err != nil {
		return &dtos.AddPersonResponseDto{Id: 0, Message: err.Error(), Successfully: false, StatusCode: 500}
	}
	return &dtos.AddPersonResponseDto{Id: int64(result), Message: "Person added successfully", Successfully: true, StatusCode: 201}
}
func (s *PersonService) RemovePersonGetById(id int) *dtos.RemovePersonGetByIdResponseDto {
	valueId, err := s.repo.RemoveGetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return &dtos.RemovePersonGetByIdResponseDto{Id: 0, Message: "Person not found", Successfully: false, StatusCode: 404}
		}
		return &dtos.RemovePersonGetByIdResponseDto{Id: 0, Message: err.Error(), Successfully: false, StatusCode: 500}
	}
	return &dtos.RemovePersonGetByIdResponseDto{Id: valueId, Message: "Person removed successfully", Successfully: true, StatusCode: 200}
}
