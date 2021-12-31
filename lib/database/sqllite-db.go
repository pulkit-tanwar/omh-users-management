package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pulkit-tanwar/omh-users-management/lib/model"
	log "github.com/sirupsen/logrus"
)

// SQLDbClient - sqlite3 Client
type SQLDbClient struct {
	db *sql.DB
}

// SetMockDBInstance - Set Mock DB instance
func (client *SQLDbClient) SetMockDBInstance(db *sql.DB) {
	client.db = db
}

//DBConnect - Create a connection to the sqlite3 db
func (client *SQLDbClient) DBConnect() error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Errorf("failed loading sqlite3 parameteres. Error :%+v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("failed to ping sqlite3 db. Error :%+v", err)
		return err
	}

	client.db = db
	log.Info("Successfully connected to sqlite3 database.")

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users ( user_name VARCHAR NOT NULL PRIMARY KEY, first_name VARCHAR NOT NULL, last_name VARCHAR NOT NULL , phone_number VARCHAR , date_created TIMESTAMP NOT NULL, date_modified TIMESTAMP NOT NULL )")
	if err != nil {
		log.Errorf("Failed to prepare db statement. Error :%+v", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Errorf("Error while executing statement. Error :%+v", err)
		return err
	}
	return nil
}

func (client *SQLDbClient) CreateUser(user model.User) error {
	query := "insert into users (user_name, first_name, last_name, phone_number, date_created, date_modified) values ($1, $2, $3, $4, $5, $6);"
	_, err := client.db.Exec(query, user.User_Name, user.First_Name, user.Last_Name, user.Phone_Number, user.DateCreated, user.DateModified)
	if err != nil {
		return err
	}
	return nil
}
