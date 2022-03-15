package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"todo-backend/postgre"
	"todo-backend/todo"
)

// repository

func TestNewRepositoryCreatedByGivenDBSuccessfully(t *testing.T) {
	db,_ := postgre.NewMockPostgreDB()
	repository := todo.NewRepository(db)
	assert.NotNil(t, repository)
}
func TestWhenICallGetTodosOnRepositoryItReturnsTodosWithoutError(t *testing.T) {
	//db,_ := postgre.NewPostgreDB(postgre.DefaultConfig)
	testTodo := todo.Todo{
		ID:   0,
		Task: "",
	}
	db,mock := postgre.NewMockPostgreDB()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "public"."todos"`)).WillReturnRows(sqlmock.NewRows([]string{"id","task"}).AddRow(testTodo.ID,testTodo.Task))
	mock.ExpectCommit()

	repository := todo.NewRepository(db)
	todos, err := repository.GetTodos()
	assert.Nil(t, err)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(todos),1)
}
func TestWhenICallAddTodoOnRepositoryWithValidRequestItReturnsAddedTodoWithoutError(t *testing.T) {
	db,mock := postgre.NewMockPostgreDB()
	repository := todo.NewRepository(db)

	mockRequest := todo.CreateTodoRequest{Task: "drink some milk"}
	expectedTodo := todo.Todo{
		ID:   2,
		Task: "drink some milk",
	}
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "public"."todos" ("task") VALUES ($1) RETURNING "id"`)).
		WithArgs(mockRequest.Task).WillReturnRows(sqlmock.NewRows([]string{"id","task"}).
		AddRow(expectedTodo.ID,expectedTodo.Task))
	mock.ExpectCommit()

	addedTodo,err := repository.AddTodo(mockRequest)
	assert.Nil(t, err)
	assert.NotEqual(t,addedTodo, todo.Todo{})

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}