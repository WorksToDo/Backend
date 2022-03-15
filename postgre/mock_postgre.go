package postgre

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockPostgreDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil
	}
	var gormDB *gorm.DB
	gormDB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}))
	if err != nil {
		return nil, nil
	}
	return gormDB, mock
}