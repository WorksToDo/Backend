package postgre

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewPostgreDbConnectionRunsSuccessfully(t *testing.T) {
	db, err := NewPostgreDB(DefaultConfig)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}