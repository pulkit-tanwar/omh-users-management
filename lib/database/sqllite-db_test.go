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

func TestRetrieveUserSuccssful(t *testing.T) {

	timeNow := time.Now().Format(time.RFC3339)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	client := &database.SQLDbClient{}
	client.SetMockDBInstance(db)

	rows := sqlmock.NewRows([]string{"user_name", "first_name", "last_name", "phone_number", "date_created", "date_modified"}).
		AddRow("123", "pulkit", "lastname", "282954", timeNow, timeNow)

	mock.ExpectQuery("^select user_name, first_name, last_name, phone_number, date_created, date_modified from users").WithArgs("123").WillReturnRows(rows)
	user, err := client.RetrieveUser("123")
	assert.Equal(t, nil, err)
	assert.Equal(t, user.User_Name, "123")
	assert.Equal(t, user.First_Name, "pulkit")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRetrieveUserDBFailed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	client := &database.SQLDbClient{}
	client.SetMockDBInstance(db)

	mock.ExpectQuery("^select user_name, first_name, last_name, phone_number, date_created, date_modified from users").WithArgs("123").WillReturnError(errors.New("error"))
	_, err = client.RetrieveUser("123")
	assert.Equal(t, errors.New("error"), err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
