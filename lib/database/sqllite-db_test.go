package database_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pulkit-tanwar/omh-users-management/lib/database"
	"github.com/pulkit-tanwar/omh-users-management/lib/model"
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

func TestCreateUserSuccessful(t *testing.T) {

	tm := time.Now()
	currentTime := tm.UTC().String()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	client := &database.SQLDbClient{}
	client.SetMockDBInstance(db)
	mock.ExpectExec("insert into users").
		WithArgs("abc", "firstname", "lastname", "282954", currentTime, currentTime).WillReturnResult(sqlmock.NewResult(1, 1))
	user := model.User{
		User_Name:    "abc",
		First_Name:   "firstname",
		Last_Name:    "lastname",
		Phone_Number: "282954",
		DateCreated:  currentTime,
		DateModified: currentTime,
	}

	if err = client.CreateUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUserDBFailed(t *testing.T) {

	tm := time.Now()
	currentTime := tm.UTC().String()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	client := &database.SQLDbClient{}
	client.SetMockDBInstance(db)
	mock.ExpectExec("insert into users").
		WithArgs("abc", "firstname", "lastname", "282954", currentTime, currentTime).WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(errors.New("error"))
	user := model.User{
		User_Name:    "abc",
		First_Name:   "firstname",
		Last_Name:    "lastname",
		Phone_Number: "282954",
		DateCreated:  currentTime,
		DateModified: currentTime,
	}

	if err = client.CreateUser(user); err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
