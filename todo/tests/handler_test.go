package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-backend/todo"
	"todo-backend/todo/mocks"
)

func TestNewHandlerCreatedByGivenService(t *testing.T){
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockService := mocks.NewMockIService(mockController)
	handler, err := todo.NewHandler(mockService)
	assert.Nil(t, err)
	assert.NotNil(t, handler)
}

func TestWhenGetTodosRequestArrivesWithValidRequestItReturnsTodosWithoutError(t *testing.T){
	request := struct {
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
	mockService := mocks.NewMockIService(mockController)
	mockService.EXPECT().GetTodos().Return(request.expectedTodos, nil)

	handler, _ := todo.NewHandler(mockService)
	app := fiber.New()
	handler.RegisterRoutes(app)

	testRequest := MakeRequestWithoutBody(http.MethodGet, "/todos")
	response, err := app.Test(testRequest)
	assert.Nil(t, err)

	AssertBodyEqual(t, response.Body, request.expectedTodos)
}

func TestWhenAddTodoRequestArrivesWithValidRequestItReturnsAddedTodoWithoutError(t *testing.T){
	test := struct {
		expectedTodo todo.Todo
		request todo.CreateTodoRequest
	}{
		expectedTodo: todo.Todo{
			ID:   1,
			Task: "buy some milk",
		},
		request: todo.CreateTodoRequest{
			Task: "buy some milk",
		},

	}
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockService := mocks.NewMockIService(mockController)
	mockService.EXPECT().AddTodo(test.request).Return(test.expectedTodo, nil)

	handler, _ := todo.NewHandler(mockService)
	app := fiber.New()
	handler.RegisterRoutes(app)

	request := MakeRequestWithBody(http.MethodPost, "/todos", test.request)
	response, err := app.Test(request)
	assert.Nil(t, err)

	AssertBodyEqual(t, response.Body, test.expectedTodo)
}

func MakeRequestWithoutBody(method, url string) *http.Request{
	req := httptest.NewRequest(method, url, http.NoBody)
	req.Header.Add("Content-type", "application/json")
	return req
}

func AssertBodyEqual(t *testing.T, responseBody io.Reader,expectedValue interface{}) {
	var actualBody interface{}
	_ = json.NewDecoder(responseBody).Decode(&actualBody)

	expectedBodyAsJSON, _ := json.Marshal(expectedValue)

	var expectedBody interface{}
	_ = json.Unmarshal(expectedBodyAsJSON, &expectedBody)
	assert.Equal(t, expectedBody, actualBody)
}

func MakeRequestWithBody(method, url string, body interface{}) *http.Request{
	bodyAsByte, _ := json.Marshal(body)
	req := httptest.NewRequest(method, url, bytes.NewReader(bodyAsByte))
	req.Header.Add("Content-type", "application/json")
	return req
}