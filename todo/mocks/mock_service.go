// Code generated by MockGen. DO NOT EDIT.
// Source: todo-backend/todo (interfaces: IService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	todo "todo-backend/todo"

	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// AddTodo mocks base method.
func (m *MockIService) AddTodo(arg0 todo.CreateTodoRequest) (todo.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTodo", arg0)
	ret0, _ := ret[0].(todo.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTodo indicates an expected call of AddTodo.
func (mr *MockIServiceMockRecorder) AddTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTodo", reflect.TypeOf((*MockIService)(nil).AddTodo), arg0)
}

// GetTodos mocks base method.
func (m *MockIService) GetTodos() ([]todo.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodos")
	ret0, _ := ret[0].([]todo.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodos indicates an expected call of GetTodos.
func (mr *MockIServiceMockRecorder) GetTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodos", reflect.TypeOf((*MockIService)(nil).GetTodos))
}
