package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"todo-backend/config"
	"todo-backend/todo"
)


type Server struct {
	app *fiber.App
	config config.Server
}

func NewServer(serverConfig config.Server,handlers []todo.IHandler) *Server {
	app := fiber.New()
	app.Use(cors.New())
	for _, item := range handlers {
		item.RegisterRoutes(app)
	}
	server := Server{app: app,config: serverConfig}
	server.AddRoutes()
	return &server
}
func healthCheck(c * fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (s *Server) AddRoutes(){
	s.app.Get("/health", healthCheck)
}

func (s Server) Run() error {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-shutdownChan
		err := s.app.Shutdown()
		if err != nil {
			log.Println("Error on shutdown gracefully")
		}
	}()
	fmt.Printf("Configuration Port is :%s",s.config.Port)
	return s.app.Listen(fmt.Sprintf(":%s", s.config.Port))
}
