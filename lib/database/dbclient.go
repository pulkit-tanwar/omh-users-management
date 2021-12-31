package database

// DB - Database Client var
var DB DBclient

// DBclient - Database Client interface
type DBclient interface {
	DBConnect() error
}
