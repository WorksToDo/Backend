package postgre

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateNewPostgreDbConnectionRunsSuccessfully(t *testing.T) {
	// change here to CI != "" after add it to pipeline.
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	db, err := NewPostgreDB(DefaultConfig)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}