package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-backend/postgre"
	"todo-backend/todo"
)

// repository

func TestNewRepositoryCreatedByGivenDBSuccessfully(t *testing.T) {
	db,_ := postgre.NewPostgreDB(postgre.DefaultConfig)
	repository := todo.NewRepository(db)
	assert.NotNil(t, repository)
}
func TestWhenICallGetTodosOnRepositoryItReturnsTodosWithoutError(t *testing.T) {
	db,_ := postgre.NewPostgreDB(postgre.DefaultConfig)
	repository := todo.NewRepository(db)
	todos, err := repository.GetTodos()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(todos),1)
}
func TestWhenICallAddTodoOnRepositoryWithValidRequestItReturnsAddedTodoWithoutError(t *testing.T) {
	db,_ := postgre.NewPostgreDB(postgre.DefaultConfig)
	repository := todo.NewRepository(db)

	mockRequest := todo.CreateTodoRequest{Task: "drink some milk"}
	addedTodo,err := repository.AddTodo(mockRequest)
	assert.Nil(t, err)
	assert.NotEqual(t,addedTodo, todo.Todo{})
}