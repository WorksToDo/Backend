package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todo-backend/config"
	"todo-backend/todo"
)

func TestCreateNewServerWithGivenConfigs(t *testing.T) {
	server := NewServer(config.Server{Port: "4000"}, nil)
	assert.NotNil(t, server)
}

func TestIsServerHealthy(t *testing.T) {
	server := NewServer(config.Server{Port: "4000"}, nil)
	req := httptest.NewRequest(fiber.MethodGet, "/health", http.NoBody)
	req.Header.Add("Content-type", "application/json")
	resp,err := server.app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode,fiber.StatusOK)
}

func TestServerRunsSuccessfullyOnSpecifiedPort(t *testing.T) {
	port,_ := freeport.GetFreePort()
	server := NewServer(config.Server{Port: fmt.Sprintf("%d",port)},[]todo.IHandler{})

	go func() {
		if err := server.Run(); err!=nil {
			t.Fail()
		}
	}()
	time.Sleep(50 * time.Millisecond)
	req, err := http.NewRequest(fiber.MethodGet,fmt.Sprintf("http://localhost:%s/health",server.config.Port),http.NoBody)
	assert.Nil(t, err)
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode,fiber.StatusOK)
}