package todo

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	RegisterRoutes(app *fiber.App)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) (*Handler, error) {
	if service == nil {
		return nil, errors.New("service failure")
	}
	return &Handler{service: service}, nil
}

func (h *Handler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.service.GetTodos()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

func (h *Handler) AddTodo(c *fiber.Ctx) error {
	var request CreateTodoRequest
	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	todo, err := h.service.AddTodo(request)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}

func (h *Handler) RegisterRoutes(app *fiber.App){
	app.Get("/todos", h.GetTodos)
	app.Post("/todos", h.AddTodo)
}


