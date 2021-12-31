package database_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pulkit-tanwar/omh-users-management/lib/database"
	"github.com/stretchr/testify/assert"
)

type mockDBClient struct {
	*sql.DB
	err error
}

func TestDBConnectSuccess(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()
	client := &database.SQLDbClient{}
	client.SetMockDBInstance(db)
	err = client.DBConnect()
	assert.Nil(t, err)
}
