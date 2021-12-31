package database

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/fatih/structs"
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

func (client *SQLDbClient) RetrieveUser(userName string) (model.User, error) {
	query := "select user_name, first_name, last_name, phone_number, date_created, date_modified from users where user_name = $1"

	row := client.db.QueryRow(query, userName)
	user := model.User{}

	err := row.Scan(
		&user.User_Name,
		&user.First_Name,
		&user.Last_Name,
		&user.Phone_Number,
		&user.DateCreated,
		&user.DateModified)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}

func (client *SQLDbClient) ModifyUserDetails(user model.User) (model.User, error) {
	query := "update users set date_modified = strftime('%d/%m/%Y, %H:%M','now'),"
	m := structs.Map(user)
	var values []interface{}
	j := 0
	response := model.User{}
	for i := range m {
		if v := m[i]; v != "" && i != "User_Name" && v != false {
			j++
			query = query + strings.ToLower(i) + "=$" + strconv.Itoa(j) + ","
			values = append(values, v)
		}
	}
	log.Debugf("Query to update users details: %s", query)

	values = append(values, m["User_Name"])
	j++
	query = query[:len(query)-1] + " WHERE " + "user_name" + "=$" + strconv.Itoa(j)
	query = query + " RETURNING " + "user_Name, first_name, last_name, phone_number"

	row := client.db.QueryRow(query, values...)
	err := row.Scan(
		&response.User_Name,
		&response.First_Name,
		&response.Last_Name,
		&response.Phone_Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}
		return response, err
	}
	return response, nil
}

func (client *SQLDbClient) DeleteUser(userName string) error {
	query := "delete from users where user_name = $1"
	_, err := client.db.Exec(query, userName)
	if err != nil {
		return err
	}
	return nil
}

//GetAllUsers - Fetch all users info
func (client *SQLDbClient) GetAllUsers() ([]model.User, error) {
	query := "SELECT user_name, first_name, last_name from users;"

	rows, err := client.db.Query(query)
	userList := []model.User{}

	if err != nil {
		return userList, err
	}

	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(
			&user.User_Name,
			&user.First_Name,
			&user.Last_Name,
		)
		if err != nil {
			return userList, err
		}

		userList = append(userList, user)
	}

	return userList, nil
}
