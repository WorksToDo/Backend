package main

import (
	"fmt"
	"os"
	"todo-backend/config"
	"todo-backend/postgre"
)


func main(){
	dbconfig := config.DB{
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		//TablePrefix: os.Getenv("DB_TABLE_PREFIX"),
	}
	db, err := postgre.NewPostgreDB(dbconfig)
	if err != nil {
		fmt.Printf("DB error: %s",err)
	}

}
