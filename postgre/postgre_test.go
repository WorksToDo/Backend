package postgre

import (
	"github.com/stretchr/testify/assert"
	"testing"
	config "todo-backend/config"
)

func TestCreateNewPostgreDbConnectionRunsSuccessfully(t *testing.T) {
	dbConfig := config.DB{
		Username: "postgres",
		Password: "todo123",
		Driver:   "postgres",
		DBName:   "postgres",
		Host:     "localhost",
		Port:     "5432",
	}
	db, err := NewPostgreDB(dbConfig)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}