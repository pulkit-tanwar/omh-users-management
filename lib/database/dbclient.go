package database

import "github.com/pulkit-tanwar/omh-users-management/lib/model"

// DB - Database Client var
var DB DBclient

// DBclient - Database Client interface
type DBclient interface {
	DBConnect() error
	CreateUser(user model.User) error
	RetrieveUser(string) (model.User, error)
	ModifyUserDetails(model.User) (model.User, error)
	DeleteUser(string) error
	GetAllUsers() ([]model.User, error)
}
