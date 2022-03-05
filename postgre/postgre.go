package postgre

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todo-backend/config"
)

var DefaultConfig = config.DB{
	Username: "postgres",
	Password: "assign1",
	Driver:   "postgres",
	DBName:   "postgres",
	Host:     "localhost",
	Port:     "5432",
}

func NewPostgreDB(config config.DB) (*gorm.DB, error){
	dsn := createDSN(config)
	db,err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDSN(config config.DB) string{
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.DBName)
}
