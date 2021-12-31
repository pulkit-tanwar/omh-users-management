# Users Management System

A go microservice that can enables us to create, modify, fetch, and delete users

## Usage

### To Run the application
```bash
  $ go run main.go serve
```

### To Run unit tests
```bash
  $ go test ./...
```

### To Run the application with custom environment variables
```bash
  $ go run main.go serve --env STAGE --host 0.0.0.0 --port 4000
```

## Development
```bash
  $ make dep           # install dependencies
  $ make test          # run unit tests
  $ make cover         # run code coverage report service (http://localhost:3001)
  $ make run           # run the service
  $ make build         # compile standalone binary for docker container
  $ make image         # build docker image  
```

## To create a user
Endpoint: http://localhost:3000/api/v1/users

HTTP Header: Content-Type : application/json


HTTP Request Body
```bash
{
    "username": "Elon@example.com",
    "firstname": "Elon",
    "lastname": "Musk"
} 
```

HTTP Response Body
```bash
{
    "userName": "Elon@example.com",
    "firstName": "Elon",
    "lastName": "Musk",
    "dateCreated": "2022-01-01T04:03:06+05:30",
    "dateModified": "2022-01-01T04:03:06+05:30"
}
```

