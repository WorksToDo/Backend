package main

import (
	"fmt"
	"os"
	"todo-backend/config"
	"todo-backend/postgre"
	"todo-backend/server"
	"todo-backend/todo"
)


func main(){
	dbConfig := config.DB{
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := postgre.NewPostgreDB(dbConfig)
	if err != nil {
		fmt.Printf("DB error: %s",err)
	}
	todoRepo := todo.NewRepository(db)
	todoService, _ := todo.NewService(todoRepo)
	todoHandler, _ := todo.NewHandler(todoService)

	handlers := []todo.IHandler{
		todoHandler,
	}
	server := server.NewServer(config.Server{Port: os.Getenv("SERVER_PORT")},handlers)
	server.Run()
}
