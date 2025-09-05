package person

type IPersonService interface {
	AddPerson(p Person) error
}

type PersonService struct {
	repo IPersonRepository
}

func NewPersonService(r IPersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (s *PersonService) AddPerson(p Person) error {
	return s.repo.Add(p)
}
