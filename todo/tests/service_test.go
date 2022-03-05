package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-backend/todo"
	"todo-backend/todo/mocks"
)

// new service olusturma testi
// get todos isteğinin testi
// post insert todo testi

func TestNewServiceCreatedByGivenRepository(t *testing.T){
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockRepo := mocks.NewMockIRepository(mockController)
	service, err := todo.NewService(mockRepo)
	assert.Nil(t, err)
	assert.NotNil(t, service)
}

func TestWhenICallGetTodosOnServiceItReturnsTodosWithoutError(t *testing.T) {
	test := struct {
		expectedTodos []todo.Todo
	}{
		expectedTodos: []todo.Todo{
			{
				ID:   0,
				Task: "buy some milk",
			},
			{
				ID:   1,
				Task: "buy some water",
			},
		},
	}
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepo := mocks.NewMockIRepository(mockController)
	mockRepo.EXPECT().GetTodos().Return(test.expectedTodos,nil)

	service,_ := todo.NewService(mockRepo)
	todos, err := service.GetTodos()

	assert.Nil(t, err)
	assert.Equal(t, todos, test.expectedTodos)
}

func TestWhenICallAddTodoOnServiceWithValidRequestItReturnsAddedTodoWithoutError(t *testing.T) {
	test := struct {
		expectedTodo todo.Todo
		request todo.CreateTodoRequest
	}{
		expectedTodo: todo.Todo{
			ID:   0,
			Task: "buy some milk",
		},
		request: todo.CreateTodoRequest{
			Task: "buy some milk",
		},
	}
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepo := mocks.NewMockIRepository(mockController)
	mockRepo.EXPECT().AddTodo(test.request).Return(test.expectedTodo,nil)

	service,_ := todo.NewService(mockRepo)
	todo, err := service.AddTodo(test.request)

	assert.Nil(t, err)
	assert.Equal(t, todo, test.expectedTodo)
}


