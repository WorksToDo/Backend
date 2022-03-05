package todo

import "errors"

type IService interface {
	GetTodos() ([]Todo, error)
	AddTodo(request CreateTodoRequest) (Todo, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) (*Service,error) {
	if repository == nil {
		return nil, errors.New("Repository is null")
	}
	return &Service{
		repository: repository,
	}, nil
}

func (s *Service) GetTodos() ([]Todo, error){
	return s.repository.GetTodos()
}

func (s *Service) AddTodo(request CreateTodoRequest) (Todo, error) {
	return s.repository.AddTodo(request)
}

