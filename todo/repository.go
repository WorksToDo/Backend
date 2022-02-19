package todo

import "gorm.io/gorm"

type IRepository interface {
	GetTodos() ([]Todo, error)
	AddTodo(request CreateTodoRequest) (Todo, error)
}

type Repository struct {
	db * gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{db:db}
}

func (r *Repository) GetTodos() ([]Todo, error) {
	var todos []Todo
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&todos).Error; err != nil {
			return err
		}
		return nil
	})
	if err !=nil  {
		return nil, err
	}
	return todos, nil
}
func (r *Repository) AddTodo(request CreateTodoRequest) (Todo, error ) {
	todo := Todo{Task: request.Task}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&todo).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}